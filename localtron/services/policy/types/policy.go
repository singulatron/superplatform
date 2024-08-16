/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package policy_svc

import (
	user "github.com/singulatron/singulatron/localtron/services/user/types"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

// Entity that is being rate limited/policed
type Entity string

const (
	EntityUserID Entity = "userId"
	EntityIP     Entity = "ip"
)

// Scope is the bucket by which we aggregate
type Scope string

const (
	ScopeEndpoint Scope = "endpoint"
	ScopeGlobal   Scope = "global"
)

type TemplateId string

const (
	TemplateIdRateLimit TemplateId = "rate-limit"
	TemplateIdBlocklist TemplateId = "rate-limit"
)

// Parameters for Rate Limit policy instance
type RateLimitParameters struct {
	MaxRequests int    `json:"maxRequests" example:"10"`
	TimeWindow  string `json:"timeWindow" example:"1m"`
	Entity      Entity `json:"entity" example:"userId"`
	Scope       Scope  `json:"scope" example:"endpoint"`
}

type BlocklistParameters struct {
	BlockedIPs []string `json:"blockedIPs"`
}

type Template struct {
	Id          TemplateId `json:"id" example:"rate-limit"`
	Name        string     `json:"name" example:"Rate Limit"`
	Description string     `json:"description" example:"Limits the number of requests based on user ID or IP address."`
}

func (t *Template) GetId() string {
	return string(t.Id)
}

var RateLimitPolicyTemplate = Template{
	Id:          TemplateIdRateLimit,
	Name:        "Rate Limit",
	Description: "Limits the number of requests. Aggregate by UserID or IP address.",
}

var BlocklistTemplate = Template{
	Id:          TemplateIdBlocklist,
	Name:        "Blocklist",
	Description: "Block access by IP, UserID and other parameters.",
}

type Instance struct {
	Id         string `json:"id"`
	Endpoint   string `json:"endpoint" example:"/user-svc/register"`
	TemplateId string `json:"templateId" example:"rate-limit" binding:"required"`

	RateLimitParameters *RateLimitParameters `json:"rateLimitParameters,omitempty"`
	BlocklistParameters *BlocklistParameters `json:"ipWhitelistParameters,omitempty"`
}

func (t *Instance) GetId() string {
	return t.Id
}

type UpsertInstanceRequest struct {
	Instance *Instance
}

type UpsertInstanceResponse struct {
}

type CheckRequest struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Ip       string `json:"ip"`
	UserId   string `json:"userId"`
}

type CheckResponse struct {
	Allowed bool `json:"allowed" binding:"required"`
}

var (
	PermissionTemplateView = user.Permission{
		Id:   "policy-svc:template:view",
		Name: "Policy Svc - Template View",
	}

	PermissionTemplateCreate = user.Permission{
		Id:   "policy-svc:template:create",
		Name: "Policy Svc - Template Create",
	}

	PermissionTemplateEdit = user.Permission{
		Id:   "policy-svc:template:edit",
		Name: "Policy Svc - Template Edit",
	}

	PermissionTemplateDelete = user.Permission{
		Id:   "policy-svc:template:delete",
		Name: "Policy Svc - Template Delete",
	}

	PermissionInstanceView = user.Permission{
		Id:   "policy-svc:instance:view",
		Name: "Policy Svc - Instance View",
	}

	PermissionInstanceCreate = user.Permission{
		Id:   "policy-svc:instance:create",
		Name: "Policy Svc - Instance Create",
	}

	PermissionInstanceEdit = user.Permission{
		Id:   "policy-svc:instance:edit",
		Name: "Policy Svc - Instance Edit",
	}

	PermissionInstanceDelete = user.Permission{
		Id:   "policy-svc:instance:delete",
		Name: "Policy Svc - Instance Delete",
	}

	PermissionCheckView = user.Permission{
		Id:   "policy-svc:check:view",
		Name: "Policy Svc - Check View",
	}

	AdminPermissions = []user.Permission{
		PermissionTemplateView,
		PermissionTemplateCreate,
		PermissionTemplateEdit,
		PermissionTemplateDelete,
		PermissionInstanceView,
		PermissionInstanceCreate,
		PermissionInstanceEdit,
		PermissionInstanceDelete,
	}

	UserPermissions = []user.Permission{
		PermissionCheckView,
	}
)
