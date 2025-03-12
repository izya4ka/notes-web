package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/izya4ka/notes-web/notes-service/handlers"
	"github.com/izya4ka/notes-web/notes-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	db    *gorm.DB
	srv   *http.Server
	mutex sync.Mutex
}

func (s *Server) InitDB(db_url string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	// Establish a connection to the PostgreSQL database using the provided DB_URL environment variable.
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically migrate the UserPostgres and Note models to the database.
	db.AutoMigrate(&models.UserPostgres{}, &models.Note{})
	log.Print("DB connection success!")
	s.db = db
}

func (s *Server) InitGIN(port string) {
	s.mutex.Lock()

	isDebug := os.Getenv("DEBUG") != ""

	if isDebug {
		gin.SetMode("debug")
	} else {
		gin.SetMode("release")
	}

	router := gin.Default()

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.GET("/notes", func(c *gin.Context) {
		handlers.GetNotes(c, s.db)
	})

	v1.POST("/notes", func(c *gin.Context) {
		handlers.PostNotes(c, s.db)
	})

	v1.GET("/notes/:id", func(c *gin.Context) {
		handlers.GetNote(c, s.db)
	})

	v1.PUT("/notes/:id", func(c *gin.Context) {
		handlers.PutNode(c, s.db)
	})

	v1.DELETE("/notes/:id", func(c *gin.Context) {
		handlers.DeleteNote(c, s.db)
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router.Handler(),
	}
	s.srv = srv
	s.mutex.Unlock()
	srv.ListenAndServe()
}

func (s *Server) Shutdown() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Starting graceful shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Gin stopping error: %v", err)
	}
	log.Println("Gin stopped!")

	sqldb, serr := s.db.DB()
	if serr == nil {
		if err := sqldb.Close(); err != nil {
			log.Fatalf("Database stopping error: %v", err)
		}
	} else {
		log.Fatalf("Database stopping error: %v", serr)
	}
	log.Println("Database connection stopped!")
	os.Exit(0)
}
