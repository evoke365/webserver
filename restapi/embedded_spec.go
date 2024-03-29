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
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Reusable authentication server from evoke365.net",
    "title": "evoke365.net OpenAPI spec",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "aus.jacy@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "api.evoke365.net",
  "basePath": "/v1",
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Check if service is healthy",
        "operationId": "healthz",
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/note": {
      "post": {
        "description": "user creates a new note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "user creates a new note",
        "operationId": "AddNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "the data of the note",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AddNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "Invalid request body supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/note/{id}": {
      "put": {
        "description": "update a note of a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "update a note",
        "operationId": "updateNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the note to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated a note",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid request data supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      },
      "delete": {
        "description": "delete a note of a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "delete a note",
        "operationId": "deleteNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the note to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid request data supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/notes/{userId}": {
      "get": {
        "description": "Retrieve notes by User ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "Get Notes of a user",
        "operationId": "getNotes",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of User",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "204": {
            "description": "user not found"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/profile/authenticate": {
      "post": {
        "description": "used for authetication between internal systems",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Autheticate a profile",
        "operationId": "autheticateProfile",
        "parameters": [
          {
            "description": "token to authenticate",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AutheticateProfileRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/profile/{id}": {
      "post": {
        "description": "used for authetication between internal systems",
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Autheticate a profile",
        "operationId": "getProfile",
        "parameters": [
          {
            "type": "string",
            "description": "ID of Profile",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "204": {
            "description": "profile not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/find/{id}": {
      "get": {
        "description": "Find a user by User ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Find user",
        "operationId": "findUser",
        "parameters": [
          {
            "type": "string",
            "description": "ID of User",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "204": {
            "description": "user not found"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/forget": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "start forget password process",
        "operationId": "forgetPassword",
        "parameters": [
          {
            "description": "User ID to reset password",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserForgetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "Log in a user",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserLoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "Invalid username/password supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/password": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "set new user password",
        "operationId": "newPassword",
        "parameters": [
          {
            "description": "set new user password",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserPasswordRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/signup": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Sign up a new user",
        "operationId": "SignupUser",
        "parameters": [
          {
            "description": "Sign up a new user",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserSignupRequests"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/verify": {
      "put": {
        "description": "This can only be done by the logged in user.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Verify a user",
        "operationId": "verifyUser",
        "parameters": [
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserVerifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "AddNoteRequest": {
      "type": "object",
      "properties": {
        "body": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "AutheticateProfileRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "UpdateNoteRequest": {
      "type": "object",
      "properties": {
        "isImportant": {
          "type": "boolean"
        }
      }
    },
    "UserForgetRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "UserLoginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "UserPasswordRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "UserSignupRequests": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "timezone": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "UserVerifyRequest": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "email": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Healthcheck endpoint",
      "name": "health"
    },
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about authentication",
      "name": "profile"
    },
    {
      "description": "Operations about note",
      "name": "note"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "Reusable authentication server from evoke365.net",
    "title": "evoke365.net OpenAPI spec",
    "termsOfService": "http://swagger.io/terms/",
    "contact": {
      "email": "aus.jacy@gmail.com"
    },
    "license": {
      "name": "Apache 2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "host": "api.evoke365.net",
  "basePath": "/v1",
  "paths": {
    "/healthz": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "health"
        ],
        "summary": "Check if service is healthy",
        "operationId": "healthz",
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/note": {
      "post": {
        "description": "user creates a new note",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "user creates a new note",
        "operationId": "AddNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "description": "the data of the note",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AddNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "Invalid request body supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/note/{id}": {
      "put": {
        "description": "update a note of a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "update a note",
        "operationId": "updateNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the note to update",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Updated a note",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UpdateNoteRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid request data supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      },
      "delete": {
        "description": "delete a note of a user",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "delete a note",
        "operationId": "deleteNote",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "format": "uuid",
            "description": "ID of the note to delete",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid request data supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/notes/{userId}": {
      "get": {
        "description": "Retrieve notes by User ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "note"
        ],
        "summary": "Get Notes of a user",
        "operationId": "getNotes",
        "parameters": [
          {
            "type": "string",
            "format": "uuid",
            "description": "authenticated token",
            "name": "token",
            "in": "header",
            "required": true
          },
          {
            "type": "string",
            "description": "ID of User",
            "name": "userId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "204": {
            "description": "user not found"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/profile/authenticate": {
      "post": {
        "description": "used for authetication between internal systems",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Autheticate a profile",
        "operationId": "autheticateProfile",
        "parameters": [
          {
            "description": "token to authenticate",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/AutheticateProfileRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/profile/{id}": {
      "post": {
        "description": "used for authetication between internal systems",
        "produces": [
          "application/json"
        ],
        "tags": [
          "profile"
        ],
        "summary": "Autheticate a profile",
        "operationId": "getProfile",
        "parameters": [
          {
            "type": "string",
            "description": "ID of Profile",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "204": {
            "description": "profile not found"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/find/{id}": {
      "get": {
        "description": "Find a user by User ID",
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Find user",
        "operationId": "findUser",
        "parameters": [
          {
            "type": "string",
            "description": "ID of User",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "204": {
            "description": "user not found"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/forget": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "start forget password process",
        "operationId": "forgetPassword",
        "parameters": [
          {
            "description": "User ID to reset password",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserForgetRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "bad request"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/login": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Logs user into the system",
        "operationId": "loginUser",
        "parameters": [
          {
            "description": "Log in a user",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserLoginRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "400": {
            "description": "Invalid username/password supplied"
          },
          "401": {
            "description": "Unauthorised user credentials"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/password": {
      "post": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "set new user password",
        "operationId": "newPassword",
        "parameters": [
          {
            "description": "set new user password",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserPasswordRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/signup": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Sign up a new user",
        "operationId": "SignupUser",
        "parameters": [
          {
            "description": "Sign up a new user",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserSignupRequests"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful operation"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    },
    "/user/verify": {
      "put": {
        "description": "This can only be done by the logged in user.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "user"
        ],
        "summary": "Verify a user",
        "operationId": "verifyUser",
        "parameters": [
          {
            "description": "Updated user object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/UserVerifyRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful Operation"
          },
          "400": {
            "description": "Invalid user supplied"
          },
          "500": {
            "description": "internal error"
          }
        }
      }
    }
  },
  "definitions": {
    "AddNoteRequest": {
      "type": "object",
      "properties": {
        "body": {
          "type": "string"
        },
        "title": {
          "type": "string"
        }
      }
    },
    "ApiResponse": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer",
          "format": "int32"
        },
        "message": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "AutheticateProfileRequest": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string"
        }
      }
    },
    "UpdateNoteRequest": {
      "type": "object",
      "properties": {
        "isImportant": {
          "type": "boolean"
        }
      }
    },
    "UserForgetRequest": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string"
        }
      }
    },
    "UserLoginRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        }
      }
    },
    "UserPasswordRequest": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "token": {
          "type": "string"
        }
      }
    },
    "UserSignupRequests": {
      "type": "object",
      "properties": {
        "email": {
          "type": "string"
        },
        "password": {
          "type": "string"
        },
        "timezone": {
          "type": "integer",
          "format": "int32"
        }
      }
    },
    "UserVerifyRequest": {
      "type": "object",
      "properties": {
        "code": {
          "type": "string"
        },
        "email": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Healthcheck endpoint",
      "name": "health"
    },
    {
      "description": "Operations about user",
      "name": "user"
    },
    {
      "description": "Operations about authentication",
      "name": "profile"
    },
    {
      "description": "Operations about note",
      "name": "note"
    }
  ],
  "externalDocs": {
    "description": "Find out more about Swagger",
    "url": "http://swagger.io"
  }
}`))
}
