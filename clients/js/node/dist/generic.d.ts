import { ClientOptions } from "./util";
import * as generic from "@singulatron/types";
export declare class GenericService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    create(object: generic.GenericObject): Promise<void>;
    find(options: generic.FindRequest): Promise<generic.FindResponse>;
    upsert(object: generic.GenericObject): Promise<void>;
    update(table: string, conditions: generic.Condition[], object: generic.GenericObject): Promise<generic.UpdateResponse>;
    delete(table: string, conditions: generic.Condition[]): Promise<generic.DeleteResponse>;
}
