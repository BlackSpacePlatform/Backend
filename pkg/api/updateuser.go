package api

import (
	"net/http"
	"strconv"

	"github.com/LensPlatform/BlackSpace/pkg/helper"
	"github.com/LensPlatform/BlackSpace/pkg/models"
)

type UpdateUserRequest struct {
	User models.User
}

type UpdateUserResponse struct {
	Error error
}

// Update user request
// swagger:parameters updateSwaggerUser
type UpdateUserRequestSwagger struct {
	// user account to create
	// in: body
	Body struct {
		// required: true
		User models.User `json:"result"`
	}
	// user id of account to update
	// in: query
	UserAccountId uint32
}

// Common operation response
// swagger:response operationResponse
type OperationResponseSwagger struct {
	// in: body
	Body struct {
		// error
		// required: true
		// example: error occured while processing request
		Error error `json:"error"`
	}
}

// swagger:route PUT /v1/user/{id} User updateSwaggerUser
//
// Update User Account
//
// Updates a user account present in the backend database
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
//      200: operationResponse
// gets a user by account id
func (s *Server) updatedUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var (
		updateUserRequest UpdateUserRequest
		updateUserResponse  UpdateUserResponse
	)

	// attempt to first obtain user as a record should exist in the database
	id := helper.ExtractIDFromRequest(r)

	_, err := s.db.GetUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = helper.DecodeJSONBody(w, r, &updateUserRequest)
	if err != nil {
		helper.ProcessMalformedRequest(w, err)
		return
	}

	if updateUserRequest.User.Email == "" || updateUserRequest.User.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and email", http.StatusBadRequest)
		return
	}

	_, err = s.db.UpdateUser(r.Context(), &updateUserRequest.User)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updateUserResponse.Error = err
	s.JSONResponse(w, r, updateUserResponse)
}

func (s *Server) ExtractJwtFromRequest(w http.ResponseWriter, r *http.Request) (uint32, error) {
	authTokenRes, err := s.ExtractJwtFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	id, err := strconv.Atoi(authTokenRes.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}

	userId := uint32(id)
	return userId, err
}
