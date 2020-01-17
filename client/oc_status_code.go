package client

// This code is taken from OpenTelemetry project in order to avoid adding the whole project as a dependency.

// https://github.com/googleapis/googleapis/blob/bee79fbe03254a35db125dc6d2f1e9b752b390fe/google/rpc/code.proto#L33-L186
const (
	OCOK                 = 0
	OCCancelled          = 1
	OCUnknown            = 2
	OCInvalidArgument    = 3
	OCDeadlineExceeded   = 4
	OCNotFound           = 5
	OCAlreadyExists      = 6
	OCPermissionDenied   = 7
	OCResourceExhausted  = 8
	OCFailedPrecondition = 9
	OCAborted            = 10
	OCOutOfRange         = 11
	OCUnimplemented      = 12
	OCInternal           = 13
	OCUnavailable        = 14
	OCDataLoss           = 15
	OCUnauthenticated    = 16
)

const (
	defaultRateLimitingBackoffSeconds = 8
	headerAccessToken                 = "X-SF-Token"
	headerRetryAfter                  = "Retry-After"
	headerContentEncoding             = "Content-Encoding"
	headerContentType                 = "Content-Type"
	headerValueGZIP                   = "gzip"
	headerValueXProtobuf              = "application/x-protobuf"
	maxHTTPBodyReadBytes              = 256 << 10
)

var httpToOCCodeMap = map[int32]int32{
	401: OCUnauthenticated,
	403: OCPermissionDenied,
	404: OCNotFound,
	429: OCResourceExhausted,
	499: OCCancelled,
	501: OCUnimplemented,
	503: OCUnavailable,
	504: OCDeadlineExceeded,
}

// OCStatusCodeFromHTTP takes an HTTP status code and return the appropriate OpenTelemetry status code
// See: https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/data-http.md
func OCStatusCodeFromHTTP(code int32) int32 {
	if code >= 100 && code < 400 {
		return OCOK
	}
	if rvCode, ok := httpToOCCodeMap[code]; ok {
		return rvCode
	}
	if code >= 400 && code < 500 {
		return OCInvalidArgument
	}
	if code >= 500 && code < 600 {
		return OCInternal
	}
	return OCUnknown
}
