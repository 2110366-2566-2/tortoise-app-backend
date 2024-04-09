// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/pets/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get filtered pets by filter params",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Get filtered pets",
                "operationId": "GetFilteredPets",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category of pet",
                        "name": "category",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Species of pet",
                        "name": "species",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sex of pet",
                        "name": "sex",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Behavior of pet",
                        "name": "behavior",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum age of pet",
                        "name": "minAge",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum age of pet",
                        "name": "maxAge",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum weight of pet",
                        "name": "minWeight",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum weight of pet",
                        "name": "maxWeight",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Minimum price of pet",
                        "name": "minPrice",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Maximum price of pet",
                        "name": "maxPrice",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PetCardResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pets/categories": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all categories that system have",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Master Data"
                ],
                "summary": "Get all categories that system have",
                "operationId": "GetCategories",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MasterDataCategoryResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pets/master": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all master data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Master Data"
                ],
                "summary": "Get all master data",
                "operationId": "GetMasterData",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllMasterDataResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pets/master/{category}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get master data by category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Master Data"
                ],
                "summary": "Get master data by category",
                "operationId": "GetMasterDataByCategory",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category of master data",
                        "name": "category",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.MasterDataResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pets/pet/{petID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get single pet by petID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Get single pet by petID",
                "operationId": "GetPetByPetID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the pet to perform the operation on",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PetResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update pet by pet ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Update pet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the pet to perform the operation on",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pet object that needs to be updated",
                        "name": "Pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "return updated pet",
                        "schema": {
                            "$ref": "#/definitions/models.PetResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete pet by pet ID and delete pet from user's pets",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Delete pet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the pet to delete",
                        "name": "petID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "return deleted count",
                        "schema": {
                            "$ref": "#/definitions/models.DeletePetResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/pets/seller/{sellerID}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get pets by sellerID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Get pets by sellerID",
                "operationId": "GetPetBySeller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the seller to perform the operation on",
                        "name": "sellerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.PetCardResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Creat a new pet",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Pets"
                ],
                "summary": "Create a new pet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID of the seller to perform the operation on",
                        "name": "sellerID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Pet object that needs to be created",
                        "name": "Pet",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PetRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "return created pet",
                        "schema": {
                            "$ref": "#/definitions/models.PetResponse"
                        }
                    },
                    "400": {
                        "description": "bad request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payment/confirm": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Confirm a payment for a transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Confirm a payment",
                "parameters": [
                    {
                        "description": "Payment object",
                        "name": "Payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PaymentIntent"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ConfirmPaymentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/payment/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create a new payment for a transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Payments"
                ],
                "summary": "Create a new payment",
                "parameters": [
                    {
                        "description": "Payment object",
                        "name": "payment",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePaymentBody"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.CreatePaymentResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AllMasterDataResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.MasterData"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.ConfirmPaymentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "transaction_id": {
                            "type": "string",
                            "example": "60163b3be1e8712c4e7f35cf"
                        }
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.CreatePaymentBody": {
            "type": "object",
            "properties": {
                "buyer_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35cf"
                },
                "pet_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35cf"
                },
                "price": {
                    "type": "integer",
                    "example": 100
                },
                "seller_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35ce"
                }
            }
        },
        "models.CreatePaymentResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "payment_id": {
                            "type": "string",
                            "example": "123456789"
                        },
                        "transaction_id": {
                            "type": "string",
                            "example": "60163b3be1e8712c4e7f35cf"
                        }
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.DeletePetResponse": {
            "type": "object",
            "properties": {
                "delete_count": {
                    "type": "integer",
                    "example": 1
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "some error message"
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "models.MasterData": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "example": "Dog"
                },
                "species": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Golden Retriever",
                        "Poodle",
                        "Bulldog",
                        "Pug",
                        "Chihuahua"
                    ]
                },
                "species_count": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "models.MasterDataCategoryResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 3
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "categories": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            },
                            "example": [
                                "Dog",
                                "Cat",
                                "Bird"
                            ]
                        }
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.MasterDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.MasterData"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.Medical_record": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "Routine checkup"
                },
                "medical_date": {
                    "type": "string",
                    "example": "2024-04-08"
                },
                "medical_id": {
                    "type": "string",
                    "example": "123456789"
                }
            }
        },
        "models.PaymentIntent": {
            "type": "object",
            "properties": {
                "payment_id": {
                    "type": "string",
                    "example": "123456789"
                },
                "payment_method": {
                    "type": "string",
                    "example": "card"
                },
                "transaction_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35cf"
                }
            }
        },
        "models.Pet": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 3
                },
                "behavior": {
                    "type": "string",
                    "example": "Friendly"
                },
                "category": {
                    "type": "string",
                    "example": "Dog"
                },
                "description": {
                    "type": "string",
                    "example": "A friendly and playful dog"
                },
                "id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35cf"
                },
                "is_sold": {
                    "type": "boolean",
                    "example": false
                },
                "media": {
                    "type": "string",
                    "example": "https://example.com/fluffy.jpg"
                },
                "medical_records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Medical_record"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "Fluffy"
                },
                "price": {
                    "type": "integer",
                    "example": 500
                },
                "seller_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35ce"
                },
                "sex": {
                    "type": "string",
                    "example": "male"
                },
                "species": {
                    "type": "string",
                    "example": "Golden Retriever"
                },
                "weight": {
                    "type": "number",
                    "example": 25.5
                }
            }
        },
        "models.PetCard": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string",
                    "example": "Dog"
                },
                "id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35cf"
                },
                "media": {
                    "type": "string",
                    "example": "https://example.com/fluffy.jpg"
                },
                "name": {
                    "type": "string",
                    "example": "Fluffy"
                },
                "price": {
                    "type": "integer",
                    "example": 100
                },
                "seller_id": {
                    "type": "string",
                    "example": "60163b3be1e8712c4e7f35ce"
                },
                "seller_name": {
                    "type": "string",
                    "example": "John"
                },
                "seller_surname": {
                    "type": "string",
                    "example": "Doe"
                },
                "species": {
                    "type": "string",
                    "example": "Golden Retriever"
                }
            }
        },
        "models.PetCardResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer",
                    "example": 1
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.PetCard"
                    }
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "models.PetRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 3
                },
                "behavior": {
                    "type": "string",
                    "example": "Friendly"
                },
                "category": {
                    "type": "string",
                    "example": "Dog"
                },
                "description": {
                    "type": "string",
                    "example": "A friendly and playful dog"
                },
                "is_sold": {
                    "type": "boolean",
                    "example": false
                },
                "media": {
                    "type": "string",
                    "example": "https://example.com/fluffy.jpg"
                },
                "medical_records": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Medical_record"
                    }
                },
                "name": {
                    "type": "string",
                    "example": "Fluffy"
                },
                "price": {
                    "type": "integer",
                    "example": 500
                },
                "sex": {
                    "type": "string",
                    "example": "male"
                },
                "species": {
                    "type": "string",
                    "example": "Golden Retriever"
                },
                "weight": {
                    "type": "number",
                    "example": 25.5
                }
            }
        },
        "models.PetResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.Pet"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "description": "Authorization by API key (Format: \"Bearer \u003cAPI_KEY\u003e\")",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{"https", "http"},
	Title:            "PetPal API",
	Description:      "PetPal API is a simple API for pet marketplaces.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
