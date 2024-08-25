package dynamic_svc

import (
	"github.com/singulatron/singulatron/sdk/go/datastore"
)

const AnyIdentifier string = "any"

type ErrorResponse struct {
	Error string `json:"error"`
}

type ObjectCreateFields struct {
	Id    string `json:"id"`
	Table string `json:"table" binding:"required"`

	// Authors is a list of user ID and organization ID who created the object.
	Authors []string `json:"authors" binding:"required" example:"[\"usr_12345\", \"org_67890\"]"`

	// Readers is a list of user IDs and role IDs that can read the object.
	Readers []string `json:"readers,omitempty"`

	// Writers is a list of user IDs and role IDs that can write the object.
	Writers []string `json:"writers,omitempty"`

	// Deleters is a list of user IDs and role IDs that can delete the object.
	Deleters []string `json:"deleters,omitempty"`

	Data map[string]interface{} `json:"data,omitempty" binding:"required"`
}

func (g ObjectCreateFields) GetId() string {
	return g.Id
}

// Object holds any kind of data, so
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
type Object struct {
	ObjectCreateFields

	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type QueryRequest struct {
	Table   string           `json:"table"`
	Query   *datastore.Query `json:"query"`
	Readers []string         `json:"readers,omitempty"`
}

type QueryOptions struct {
	Table string
	Query *datastore.Query
}

type QueryResponse struct {
	Objects []*Object `json:"objects,omitempty"`
}

type CreateObjectRequest struct {
	Object *ObjectCreateFields `json:"object,omitempty"`
}

type CreateObjectResponse struct {
	Object *Object `json:"object,omitempty"`
}

type CreateManyRequest struct {
	Objects []*ObjectCreateFields `json:"objects,omitempty"`
}

type UpsertObjectRequest struct {
	Object *ObjectCreateFields `json:"object,omitempty"`
}

type UpsertObjectResponse struct {
	Object *Object `json:"object,omitempty"`
}

type UpsertManyRequest struct {
	Objects []*ObjectCreateFields `json:"objects,omitempty"`
}

type UpsertManyResponse struct {
	Objects []*Object `json:"objects,omitempty"`
}

type DeleteObjectRequest struct {
	Table   string             `json:"table"`
	Filters []datastore.Filter `json:"conditions"`
}

type DeleteObjectResponse struct {
}

type UpdateObjectRequest struct {
	Table   string             `json:"table,omitempty"`
	Filters []datastore.Filter `json:"conditions,omitempty"`
	Object  *Object            `json:"object,omitempty"`
}

type UpdateObjectResponse struct {
}
