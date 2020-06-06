package api

import (
	"net/http"

	"github.com/LensPlatform/BlackSpace/pkg/helper"
)

type OperationResponse struct {
	Error error `json:"error"`
}

// Delete user by id request
// swagger:parameters deleteUser
type DeleteUserByIdRequestSwagger struct {
	// id of the user account to delete
	// in: query
	// required: true
	Id uint32 `json:"result"`
}

// swagger:route DELETE /v1/user/{id} User deleteUser
// Delete User Account
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
// deletes a user by account id
func (s *Server) deleteUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	var response OperationResponse

	id := helper.ExtractIDFromRequest(r)

	// Ensure the user account actually exists
	exist, _, err := s.db.GetUserIfExists(r.Context(),id, "", "")
	if !exist {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.db.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response.Error = err
	s.JSONResponse(w, r, response)
}
