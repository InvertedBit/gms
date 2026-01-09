package main

import (
	"github.com/invertedbit/gms/wasm/events"
)

func main() {

	events.RegisterEventListeners()

	select {}
}
