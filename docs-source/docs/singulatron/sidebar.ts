import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
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
          id: "singulatron/create-generic-object",
          label: "Create Generic Object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-generic-object",
          label: "Delete Generic Object",
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
          id: "singulatron/upsert-generic-object",
          label: "Upsert Generic Object",
          className: "api-method post",
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
