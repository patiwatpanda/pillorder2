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
    EntDentistEdges,
    EntDentistEdgesFromJSON,
    EntDentistEdgesFromJSONTyped,
    EntDentistEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntDentist
 */
export interface EntDentist {
    /**
     * DentistName holds the value of the "Dentist_name" field.
     * @type {string}
     * @memberof EntDentist
     */
    dentistName?: string;
    /**
     * 
     * @type {EntDentistEdges}
     * @memberof EntDentist
     */
    edges?: EntDentistEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntDentist
     */
    id?: number;
}

export function EntDentistFromJSON(json: any): EntDentist {
    return EntDentistFromJSONTyped(json, false);
}

export function EntDentistFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntDentist {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dentistName': !exists(json, 'Dentist_name') ? undefined : json['Dentist_name'],
        'edges': !exists(json, 'edges') ? undefined : EntDentistEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntDentistToJSON(value?: EntDentist | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Dentist_name': value.dentistName,
        'edges': EntDentistEdgesToJSON(value.edges),
        'id': value.id,
    };
}

