import { ClientOptions, call } from "./util";
import * as config from "@singulatron/types";

export class ConfigService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options.address!, this.options.apiKey!, endpoint, request);
  }

  async configGet(): Promise<config.ConfigGetResponse> {
    return await this.call("/config/get", {});
  }
}
