/*
Package tarmac is a client package for WASM functions running within a Tarmac server.

This package provides user-friendly functions that wrap the Web Assembly Procedure Call (waPC) based functions of
Tarmac. Guest WASM functions running inside Tarmac can use this library to call back the Tarmac host and perform
host-level actions such as storing data within the database, logging specific data, or looking up configurations.

*/
package tarmac

// ServerRequest is used to create the Payload JSON sent to Tarmac WASM functions for incoming requests. A
// "request:handler" function will receive this Request type in JSON format and will need to parse it accordingly.
type ServerRequest struct {

	// Headers are request headers such as HTTP headers, or other metadata depending on the protocol the request as
	// received with.
	Headers map[string]string `json:"headers"`

	// Payload is a []byte that has been base64 encoded.
	Payload string `json:"payload"`
}

// ServerResponse is used to create a Response Payload sent from Tarmac WASM functions to the Tarmac host. A
// "request:handler" function will return this Response type in JSON format. The host will parse it accordingly and use
// the response details to reply to clients.
type ServerResponse struct {

	// Headers are response headers such as HTTP headers, or other metadata depending on the protocol. Any value in
	// this map will be appended to existing response headers provided by the Tarmac server.
	Headers map[string]string `json:"headers"`

	// Status is used within Response messages from Tarmac, it indicates either failure or success for both Host
	// Callbacks, and request handler calls. Status codes used should follow HTTP status code conventions, even if the call is
	// non-HTTP, following a common return code will enable cross platform execution.
	Status Status `json:"status"`

	// Payload is a []byte that should be base64 encoded. This payload is the response payload from the WASM Module
	// itself. The contents of this field will be decoded and returned to the client.
	Payload string `json:"payload"`
}

// Status is used within Response messages from Tarmac, it indicates either failure or success for both Host Callbacks,
// and request handler calls. Status codes used should follow HTTP status code conventions, even if the call is non-HTTP,
// following a common return code will enable cross platform execution.
type Status struct {
	// Code is the HTTP status code based return code for function execution.
	Code int `json:"code"`

	// Status is the human readible error message or success message for function execution.
	Status string `json:"status"`
}

// KVStoreGet is a structure used to create Get request callbacks to the Tarmac KVStore interface. This structure is a
// general request type used for all KVStore types provided by Tarmac.
type KVStoreGet struct {
	// Key is the index key to use when accessing the key:value store.
	Key string `json:"key"`
}

// KVStoreGetResponse is a structure supplied as response messages to KVStore Get requests. This response is a general
// response type used for all KVStore types provided by Tarmac.
type KVStoreGetResponse struct {
	// Status is the human readible error message or success message for function execution.
	Status Status `json:"status"`

	// Data is the response data provided by the key:value store. This data is base64 encoded to provide a simple
	// JSON-friendly field for arbitrary data.
	Data string `json:"data"`
}

// KVStoreSet is a structure used to create a Set request callback to the Tarmac KVStore interface. This structure is a
// general request type used for all KVStore types provided by Tarmac.
type KVStoreSet struct {
	// Key is the index key used to store the data.
	Key string `json:"key"`

	// Data is the user-supplied key:value data. This field should contain a base64 encoded byte slice. Tarmac
	// expects this field to base base64 encoded, and neglecting to do so will result in an error returned from the callback
	// function.
	Data string `json:"data"`
}

// KVStoreSetResponse is a structure supplied as a response message to the KVStore Set callback function. This response
// is a general response type used for all KVStore types provided by Tarmac.
type KVStoreSetResponse struct {
	// Status is the human readible error message or success message for function execution.
	Status Status `json:"status"`
}

// KVStoreDelete is a structure used to create Delete callback requests to the Tarmac KVStore interface. This structure
// is a general request type used for all KVStore types provided by Tarmac.
type KVStoreDelete struct {
	// Key is the index key used to store the data.
	Key string `json:"key"`
}

// KVStoreDeleteResponse is a structure supplied as a response message to the KVStore Delete callback function. This
// response is a general response type used for all KVStore types provided by Tarmac.
type KVStoreDeleteResponse struct {
	// Status is the human readible error message or success message for function execution.
	Status Status `json:"status"`
}

// KVStoreKeysResponse is a structure supplied as a response message to the KVStore Keys callback function. This
// response is a general response type used for all KVStore types provided by Tarmac.
type KVStoreKeysResponse struct {
	// Keys is a list of keys available within the KV Store.
	Keys []string `json:"keys"`

	// Status is the human readible error message or success message for function execution.
	Status Status `json:"status"`
}

// MetricsCounter is a structure used to create Counter metrics callback requests to the Tarmac Metrics interface.
type MetricsCounter struct {
	// Name is the name of the metric as exposed via the metrics HTTP end-point. Name must be unique across
	// all metrics; duplicate names will create a panic.
	Name string `json:"name"`
}

// MetricsGauge is a structure used to create Gauge metrics callback requests to the Tarmacs Metrics interface.
type MetricsGauge struct {
	// Name is the name of the metric as exposed via the metrics HTTP end-point. Name must be unique across all
	// metrics; duplicate names will create a panic.
	Name string `json:"name"`

	// Action is the action to be performed for the Gauge metric. Valid options are inc (Increment) and dec (Decrement).
	Action string `json:"action"`
}

// MetricsHistogram is a structure used to create Histogram metrics callback requests to the Tarmacs Metrics interface.
type MetricsHistogram struct {
	// Name is the name of the metric as exposed via the metrics HTTP end-point. Name must be unique across all
	// metrics; duplicate names will create a panic.
	Name string `json:"name"`

	// Value is the value to Observe for the Histogram metric.
	Value float64 `json:"value"`
}

// HTTPCall is a structure used to create HTTP calls to remote systems.
type HTTPCall struct {
	// Method is the HTTP method type for the HTTP request; valid options are GET, POST, PUT, PATCH, HEAD, & DELETE.
	Method string `json:"method"`

	// Headers are the HTTP headers to include in the HTTP request.
	Headers map[string]string `json:"headers"`

	// URL is the HTTP URL to call.
	URL string `json:"url"`

	// Body is the user-supplied HTTP body data. This field should contain a base64 encoded string. Tarmac expects this
	// field to be based64 encoded, and neglecting to do so will result in an error from the callback function.
	Body string `json:"body"`

	// Insecure will disable TLS host verification; this is common with self-signed certificates; however, use caution.
	Insecure bool `json:"insecure"`
}

// HTTPCallResponse is a structure supplied as a response message to a remote HTTP call callback function.
type HTTPCallResponse struct {
	// Status is the human readible error message or success message for function execution.
	Status Status

	// Code is the HTTP Status Code returned from the target server.
	Code int `json:"status"`

	// Headers are the HTTP headers returned from the HTTP request.
	Headers map[string]string `json:"headers"`

	// Body is the server-supplied HTTP payload data. The server-supplied payload will be base64 encoded to provide a
	// simple JSON-friendly field for arbitrary data.
	Body string `json:"body"`
}
