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

type Template struct {
	// Unique identifier for the policy template.
	Id string `json:"id" example:"rate-limit"`

	// Human-readable name of the policy.
	Name string `json:"name" example:"Rate Limit by IP"`

	// Description of the policy.
	Description string `json:"description" example:"Limits the number of requests from a single IP within a certain time window."`

	// Default parameters for the policy.
	Parameters map[string]interface{} `json:"parameters"`
}

func (t *Template) GetId() string {
	return t.Id
}

type Instance struct {
	Id string `json:"id"`

	// The endpoint to which the policy is applied.
	Endpoint string `json:"endpoint" example:"/user-svc/register"`

	// HTTP method (e.g., "GET", "POST").
	Method string `json:"method" example:"POST"`

	// The ID of the policy template.
	TemplateId string `json:"templateId" example:"rate-limit" binding:"required"`

	// Additional parameters or overrides for the policy.
	Parameters map[string]interface{} `json:"parameters"`
}

func (t *Instance) GetId() string {
	return t.Id
}

type UpsertTemplateRequest struct {
	Template
}

type UpsertTemplateResponse struct {
}

type UpsertInstanceRequest struct {
	Instance
}

type CheckRequest struct {
	Endpoint string `json:"endpoint"`
	Method   string `json:"method"`
	Ip       string `json:"ip"`
	UserId   string `json:"userId"`
}

type UpsertPolicyInstanceResponse struct {
	InstanceId string `json:"instanceId"`
}

var (
	RateLimitPolicyTemplate = Template{
		Id:          "rate-limit",
		Name:        "Rate Limit by IP",
		Description: "Limits the number of requests from a single IP within a certain time window.",
		Parameters: map[string]interface{}{
			"maxRequests": 10,
			"timeWindow":  "1m",
		},
	}

	IPWhitelistPolicyTemplate = Template{
		Id:          "ip-whitelist",
		Name:        "IP Whitelist",
		Description: "Allows access only from specific IP addresses.",
		Parameters: map[string]interface{}{
			"allowedIPs": []string{"192.168.1.1", "192.168.1.2"},
		},
	}
)

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
