import { ClientOptions } from "./util";
import * as generic from "@singulatron/types";
export declare class DynamicService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    create(request: generic.CreateRequest): Promise<void>;
    find(options: generic.FindRequest): Promise<generic.FindResponse>;
    upsert(object: generic.GenericObjectCreateFields): Promise<void>;
    update(table: string, conditions: generic.Condition[], object: generic.GenericObject): Promise<generic.UpdateResponse>;
    delete(table: string, conditions: generic.Condition[]): Promise<generic.DeleteResponse>;
}
