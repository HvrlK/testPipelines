package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	v1 "testPipelines/app/v1"
)

type Server struct {
	engine     *gin.Engine
	db         *gorm.DB
	Router     *v1.MainRouter
	Controller *v1.MainController
}

func (s *Server) Start() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	ConnectToInfura()

	_ = s.engine.Run(":" + port)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.engine.ServeHTTP(w, req)
}

func NewServer(db *gorm.DB) *Server {
	ginServer := gin.Default()
	controller := v1.NewMainController(ginServer, db)
	router := v1.NewMainRouter(controller, ginServer)
	return &Server{ginServer, db, router, controller}
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	setTestDatabaseData(db)
	return db
}