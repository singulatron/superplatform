package generictypes

import (
	"github.com/singulatron/singulatron/localtron/datastore"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type GenericObjectCreateFields struct {
	Id    string `json:"id"`
	Table string `json:"table" binding:"required"`

	// Public is true when the object is visible to all users.
	// When it's false the entry is only visible to the user who created it.
	Public bool `json:"public,omitempty"`

	Data   map[string]any `json:"data,omitempty" binding:"required"`
	UserId string         `json:"userId,omitempty"`
}

func (g GenericObjectCreateFields) GetId() string {
	return g.Id
}

// GenericObject holds any kind of data, so
// we don't have to implement simple CRUD for
// any new simple entity.
//
// When JSON marshalled, might look like this:
//
//	{
//		"id": "id1",
//		"createdAt": "2024-05-30 13:53:22",
//		"createdAt": "2024-05-30 13:53:22",
//		"userId": "userid1",
//		"data": {
//			"anyfield1": "anyvalue",
//			"anyfield2": 42
//		}
//	}
type GenericObject struct {
	GenericObjectCreateFields

	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type FindRequest struct {
	Table  string           `json:"table"`
	Query  *datastore.Query `json:"query"`
	Public bool             `json:"public"`
}

type FindOptions struct {
	Table  string
	Query  *datastore.Query
	UserId string
	Public bool
}

type FindResponse struct {
	Objects []*GenericObject `json:"objects,omitempty"`
}

type CreateRequest struct {
	GenericObjectCreateFields `json:"object,omitempty"`
}

type CreateManyRequest struct {
	Objects []*GenericObject `json:"objects,omitempty"`
}

type UpsertRequest struct {
	Object *GenericObject `json:"object,omitempty"`
}

type UpsertManyRequest struct {
	Objects []*GenericObject `json:"objects,omitempty"`
}

type DeleteRequest struct {
	Table      string                `json:"table"`
	Conditions []datastore.Condition `json:"conditions"`
}

type DeleteResponse struct {
}

type UpdateRequest struct {
	Table      string                `json:"table,omitempty"`
	Conditions []datastore.Condition `json:"conditions,omitempty"`
	Object     *GenericObject        `json:"object,omitempty"`
}

type UpdateResponse struct {
}
