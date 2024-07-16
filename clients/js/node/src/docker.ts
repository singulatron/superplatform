import { ClientOptions, call } from "./util";
import * as docker from "@singulatron/types";

export class DockerService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options.address!, this.options.apiKey!, endpoint, request);
  }

  async dockerInfo(): Promise<docker.DockerInfoResponse> {
    return this.call("/docker/info", {});
  }
}
