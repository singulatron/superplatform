export interface FieldSelector {
    field?: string;
    oneOf?: string[];
    any?: boolean;
}
export declare function field(fieldName: string): FieldSelector;
export declare function fields(fieldNames: string[]): FieldSelector;
export declare function anyField(): FieldSelector;
export interface Query {
    conditions?: Condition[];
    after?: any[];
    limit?: number;
    orderBys?: OrderBy[];
    count?: boolean;
}
export interface OrderBy {
    /** The field by which to order the results */
    field?: string;
    /** Indicates whether the sorting should be in descending order. */
    desc?: boolean;
    /** When set to true, indicates that the results should be randomized instead of ordered by the Field and Desc criteria. */
    randomize?: boolean;
}
export declare function orderByRandom(): OrderBy;
export declare function orderByField(field: string, desc: boolean): OrderBy;
export interface Condition {
    equal?: EqualCondition;
    all?: AllCondition;
    startsWith?: StartsWithCondition;
    contains?: ContainsCondition;
}
export declare function conditionFieldIs(condition: Condition, fieldName: string): boolean;
export declare function conditionField(condition: Condition): string;
export declare function conditionsToKeyValue(conditions: Condition[]): {
    [key: string]: any;
};
export declare function conditionValue(condition: Condition): any;
export declare function queryHasFieldCondition(query: Query, fieldName: string): boolean;
export interface EqualCondition {
    selector: FieldSelector;
    value: any;
}
export interface StartsWithCondition {
    selector: FieldSelector;
    value: any;
}
export interface ContainsCondition {
    selector: FieldSelector;
    value: any;
}
export interface AllCondition {
}
export declare function equal(selector: FieldSelector, value: any): Condition;
export declare function startsWith(selector: FieldSelector, value: any): Condition;
export declare function contains(selector: FieldSelector, value: any): Condition;
export declare function all(): Condition;
export declare function id(id: string): Condition;
export declare function userId(id: string): Condition;
export interface GenericObject {
    id: string;
    createdAt: string;
    updatedAt: string;
    userId?: string;
    table: string;
    data: any;
    public?: boolean;
}
export interface CreateRequest {
    object: GenericObject;
}
export interface CreateResponse {
}
export interface UpdateRequest {
    table: string;
    conditions: Condition[];
    object: GenericObject;
}
export interface UpdateResponse {
}
export interface DeleteRequest {
    table: string;
    conditions: Condition[];
}
export interface DeleteResponse {
}
export interface FindRequest {
    table?: string;
    conditions?: Condition[];
    public?: boolean;
    orderBys?: OrderBy[];
}
export interface FindResponse {
    objects: GenericObject[];
}
export interface UpsertRequest {
    object: GenericObject;
}
export interface UpsertResponse {
}
