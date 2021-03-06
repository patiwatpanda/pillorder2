/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPillorder,
    EntPillorderFromJSON,
    EntPillorderFromJSONTyped,
    EntPillorderToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntEmployeeEdges
 */
export interface EntEmployeeEdges {
    /**
     * Pillorders holds the value of the pillorders edge.
     * @type {Array<EntPillorder>}
     * @memberof EntEmployeeEdges
     */
    pillorders?: Array<EntPillorder>;
}

export function EntEmployeeEdgesFromJSON(json: any): EntEmployeeEdges {
    return EntEmployeeEdgesFromJSONTyped(json, false);
}

export function EntEmployeeEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntEmployeeEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pillorders': !exists(json, 'pillorders') ? undefined : ((json['pillorders'] as Array<any>).map(EntPillorderFromJSON)),
    };
}

export function EntEmployeeEdgesToJSON(value?: EntEmployeeEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pillorders': value.pillorders === undefined ? undefined : ((value.pillorders as Array<any>).map(EntPillorderToJSON)),
    };
}


