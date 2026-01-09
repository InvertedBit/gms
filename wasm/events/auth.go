package events

import (
	"log"

	"honnef.co/go/js/dom/v2"
)

func RegisterEventListeners() {
	body := dom.GetWindow().Document().QuerySelector("body")

	body.AddEventListener("gms:login-failed", false, HandleLoginFailed)
}

func HandleLoginFailed(event dom.Event) {
	log.Println("Login failed event received")
	message, err := TryGetEventDetail(event, "message")
	if err != nil {
		log.Println("Error getting event detail:", err)
	} else {
		log.Println("Event detail:", message)
	}
	log.Println(event.Target().OuterHTML())
	inputs := event.Target().QuerySelectorAll(".input-accent")
	for _, input := range inputs {
		input.Class().Remove("input-accent")
		input.Class().Add("input-error")
		// log.Println("Input field:", input.OuterHTML())
	}
	errorMessageContainer := event.Target().QuerySelector(".login-error-message")
	errorMessageContainer.SetInnerHTML(message)
	errorMessageContainer.Class().Remove("hidden")

}
