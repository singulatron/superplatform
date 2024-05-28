package modeltypes

import "sync"

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

type StatusRequest struct{}

type StatusResponse struct {
	Status *Status `json:"status"`
}
