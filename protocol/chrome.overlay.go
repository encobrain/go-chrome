package protocol

import (
	"encoding/json"

	overlay "github.com/mkenney/go-chrome/protocol/overlay"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Overlay is a struct that provides a namespace for the Chrome Overlay protocol methods.

The Overlay protocol provides various functionality related to drawing atop the inspected page.
EXPERIMENTAL

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/
*/
type Overlay struct{}

/*
Disable disables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-disable
*/
func (Overlay) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Overlay.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables domain notifications.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-enable
*/
func (Overlay) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Overlay.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetHighlightObjectForTest is for testing.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-getHighlightObjectForTest
*/
func (Overlay) GetHighlightObjectForTest(
	socket sock.Socketer,
	params *overlay.GetHighlightObjectForTestParams,
) (overlay.GetHighlightObjectForTestResult, error) {
	command := sock.NewCommand("Overlay.getHighlightObjectForTest", params)
	result := overlay.GetHighlightObjectForTestResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	if nil != command.Result() {
		resultData, err := json.Marshal(command.Result())
		if nil != err {
			return result, err
		}

		err = json.Unmarshal(resultData, &result)
		if nil != err {
			return result, err
		}
	}

	return result, command.Error()
}

/*
HideHighlight hides any highlight.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-hideHighlight
*/
func (Overlay) HideHighlight(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Overlay.hideHighlight", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
HighlightFrame highlights owner element of the frame with given ID.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightFrame
*/
func (Overlay) HighlightFrame(
	socket sock.Socketer,
	params *overlay.HighlightFrameParams,
) error {
	command := sock.NewCommand("Overlay.highlightFrame", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
HighlightNode highlights DOM node with given ID or with the given JavaScript object wrapper. Either
nodeID or objectID must be specified.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightNode
*/
func (Overlay) HighlightNode(
	socket sock.Socketer,
	params *overlay.HighlightNodeParams,
) error {
	command := sock.NewCommand("Overlay.highlightNode", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
HighlightQuad highlights given quad. Coordinates are absolute with respect to the main frame
viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightQuad
*/
func (Overlay) HighlightQuad(
	socket sock.Socketer,
	params *overlay.HighlightQuadParams,
) error {
	command := sock.NewCommand("Overlay.highlightQuad", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
HighlightRect highlights given rectangle. Coordinates are absolute with respect to the main frame
viewport.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-highlightRect
*/
func (Overlay) HighlightRect(
	socket sock.Socketer,
	params *overlay.HighlightRectParams,
) error {
	command := sock.NewCommand("Overlay.highlightRect", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetInspectMode enters the 'inspect' mode. In this mode, elements that user is hovering over are
highlighted. Backend then generates 'inspectNodeRequested' event upon element selection.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setInspectMode
*/
func (Overlay) SetInspectMode(
	socket sock.Socketer,
	params *overlay.SetInspectModeParams,
) error {
	command := sock.NewCommand("Overlay.setInspectMode", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetPausedInDebuggerMessage sets the paused message

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setPausedInDebuggerMessage
*/
func (Overlay) SetPausedInDebuggerMessage(
	socket sock.Socketer,
	params *overlay.SetPausedInDebuggerMessageParams,
) error {
	command := sock.NewCommand("Overlay.setPausedInDebuggerMessage", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetShowDebugBorders requests that backend shows debug borders on layers.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowDebugBorders
*/
func (Overlay) SetShowDebugBorders(
	socket sock.Socketer,
	params *overlay.SetShowDebugBordersParams,
) error {
	command := sock.NewCommand("Overlay.setShowDebugBorders", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetShowFPSCounter requests that backend shows the FPS counter.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowFPSCounter
*/
func (Overlay) SetShowFPSCounter(
	socket sock.Socketer,
	params *overlay.SetShowFPSCounterParams,
) error {
	command := sock.NewCommand("Overlay.setShowFPSCounter", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetShowPaintRects that backend shows paint rectangles.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowPaintRects
*/
func (Overlay) SetShowPaintRects(
	socket sock.Socketer,
	params *overlay.SetShowPaintRectsParams,
) error {
	command := sock.NewCommand("Overlay.setShowPaintRects", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetShowScrollBottleneckRects requests that backend shows scroll bottleneck rects.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowScrollBottleneckRects
*/
func (Overlay) SetShowScrollBottleneckRects(
	socket sock.Socketer,
	params *overlay.SetShowScrollBottleneckRectsParams,
) error {
	command := sock.NewCommand("Overlay.setShowScrollBottleneckRects", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetShowViewportSizeOnResize paints viewport size upon main frame resize.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setShowViewportSizeOnResize
*/
func (Overlay) SetShowViewportSizeOnResize(
	socket sock.Socketer,
	params *overlay.SetShowViewportSizeOnResizeParams,
) error {
	command := sock.NewCommand("Overlay.setShowViewportSizeOnResize", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
SetSuspended sets the suspended state

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#method-setSuspended
*/
func (Overlay) SetSuspended(
	socket sock.Socketer,
	params *overlay.SetSuspendedParams,
) error {
	command := sock.NewCommand("Overlay.setSuspended", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnInspectNodeRequested adds a handler to the Overlay.inspectNodeRequested event.
Overlay.inspectNodeRequested fires when the node should be inspected. This happens after call to
`setInspectMode` or when user manually inspects an element.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-inspectNodeRequested
*/
func (Overlay) OnInspectNodeRequested(
	socket sock.Socketer,
	callback func(event *overlay.InspectNodeRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.inspectNodeRequested",
		func(response *sock.Response) {
			event := &overlay.InspectNodeRequestedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnNodeHighlightRequested adds a handler to the Overlay.nodeHighlightRequested event.
Overlay.nodeHighlightRequested fires when the node should be highlighted. This happens after call to
`setInspectMode`.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-nodeHighlightRequested
*/
func (Overlay) OnNodeHighlightRequested(
	socket sock.Socketer,
	callback func(event *overlay.NodeHighlightRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.nodeHighlightRequested",
		func(response *sock.Response) {
			event := &overlay.NodeHighlightRequestedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnScreenshotRequested adds a handler to the Overlay.screenshotRequested event.
Overlay.screenshotRequested fires when user asks to capture screenshot of some area on the page.

https://chromedevtools.github.io/devtools-protocol/tot/Overlay/#event-screenshotRequested
*/
func (Overlay) OnScreenshotRequested(
	socket sock.Socketer,
	callback func(event *overlay.ScreenshotRequestedEvent),
) {
	handler := sock.NewEventHandler(
		"Overlay.screenshotRequested",
		func(response *sock.Response) {
			event := &overlay.ScreenshotRequestedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
