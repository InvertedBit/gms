package htmx

type HXHeader struct {
	Boosted               string `header:"HX-Boosted"`
	CurrentURL            string `header:"HX-Current-URL"`
	HistoryRestoreRequest string `header:"HX-History-Restore-Request"`
	Prompt                string `header:"HX-Prompt"`
	Request               string `header:"HX-Request"`
	Target                string `header:"HX-Target"`
	Trigger               string `header:"HX-Trigger"`
	TriggerName           string `header:"HX-Trigger-Name"`
}

func (h *HXHeader) IsHTMXRequest() bool {
	return h.Request != "" && h.Request == "true"
}

func (h *HXHeader) IsBoosted() bool {
	return h.Boosted == "true"
}

func (h *HXHeader) IsHistoryRestoreRequest() bool {
	return h.HistoryRestoreRequest == "true"
}

func (h *HXHeader) IsPrompt() bool {
	return h.Prompt == "true"
}

func (h *HXHeader) GetCurrentURL() string {
	return h.CurrentURL
}

func (h *HXHeader) GetTarget() string {
	return h.Target
}

func (h *HXHeader) GetTrigger() string {
	return h.Trigger
}

func (h *HXHeader) GetTriggerName() string {
	return h.TriggerName
}
