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
    EntPatientEdges,
    EntPatientEdgesFromJSON,
    EntPatientEdgesFromJSONTyped,
    EntPatientEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatient
 */
export interface EntPatient {
    /**
     * PatientAge holds the value of the "Patient_age" field.
     * @type {number}
     * @memberof EntPatient
     */
    patientAge?: number;
    /**
     * PatientName holds the value of the "Patient_name" field.
     * @type {string}
     * @memberof EntPatient
     */
    patientName?: string;
    /**
     * 
     * @type {EntPatientEdges}
     * @memberof EntPatient
     */
    edges?: EntPatientEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatient
     */
    id?: number;
}

export function EntPatientFromJSON(json: any): EntPatient {
    return EntPatientFromJSONTyped(json, false);
}

export function EntPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patientAge': !exists(json, 'Patient_age') ? undefined : json['Patient_age'],
        'patientName': !exists(json, 'Patient_name') ? undefined : json['Patient_name'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntPatientToJSON(value?: EntPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Patient_age': value.patientAge,
        'Patient_name': value.patientName,
        'edges': EntPatientEdgesToJSON(value.edges),
        'id': value.id,
    };
}


