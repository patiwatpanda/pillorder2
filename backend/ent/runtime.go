// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/patiwatpanda/app/ent/dentist"
	"github.com/patiwatpanda/app/ent/employee"
	"github.com/patiwatpanda/app/ent/patient"
	"github.com/patiwatpanda/app/ent/pillorder"
	"github.com/patiwatpanda/app/ent/pillorderitem"
	"github.com/patiwatpanda/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	dentistFields := schema.Dentist{}.Fields()
	_ = dentistFields
	// dentistDescDentistName is the schema descriptor for Dentist_name field.
	dentistDescDentistName := dentistFields[0].Descriptor()
	// dentist.DentistNameValidator is a validator for the "Dentist_name" field. It is called by the builders before save.
	dentist.DentistNameValidator = dentistDescDentistName.Validators[0].(func(string) error)
	employeeFields := schema.Employee{}.Fields()
	_ = employeeFields
	// employeeDescEmployeeName is the schema descriptor for employee_name field.
	employeeDescEmployeeName := employeeFields[0].Descriptor()
	// employee.EmployeeNameValidator is a validator for the "employee_name" field. It is called by the builders before save.
	employee.EmployeeNameValidator = employeeDescEmployeeName.Validators[0].(func(string) error)
	// employeeDescEmployeeEmail is the schema descriptor for employee_email field.
	employeeDescEmployeeEmail := employeeFields[1].Descriptor()
	// employee.EmployeeEmailValidator is a validator for the "employee_email" field. It is called by the builders before save.
	employee.EmployeeEmailValidator = employeeDescEmployeeEmail.Validators[0].(func(string) error)
	// employeeDescEmployeePassword is the schema descriptor for employee_password field.
	employeeDescEmployeePassword := employeeFields[2].Descriptor()
	// employee.EmployeePasswordValidator is a validator for the "employee_password" field. It is called by the builders before save.
	employee.EmployeePasswordValidator = employeeDescEmployeePassword.Validators[0].(func(string) error)
	patientFields := schema.Patient{}.Fields()
	_ = patientFields
	// patientDescPatientName is the schema descriptor for Patient_name field.
	patientDescPatientName := patientFields[0].Descriptor()
	// patient.PatientNameValidator is a validator for the "Patient_name" field. It is called by the builders before save.
	patient.PatientNameValidator = patientDescPatientName.Validators[0].(func(string) error)
	// patientDescPatientAge is the schema descriptor for Patient_age field.
	patientDescPatientAge := patientFields[1].Descriptor()
	// patient.PatientAgeValidator is a validator for the "Patient_age" field. It is called by the builders before save.
	patient.PatientAgeValidator = patientDescPatientAge.Validators[0].(func(int) error)
	pillorderFields := schema.Pillorder{}.Fields()
	_ = pillorderFields
	// pillorderDescPillorderNameID is the schema descriptor for PillorderNameID field.
	pillorderDescPillorderNameID := pillorderFields[0].Descriptor()
	// pillorder.PillorderNameIDValidator is a validator for the "PillorderNameID" field. It is called by the builders before save.
	pillorder.PillorderNameIDValidator = pillorderDescPillorderNameID.Validators[0].(func(string) error)
	pillorderitemFields := schema.PillorderItem{}.Fields()
	_ = pillorderitemFields
	// pillorderitemDescPillorderItemName is the schema descriptor for PillorderItem_name field.
	pillorderitemDescPillorderItemName := pillorderitemFields[0].Descriptor()
	// pillorderitem.PillorderItemNameValidator is a validator for the "PillorderItem_name" field. It is called by the builders before save.
	pillorderitem.PillorderItemNameValidator = pillorderitemDescPillorderItemName.Validators[0].(func(string) error)
}