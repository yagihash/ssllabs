package ssllabstest

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func NewRouterGetStatusCodes(t *testing.T) *mux.Router {
	t.Helper()

	h := NewHandler(t, SampleResponseGetStatusCodesBytes())
	r := NewRouter(t, http.MethodGet, "/getStatusCodes", h)

	return r
}

func SampleResponseGetStatusCodesBytes() []byte {
	return []byte(`
{
  "statusDetails": {
    "TESTING_STRICT_RI": "Testing Strict Renegotiation",
    "TESTING_PROTOCOL_INTOLERANCE_304": "Testing Protocol Intolerance (TLS 1.3)",
    "TESTING_HANDSHAKE_SIMULATION": "Simulating handshakes",
    "TESTING_CVE_2014_0224": "Testing CVE-2014-0224",
    "TESTING_ZERO_RTT": "Testing 0-RTT",
    "TESTING_PROTO_3_2_V2H": "Testing TLS 1.1 (v2 handshake)",
    "TESTING_HEARTBLEED": "Testing Heartbleed",
    "TESTING_RENEGOTIATION": "Testing renegotiation",
    "TESTING_PROTOCOL_INTOLERANCE_300": "Testing Protocol Intolerance (SSL 3.0)",
    "TESTING_ECDHE_PARAMETER_REUSE": "Testing ECDHE parameter reuse",
    "TESTING_SUITES_BULK": "Bulk-testing less common cipher suites",
    "TESTING_PROTO_3_1_V2H": "Testing TLS 1.0 (v2 handshake)",
    "TESTING_PROTOCOL_INTOLERANCE_301": "Testing Protocol Intolerance (TLS 1.0)",
    "TESTING_PROTOCOL_INTOLERANCE_302": "Testing Protocol Intolerance (TLS 1.1)",
    "BUILDING_TRUST_PATHS": "Building trust paths",
    "TESTING_PROTOCOL_INTOLERANCE_303": "Testing Protocol Intolerance (TLS 1.2)",
    "TESTING_PROTO_3_0": "Testing SSL 3.0",
    "TESTING_DROWN": "Testing for DROWN",
    "TESTING_PROTO_3_1": "Testing TLS 1.0",
    "TESTING_PROTO_3_3_V2H": "Testing TLS 1.1 (v2 handshake)",
    "TESTING_SUITE_PREFERENCE": "Determining cipher suite preference",
    "TESTING_TLS_VERSION_INTOLERANCE": "Testing TLS version intolerance",
    "VALIDATING_TRUST_PATHS": "Validating trust paths",
    "TESTING_LONG_HANDSHAKE": "Testing Long Handshake (might take a while)",
    "TESTING_SUITES_DEPRECATED": "Testing deprecated cipher suites",
    "TESTING_TICKETBLEED": "Testing Ticketbleed",
    "RETRIEVING_CERT_V3__SNI_APEX": "Retrieving certificate",
    "TESTING_SESSION_TICKETS": "Testing Session Ticket support",
    "TESTING_PROTO_3_4": "Testing TLS 1.3",
    "TESTING_PROTOCOL_INTOLERANCE_499": "Testing Protocol Intolerance (TLS 2.152)",
    "TESTING_PROTO_3_2": "Testing TLS 1.1",
    "TESTING_PROTO_3_3": "Testing TLS 1.2",
    "RETRIEVING_CERT_V3__SNI_WWW": "Retrieving certificate",
    "TESTING_SUITES_NO_SNI": "Observed extra suites during simulation, Testing cipher suites without SNI support",
    "TESTING_CAPABILITIES": "Determining server capabilities",
    "TESTING_EXTENSION_INTOLERANCE": "Testing Extension Intolerance (might take a while)",
    "TESTING_EC_NAMED_CURVES": "Determining supported named groups",
    "TESTING_NPN": "Testing NPN",
    "TESTING_POODLE_TLS": "Testing POODLE against TLS",
    "CHECKING_REVOCATION": "Checking for revoked certificates",
    "TESTING_BEAST": "Testing for BEAST",
    "TESTING_COMPRESSION": "Testing compression",
    "RETRIEVING_CERT_V3__NO_SNI": "Retrieving certificate",
    "TESTING_ZOMBIE_POODLE_AND_GOLDENDOODLE": "Testing Zombie POODLE and GOLDENDOODLE",
    "RETRIEVING_CERT_TLS13": "Retrieving certificate",
    "TESTING_PROTO_2_0": "Testing SSL 2.0",
    "TESTING_ALPN": "Determining supported ALPN protocols",
    "TESTING_OCSP_STAPLING_PRIME": "Trying to prime OCSP stapling",
    "TESTING_SESSION_RESUMPTION": "Testing session resumption",
    "TESTING_OCSP_STAPLING": "Testing OCSP stapling",
    "TESTING_PROTOCOL_INTOLERANCE_399": "Testing Protocol Intolerance (TLS 1.152)",
    "TESTING_STRICT_SNI": "Testing Strict SNI",
    "TESTING_HTTPS": "Sending one complete HTTPS request",
    "PREPARING_REPORT": "Preparing the report",
    "TESTING_SSL2_SUITES": "Checking if SSL 2.0 has any ciphers enabled",
    "TESTING_SUITES": "Determining available cipher suites",
    "TESTING_CVE_2016_2107": "Testing CVE-2016-2107",
    "TESTING_PROTO_3_0_V2H": "Testing SSL 3.0 (v2 handshake)",
    "TESTING_V2H_HANDSHAKE": "Testing v2 handshake",
    "TESTING_BLEICHENBACHER": "Testing Bleichenbacher"
  }
}`)
}
