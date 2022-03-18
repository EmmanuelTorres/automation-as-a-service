// Package classification of User API
//
// Documentation for User API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package app

import "github.com/automation-as-a-service/internal/datastruct"

// NOTE: Types defined here are purely for documentation purposes

// A single id of the newly-created object
// swagger:response objectCreatedResponse
type objectCreatedResponseWrapper struct {
	// An id of the newly-created object
	// in: body
	Body struct {
		Id int64 `json:"id"`
	}
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// A list of users
// swagger:response usersResponse
type usersResponseWrapper struct {
	// All users
	// in: body
	Body []datastruct.Person
}

// A single user
// swagger:response userResponse
type userResponseWrapper struct {
	// A single user
	// in: body
	Body datastruct.Person
}

// swagger:parameters GetUser DeleteUser
type usernameParamsWrapper struct {
	// The username of the user for which the operation relates
	// in: path
	// required: true
	Username string `json:"username"`
}
