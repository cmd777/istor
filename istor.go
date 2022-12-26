package istor

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

/*
Onionoo provides the Last-Modified http date, Name and Fingerprint of a relay.

int is a response code, 0-10 (0 being Not a TOR Relay.)

error is a generic error interface. (nil being Not a TOR Relay)

# Use IfModifiedSince to cache

# Caching is not required, but it is useful for multiple requests to the same ip

Basic Example (No Caching):

	_, ResponseCode, err := istor.IsRelay("1.2.3.4", "")
	switch ResponseCode {
	case istor.IP_NOT_TOR: // 0
		fmt.Println("Not a TOR Relay")
	case istor.IP_TOR_RELAY: // 10
		fmt.Println("TOR Relay")
	default: // 1 - 9 are error codes
		fmt.Println("An error occoured...", err.Error())
	}
*/
func IsRelay(ip, IfModifiedSince string) (Onionoo, int, error) {
	var O Onionoo

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://onionoo.torproject.org/summary?limit=1&search=%v", ip), nil)
	if err != nil {
		return O, NEWREQUEST_FAIL, err
	}

	req.Header.Set("Accept-Encoding", "gzip")

	if len(IfModifiedSince) != 0 {
		req.Header.Set("If-Modified-Since", IfModifiedSince)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return O, CLIENT_DO_FAIL, err
	}

	O.LastModified = resp.Header.Get("Last-Modified")
	switch resp.StatusCode {
	case 304:
		return O, CONTENT_NOT_MODIFIED, fmt.Errorf(fmt.Sprintf("content not modified since %v", IfModifiedSince))
	case 400:
		return O, BAD_REQUEST, errors.New("the request for a known resource could not be processed because of bad syntax")
	case 404:
		return O, NOT_AVAILABLE, errors.New("the request could not be processed because the requested resource could not be found")
	case 500:
		return O, INTERNAL_SERVER_ERROR, errors.New("there is an unspecific problem with the server which the service operator may not yet be aware of")
	case 503:
		return O, SERVICE_UNAVAILABLE, errors.New("the server is temporarily down for maintenance, or there is a temporary problem with the server that the service operator is already aware of")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return O, IO_READ_FAIL, err
	}

	uerr := json.Unmarshal(body, &O)
	if uerr != nil {
		return O, JSON_UNMARSHAL_FAIL, uerr
	}

	if len(O.Relays) != 0 {
		return O, IP_TOR_RELAY, errors.New("ip is a tor relay")
	}

	return O, IP_NOT_TOR, nil
}
