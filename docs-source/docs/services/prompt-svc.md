---
sidebar_position: 3
tags:
  - prompt-svc
  - prompts
  - ai
  - service
---

# Prompt Svc

The prompt service provides an easy to use interface to prompt LLMs and use AI models. Aims to serve humans and machines alike with its resilient queue based architecture.

> This page is a high level overview of the `Prompt Svc`. For API documentation, please see the [Prompt Svc API documentation](/docs/singulatron/upsert-instance).

## How It Works

The prompt service accepts prompts,maintains a list of prompts and processes them as soon as it's able to. It handles retries of prompts that errored with an exponential backoff.

It's able to stream back LLM responses, or it can respond syncronously if that's what the caller wants, for details see the [`/prompt-svc/prompt`](/docs/singulatron/add-prompt).

It talks to two services primarily:

- [Chat Svc] to save prompt responses to chat threads and messages.
- [Model Svc] to get the address of the running AI models, see their status etc.
