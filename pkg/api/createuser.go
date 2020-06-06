package api

import (
	"net/http"
	"gitlab.com/yoanyombapro/CubeMicroservices/podinfo/pkg/utils/requestutils"

	"github.com/LensPlatform/BlackSpace/pkg/helper"
	"github.com/LensPlatform/BlackSpace/pkg/models"
)

type SignUpUserRequest struct {
	User models.User `json:"user"`
}

type SignUpUserResponse struct {
	Error error  `json:"error"`
	Id    uint32 `json:"id"`
	JwtToken string `json:"jwt"`
}

// Sign up user request
// swagger:parameters signUpUserReq
type signUpUserRequest struct {
	// in: body
	Body struct {
		// user account to create
		// required: true
		User models.User `json:"user"`
	}
}

// User Successfully signed up
// swagger:response signUpResp
type signUpResponse struct {
	// in: body
	Body struct {
		// user account id
		// required: true
		// example: 20
		Id uint32 `json:"id"`
		// error
		// required: true
		// example: user already exists
		Error error `json:"error"`
		// JWT Token
		// required: true
		// example: askhdkjashjd.sajkhdjashjfdsa.askjdhjkashdja
		JwtToken string `json:"jwt"`
	}
}

// Error occured during request lifecycle
// swagger:response genericError
type genericErrorResponseSwagger struct {
	// in: body
	Body struct {
		// required: true
		Error error `json:"error"`
	}
}

// swagger:route POST /v1/user User signUpUserReq
//
// Create User Account
//
// creates a user account object in the backend database
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
//
//     Security:
//       api_key:
//       oauth: read, write
// responses:
//      200: signUpResp
// 400: genericError
// 404: genericError
// 403: genericError
// 406: genericError
// 401: genericError
// 500: genericError
// creates a user account
func (s *Server) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var (
		signUpUserRequest SignUpUserRequest
		userAccount       models.User
	)

	// extract the authn_id from the jwt token in the request header
	/*
	authTokenRes, err := s.ExtractJwtFromHeader(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	*/

	err:= helper.DecodeJSONBody(w, r, &signUpUserRequest)
	if err != nil {
		requestutils.ProcessMalformedRequest(w, err)
		return
	}

	userAccount = signUpUserRequest.User
	if userAccount.Email == "" || userAccount.Username == "" {
		http.Error(w, "invalid input parameters. please specify a username and email", http.StatusBadRequest)
		return
	}

	createdUserAccount, err := s.db.CreateUser(r.Context(), &userAccount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := s.GenerateAndSignJwtToken(createdUserAccount.Id, createdUserAccount)
	if err != nil {
		s.ErrorResponse(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	response := SignUpUserResponse{Id: createdUserAccount.Id, Error: err, JwtToken: t}
	s.JSONResponse(w, r, response)
}


