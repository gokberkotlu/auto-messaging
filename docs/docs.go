// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
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
        "/api/v1/message/get-sent-messages": {
            "get": {
                "description": "Retrieve a list of sent messages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "List messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Message"
                            }
                        }
                    }
                }
            }
        },
        "/api/v1/message/switch-auto-messaging-mode/{active}": {
            "get": {
                "description": "Start/Stop automatic message sending",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "messages"
                ],
                "summary": "Start/Stop Messaging",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Active Mode",
                        "name": "active",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.SwitchResDTO"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.SwitchResDTO": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "entity.Message": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "status": {
                    "$ref": "#/definitions/entity.StatusType"
                },
                "to": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "entity.StatusType": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "Active",
                "Sent"
            ]
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server celler server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
