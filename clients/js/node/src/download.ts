import { ClientOptions, call } from "./util";
import * as download from "@singulatron/types";

export class DownloadService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any, method: string = "POST"): Promise<any> {
    return call(this.options, endpoint, request);
  }

  async download(url: string) {
    this.call("/download-service/download", { url: url }, "PUT");
  }

  async pause(url: string) {
    this.call(`/download-service/${encodeURIComponent(url)}/pause`, {}, "PUT");
  }

  async list(): Promise<download.ListResponse> {
    return this.call("/download-service/downloads", {});
  }

  async get(url: string): Promise<download.GetResponse> {
    return this.call(
      `/download-service/download/${encodeURIComponent(url)}`,
      {},
      "GET"
    );
  }
}
