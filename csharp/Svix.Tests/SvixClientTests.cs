﻿using System;
using System.Collections.Generic;
using Microsoft.Extensions.Logging.Abstractions;
using Svix.Models;
using Xunit;

namespace Svix.Tests;

public class SvixClientTests
{
    [Fact]
    public void Constructor_WhenCalled_DoesNotNeedLogger()
    {
        var sut = new SvixClient("", new SvixOptions("http://some.url"));

        Assert.NotNull(sut);
    }

    [Fact]
    public void Constructor_WhenCalled_AcceptsLogger()
    {
        var sut = new SvixClient("", new SvixOptions("http://some.url"), new NullLogger<SvixClient>());

        Assert.NotNull(sut);
    }

    [IgnoreIfClientVarsUnset]
    public void KitchenSink_SeemsToWorkOkay()
    {
        var token = System.Environment.GetEnvironmentVariable("SVIX_TOKEN");
        var url = System.Environment.GetEnvironmentVariable("SVIX_SERVER_URL");
        var client = new SvixClient(token, new SvixOptions(url));

        var app = client.Application.Create(new ApplicationIn { Name = "App1" });
        try
        {
            client.EventType.Create(new EventTypeIn { Name = "event.started", Description = "Something started" });
        }
        catch (ApiException e)
        {
            // We expect conflicts, but any other status is an error
            Assert.Equal(409, e.ErrorCode);
        }

        try
        {
            client.EventType.Create(new EventTypeIn { Name = "event.ended", Description = "Something started" });
        }
        catch (ApiException e)
        {
            // We expect conflicts, but any other status is an error
            Assert.Equal(409, e.ErrorCode);
        }

        var ep = client.Endpoint.Create(app.Id,
            new EndpointIn { Url = "https://example.svix.com/", Channels = new List<string> { "ch0", "ch1" } });

        ep.Channels.Sort();
        Assert.Equal(new List<string> { "ch0", "ch1" }, ep.Channels);
        Assert.Null(ep.FilterTypes);

        var epPatched = client.Endpoint.Patch(app.Id, ep.Id,
            new EndpointPatch { FilterTypes = new List<string> { "event.started", "event.ended" } });
        epPatched.Channels.Sort();
        epPatched.FilterTypes.Sort();
        Assert.Equal(new List<string> { "ch0", "ch1" }, epPatched.Channels);
        Assert.Equal(new List<string> { "event.ended", "event.started" }, epPatched.FilterTypes);

        // Should not throw an exception if the serialization code handles empty bodies properly
        client.Endpoint.Delete(app.Id, ep.Id);
    }
}