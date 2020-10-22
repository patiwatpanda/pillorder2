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
 * @interface EntDentistEdges
 */
export interface EntDentistEdges {
    /**
     * Pillorders holds the value of the pillorders edge.
     * @type {Array<EntPillorder>}
     * @memberof EntDentistEdges
     */
    pillorders?: Array<EntPillorder>;
}

export function EntDentistEdgesFromJSON(json: any): EntDentistEdges {
    return EntDentistEdgesFromJSONTyped(json, false);
}

export function EntDentistEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentistEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pillorders': !exists(json, 'pillorders') ? undefined : ((json['pillorders'] as Array<any>).map(EntPillorderFromJSON)),
    };
}

export function EntDentistEdgesToJSON(value?: EntDentistEdges | null): any {
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


