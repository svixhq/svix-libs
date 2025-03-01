# frozen_string_literal: true
# This file is @generated
require "json"

module Svix
  class IngestEndpointSecretOut
    attr_accessor :key

    ALL_FIELD ||= ["key"].freeze
    private_constant :ALL_FIELD

    def initialize(attributes = {})
      unless attributes.is_a?(Hash)
        fail(
          ArgumentError,
          "The input argument (attributes) must be a hash in `Svix::IngestEndpointSecretOut` new method"
        )
      end

      attributes.each do |k, v|
        unless ALL_FIELD.include?(k.to_s)
          fail(ArgumentError, "The field #{k} is not part of Svix::IngestEndpointSecretOut")
        end

        instance_variable_set("@#{k}", v)
        instance_variable_set("@__#{k}_is_defined", true)
      end
    end

    def self.deserialize(attributes = {})
      attributes = attributes.transform_keys(&:to_s)
      attrs = Hash.new
      attrs["key"] = attributes["key"]
      new(attrs)
    end

    def serialize
      out = Hash.new
      out["key"] = Svix::serialize_primitive(@key) if @key
      out
    end

    # Serializes the object to a json string
    # @return String
    def to_json
      JSON.dump(serialize)
    end
  end
end
