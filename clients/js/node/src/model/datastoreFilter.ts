/**
 * Singulatron
 * AI management and development platform.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: sales@singulatron.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { DatastoreOp } from './datastoreOp';

export class DatastoreFilter {
    'fields'?: Array<string>;
    'op'?: DatastoreOp;
    'values'?: Array<object>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "fields",
            "baseName": "fields",
            "type": "Array<string>"
        },
        {
            "name": "op",
            "baseName": "op",
            "type": "DatastoreOp"
        },
        {
            "name": "values",
            "baseName": "values",
            "type": "Array<object>"
        }    ];

    static getAttributeTypeMap() {
        return DatastoreFilter.attributeTypeMap;
    }
}

export namespace DatastoreFilter {
}