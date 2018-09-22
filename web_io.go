// +build wasm

package main

import (
	"log"

	"syscall/js"

	"github.com/djhworld/gomeboycolor-wasm/webworker"
	"github.com/djhworld/gomeboycolor/inputoutput"
	"github.com/djhworld/gomeboycolor/types"
)

const (
	JsKeyUp    = 38
	JsKeyDown  = 40
	JsKeyLeft  = 37
	JsKeyRight = 39
	JsKeyZ     = 90
	JsKeyX     = 88
	JsKeyA     = 65
	JsKeyS     = 83
	ALPHA_FULL = 255
	FOUR       = 4
)

var DefaultControlScheme inputoutput.ControlScheme = inputoutput.ControlScheme{
	UP:     JsKeyUp,
	DOWN:   JsKeyDown,
	LEFT:   JsKeyLeft,
	RIGHT:  JsKeyRight,
	A:      JsKeyZ,
	B:      JsKeyX,
	START:  JsKeyA,
	SELECT: JsKeyS,
}

// WebIO is for running the emulator in a web environment
type WebIO struct {
	*inputoutput.CoreIO
	html5Display *html5CanvasDisplay
}

func NewWebIO(headless bool, displayFps bool) *WebIO {
	log.Println("Creating Web based IO Handler")
	html5Display := new(html5CanvasDisplay)

	frameRateReporter := func(v float32) {
		if displayFps {
			log.Printf("Average frame rate\t%.2f\tfps", v)
			webworker.SendFrameRate(v)
		}
	}

	return &WebIO{
		inputoutput.NewCoreIO(headless, frameRateReporter, html5Display),
		html5Display,
	}
}

func (i *WebIO) Init(title string, screenSize int, onCloseHandler func()) error {
	var err error = nil
	i.OnCloseHandler = onCloseHandler

	var messageCB js.Callback
	messageCB = js.NewCallback(func(args []js.Value) {
		input := args[0].Get("data")
		switch input.Index(0).String() {
		case "ku":
			i.KeyHandler.KeyUp(input.Index(1).Int())
		case "kd":
			i.KeyHandler.KeyDown(input.Index(1).Int())
		case "stop":
			messageCB.Release()
			i.StopChannel <- 1
		}
	})

	err = i.html5Display.init(title, i.Headless)
	if err != nil {
		return err
	}

	i.KeyHandler.Init(DefaultControlScheme) //TODO: allow user to define controlscheme

	self := js.Global().Get("self")
	self.Call("addEventListener", "message", messageCB, false)

	return err
}

type html5CanvasDisplay struct {
	Name      string
	imageData []uint8
	headless  bool
}

func (s *html5CanvasDisplay) init(title string, headless bool) error {
	s.Name = inputoutput.PREFIX + "-SCREEN"
	log.Printf("%s: Initialising display", s.Name)

	imageDataLen := inputoutput.SCREEN_WIDTH * inputoutput.SCREEN_HEIGHT * 4
	s.imageData = make([]uint8, imageDataLen, imageDataLen)
	s.headless = headless

	return nil
}

func (s *html5CanvasDisplay) Stop() {
	// noop
}

func (s *html5CanvasDisplay) DrawFrame(screenData *types.Screen) {
	i := 0

	for y := 0; y < inputoutput.SCREEN_HEIGHT; y++ {
		for x := 0; x < inputoutput.SCREEN_WIDTH; x++ {
			var pixel types.RGB = screenData[y][x]
			s.imageData[i] = pixel.Red
			s.imageData[i+1] = pixel.Green
			s.imageData[i+2] = pixel.Blue
			s.imageData[i+3] = ALPHA_FULL
			i += FOUR
		}
	}

	// TODO this is probably a performance bottleneck
	if !s.headless {
		webworker.SendScreenUpdate(s.imageData)
	}
}
