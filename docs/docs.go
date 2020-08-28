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
            "name": "NuOrder API Support",
            "email": "flavio.costa@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.nuorder.com/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/faq": {
            "get": {
                "description": "Get of all Questions and Answers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faq"
                ],
                "summary": "Get a list to all questions and answers from API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Faq"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new Question and Answer with the input paylod",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faq"
                ],
                "summary": "Create a new Question and Answer item",
                "parameters": [
                    {
                        "description": "Create",
                        "name": "faq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.FaqRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Faq"
                        }
                    }
                }
            }
        },
        "/api/v1/faq/{id}": {
            "get": {
                "description": "Get a question and answer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "faq"
                ],
                "summary": "Get one question and answer item from the API",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ObjectId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Faq"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Faq": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "To remove an item click on RemoveItem button"
                },
                "id": {
                    "type": "string",
                    "example": "5f484f697ee3881a0ca9a037"
                },
                "question": {
                    "type": "string",
                    "example": "How can I remove an item?"
                }
            }
        },
        "models.FaqRequest": {
            "type": "object",
            "properties": {
                "answer": {
                    "type": "string",
                    "example": "To remove an item click on RemoveItem button"
                },
                "question": {
                    "type": "string",
                    "example": "How can I remove an item?"
                }
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
	Host:        "localhost:3000",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "FAQ API",
	Description: "This is a MVP for Questions and Answers https://www.nuorder.com/ page",
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
