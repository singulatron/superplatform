import { ClientOptions, call } from "./util";
import * as generic from "@singulatron/types";

export interface FindOptions {
  table: string;
  conditions: generic.Condition[];
  public: boolean;
  orderBys: generic.OrderBy[];
}

export class GenericService {
  private options: ClientOptions;

  constructor(options: ClientOptions) {
    this.options = options;
  }

  call(endpoint: string, request: any): Promise<any> {
    return call(this.options, endpoint, request);
  }

  async create(object: generic.GenericObject): Promise<void> {
    const request: generic.CreateRequest = {
      object: object,
    };

    return this.call("/generic/create", request);
  }

  async find(options: FindOptions): Promise<generic.FindResponse> {
    const request: generic.FindRequest = options;

    return this.call("/generic/find", request);
  }

  async upsert(object: generic.GenericObject): Promise<void> {
    const request: generic.UpsertRequest = {
      object: object,
    };

    return this.call("/generic/upsert", request);
  }

  async update(
    table: string,
    conditions: generic.Condition[],
    object: generic.GenericObject
  ): Promise<generic.UpdateResponse> {
    const request: generic.UpdateRequest = {
      table: table,
      conditions: conditions,
      object: object,
    };

    return this.call("/generic/update", request);
  }

  async delete(
    table: string,
    conditions: generic.Condition[]
  ): Promise<generic.DeleteResponse> {
    const request: generic.DeleteRequest = {
      table: table,
      conditions: conditions,
    };

    return this.call("/generic/delete", request);
  }
}
