package events

import (
	"errors"

	"honnef.co/go/js/dom/v2"
)

func TryGetEventDetail(event dom.Event, key string) (string, error) {
	if customEvent, ok := event.(*dom.CustomEvent); ok {
		if value := customEvent.BasicEvent.Value.Get("detail").Get(key); !value.IsUndefined() {
			return value.String(), nil
		}
	}
	return "", errors.New("key not found in event detail")
}
