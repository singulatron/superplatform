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
	Port   int    `json:"port"`
	Images Images `json:"images"`
}

type Images struct {
	Default string `json:"default"`
	Cuda    string `json:"cuda,omitempty"`
}

type Assets map[string]string

type Model struct {
	Id             string
	Platform       Platform
	Name           string
	Parameters     string
	Flavour        string
	Version        string
	Quality        string
	Extension      string
	FullName       string
	Tags           []string
	Mirrors        []string
	Size           float64
	Uncensored     bool
	MaxRam         float64
	Description    string
	PromptTemplate string
	QuantComment   string
	MaxBits        int
	Bits           int
	Assets         map[string]string
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

type Status struct {
	SelectedExists bool   `json:"selectedExists"`
	CurrentModelId string `json:"currentModelId"`
	/* Running triggers onModelLaunch on the frontend.
	Running is true when the model is both running and answering
	- fully loaded. */
	Running      bool   `json:"running"`
	ModelAddress string `json:"modelAddress"`
}

type StatusRequest struct {
	Url string `json:"url"`
}

type StatusResponse struct {
	Status *Status `json:"status"`
}

type StartRequest struct {
	Platform Platform `json:"assets"`
	Assets   Assets   `json:"assets"`
}

type StartResponse struct {
}

type MakeDefaultRequest struct {
	Url string `json:"url"`
}

type MakeDefaultResponse struct {
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
