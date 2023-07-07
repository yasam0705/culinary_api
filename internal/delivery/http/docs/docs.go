// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/recipe": {
            "get": {
                "description": "Recipe List",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregator"
                ],
                "summary": "Recipe List",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Recipe"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorBadRequest"
                        }
                    }
                }
            },
            "post": {
                "description": "Create recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregator"
                ],
                "summary": "Create recipe",
                "parameters": [
                    {
                        "description": "limit",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateAggregatorRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateAggregatorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorBadRequest"
                        }
                    }
                }
            }
        },
        "/v1/recipe/{id}": {
            "get": {
                "description": "Recipe by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregator"
                ],
                "summary": "Recipe by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "recipe_id",
                        "name": "id",
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
                                "$ref": "#/definitions/models.CulinaryAggregator"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorBadRequest"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete recipe by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregator"
                ],
                "summary": "Delete recipe by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "recipe_id",
                        "name": "id",
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
                                "$ref": "#/definitions/models.CulinaryAggregator"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorBadRequest"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CookingSteps": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "order_number": {
                    "type": "integer"
                },
                "recipe_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateAggregatorRequest": {
            "type": "object",
            "properties": {
                "cooking_steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CreateCookingStepsRequest"
                    }
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CreateIngredientsRequest"
                    }
                },
                "recipe": {
                    "$ref": "#/definitions/models.CreateRecipeRequest"
                }
            }
        },
        "models.CreateAggregatorResponse": {
            "type": "object",
            "properties": {
                "guid": {
                    "type": "string"
                }
            }
        },
        "models.CreateCookingStepsRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "order_number": {
                    "type": "integer"
                }
            }
        },
        "models.CreateIngredientsRequest": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "dimension": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.CreateRecipeRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.CulinaryAggregator": {
            "type": "object",
            "properties": {
                "cooking_steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CookingSteps"
                    }
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Ingredients"
                    }
                },
                "recipe": {
                    "$ref": "#/definitions/models.Recipe"
                }
            }
        },
        "models.ErrorBadRequest": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "models.Ingredients": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "number"
                },
                "dimension": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Recipe": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
