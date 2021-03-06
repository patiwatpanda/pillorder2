// Code generated by entc, DO NOT EDIT.

package pillorder

const (
	// Label holds the string label denoting the pillorder type in the database.
	Label = "pillorder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPillorderNameID holds the string denoting the pillordernameid field in the database.
	FieldPillorderNameID = "pillorder_name_id"
	// FieldPillorderDate holds the string denoting the pillorderdate field in the database.
	FieldPillorderDate = "pillorder_date"

	// EdgeEmployee holds the string denoting the employee edge name in mutations.
	EdgeEmployee = "employee"
	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeDentist holds the string denoting the dentist edge name in mutations.
	EdgeDentist = "dentist"
	// EdgePillorderitem holds the string denoting the pillorderitem edge name in mutations.
	EdgePillorderitem = "pillorderitem"

	// Table holds the table name of the pillorder in the database.
	Table = "pillorders"
	// EmployeeTable is the table the holds the employee relation/edge.
	EmployeeTable = "pillorders"
	// EmployeeInverseTable is the table name for the Employee entity.
	// It exists in this package in order to avoid circular dependency with the "employee" package.
	EmployeeInverseTable = "employees"
	// EmployeeColumn is the table column denoting the employee relation/edge.
	EmployeeColumn = "employee_id"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "pillorders"
	// PatientInverseTable is the table name for the Patient entity.
	// It exists in this package in order to avoid circular dependency with the "patient" package.
	PatientInverseTable = "patients"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
	// DentistTable is the table the holds the dentist relation/edge.
	DentistTable = "pillorders"
	// DentistInverseTable is the table name for the Dentist entity.
	// It exists in this package in order to avoid circular dependency with the "dentist" package.
	DentistInverseTable = "dentists"
	// DentistColumn is the table column denoting the dentist relation/edge.
	DentistColumn = "dentist_id"
	// PillorderitemTable is the table the holds the pillorderitem relation/edge.
	PillorderitemTable = "pillorders"
	// PillorderitemInverseTable is the table name for the PillorderItem entity.
	// It exists in this package in order to avoid circular dependency with the "pillorderitem" package.
	PillorderitemInverseTable = "pillorder_items"
	// PillorderitemColumn is the table column denoting the pillorderitem relation/edge.
	PillorderitemColumn = "pillorderitem_id"
)

// Columns holds all SQL columns for pillorder fields.
var Columns = []string{
	FieldID,
	FieldPillorderNameID,
	FieldPillorderDate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Pillorder type.
var ForeignKeys = []string{
	"dentist_id",
	"employee_id",
	"patient_id",
	"pillorderitem_id",
}

var (
	// PillorderNameIDValidator is a validator for the "PillorderNameID" field. It is called by the builders before save.
	PillorderNameIDValidator func(string) error
)
