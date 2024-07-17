import { ClientOptions } from "./util";
import * as generic from "@singulatron/types";
export declare class GenericService {
    private options;
    constructor(options: ClientOptions);
    call(endpoint: string, request: any): Promise<any>;
    create(table: string, object: generic.GenericObject): Promise<void>;
    find(table: string, conditions: generic.Condition[], _public?: boolean): Promise<generic.FindResponse>;
    upsert(table: string, object: generic.GenericObject): Promise<void>;
    update(table: string, conditions: generic.Condition[], object: generic.GenericObject): Promise<generic.UpdateResponse>;
    delete(table: string, conditions: generic.Condition[]): Promise<generic.DeleteResponse>;
}
