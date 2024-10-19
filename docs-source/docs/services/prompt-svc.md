---
sidebar_position: 70
tags:
  - prompt-svc
  - prompts
  - ai
  - services
---

# Prompt Svc

The prompt service provides an easy to use interface to prompt LLMs and use AI models. Aims to serve humans and machines alike with its resilient queue based architecture.

> This page is a high level overview of the `Prompt Svc`. For more details, please see the [Prompt Svc API documentation](/docs/superplatform/upsert-instance).

## Responsibilities

The prompt service:

- Accepts prompts
- Maintains a list of prompts
- Processes prompts as soon as it's able to
- Streams prompt answers
- Handles retries of prompts that errored with an exponential backoff

It's able to stream back LLM responses, or it can respond syncronously if that's what the caller wants, for details see the [Add Prompt (`/prompt-svc/prompt`) Endpoint](/docs/superplatform/add-prompt).

## Dependencies

- [Chat Svc](/docs/services/chat-svc) to save prompt responses to chat threads and messages
- [Model Svc](/docs/services/model-svc) to get the address of the running AI models, see their status etc.

## Current Limitations

There are planned improvements for this service:

- It should manage models: start needed ones and stop unneeded ones based on the volume of prompts in the backlog
