package main

import (
	"github.com/invertedbit/gms/wasm/events"
)

func main() {

	// body.SetInnerHTML("<h1>Hello from WebAssembly!</h1>")

	events.RegisterEventListeners()

	select {}
}
