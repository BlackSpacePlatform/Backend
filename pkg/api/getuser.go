package api

import (
	"net/http"

	"github.com/LensPlatform/BlackSpace/pkg/helper"
	"github.com/LensPlatform/BlackSpace/pkg/models"
)

type GetUserResponse struct {
	User models.UserORM `json:"user"`
	Error error `json:"error"`
}

// Get user by id request
// swagger:parameters getUserRequest
type GetUserRequestSwagger struct {
	// user account to create
	// in: body
	Body struct {
		// id of the user account to get
		// in: query
		// required: true
		Id uint32 `json:"result"`
	}
}

// Common operation response
// swagger:response getUserResponse
type GetUserResponseSwagger struct {
	// in: body
	Body struct {
		// error
		// required: true
		// example: error occured while processing request
		Error error `json:"error"`
		// User
		// required: true
		User models.UserORM `json:"user"`
	}
}

// swagger:route GET /v1/user/{id} User getUserRequest
//
// Get User Account By ID
//
// Returns a user account by id
//
//     Consumes:
//     - application/json
//     - application/x-protobuf
//
//     Produces:
//     - application/json
//     - application/x-protobuf
//
//     Schemes: http, https, ws, wss
//
//     Security:
//       api_key:
//       oauth: read, write
// responses:
//      200: getUserResponse
// gets a user by account id
func (s *Server) getUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		userAccount *models.UserORM
		response    GetUserResponse
	)

	id := helper.ExtractIDFromRequest(r)

	userAccount, err := s.db.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Error = err
	response.User = *userAccount

	// store the request in the database
	s.JSONResponse(w, r, response)
}
