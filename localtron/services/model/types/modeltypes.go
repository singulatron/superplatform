package modeltypes

import "sync"

/*
Platform roughly represents an AI container + its settings.
*/
type Platform struct {
	ID        string    `json:"id"`
	Name      *string   `json:"name,omitempty"`
	Version   *int      `json:"version,omitempty"`
	Container Container `json:"container"`
}

type Container struct {
	/* Port is the internal port of the Container */
	Port   int    `json:"port"`
	Images Images `json:"images"`
}

type Images struct {
	Default string `json:"default"`
	Cuda    string `json:"cuda,omitempty"`
}

type Assets map[string]string

type Model struct {
	Id             string            `json:"id"`
	Platform       Platform          `json:"platform"`
	Name           string            `json:"name"`
	Parameters     string            `json:"parameters"`
	Flavour        string            `json:"flavour"`
	Version        string            `json:"version"`
	Quality        string            `json:"quality"`
	Extension      string            `json:"extension"`
	FullName       string            `json:"full_name"`
	Tags           []string          `json:"tags"`
	Mirrors        []string          `json:"mirrors"`
	Size           float64           `json:"size"`
	Uncensored     bool              `json:"uncensored"`
	MaxRam         float64           `json:"max_ram"`
	Description    string            `json:"description"`
	PromptTemplate string            `json:"prompt_template"`
	QuantComment   string            `json:"quant_comment"`
	MaxBits        int               `json:"max_bits"`
	Bits           int               `json:"bits"`
	Assets         map[string]string `json:"assets"`
}

/* Internal type for ModelService */
type ModelState struct {
	sync.Mutex
	Answering         bool
	HasCheckerRunning bool
}

// Setter methods for each field
func (m *ModelState) SetAnswering(v bool) {
	m.Lock()
	defer m.Unlock()
	m.Answering = v
}

func (m *ModelState) SetHasCheckerRunning(v bool) {
	m.Lock()
	defer m.Unlock()
	m.HasCheckerRunning = v
}

type ModelStatus struct {
	AssetsReady bool `json:"assetsReady"`
	/* Running triggers onModelLaunch on the frontend.
	Running is true when the model is both running and answering
	- fully loaded. */
	Running bool   `json:"running"`
	Address string `json:"address"`
}

type StatusRequest struct {
	Url string `json:"url"`
}

type StatusResponse struct {
	Status *ModelStatus `json:"status"`
}

type StartRequest struct {
	ModelId string `json:"status,omitempty"`
}

type StartResponse struct {
}

type MakeDefaultRequest struct {
	Url string `json:"url"`
}

type MakeDefaultResponse struct {
}

type GetModelsResponse struct {
	Models []*Model `json:"models,omitempty"`
}

//
// Events
//

const EventModelReadyName = "modelReady"

type EventModelReady struct {
	ThreadId string `json:"threadId"`
}

func (e EventModelReady) Name() string {
	return EventModelReadyName
}
