import { ClientOptions, call, uuid } from "./util";
import * as prompt from "@singulatron/types";

export class PromptService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options, endpoint, request);
  }

  async promptAdd(prompt: prompt.AddPromptRequest): Promise<void> {
    if (!prompt.id) {
      prompt.id = uuid();
    }

    return this.call("'/prompt-service/add", prompt);
  }

  async promptRemove(promptId: string): Promise<void> {
    const request: prompt.RemovePromptRequest = { promptId: promptId };
    return this.call("'/prompt-service/remove", request);
  }

  async promptList(
    request: prompt.ListPromptsRequest
  ): Promise<prompt.ListPromptsResponse> {
    return this.call("'/prompt-service/list", request);
  }
}
