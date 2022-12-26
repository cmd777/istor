package istor

type Onionoo struct {
	LastModified string
	Relays       []struct {
		Name        string `json:"n"`
		Fingerprint string `json:"f"`
	} `json:"relays"`
}

const (
	IP_NOT_TOR = 0 + iota
	NEWREQUEST_FAIL
	CLIENT_DO_FAIL
	CONTENT_NOT_MODIFIED
	BAD_REQUEST
	NOT_AVAILABLE
	INTERNAL_SERVER_ERROR
	SERVICE_UNAVAILABLE
	IO_READ_FAIL
	JSON_UNMARSHAL_FAIL
	IP_TOR_RELAY
)
