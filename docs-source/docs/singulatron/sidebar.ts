import type { SidebarsConfig } from "@docusaurus/plugin-content-docs";

const sidebar: SidebarsConfig = {
  apisidebar: [
    {
      type: "category",
      label: "chat",
      items: [
        {
          type: "doc",
          id: "singulatron/send-a-new-message-to-a-chat-thread",
          label: "Send a new message to a chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-a-message-from-a-chat-thread",
          label: "Remove a message from a chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/retrieve-messages-from-a-chat-thread",
          label: "Retrieve messages from a chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/retrieve-details-of-a-chat-thread",
          label: "Retrieve details of a chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/create-a-new-chat-thread",
          label: "Create a new chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-a-chat-thread",
          label: "Remove a chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/update-an-existing-chat-thread",
          label: "Update an existing chat thread",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/retrieve-a-list-of-chat-threads-for-a-user",
          label: "Retrieve a list of chat threads for a user",
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
          id: "singulatron/retrieve-the-current-configuration",
          label: "Retrieve the current configuration",
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
          id: "singulatron/initiate-a-file-download",
          label: "Initiate a file download",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/retrieve-download-details",
          label: "Retrieve download details",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/pause-an-ongoing-download",
          label: "Pause an ongoing download",
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
          id: "singulatron/subscribe-to-firehose-events",
          label: "Subscribe to firehose events",
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
          id: "singulatron/create-a-new-generic-object",
          label: "Create a new generic object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/delete-a-generic-object",
          label: "Delete a generic object",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/retrieve-generic-objects-based-on-criteria",
          label: "Retrieve generic objects based on criteria",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/update-generic-objects-based-on-conditions",
          label: "Update generic objects based on conditions",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/create-or-update-a-generic-object",
          label: "Create or update a generic object",
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
          label: "List prompts",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/remove-a-prompt",
          label: "Remove a prompt",
          className: "api-method post",
        },
        {
          type: "doc",
          id: "singulatron/subscribe-to-prompt-responses",
          label: "Subscribe to prompt responses",
          className: "api-method get",
        },
      ],
    },
  ],
};

export default sidebar.apisidebar;
