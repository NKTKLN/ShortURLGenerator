// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "url": "https://nktkln.com",
            "email": "nktkln@nktkln.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/add": {
            "post": {
                "description": "Reference abbreviation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL",
                        "name": "url",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/others.ApiUrl"
                        }
                    }
                }
            }
        },
        "/archive": {
            "get": {
                "description": "Reference abbreviation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Archive",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/others.Api"
                        }
                    }
                }
            }
        },
        "/delete": {
            "post": {
                "description": "Delete url from db",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete url",
                "parameters": [
                    {
                        "type": "string",
                        "description": "URL id",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Password",
                        "name": "password",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/others.Api"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "others.Api": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                }
            }
        },
        "others.ApiUrl": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
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
	Version:     "2.0",
	Host:        "",
	BasePath:    "/api",
	Schemes:     []string{},
	Title:       "Short URL Generator API",
	Description: "This api will help you control your links from outside the site",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
