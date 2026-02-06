package htmx

import (
	"encoding/json"

	"github.com/gofiber/fiber/v3"
)

type HXResponseHeader string

const (
	HXLocation           HXResponseHeader = "HX-Location"
	HXPushUrl            HXResponseHeader = "HX-Push-Url"
	HXRedirect           HXResponseHeader = "HX-Redirect"
	HXRefresh            HXResponseHeader = "HX-Refresh"
	HXReplaceUrl         HXResponseHeader = "HX-Replace-Url"
	HXReswap             HXResponseHeader = "HX-Reswap"
	HXRetarget           HXResponseHeader = "HX-Retarget"
	HXReselect           HXResponseHeader = "HX-Reselect"
	HXTrigger            HXResponseHeader = "HX-Trigger"
	HXTriggerAfterSettle HXResponseHeader = "HX-Trigger-After-Settle"
	HXTriggerAfterSwap   HXResponseHeader = "HX-Trigger-After-Swap"
)

func (h HXResponseHeader) String() string {
	return string(h)
}

func (h HXResponseHeader) Set(c fiber.Ctx, value string) {
	c.Set(h.String(), value)
}

func (h HXResponseHeader) SetJson(c fiber.Ctx, value interface{}) string {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	c.Set(h.String(), string(jsonValue))
	return string(jsonValue)
}
