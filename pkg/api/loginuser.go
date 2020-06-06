package api

import (
	"net/http"

	"github.com/LensPlatform/BlackSpace/pkg/helper"
)

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JwtToken string `json:"token"`
	Error    error  `json:"error"`
}

// Log in user
// swagger:parameters loginUserRequest
type loginUserRequest struct {
	// in: body
	Body struct {
		// user name
		// required: true
		// example: test-username
		Username string `json:"username"`
		// password
		// required: true
		// example: test-password
		Password string `json:"password"`
	}
}

// swagger:response loginUserResponse
type loginUserResponse struct {
	// in: body
	Body struct {
		// Jwt Token
		// required: true
		// example: kBxbjzKVDjvasgvds.askdhjaskjdgsagjcdgc.asjdjkasfgdas
		JwtToken string `json:"token"`
		// error
		// required: true
		// example: unable to get token
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/user/login User loginUserRequest
//
// Log in user
//
// Logs in a user into the system
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
//      200: loginUserResponse
// logs in a team into the system
func (s *Server) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var (
		loginUserRequest LoginUserRequest
	)

	// decode the data present in the body
	err := helper.DecodeJSONBody(w, r, &loginUserRequest)
	if err != nil {
		helper.ProcessMalformedRequest(w, err)
		return
	}

	if loginUserRequest.Password == "" || loginUserRequest.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and password", http.StatusBadRequest)
		return
	}

	_, user, err := s.db.GetUserIfExists(r.Context(), 0, loginUserRequest.Username, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !s.db.ComparePasswords(user.Password, []byte(loginUserRequest.Password)){
		http.Error(w, "invalid password", http.StatusBadRequest)
	}

	token, err := s.GenerateAndSignJwtToken(user.Id, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	loginResponse := LoginResponse{
		JwtToken: token,
		Error:    nil,
	}

	s.JSONResponse(w, r, loginResponse)
}
