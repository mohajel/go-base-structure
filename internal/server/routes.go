package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	_ "base-1/docs"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	// r.GET("/", s.HelloWorldHandler)

	// r.GET("/health", s.healthHandler)

	r.GET("/oik", GetStructArrayByString)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// @Description get struct array by ID
// @Accept  json
// @Produce  json
// @Param   some_id     path    string     true        "Some ID"
// @Param   offset     query    int     true        "Offset"
// @Param   limit      query    int     true        "Limit"
// @Success 200 {string} string	"ok"
// @Router /oik [get]
func GetStructArrayByString(c *gin.Context) {
}

// // @Summary Add a new pet to the store
// // @Description get string by ID
// // @Accept  json
// // @Produce  jsonv2/testapi/get-struct-array-by-string/:some_id
// // @Success 200 {string} string	"ok"
// // @Failure 400 {object} map[string]string "We need ID!!"
// // @Failure 404 {object} map[string]string"Can not find ID"
// // @Router / [get]

// func (s *Server) HelloWorldHandler(c *gin.Context) {
// 	resp := make(map[string]string)
// 	resp["message"] = "Hello World"

// 	c.JSON(http.StatusOK, resp)
// }

// func (s *Server) healthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, s.db.Health())
// }
