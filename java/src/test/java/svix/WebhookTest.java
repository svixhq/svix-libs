package svix;

import static org.junit.Assert.assertThrows;
import static org.junit.Assert.assertTrue;

import java.net.http.HttpHeaders;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.function.BiPredicate;
import org.junit.Test;
import org.junit.function.ThrowingRunnable;

public class WebhookTest {
	private static final String DEFAULT_SECRET = "MfKQ9r8GKYqrTwjUPD8ILPZIo2LaLaSw";
	private static final String VALID_PAYLOAD = "{\"test\": 2432232314}";

	private final Webhook webhook = new Webhook(DEFAULT_SECRET);

	@Test
	public void verify_validPayloadAndheader_returnTrue() throws WebhookVerificationError {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		HttpHeaders headers = HttpHeaders.of(headerMap, new DefaultBiPredicate());

		assertTrue(webhook.verify(VALID_PAYLOAD, headers));
	}

	@Test
	public void verify_validPayloadAndheaderWithMultiplePayloads_returnTrue() throws WebhookVerificationError {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.put(Webhook.MSG_SIGNATURE_KEY,
		    Arrays.asList("v2,tmIKtSyZlE3uFJELVlNIOLJ1OE=", "v1,g0hM9SsE+OTPJTGt/tmIKtSyZlE3uFJELVlNIOLJ1OE="));

		HttpHeaders headers = HttpHeaders.of(headerMap, new DefaultBiPredicate());

		assertTrue(webhook.verify(VALID_PAYLOAD, headers));
	}

	@Test
	public void verify_missingId_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.remove(Webhook.MSG_ID_KEY);

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	@Test
	public void verify_missingTimestamp_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.remove(Webhook.MSG_TIMESTAMP_KEY);

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	@Test
	public void verify_missingSignature_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.remove(Webhook.MSG_SIGNATURE_KEY);

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	@Test
	public void verify_signatureWithDifferentVersion_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.put(Webhook.MSG_SIGNATURE_KEY, Arrays.asList("v2,g0hM9SsE+OTPJTGt/tmIKtSyZlE3uFJELVlNIOLJ1OE="));

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	@Test
	public void verify_missingPartsInSignature_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.put(Webhook.MSG_SIGNATURE_KEY, Arrays.asList("invalid_signature"));

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	@Test
	public void verify_signatureMismatch_throwsException() {
		Map<String, List<String>> headerMap = createValidHeadersMap();
		headerMap.put(Webhook.MSG_SIGNATURE_KEY, Arrays.asList("v1,invalid_signature"));

		assertThrows(WebhookVerificationError.class, verify(headerMap));
	}

	private Map<String, List<String>> createValidHeadersMap() {
		HashMap<String, List<String>> map = new HashMap<String, List<String>>();
		map.put(Webhook.MSG_ID_KEY, Arrays.asList("msg_p5jXN8AQM9LWM0D4loKWxJek"));
		map.put(Webhook.MSG_TIMESTAMP_KEY, Arrays.asList("1614265330"));
		map.put(Webhook.MSG_SIGNATURE_KEY, Arrays.asList("v1,g0hM9SsE+OTPJTGt/tmIKtSyZlE3uFJELVlNIOLJ1OE="));
		return map;
	}

	private ThrowingRunnable verify(final Map<String, List<String>> headerMap) {
		return new ThrowingRunnable() {
			@Override
			public void run() throws WebhookVerificationError {
				HttpHeaders headers = HttpHeaders.of(headerMap, new DefaultBiPredicate());
				webhook.verify(VALID_PAYLOAD, headers);
			}
		};
	}

	private class DefaultBiPredicate implements BiPredicate<String, String> {
		@Override
		public boolean test(String arg0, String arg1) {
			return true;
		}
	}
}