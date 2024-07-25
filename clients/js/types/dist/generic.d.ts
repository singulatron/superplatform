export interface FieldSelector {
    field?: string;
    oneOf?: string[];
    any?: boolean;
}
export declare function field(fieldName: string): FieldSelector;
export declare function fields(fieldNames: string[]): FieldSelector;
export declare function anyField(): FieldSelector;
export interface Query {
    /** Conditions are filtering options of a query. */
    conditions?: Condition[];
    /** After is used for paginations. Instead of offset-based pagination,
     * we support cursor-based pagination because it works better in a scalable,
     * distributed environment.
     */
    after?: any[];
    /** Limit the number of records in the result set. */
    limit?: number;
    /** OrderBys order the result set. */
    orderBys?: OrderBy[];
    /** Count true means return the count of the dataset filtered by Conditions
     * without after or limit.
     */
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
export interface GenericObjectCreateFields {
    id?: string;
    table: string;
    data: any;
    /** Public determines if the object is visible to all users.
     * When it's false the entry is only visible to the user who created it.
     * When it's true the entry is visible to everyone.
     */
    public?: boolean;
}
export interface GenericObject extends GenericObjectCreateFields {
    createdAt: string;
    updatedAt: string;
    userId?: string;
}
export interface CreateRequest {
    object: GenericObjectCreateFields;
}
export interface CreateResponse {
    object: GenericObject;
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
    query?: Query;
    public?: boolean;
}
export interface FindResponse {
    objects: GenericObject[];
}
export interface UpsertRequest {
    object: GenericObjectCreateFields;
}
export interface UpsertResponse {
}
