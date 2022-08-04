// Package classification BuddyAPI.
//
// Documentation of BuddyAPI.
//
//     Schemes: http
//     BasePath: /
//     Version: 1.0.0
//     Host: http://localhost:8000
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - key: []
//
//    SecurityDefinitions:
//    key:
//      type: apiKey
//      in: header
//      name: jwt-token
//
// swagger:meta
package docs

import (
	pb_article "github.com/maslow123/api-gateway/pkg/articles/pb"
	article_routes "github.com/maslow123/api-gateway/pkg/articles/routes"
	"github.com/maslow123/api-gateway/pkg/users/pb"
	user_routes "github.com/maslow123/api-gateway/pkg/users/routes"
)

// swagger:route POST /users/create Users RegisterEndpoint
// Register for create new user or publisher
// responses:
//   201: registerResponse

// This text will appear as description of your response body.
// swagger:response registerResponse
type registerResponseWrapper struct {
	// in:body
	RegisterResponse pb.GenericResponse
}

// swagger:parameters RegisterEndpoint
type registerParamsWrapper struct {
	// request body for register or create new user / publisher
	// in:body
	// required: true
	RegisterRequest pb.RegisterRequest
}

// swagger:route POST /users/login Users LoginEndpoint
// For login user company / publisher
// responses:
//   200: loginResponse

// This is response from login
// swagger:response loginResponse
type loginResponseWrapper struct {
	// in:body
	LoginResponse pb.LoginResponse
}

// swagger:parameters LoginEndpoint
type loginParamsWrapper struct {
	// request login for authentication
	// in:body
	// required: true
	LoginRequest pb.LoginRequest
}

/*      WITH AUTH       */

// swagger:route GET /users/list Users ShowRegisterUserEndpoint
// For show all user / publisher
// responses:
//   200: ShowRegisterUserResponse

// List registered user
// swagger:response ShowRegisterUserResponse
type showRegisterUserResponseWrapper struct {
	// in:body
	ShowRegisterUserResponse pb.ShowUserResponse
}

// swagger:parameters ShowRegisterUserEndpoint
type showRegisterUserParamsWrapper struct {
	// in: query
	// required: true
	// example: 1
	Limit int32
	// required: true
	// example: 10
	Page int32
}

// swagger:route GET /users/point Users ShowUserPoint
// For show all user / publisher point
// responses:
//   200: ShowAllUserPoint

// List all user point
// swagger:response ShowAllUserPoint
type showAllUserPointResponseWrapper struct {
	// in:body
	showAllUserPointResponse pb.ShowUserResponse
}

// swagger:parameters ShowUserPoint
type showAllUserPointParamsWrapper struct {
	// in: query
	// required: true
	// example: 1
	Limit int32
	// required: true
	// example: 10
	Page int32
}

// swagger:route PUT /users/set-point Users SetArticlePoint
// For set article point based on ArticleID
// responses:
//   200: SetArticlePoint

// Response for set article point
// swagger:response SetArticlePoint
type setArticlePointResponseWrapper struct {
	// in:body
	SetArticlePointResponse pb.GenericResponse
}

// swagger:parameters SetArticlePoint
type setArticlePointParamsWrapper struct {
	// in:body
	// required: true
	SetArticlePointPointRequest user_routes.SetArticlePointRequestBody
}

/* ARTICLES */
// swagger:route POST /articles/create Articles CreateArticleEndpoint
// For create new article
// responses:
//   200: CreateArticleEndpoint

// Response for create article
// swagger:response CreateArticleEndpoint
type createArticleResponseWrapper struct {
	// in:body
	CreateArticleResponse pb_article.GenericResponse
}

// swagger:parameters CreateArticleEndpoint
type createArticleParamsWrapper struct {
	// in:body
	// required: true
	CreateArticleRequest article_routes.CreateArticleRequestBody
}

// swagger:route GET /articles/detail/{article_id} Articles DetailArticleEndpoint
// For get detail article
// responses:
//   200: DetailArticleEndpoint

// Response for detail article
// swagger:response DetailArticleEndpoint
type detailArticleResponseWrapper struct {
	// in:body
	DetailArticleResponse pb_article.GenericResponse
}

// swagger:parameters DetailArticleEndpoint
type detailArticleParamsWrapper struct {
	// in:path
	// required: true
	ArticleID int32 `json:"article_id"`
}

// swagger:route GET /articles/point Articles GetUserArticlePointEndpoint
// For get total user point based on user id (get from token)
// responses:
//   200: GetUserArticlePointEndpoint

// Response for get user article point
// swagger:response GetUserArticlePointEndpoint
type getUserArticlePointResponseWrapper struct {
	// in:body
	GetUserArticlePointResponse pb_article.UserArticlePointResponse
}

// swagger:parameters GetUserArticlePointEndpoint
type getUserArticlePointParamsWrapper struct {
}
