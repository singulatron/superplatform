import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "doc",
      id: "singulatron/singulatron",
    },
    {
      type: "category",
      label: "Chat Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/delete-message",
          label: "Delete a Message",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/add-thread",
          label: "Add Thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-thread",
          label: "Delete a Thread",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/get-thread",
          label: "Get Thread",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/update-thread",
          label: "Update Thread",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/add-message",
          label: "Add Message",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-messages",
          label: "List Messages",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-threads",
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
          id: "singulatron/get-config",
          label: "Get Config",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/save-config",
          label: "Save Config",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "Docker Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/launch-container",
          label: "Launch a Docker Container",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/is-running",
          label: "Check If a Container Is Running",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-container-summary",
          label: "Get Container Summary",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-host",
          label: "Get Docker Host",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-info",
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
          id: "singulatron/download",
          label: "Download a File",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-download",
          label: "Get a Download",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/pause",
          label: "Pause a Download",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/list-downloads",
          label: "List Downloads",
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
          id: "singulatron/publish-an-event",
          label: "Publish an Event",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/subscribe-to-the-event-stream",
          label: "Subscribe to the Event Stream",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Generic Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/create-object",
          label: "Create a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/upsert-object",
          label: "Upsert a Generic Object",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/query",
          label: "Query Generic Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-objects",
          label: "Delete a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/update-objects",
          label: "Update Generic Objects",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Model Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/start-default-model",
          label: "Start the Default Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-default-model-status",
          label: "Get Default Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-model",
          label: "Get a Model",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/make-default",
          label: "Make a Model Default",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/start-model",
          label: "Start a Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-model-status",
          label: "Get Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/list-models",
          label: "List Models",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Node Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/list-nodes",
          label: "List Nodes",
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
          id: "singulatron/check",
          label: "Check",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/upsert-instance",
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
          id: "singulatron/subscribe",
          label: "Subscribe to Prompt",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/add-prompt",
          label: "Add Prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-prompts",
          label: "List Prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-prompt",
          label: "Remove Prompt",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "User Svc",
      items: [
        {
          type: "doc",
          id: "singulatron/change-password",
          label: "Change User Password",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/change-password-admin",
          label: "Change User Password (Admin)",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/login",
          label: "Login",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/create-organization",
          label: "Create an Organization",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/add-user-to-organization",
          label: "Add a User to an Organization",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-user-from-organization",
          label: "Remove a User from an Organization",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/upsert-permission",
          label: "Upsert a Permission",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/is-authorized",
          label: "Is Authorized",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-public-key",
          label: "Get Public Key",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/register",
          label: "Register",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/create-role",
          label: "Create a New Role",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-role",
          label: "Delete a Role",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/add-permission-to-role",
          label: "Add Permission to Role",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-permissions-by-role",
          label: "Get Permissions by Role",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/set-role-permission",
          label: "Set Role Permissions",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-roles",
          label: "Get all Roles",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/create-user",
          label: "Create a New User",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-user",
          label: "Delete a User",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/save-user-profile",
          label: "Save User Profile",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/read-user-by-token",
          label: "Read User by Token",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-users",
          label: "List Users",
          className: "api-method post",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
