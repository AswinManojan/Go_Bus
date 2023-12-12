package server

import "github.com/gin-gonic/gin"

// Serverstruct struct is used to intialize the gin Engine and other related methods
type Serverstruct struct {
	R *gin.Engine
}

// StartServer is used to start the Server
func (s *Serverstruct) StartServer() {
	s.R.LoadHTMLGlob("templates/*.html")
	s.R.Run(":8080")
}

// NewServer is used to create a initialize and connect to a Server
func NewServer() *Serverstruct {
	router := gin.Default()
	return &Serverstruct{
		R: router,
	}
}
