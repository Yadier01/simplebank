package api

import (
	"fmt"

	db "github.com/Yadier01/simplebank/db/sqlc"
	"github.com/Yadier01/simplebank/token"
	"github.com/Yadier01/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
	}
	//validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurency)
	}

	//routes
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts/", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.POST("/transfer", server.createTransfer)
	server.router = router
}
func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
