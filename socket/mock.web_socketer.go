package socket

import (
	"encoding/json"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
MockWebSocket implements WebSocketer for mocking
*/
type MockWebSocket struct{}

/*
NewMockWebsocket returns a mock websocket connection
*/
func NewMockWebsocket(socketURL *url.URL) (WebSocketer, error) {
	log.Infof("Mock websocket connection to %s established", socketURL.String())
	return &MockWebSocket{}, nil
}

/*
Close implements WebSocketer
*/
func (sock *MockWebSocket) Close() error {
	return nil
}

/*
ReadJSON implements WebSocketer

This uses a stack of responses to attempt to emulate Chromium behavior for
testsing. To use, add a response to the stack with addMockWebsocketResponse().

There is a potential timing issue when emulating command. "Commands" are structs
that implement the Commander interface and are a type of event handler that
makes a request to the socket and only handles responses to that request.

Due to the nature the socket read loop, which is a stream, Commander objects use
`sync.WaitGroup` to wait for the socket response for the request that was
submitted. Because of this, your mock data must be added to the response stack
BEFORE the Commander handler is executed or the test routine will lock forever
preventing you from adding your mock data.

That means that it's possible for the socket read loop to receive your mock
response before Commanderthe handler command is registered with the socket. If
that happens and a handler isn't present to receive it then the response is
discarded, and then when the Commander handler is executed the routine will lock
forever and the test won't finish.

This is only a problem with mocking the socket stream data for unit tests. It
does impact interacting Chrome in any way.
*/
func (sock *MockWebSocket) ReadJSON(v interface{}) error {
	var data interface{}

	if len(_mockWebsocketResponses) > 0 {
		time.Sleep(time.Second * 1)
		data = _mockWebsocketResponses[0]
		_mockWebsocketResponses = _mockWebsocketResponses[1:]

	} else {
		time.Sleep(time.Millisecond * 100)
		data = &Response{
			Error:  &Error{},
			ID:     0,
			Method: "Unknown.event",
		}
	}

	jsonBytes, err := json.Marshal(data)
	log.Debugf("ReadJSON(): returning data %s", jsonBytes)
	err = json.Unmarshal(jsonBytes, &v)
	return err
}

// MockJSONData flag for mocking ReadJSON()
var MockJSONData []byte

// MockJSONRead flag for mocking ReadJSON()
var MockJSONRead = false

// MockJSONType flag for mocking ReadJSON()
var MockJSONType = "command"

// MockJSONError flag for mocking ReadJSON()
var MockJSONError = true

// MockJSONThrowError flag for mocking ReadJSON()
var MockJSONThrowError = false

/*
WriteJSON implements WebSocketer
*/
func (sock *MockWebSocket) WriteJSON(v interface{}) error {
	return nil
}

func addMockWebsocketResponse(id int, error *Error, method string, data ...interface{}) {
	response := &Response{
		Error:  error,
		ID:     id,
		Method: method,
	}
	if len(data) > 0 {
		response.Result, _ = json.Marshal(data[0])
	}
	if len(data) > 1 {
		response.Params, _ = json.Marshal(data[1])
	}

	_mockWebsocketResponses = append(_mockWebsocketResponses, response)
}

var _mockWebsocketResponses = make([]*Response, 0)