import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "doc",
      id: "superplatform/superplatform",
    },
    {
      type: "category",
      label: "Chat Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/events",
          label: "Events",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/delete-message",
          label: "Delete a Message",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/add-thread",
          label: "Add Thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/delete-thread",
          label: "Delete a Thread",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/get-thread",
          label: "Get Thread",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/update-thread",
          label: "Update Thread",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/add-message",
          label: "Add Message",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/get-messages",
          label: "List Messages",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/get-threads",
          label: "Get Threads",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Config Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/get-config",
          label: "Get Config",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/save-config",
          label: "Save Config",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Deploy Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/save-deployment",
          label: "Save Deployment",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/list-deployments",
          label: "List Deployments",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Docker Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/launch-container",
          label: "Launch a Container",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/is-running",
          label: "Check If a Container Is Running",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/get-container-summary",
          label: "Get Container Summary",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/get-host",
          label: "Get Docker Host",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/get-info",
          label: "Get Docker Service Information",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Download Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/download",
          label: "Download a File",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/get-download",
          label: "Get a Download",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/pause",
          label: "Pause a Download",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/list-downloads",
          label: "List Downloads",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Dynamic Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/create-object",
          label: "Create a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/upsert-object",
          label: "Upsert a Generic Object",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/query",
          label: "Query Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/delete-objects",
          label: "Delete a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/update-objects",
          label: "Update Objects",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Firehose Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/publish-event",
          label: "Publish an Event",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/subscribe-to-events",
          label: "Subscribe to the Event Stream",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Model Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/start-default-model",
          label: "Start the Default Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/get-default-model-status",
          label: "Get Default Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/get-model",
          label: "Get a Model",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/make-default",
          label: "Make a Model Default",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/start-model",
          label: "Start a Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/get-model-status",
          label: "Get Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/list-models",
          label: "List Models",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Policy Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/check",
          label: "Check",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/upsert-instance",
          label: "Upsert an Instance",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Prompt Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/add-prompt",
          label: "Add Prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/list-prompts",
          label: "List Prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/subscribe-to-prompt-responses",
          label: "Subscribe to Prompt Responses by Thread",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/remove-prompt",
          label: "Remove Prompt",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Registry Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/list-nodes",
          label: "List Nodes",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/save-service-definition",
          label: "Register Service Definition",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/delete-service-definition",
          label: "Delete Service Definition",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/register-service-instance",
          label: "Register Service Instance",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/remove-service-instance",
          label: "Remove Service Instance",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/query-service-instances",
          label: "List Service Instances",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "User Svc",
      items: [
        {
          type: "doc",
          id: "superplatform/change-password",
          label: "Change User Password",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/change-password-admin",
          label: "Change User Password (Admin)",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/login",
          label: "Login",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/create-organization",
          label: "Create an Organization",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/add-user-to-organization",
          label: "Add a User to an Organization",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/remove-user-from-organization",
          label: "Remove a User from an Organization",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/upsert-permission",
          label: "Upsert a Permission",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/is-authorized",
          label: "Is Authorized",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/get-public-key",
          label: "Get Public Key",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/register",
          label: "Register",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/create-role",
          label: "Create a New Role",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/delete-role",
          label: "Delete a Role",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/add-permission-to-role",
          label: "Add Permission to Role",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/get-permissions-by-role",
          label: "Get Permissions by Role",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/set-role-permission",
          label: "Set Role Permissions",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/get-roles",
          label: "Get all Roles",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "superplatform/create-user",
          label: "Create a New User",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/delete-user",
          label: "Delete a User",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "superplatform/save-user-profile",
          label: "Save User Profile",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "superplatform/read-user-by-token",
          label: "Read User by Token",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "superplatform/get-users",
          label: "List Users",
          className: "api-method post",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
