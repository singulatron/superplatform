import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "doc",
      id: "singulatron/singulatron",
    },
    {
      type: "category",
      label: "chat",
      items: [
        {
          type: "doc",
          id: "singulatron/add-message",
          label: "Add Message",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-message",
          label: "Delete Message",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-messages",
          label: "Get Messages",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-thread",
          label: "Get Thread",
          className: "api-method post",
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
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/update-thread",
          label: "Update Thread",
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
      label: "config",
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
      label: "download",
      items: [
        {
          type: "doc",
          id: "singulatron/do",
          label: "Do",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/get-download",
          label: "Get Download",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/list",
          label: "List",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/pause",
          label: "Pause",
          className: "api-method post",
        },
      ],
    },
    {
      type: "category",
      label: "firehose",
      items: [
        {
          type: "doc",
          id: "singulatron/subscribe",
          label: "Subscribe",
          className: "api-method get",
        },
      ],
    },
    {
      type: "category",
      label: "generic",
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
          className: "api-method post",
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
      label: "model",
      items: [
        {
          type: "doc",
          id: "singulatron/list-models",
          label: "List Models",
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
          id: "singulatron/get-model-status",
          label: "Get Model Status",
          className: "api-method get",
        },
        {
          type: "doc",
          id: "singulatron/start-the-default-model",
          label: "Start the Default Model",
          className: "api-method put",
        },
      ],
    },
    {
      type: "category",
      label: "prompts",
      items: [
        {
          type: "doc",
          id: "singulatron/add-prompt",
          label: "Add Prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/list-prompts",
          label: "List Prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-prompt",
          label: "Remove Prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/subscribe-to-prompt",
          label: "Subscribe to Prompt",
          className: "api-method get",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
