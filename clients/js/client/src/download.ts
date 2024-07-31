import { ClientOptions, call } from "./util";
import * as download from "@singulatron/types";

export class DownloadService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options, endpoint, request);
  }

  async do(url: string) {
    this.call("/download/do", { url: url });
  }

  async pause(url: string) {
    this.call("/download/pause", { url: url });
  }

  async list(): Promise<download.ListResponse> {
    return this.call("/download/list", {});
  }

  async get(url: string): Promise<download.GetResponse> {
    return this.call("/download/get", {
      url: url,
    });
  }
}
