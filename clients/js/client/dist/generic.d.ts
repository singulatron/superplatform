import { ClientOptions } from "./util";
import * as generic from "@singulatron/types";
export declare class GenericService {
    private options;
    constructor(options: ClientOptions);
    private call;
    create(request: generic.CreateRequest): Promise<void>;
    find(options: generic.FindRequest): Promise<generic.FindResponse>;
    upsert(object: generic.GenericObjectCreateFields): Promise<void>;
    update(table: string, conditions: generic.Condition[], object: generic.GenericObject): Promise<generic.UpdateResponse>;
    delete(table: string, conditions: generic.Condition[]): Promise<generic.DeleteResponse>;
}
