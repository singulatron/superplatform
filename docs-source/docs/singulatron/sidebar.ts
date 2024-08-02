import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "doc",
      id: "singulatron/singulatron",
    },
    {
      type: "category",
      label: "Chat Service",
      items: [
        {
          type: "doc",
          id: "singulatron/delete-message",
          label: "Delete Message",
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
          label: "Delete Thread",
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
          label: "Get Messages",
          className: "api-method get",
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
      label: "Config Service",
      items: [
        {
          type: "doc",
          id: "singulatron/get",
          label: "Get",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/save",
          label: "Save",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Docker Service",
      items: [
        {
          type: "doc",
          id: "singulatron/launch-a-docker-container",
          label: "Launch a Docker Container",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/check-if-a-container-is-running",
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
          id: "singulatron/get-docker-host",
          label: "Get Docker Host",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-docker-service-information",
          label: "Get Docker Service Information",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Download Service",
      items: [
        {
          type: "doc",
          id: "singulatron/download-a-file",
          label: "Download a File",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/get-a-download",
          label: "Get a Download",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/pause-a-download",
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
      label: "Firehose Service",
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
      label: "Generic Service",
      items: [
        {
          type: "doc",
          id: "singulatron/create-a-generic-object",
          label: "Create a Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-a-generic-object",
          label: "Delete a Generic Object",
          className: "api-method delete",
        },
        {
          type: "doc",
          id: "singulatron/find-generic-objects",
          label: "Find Generic Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/update-generic-objects",
          label: "Update Generic Objects",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/upsert-a-generic-object",
          label: "Upsert a Generic Object",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "Model Service",
      items: [
        {
          type: "doc",
          id: "singulatron/get-model-status",
          label: "Get Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/get-a-model",
          label: "Get a Model",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/make-a-model-default",
          label: "Make a Model Default",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/start-a-model",
          label: "Start a Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/start-the-default-model",
          label: "Start the Default Model",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/list-models",
          label: "List Models",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "Prompt Service",
      items: [
        {
          type: "doc",
          id: "singulatron/subscribe-to-prompt",
          label: "Subscribe to Prompt",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/list-prompts",
          label: "List Prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/add-prompt",
          label: "Add Prompt",
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
      label: "User Service",
      items: [
        {
          type: "doc",
          id: "singulatron/change-user-password",
          label: "Change User Password",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/change-user-password-admin",
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
          id: "singulatron/upsert-a-permission",
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
          id: "singulatron/register-a-new-user",
          label: "Register a New User",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/create-a-new-role",
          label: "Create a New Role",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/add-permission-to-role",
          label: "Add Permission to Role",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/set-role-permissions",
          label: "Set Role Permissions",
          className: "api-method put",
        },
        {
          type: "doc",
          id: "singulatron/create-a-new-user",
          label: "Create a New User",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/read-user-by-token",
          label: "Read User by Token",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/list-users",
          label: "List Users",
          className: "api-method post",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
