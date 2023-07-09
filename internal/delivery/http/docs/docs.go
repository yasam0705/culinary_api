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
        "/v1/cooking-step": {
            "put": {
                "description": "Update cooking step",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cooking-step"
                ],
                "summary": "Update cooking step",
                "parameters": [
                    {
                        "description": "data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateStepRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateStepResponse"
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
                "description": "Create cooking step",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cooking-step"
                ],
                "summary": "Create cooking step",
                "parameters": [
                    {
                        "description": "data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateStepRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateStepResponse"
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
        "/v1/cooking-step/{id}": {
            "delete": {
                "description": "Delete cooking step by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cooking-step"
                ],
                "summary": "Delete cooking step by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
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
                                "$ref": "#/definitions/models.DeleteStepResponse"
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
        },
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
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "cooking_time_from",
                        "name": "cooking_time_from",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "cooking_time_to",
                        "name": "cooking_time_to",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "guid,guid",
                        "name": "ingridients",
                        "in": "query"
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
            "put": {
                "description": "Update recipe",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "aggregator"
                ],
                "summary": "Update recipe",
                "parameters": [
                    {
                        "description": "data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateRecipeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateRecipeResponse"
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
                        "description": "data",
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
        "/v1/recipe-ingridient": {
            "get": {
                "description": "ingridient list",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe-ingridient"
                ],
                "summary": "ingridient list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
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
                                "$ref": "#/definitions/models.Ingridient"
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
            "put": {
                "description": "Update ingridient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe-ingridient"
                ],
                "summary": "Update ingridient",
                "parameters": [
                    {
                        "description": "data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateIngredientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateIngredientResponse"
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
                "description": "Create ingridient",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe-ingridient"
                ],
                "summary": "Create ingridient",
                "parameters": [
                    {
                        "description": "data",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateIngredientRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateIngredientResponse"
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
                "description": "Delete ingridient by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "recipe-ingridient"
                ],
                "summary": "Delete ingridient by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "recipe_id",
                        "name": "recipe_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "ingredient_id",
                        "name": "ingredient_id",
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
                                "$ref": "#/definitions/models.DeleteIngredientResponse"
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
                                "$ref": "#/definitions/models.DeleteRecipeResponse"
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
                "cooking_time": {
                    "type": "number"
                },
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
        "models.CookingStepsUpdate": {
            "type": "object",
            "properties": {
                "cooking_time": {
                    "type": "number"
                },
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
                "cooking_time": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "order_number": {
                    "type": "integer"
                }
            }
        },
        "models.CreateIngredientRequest": {
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
                },
                "recipe_id": {
                    "type": "string"
                }
            }
        },
        "models.CreateIngredientResponse": {
            "type": "object",
            "properties": {
                "guid": {
                    "type": "string"
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
        "models.CreateStepRequest": {
            "type": "object",
            "properties": {
                "cooking_time": {
                    "type": "number"
                },
                "description": {
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
        "models.CreateStepResponse": {
            "type": "object",
            "properties": {
                "guid": {
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
        "models.DeleteIngredientResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.DeleteRecipeResponse": {
            "type": "object",
            "properties": {
                "succress": {
                    "type": "boolean"
                }
            }
        },
        "models.DeleteStepResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
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
        "models.IngredientsUpdate": {
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
        "models.Ingridient": {
            "type": "object",
            "properties": {
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
                "cooking_time": {
                    "type": "number"
                },
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
        },
        "models.RecipeUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "guid": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.UpdateIngredientRequest": {
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
                },
                "recipe_id": {
                    "type": "string"
                }
            }
        },
        "models.UpdateIngredientResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.UpdateRecipeRequest": {
            "type": "object",
            "properties": {
                "cooking_steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.CookingStepsUpdate"
                    }
                },
                "ingredients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.IngredientsUpdate"
                    }
                },
                "recipe": {
                    "$ref": "#/definitions/models.RecipeUpdate"
                }
            }
        },
        "models.UpdateRecipeResponse": {
            "type": "object",
            "properties": {
                "succress": {
                    "type": "boolean"
                }
            }
        },
        "models.UpdateStepRequest": {
            "type": "object",
            "properties": {
                "cooking_time": {
                    "type": "number"
                },
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
        "models.UpdateStepResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean"
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
