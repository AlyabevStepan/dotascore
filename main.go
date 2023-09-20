package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"github.com/AlyabevStepan/dotascore/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

type Server struct {
	httpServer *http.Server
}

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "postgres"
	port_server= "8080"
)

func main(){
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env :%s", err.Error())
	}


  r := gin.Default()

  r.PUT("/put_game", handler.put_game )
  r.PUT("/put_games", handler.put_games)
  r.GET("/get_players",handler.get_players)
  r.Run(":8080")
  
  
  srv := new(Server)
	if err := srv.Run(port_server, r); err != nil {
		fmt.Printf("this is handlerbag: %s", err.Error())
	}
	
}


func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		MaxHeaderBytes: 1 << 20,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func TestPingRoute(t *testing.T) {
	router := handler.setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}