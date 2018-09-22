// +build wasm

package webworker

import (
	"fmt"
	"syscall/js"
)

func SendLaunchOK() {
	postMessage("launch-ok", "")
}

func SendInitOK() {
	postMessage("init-ok", "")
}

func SendInitFailed(err error) {
	postMessage("init-failed", fmt.Sprintf("Error initialising the emulator: %v", err))
}

func SendGotGameId(gameId string) {
	postMessage("got-game-id", gameId)
}

func SendLoadSaveOK() {
	postMessage("load-save-ok", "")
}

func SendStopOK() {
	postMessage("stop-ok", "")
}

func SendSaveState(gameId, state string) {
	postMessage("save-state", []interface{}{gameId, state})
}

func SendScreenUpdate(screenData []uint8) {
	typedArray := js.TypedArrayOf(screenData)
	defer typedArray.Release()

	clamped := js.Global().Get("Uint8ClampedArray").New(typedArray)
	transferable := js.Global().Get("Array").New(clamped).Get("buffer")
	js.Global().Call("postMessage", []interface{}{"screen-update", clamped}, transferable)
}

func SendFrameRate(rate float32) {
	postMessage("frame-rate-report", rate)
}

func postMessage(action string, body interface{}) {
	js.Global().Call("postMessage", []interface{}{action, body})
}
