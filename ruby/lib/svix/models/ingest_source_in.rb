# frozen_string_literal: true
# This file is @generated
require "json"

require_relative "./adobe_sign_config"
require_relative "./cron_config"
require_relative "./docusign_config"
require_relative "./github_config"
require_relative "./hubspot_config"
require_relative "./segment_config"
require_relative "./shopify_config"
require_relative "./slack_config"
require_relative "./stripe_config"
require_relative "./svix_config"
require_relative "./zoom_config"

module Svix
  class IngestSourceInConfig
    class GenericWebhook

      def serialize
        Hash.new
      end

      def self.deserialize(attributes = {})
        new
      end
      # Serializes the object to a json string
      # @return String
      def to_json
        JSON.dump(serialize)
      end
    end

    class Cron < CronConfig
    end

    class AdobeSign < AdobeSignConfig
    end

    class Beehiiv < SvixConfig
    end

    class Brex < SvixConfig
    end

    class Clerk < SvixConfig
    end

    class Docusign < DocusignConfig
    end

    class Github < GithubConfig
    end

    class Guesty < SvixConfig
    end

    class Hubspot < HubspotConfig
    end

    class IncidentIo < SvixConfig
    end

    class Lithic < SvixConfig
    end

    class Nash < SvixConfig
    end

    class Pleo < SvixConfig
    end

    class Replicate < SvixConfig
    end

    class Resend < SvixConfig
    end

    class Safebase < SvixConfig
    end

    class Sardine < SvixConfig
    end

    class Segment < SegmentConfig
    end

    class Shopify < ShopifyConfig
    end

    class Slack < SlackConfig
    end

    class Stripe < StripeConfig
    end

    class Stych < SvixConfig
    end

    class Svix < SvixConfig
    end

    class Zoom < ZoomConfig
    end
  end

  class IngestSourceIn
    attr_accessor :name
    attr_accessor :uid
    attr_accessor :config

    ALL_FIELD ||= ["name", "uid", "config"].freeze
    private_constant :ALL_FIELD
    TYPE_TO_NAME = {
      IngestSourceInConfig::GenericWebhook => "generic-webhook",
      IngestSourceInConfig::Cron => "cron",
      IngestSourceInConfig::AdobeSign => "adobe-sign",
      IngestSourceInConfig::Beehiiv => "beehiiv",
      IngestSourceInConfig::Brex => "brex",
      IngestSourceInConfig::Clerk => "clerk",
      IngestSourceInConfig::Docusign => "docusign",
      IngestSourceInConfig::Github => "github",
      IngestSourceInConfig::Guesty => "guesty",
      IngestSourceInConfig::Hubspot => "hubspot",
      IngestSourceInConfig::IncidentIo => "incident-io",
      IngestSourceInConfig::Lithic => "lithic",
      IngestSourceInConfig::Nash => "nash",
      IngestSourceInConfig::Pleo => "pleo",
      IngestSourceInConfig::Replicate => "replicate",
      IngestSourceInConfig::Resend => "resend",
      IngestSourceInConfig::Safebase => "safebase",
      IngestSourceInConfig::Sardine => "sardine",
      IngestSourceInConfig::Segment => "segment",
      IngestSourceInConfig::Shopify => "shopify",
      IngestSourceInConfig::Slack => "slack",
      IngestSourceInConfig::Stripe => "stripe",
      IngestSourceInConfig::Stych => "stych",
      IngestSourceInConfig::Svix => "svix",
      IngestSourceInConfig::Zoom => "zoom"
    }
    private_constant :TYPE_TO_NAME
    NAME_TO_TYPE = TYPE_TO_NAME.invert
    private_constant :NAME_TO_TYPE

    def initialize(attributes = {})
      unless attributes.is_a?(Hash)
        fail(
          ArgumentError,
          "The input argument (attributes) must be a hash in `Svix::IngestSourceIn` new method"
        )
      end

      attributes.each do |k, v|
        unless ALL_FIELD.include?(k.to_s)
          fail(ArgumentError, "The field #{k} is not part of Svix::IngestSourceIn")
        end

        if k == "config"
          unless TYPE_TO_NAME.key?(v.class)
            fail(ArgumentError, "The field #{k} can't be a `#{v.class}` expected one of #{TYPE_TO_NAME.keys}")
          end

          instance_variable_set("@__enum_discriminator", TYPE_TO_NAME[v.class])
        end

        instance_variable_set("@#{k}", v)
        instance_variable_set("@__#{k}_is_defined", true)
      end

      if @__enum_discriminator.nil?
        fail(ArgumentError, "Required config field was not set")
      end
    end

    def self.deserialize(attributes = {})
      attributes = attributes.transform_keys(&:to_s)
      attrs = Hash.new
      attrs["name"] = attributes["name"]
      attrs["uid"] = attributes["uid"]
      unless NAME_TO_TYPE.key?(attributes["type"])
        fail(ArgumentError, "Invalid type `#{attributes["type"]}` expected on of #{NAME_TO_TYPE.keys}")
      end

      unless attributes.key?("config")
        fail(ArgumentError, "Missing required field config")
      end

      attrs["config"] = NAME_TO_TYPE[attributes["type"]].deserialize(attributes["config"])
      new(attrs)
    end

    def serialize
      out = Hash.new
      out["name"] = Svix::serialize_primitive(@name) if @name
      out["uid"] = Svix::serialize_primitive(@uid) if @uid
      out["type"] = @__enum_discriminator
      out["config"] = @config.serialize
      out
    end

    # Serializes the object to a json string
    # @return String
    def to_json
      JSON.dump(serialize)
    end

  end
end
