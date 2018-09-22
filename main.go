// +build wasm
package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"strings"
	"syscall/js"
	"time"

	"github.com/djhworld/gomeboycolor-wasm/webworker"
	"github.com/djhworld/gomeboycolor/cartridge"
	"github.com/djhworld/gomeboycolor/config"
	"github.com/djhworld/gomeboycolor/gbc"
	"github.com/djhworld/gomeboycolor/saves"
)

const TITLE string = "gomeboycolor"

func main() {
	fmt.Printf("%s.\n", TITLE)
	fmt.Println("Copyright (c) 2013-2018. Daniel James Harper.")
	fmt.Println("http://djhworld.github.io/gomeboycolor")
	fmt.Println(strings.Repeat("*", 120))

	var emulatorSetup *EmulatorSetup = new(EmulatorSetup)
	webworker.SendLaunchOK()

	awaitSetup(emulatorSetup)

	emulator, err := gbc.Init(
		emulatorSetup.cart,
		emulatorSetup.saveStore,
		emulatorSetup.config,
		NewWebIO(
			emulatorSetup.config.Headless,
			emulatorSetup.config.DisplayFPS,
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	//Starts emulator code
	go emulator.Run(emulatorSetup.config.FrameRateLock)

	//set the IO controller to run indefinitely (it waits for screen updates)
	emulator.RunIO()

	webworker.SendStopOK()
	log.Println("Goodbye!")
}

func awaitSetup(emulatorSetup *EmulatorSetup) {
	var done bool = false

	var messageCB js.Callback
	messageCB = js.NewCallback(func(args []js.Value) {
		input := args[0].Get("data")
		switch input.Index(0).String() {
		case "init":
			if err := emulatorSetup.handleInit(input.Index(1)); err != nil {
				webworker.SendInitFailed(err)
			} else {
				webworker.SendInitOK()
			}
		case "get-game-id":
			gameId := emulatorSetup.handleGetGameId()
			webworker.SendGotGameId(gameId)
		case "load-save":
			emulatorSetup.handleLoadGame(input.Index(1))
			webworker.SendLoadSaveOK()
		case "start":
			done = true
		}
	})

	js.Global().Get("self").Call("addEventListener", "message", messageCB, false)

	defer func() {
		log.Println("Removing temporary message event handler")
		js.Global().Get("self").Call("removeEventListener", "message", messageCB, false)
		messageCB.Release()
	}()

	for !done {
		time.Sleep(50 * time.Millisecond)
	}
}

type EmulatorSetup struct {
	saveStore saves.Store
	cart      *cartridge.Cartridge
	config    *config.Config
}

func (e *EmulatorSetup) handleInit(initData js.Value) error {
	log.Println("Initialising")

	initConfig := initData.Get("config")
	e.config = &config.Config{
		Title:         TITLE,
		ScreenSize:    1,
		SkipBoot:      initConfig.Get("skipBoot").Bool(),
		DisplayFPS:    initConfig.Get("showfps").Bool(),
		ColorMode:     initConfig.Get("colorMode").Bool(),
		Debug:         false,
		BreakOn:       "0x0000",
		DumpState:     false,
		Headless:      initConfig.Get("headless").Bool(),
		FrameRateLock: int64(initConfig.Get("fpsLock").Int()),
	}

	romData := initData.Get("romData")
	romName := romData.Get("name").String()
	romBase64 := romData.Get("base64Bytes").String()
	if romBytes, err := base64.StdEncoding.DecodeString(romBase64); err != nil {
		return err
	} else {
		cart, err := cartridge.NewCartridge(romName, romBytes)
		if err != nil {
			return err
		}

		e.cart = cart
	}

	return nil
}

func (e *EmulatorSetup) handleGetGameId() string {
	return e.cart.ID
}

func (e *EmulatorSetup) handleLoadGame(saveData js.Value) error {
	webStore := NewWebStore()
	e.saveStore = webStore

	if saveDataStr := saveData.String(); saveDataStr != "" {
		return webStore.PutSave(e.cart.ID, saveDataStr)
	}

	return nil
}
