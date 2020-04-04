// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple server for my homework. It can make request towards a PSQL database, and get full tables of data from it.",
    "title": "Homework server",
    "contact": {
      "email": "szmgepesz@gmail.com"
    },
    "version": "2.0.0"
  },
  "host": "petstore.swagger.io",
  "basePath": "/",
  "paths": {
    "/CreateTable": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Create a table with the given code.",
        "operationId": "CreateTable",
        "parameters": [
          {
            "type": "integer",
            "description": "Type of the table that have to be created.",
            "name": "TableType",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Table is ready."
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/GetDataPointDescriptionTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetDataPointDescriptionTable",
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DataPointDescription"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    },
    "/GetDataPointTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetDataPointTable",
        "parameters": [
          {
            "type": "integer",
            "description": "Type of the table that have to be created.",
            "name": "TableType",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DataPoint"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    },
    "/GetModuleTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetModule",
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Module"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    }
  },
  "definitions": {
    "DataPoint": {
      "type": "object",
      "properties": {
        "DataPointDescription": {
          "type": "object",
          "format": "$ref:\"#/definitions/DataPointDescription\""
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Module": {
          "type": "object",
          "format": "$ref:\"#/definitions/Module\""
        },
        "Timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "Value": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "DataPointDescription": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Max": {
          "type": "integer",
          "format": "int64"
        },
        "Min": {
          "type": "integer",
          "format": "int64"
        },
        "Name": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    },
    "Module": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Max": {
          "type": "integer"
        },
        "Name": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "The model that handle the requests",
      "name": "SQLWebHandler"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is a simple server for my homework. It can make request towards a PSQL database, and get full tables of data from it.",
    "title": "Homework server",
    "contact": {
      "email": "szmgepesz@gmail.com"
    },
    "version": "2.0.0"
  },
  "host": "petstore.swagger.io",
  "basePath": "/",
  "paths": {
    "/CreateTable": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Create a table with the given code.",
        "operationId": "CreateTable",
        "parameters": [
          {
            "type": "integer",
            "description": "Type of the table that have to be created.",
            "name": "TableType",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Table is ready."
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/GetDataPointDescriptionTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetDataPointDescriptionTable",
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DataPointDescription"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    },
    "/GetDataPointTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetDataPointTable",
        "parameters": [
          {
            "type": "integer",
            "description": "Type of the table that have to be created.",
            "name": "TableType",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/DataPoint"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    },
    "/GetModuleTable": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "SQLWebHandler"
        ],
        "summary": "Get the specified table from the SQL server.",
        "operationId": "GetModule",
        "responses": {
          "200": {
            "description": "JSON object containing table information",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Module"
              }
            }
          },
          "404": {
            "description": "There is no data to send"
          }
        }
      }
    }
  },
  "definitions": {
    "DataPoint": {
      "type": "object",
      "properties": {
        "DataPointDescription": {
          "type": "object",
          "format": "$ref:\"#/definitions/DataPointDescription\""
        },
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Module": {
          "type": "object",
          "format": "$ref:\"#/definitions/Module\""
        },
        "Timestamp": {
          "type": "string",
          "format": "date-time"
        },
        "Value": {
          "type": "number",
          "format": "float"
        }
      }
    },
    "DataPointDescription": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Max": {
          "type": "integer",
          "format": "int64"
        },
        "Min": {
          "type": "integer",
          "format": "int64"
        },
        "Name": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    },
    "Module": {
      "type": "object",
      "properties": {
        "ID": {
          "type": "integer",
          "format": "int64"
        },
        "Max": {
          "type": "integer"
        },
        "Name": {
          "type": "string"
        },
        "Title": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "The model that handle the requests",
      "name": "SQLWebHandler"
    }
  ]
}`))
}
