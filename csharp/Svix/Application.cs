﻿// this file is @generated
using System;
using System.Collections.Generic;
using System.Net;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Logging;
using Svix.Abstractions;
using Svix.Api;
using Svix.Client;
using Svix.Model;
using Svix.Models;

namespace Svix
{
    public partial class ApplicationListOptions
    {
        public int? Limit { get; set; }
        public string? Iterator { get; set; }
        public Ordering? Order { get; set; }
    }

    public partial class ApplicationCreateOptions
    {
        public string? IdempotencyKey { get; set; }
    }

    public sealed class Application : SvixResourceBase
    {
        private readonly ApplicationApi _applicationApi;

        public Application(ISvixClient svixClient, ApplicationApi applicationApi)
            : base(svixClient)
        {
            _applicationApi =
                applicationApi ?? throw new ArgumentNullException(nameof(applicationApi));
        }

        public async Task<ListResponseApplicationOut> ListAsync(
            ApplicationListOptions options = null,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                var response = await _applicationApi.V1ApplicationListWithHttpInfoAsync(
                    limit: options?.Limit,
                    iterator: options?.Iterator,
                    order: options?.Order,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(ListAsync)} failed");

                if (Throw)
                    throw;
                return new ListResponseApplicationOut();
            }
        }

        public ListResponseApplicationOut List(ApplicationListOptions options = null)
        {
            try
            {
                var response = _applicationApi.V1ApplicationListWithHttpInfo(
                    limit: options?.Limit,
                    iterator: options?.Iterator,
                    order: options?.Order
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(List)} failed");

                if (Throw)
                    throw;
                return new ListResponseApplicationOut();
            }
        }

        public async Task<ApplicationOut> CreateAsync(
            ApplicationIn applicationIn,
            ApplicationCreateOptions options = null,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                applicationIn =
                    applicationIn ?? throw new ArgumentNullException(nameof(applicationIn));
                var response = await _applicationApi.V1ApplicationCreateWithHttpInfoAsync(
                    applicationIn: applicationIn,
                    idempotencyKey: options?.IdempotencyKey,
                    getIfExists: false,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(CreateAsync)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public ApplicationOut Create(
            ApplicationIn applicationIn,
            ApplicationCreateOptions options = null
        )
        {
            try
            {
                applicationIn =
                    applicationIn ?? throw new ArgumentNullException(nameof(applicationIn));
                var response = _applicationApi.V1ApplicationCreateWithHttpInfo(
                    applicationIn: applicationIn,
                    idempotencyKey: options?.IdempotencyKey,
                    getIfExists: false
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Create)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public async Task<ApplicationOut> GetOrCreateAsync(
            ApplicationIn applicationIn,
            ApplicationCreateOptions options = null,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                var response = await _applicationApi.V1ApplicationCreateWithHttpInfoAsync(
                    applicationIn: applicationIn,
                    idempotencyKey: options?.IdempotencyKey,
                    getIfExists: true,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(CreateAsync)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public ApplicationOut GetOrCreate(
            ApplicationIn applicationIn,
            ApplicationCreateOptions options = null
        )
        {
            try
            {
                var response = _applicationApi.V1ApplicationCreateWithHttpInfo(
                    applicationIn: applicationIn,
                    idempotencyKey: options?.IdempotencyKey,
                    getIfExists: true
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Create)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public async Task<ApplicationOut> GetAsync(
            string appId,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                var response = await _applicationApi.V1ApplicationGetWithHttpInfoAsync(
                    appId,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(GetAsync)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public ApplicationOut Get(string appId)
        {
            try
            {
                var response = _applicationApi.V1ApplicationGetWithHttpInfo(appId);
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Get)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public async Task<ApplicationOut> UpdateAsync(
            string appId,
            ApplicationIn applicationIn,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                applicationIn =
                    applicationIn ?? throw new ArgumentNullException(nameof(applicationIn));
                var response = await _applicationApi.V1ApplicationUpdateWithHttpInfoAsync(
                    appId,
                    applicationIn: applicationIn,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(UpdateAsync)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public ApplicationOut Update(string appId, ApplicationIn applicationIn)
        {
            try
            {
                applicationIn =
                    applicationIn ?? throw new ArgumentNullException(nameof(applicationIn));
                var response = _applicationApi.V1ApplicationUpdateWithHttpInfo(
                    appId,
                    applicationIn: applicationIn
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Update)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public async Task<bool> DeleteAsync(
            string appId,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                var response = await _applicationApi.V1ApplicationDeleteWithHttpInfoAsync(
                    appId,
                    cancellationToken: cancellationToken
                );
                return response.StatusCode == HttpStatusCode.NoContent;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(DeleteAsync)} failed");

                if (Throw)
                    throw;
                return false;
            }
        }

        public bool Delete(string appId)
        {
            try
            {
                var response = _applicationApi.V1ApplicationDeleteWithHttpInfo(appId);
                return response.StatusCode == HttpStatusCode.NoContent;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Delete)} failed");

                if (Throw)
                    throw;
                return false;
            }
        }

        public async Task<ApplicationOut> PatchAsync(
            string appId,
            ApplicationPatch applicationPatch,
            CancellationToken cancellationToken = default
        )
        {
            try
            {
                applicationPatch =
                    applicationPatch ?? throw new ArgumentNullException(nameof(applicationPatch));
                var response = await _applicationApi.V1ApplicationPatchWithHttpInfoAsync(
                    appId,
                    applicationPatch: applicationPatch,
                    cancellationToken: cancellationToken
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(PatchAsync)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }

        public ApplicationOut Patch(string appId, ApplicationPatch applicationPatch)
        {
            try
            {
                applicationPatch =
                    applicationPatch ?? throw new ArgumentNullException(nameof(applicationPatch));
                var response = _applicationApi.V1ApplicationPatchWithHttpInfo(
                    appId,
                    applicationPatch: applicationPatch
                );
                return response.Data;
            }
            catch (ApiException e)
            {
                Logger?.LogError(e, $"{nameof(Patch)} failed");

                if (Throw)
                    throw;
                return new ApplicationOut();
            }
        }
    }
}
