// This file is @generated
package com.svix.models;

import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.svix.Utils;
import lombok.*;

import java.util.HashMap;
import java.util.Map;

@ToString
@EqualsAndHashCode
public abstract class IngestSourceInConfig {
    @JsonIgnore
    public String getVariantName() {
        VariantName annotation = this.getClass().getAnnotation(VariantName.class);
        return annotation != null ? annotation.value() : null;
    }

    public abstract JsonNode toJsonNode();

    @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("generic-webhook")
    public static class GenericWebhook extends IngestSourceInConfig {
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().createObjectNode();
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("cron")
    public static class Cron extends IngestSourceInConfig {
        private final CronConfig cron;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(cron);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("adobe-sign")
    public static class AdobeSign extends IngestSourceInConfig {
        private final AdobeSignConfig adobeSign;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(adobeSign);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("beehiiv")
    public static class Beehiiv extends IngestSourceInConfig {
        private final SvixConfig beehiiv;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(beehiiv);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("brex")
    public static class Brex extends IngestSourceInConfig {
        private final SvixConfig brex;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(brex);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("clerk")
    public static class Clerk extends IngestSourceInConfig {
        private final SvixConfig clerk;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(clerk);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("docusign")
    public static class Docusign extends IngestSourceInConfig {
        private final DocusignConfig docusign;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(docusign);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("github")
    public static class Github extends IngestSourceInConfig {
        private final GithubConfig github;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(github);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("guesty")
    public static class Guesty extends IngestSourceInConfig {
        private final SvixConfig guesty;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(guesty);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("hubspot")
    public static class Hubspot extends IngestSourceInConfig {
        private final HubspotConfig hubspot;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(hubspot);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("incident-io")
    public static class IncidentIo extends IngestSourceInConfig {
        private final SvixConfig incidentIo;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(incidentIo);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("lithic")
    public static class Lithic extends IngestSourceInConfig {
        private final SvixConfig lithic;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(lithic);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("nash")
    public static class Nash extends IngestSourceInConfig {
        private final SvixConfig nash;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(nash);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("pleo")
    public static class Pleo extends IngestSourceInConfig {
        private final SvixConfig pleo;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(pleo);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("replicate")
    public static class Replicate extends IngestSourceInConfig {
        private final SvixConfig replicate;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(replicate);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("resend")
    public static class Resend extends IngestSourceInConfig {
        private final SvixConfig resend;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(resend);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("safebase")
    public static class Safebase extends IngestSourceInConfig {
        private final SvixConfig safebase;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(safebase);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("sardine")
    public static class Sardine extends IngestSourceInConfig {
        private final SvixConfig sardine;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(sardine);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("segment")
    public static class Segment extends IngestSourceInConfig {
        private final SegmentConfig segment;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(segment);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("shopify")
    public static class Shopify extends IngestSourceInConfig {
        private final ShopifyConfig shopify;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(shopify);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("slack")
    public static class Slack extends IngestSourceInConfig {
        private final SlackConfig slack;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(slack);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("stripe")
    public static class Stripe extends IngestSourceInConfig {
        private final StripeConfig stripe;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(stripe);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("stych")
    public static class Stych extends IngestSourceInConfig {
        private final SvixConfig stych;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(stych);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("svix")
    public static class Svix extends IngestSourceInConfig {
        private final SvixConfig svix;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(svix);
            }
    }
    @Getter
        @Setter
        @AllArgsConstructor
        @ToString
    @EqualsAndHashCode(callSuper = false)
    @VariantName("zoom")
    public static class Zoom extends IngestSourceInConfig {
        private final ZoomConfig zoom;
        @Override public JsonNode toJsonNode() {
            return Utils.getObjectMapper().valueToTree(zoom);
            }
    }
    @FunctionalInterface
    private interface TypeFactory {
        IngestSourceInConfig create(JsonNode config);
    }
    private static final Map<String, TypeFactory> TY_M = new HashMap<>();
    private static final ObjectMapper m = Utils.getObjectMapper();
    static {
        TY_M.put("generic-webhook", c -> new GenericWebhook());
            TY_M.put("cron", c -> new Cron(m.convertValue(c, CronConfig.class)));
            TY_M.put("adobe-sign", c -> new AdobeSign(m.convertValue(c, AdobeSignConfig.class)));
            TY_M.put("beehiiv", c -> new Beehiiv(m.convertValue(c, SvixConfig.class)));
            TY_M.put("brex", c -> new Brex(m.convertValue(c, SvixConfig.class)));
            TY_M.put("clerk", c -> new Clerk(m.convertValue(c, SvixConfig.class)));
            TY_M.put("docusign", c -> new Docusign(m.convertValue(c, DocusignConfig.class)));
            TY_M.put("github", c -> new Github(m.convertValue(c, GithubConfig.class)));
            TY_M.put("guesty", c -> new Guesty(m.convertValue(c, SvixConfig.class)));
            TY_M.put("hubspot", c -> new Hubspot(m.convertValue(c, HubspotConfig.class)));
            TY_M.put("incident-io", c -> new IncidentIo(m.convertValue(c, SvixConfig.class)));
            TY_M.put("lithic", c -> new Lithic(m.convertValue(c, SvixConfig.class)));
            TY_M.put("nash", c -> new Nash(m.convertValue(c, SvixConfig.class)));
            TY_M.put("pleo", c -> new Pleo(m.convertValue(c, SvixConfig.class)));
            TY_M.put("replicate", c -> new Replicate(m.convertValue(c, SvixConfig.class)));
            TY_M.put("resend", c -> new Resend(m.convertValue(c, SvixConfig.class)));
            TY_M.put("safebase", c -> new Safebase(m.convertValue(c, SvixConfig.class)));
            TY_M.put("sardine", c -> new Sardine(m.convertValue(c, SvixConfig.class)));
            TY_M.put("segment", c -> new Segment(m.convertValue(c, SegmentConfig.class)));
            TY_M.put("shopify", c -> new Shopify(m.convertValue(c, ShopifyConfig.class)));
            TY_M.put("slack", c -> new Slack(m.convertValue(c, SlackConfig.class)));
            TY_M.put("stripe", c -> new Stripe(m.convertValue(c, StripeConfig.class)));
            TY_M.put("stych", c -> new Stych(m.convertValue(c, SvixConfig.class)));
            TY_M.put("svix", c -> new Svix(m.convertValue(c, SvixConfig.class)));
            TY_M.put("zoom", c -> new Zoom(m.convertValue(c, ZoomConfig.class)));
            }

    public static IngestSourceInConfig fromTypeAndConfig(String type, JsonNode config) {
        TypeFactory factory = TY_M.get(type);
        if (factory == null) {
            throw new IllegalArgumentException("Unknown type: " + type);
        }
        return factory.create(config);
    }

}
