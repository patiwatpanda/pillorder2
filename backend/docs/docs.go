// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/dentists": {
            "get": {
                "description": "list dentist entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List dentist entities",
                "operationId": "list-dentist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Dentist"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create dentist",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create dentist",
                "operationId": "create-dentist",
                "parameters": [
                    {
                        "description": "Dentist entity",
                        "name": "dentist",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Dentist"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Dentist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/dentists/{id}": {
            "get": {
                "description": "get dentist by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a dentist entity by ID",
                "operationId": "get-dentist",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Dentist"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees": {
            "get": {
                "description": "list employee entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List employee entities",
                "operationId": "list-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Employee"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create employee",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create employee",
                "operationId": "create-employee",
                "parameters": [
                    {
                        "description": " Employee entity",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "description": "get employee by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a Employee entity by ID",
                "operationId": "get-employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Employee"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/patients": {
            "get": {
                "description": "list patient entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List patient entities",
                "operationId": "list-patient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Patient"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create patient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create patient",
                "operationId": "create-patient",
                "parameters": [
                    {
                        "description": "Patient entity",
                        "name": "patient",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Patient"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Patient"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/patients/{id}": {
            "get": {
                "description": "get patient by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a patient entity by ID",
                "operationId": "get-patient",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Patient"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/pillorderItems": {
            "post": {
                "description": "Create pillorderitem",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create pillorderitem",
                "operationId": "create-pillorderitem",
                "parameters": [
                    {
                        "description": "PillorderItem entity",
                        "name": "pillorderitem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.PillorderItem"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.PillorderItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/pillorderitems": {
            "get": {
                "description": "list pillorderitem entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List pillorderitem entities",
                "operationId": "list-pillorderitem",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.PillorderItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/pillorderitems/{id}": {
            "get": {
                "description": "get PillorderItem by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a PillorderItem entity by ID",
                "operationId": "get-pillorderitem",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "PillorderItem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.PillorderItem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/pillorders": {
            "get": {
                "description": "list pillorder entities",
                "produces": [
                    "application/json"
                ],
                "summary": "List pillorder entities",
                "operationId": "list-pillorder",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ent.Pillorder"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            },
            "post": {
                "description": "Create pillorder",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create pillorder",
                "operationId": "create-pillorder",
                "parameters": [
                    {
                        "description": "Pillorder entity",
                        "name": "pillorder",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ent.Pillorder"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Pillorder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/pillorders/{id}": {
            "get": {
                "description": "get pillorder by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Get a Pillorder entity by ID",
                "operationId": "get-pillorder",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Pillorder ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ent.Pillorder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ent.Dentist": {
            "type": "object",
            "properties": {
                "Dentist_name": {
                    "description": "DentistName holds the value of the \"Dentist_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the DentistQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.DentistEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.DentistEdges": {
            "type": "object",
            "properties": {
                "pillorders": {
                    "description": "Pillorders holds the value of the pillorders edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Pillorder"
                    }
                }
            }
        },
        "ent.Employee": {
            "type": "object",
            "properties": {
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the EmployeeQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.EmployeeEdges"
                },
                "employee_email": {
                    "description": "EmployeeEmail holds the value of the \"employee_email\" field.",
                    "type": "string"
                },
                "employee_name": {
                    "description": "EmployeeName holds the value of the \"employee_name\" field.",
                    "type": "string"
                },
                "employee_password": {
                    "description": "EmployeePassword holds the value of the \"employee_password\" field.",
                    "type": "string"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.EmployeeEdges": {
            "type": "object",
            "properties": {
                "pillorders": {
                    "description": "Pillorders holds the value of the pillorders edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Pillorder"
                    }
                }
            }
        },
        "ent.Patient": {
            "type": "object",
            "properties": {
                "Patient_age": {
                    "description": "PatientAge holds the value of the \"Patient_age\" field.",
                    "type": "integer"
                },
                "Patient_name": {
                    "description": "PatientName holds the value of the \"Patient_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PatientQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PatientEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PatientEdges": {
            "type": "object",
            "properties": {
                "pillorders": {
                    "description": "Pillorders holds the value of the pillorders edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Pillorder"
                    }
                }
            }
        },
        "ent.Pillorder": {
            "type": "object",
            "properties": {
                "PillorderDate": {
                    "description": "PillorderDate holds the value of the \"PillorderDate\" field.",
                    "type": "string"
                },
                "PillorderNameID": {
                    "description": "PillorderNameID holds the value of the \"PillorderNameID\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PillorderQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PillorderEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PillorderEdges": {
            "type": "object",
            "properties": {
                "dentist": {
                    "description": "Dentist holds the value of the dentist edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Dentist"
                },
                "employee": {
                    "description": "Employee holds the value of the employee edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Employee"
                },
                "patient": {
                    "description": "Patient holds the value of the patient edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.Patient"
                },
                "pillorderitem": {
                    "description": "Pillorderitem holds the value of the pillorderitem edge.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PillorderItem"
                }
            }
        },
        "ent.PillorderItem": {
            "type": "object",
            "properties": {
                "PillorderItem_name": {
                    "description": "PillorderItemName holds the value of the \"PillorderItem_name\" field.",
                    "type": "string"
                },
                "edges": {
                    "description": "Edges holds the relations/edges for other nodes in the graph.\nThe values are being populated by the PillorderItemQuery when eager-loading is set.",
                    "type": "object",
                    "$ref": "#/definitions/ent.PillorderItemEdges"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "integer"
                }
            }
        },
        "ent.PillorderItemEdges": {
            "type": "object",
            "properties": {
                "pillorders": {
                    "description": "Pillorders holds the value of the pillorders edge.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ent.Pillorder"
                    }
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": true
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        },
        "OAuth2AccessCode": {
            "type": "oauth2",
            "flow": "accessCode",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information"
            }
        },
        "OAuth2Application": {
            "type": "oauth2",
            "flow": "application",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://example.com/oauth/authorize",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "write": " Grants write access"
            }
        },
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "https://example.com/oauth/token",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "read": " Grants read access",
                "write": " Grants write access"
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "SUT SA Example API Playlist Vidoe",
	Description: "This is a sample server for SUT SE 2563",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}