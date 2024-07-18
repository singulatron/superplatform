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

  async promptAdd(prompt: prompt.Prompt): Promise<void> {
    if (!prompt.id) {
      prompt.id = uuid();
    }
    const request: prompt.AddPromptRequest = { prompt: prompt };
    return this.call("/prompt/add", request);
  }

  async promptRemove(prompt: prompt.Prompt): Promise<void> {
    const request: prompt.RemovePromptRequest = { prompt: prompt };
    return this.call("/prompt/remove", request);
  }

  async promptList(
    request: prompt.ListPromptsRequest
  ): Promise<prompt.ListPromptsResponse> {
    return this.call("/prompt/list", request);
  }
}
