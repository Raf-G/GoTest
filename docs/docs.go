// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Rovshan Gasanov",
            "email": "rovshan27121@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/baskets/product": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Baskets"
                ],
                "summary": "Add product to basket",
                "parameters": [
                    {
                        "description": "new product added",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.BasketProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.BasketProduct"
                        }
                    }
                }
            }
        },
        "/baskets/product/{basket_id}": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Baskets"
                ],
                "summary": "Delete product to basket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BasketID",
                        "name": "basket_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/baskets/product/{product_id}/{basket_id}": {
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Baskets"
                ],
                "summary": "Decrease quantity product to basket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "BasketID",
                        "name": "basket_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new product added",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.BasketProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.BasketProduct"
                        }
                    }
                }
            }
        },
        "/baskets/{basket_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Baskets"
                ],
                "summary": "Get basket",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "BasketID",
                        "name": "basket_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Basket"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Order"
                            }
                        }
                    }
                }
            }
        },
        "/orders/{order_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Get order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "OrderID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Delete order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "OrderID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/orders/{user_id}": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Add order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Order"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get products",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Product"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Add product",
                "parameters": [
                    {
                        "description": "new product",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                }
            }
        },
        "/products/{product_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Edit product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "edit product",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Product"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/jsondelete"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/reviews": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Get reviews",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Review"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Add review",
                "parameters": [
                    {
                        "description": "new review",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Review"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Review"
                        }
                    }
                }
            }
        },
        "/reviews/product/{product_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Get reviews product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ProductID",
                        "name": "product_id",
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
                                "$ref": "#/definitions/domain.Review"
                            }
                        }
                    }
                }
            }
        },
        "/reviews/{review_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Get review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ReviewID",
                        "name": "review_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Review"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Edit review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ReviewID",
                        "name": "review_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "new review",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Review"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Review"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Delete review",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ReviewID",
                        "name": "review_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/roles": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "summary": "Get roles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Role"
                            }
                        }
                    }
                }
            }
        },
        "/roles/{role_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Roles"
                ],
                "summary": "Get role",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "RoleID",
                        "name": "role_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Role"
                        }
                    }
                }
            }
        },
        "/statuses": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statuses"
                ],
                "summary": "Get statuses",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Status"
                            }
                        }
                    }
                }
            }
        },
        "/statuses/{status_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statuses"
                ],
                "summary": "Get status",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "StatusID",
                        "name": "status_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Status"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.User"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Add user",
                "parameters": [
                    {
                        "description": "new user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            }
        },
        "/users/{user_id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Edit user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "edit user",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "UserID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Basket": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.BasketProduct"
                    }
                },
                "total_price": {
                    "type": "integer",
                    "example": 400
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "domain.BasketProduct": {
            "type": "object",
            "properties": {
                "basket_id": {
                    "type": "integer",
                    "example": 1
                },
                "count": {
                    "type": "integer",
                    "example": 2
                },
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "product_id": {
                    "type": "integer",
                    "example": 1
                },
                "total_price": {
                    "type": "integer",
                    "example": 400
                }
            }
        },
        "domain.Order": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status_id": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Product": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "test description"
                },
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "name": {
                    "type": "string",
                    "example": "testProduct"
                },
                "price": {
                    "type": "integer",
                    "example": 300
                }
            }
        },
        "domain.Review": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "test description"
                },
                "grade": {
                    "type": "integer",
                    "example": 5
                },
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "product_id": {
                    "type": "integer",
                    "example": 1
                },
                "user_id": {
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "domain.Role": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "administrator"
                }
            }
        },
        "domain.Status": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "В обработке"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "login": {
                    "type": "string",
                    "example": "testLogin"
                },
                "name": {
                    "type": "string",
                    "example": "nameTest"
                },
                "password": {
                    "type": "string",
                    "example": "qweqwe122"
                },
                "role": {
                    "type": "integer",
                    "example": 1
                },
                "surname": {
                    "type": "string",
                    "example": "surnameTest"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8181",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "GoTest Swagger API",
	Description:      "Swagger API for Golang GoTest.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
