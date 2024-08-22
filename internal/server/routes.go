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

	r.GET("/", s.HelloWorldHandler)

	// r.GET("/health", s.healthHandler)

	r.GET("/ok", GetStructArrayByString)

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
// @Router /ok [get]
func GetStructArrayByString(c *gin.Context) {
}

type HelloWorldResp struct {
	Data string `json:"data"`
}

// @Summary HelloWorldHandler Summary
// @Description HelloWorldHandler Description
// @Accept  json
// @Success 200 {object} HelloWorldResp
// @Router / [get]
func (s *Server) HelloWorldHandler(c *gin.Context) {
	// resp := make(map[string]string)
	// resp["message"] = "Hello World"
	resp := HelloWorldResp{Data: "Hello World"}

	c.JSON(http.StatusOK, resp)
}

// func (s *Server) healthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, s.db.Health())
// }
