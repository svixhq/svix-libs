# frozen_string_literal: true

module Svix
    class Webhook
        def initialize(secret)
            @secret = secret
        end

        def verify(payload, headers)
            msgId = headers["svix-id"]
            msgSignature = headers["svix-signature"]
            msgTimestamp = headers["svix-timestamp"]
            puts msgId
            if !msgSignature || !msgId || !msgTimestamp
                raise WebhookVerificationError, "Missing required headers"
            end

            toSign = "${msgId}.${msgTimestamp}.${payload}"
            signature = OpenSSL::HMAC.hexdigest(OpenSSL::Digest.new("sha256"), secret, toSign)

            passedSignatures = msgSignature.split(" ");

            passedSignatures.each do |versionedSignature|
                version, signature = versionedSignature.split(',', 2)
                if version != "v1" 
                    next
                end
                if signature == expectedSignature
                    return JSON.parse(payload, symbolize_names: true)
                end
            end
            raise WebhookVerificationError, "No matching signature found"
        end
    end
end