package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
	"context"
	"time"
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
  
  r.PUT("/put_game", put_game )
  r.PUT("/put_games", put_games)
  r.GET("/get_players",get_players)
  
  r.Run()
  srv := new(Server)
	if err := srv.Run(port_server, r); err != nil {
		fmt.Printf("this is handlerbag: %s", err.Error())
	}
}


func put_game(c *gin.Context){
	password := os.Getenv("DB_PASSWORD");
	psqlconn := fmt.Sprintf("host= %s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",psqlconn)
	if err!=nil{
		fmt.Println(err)	
	}
	defer db.Close()
	insertDynStmt := `insert into test (name, empid) values('1', '2')`
	_, e := db.Exec(insertDynStmt)
	if e!=nil{
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, gin.H{})
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



func put_games(c *gin.Context){
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env :%s", err.Error())
	}
	password := os.Getenv("DB_PASSWORD");
	psqlconn := fmt.Sprintf("host= %s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",psqlconn)
	if err!=nil{
		fmt.Println(err)	
	}
	defer db.Close()
	insertDynStmt := `insert into test (name, empid) values('1', '2')`
	_, e := db.Exec(insertDynStmt)
	if e!=nil{
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, gin.H{})
}


func get_players(c *gin.Context){
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error loading env :%s", err.Error())
	}
	password := os.Getenv("DB_PASSWORD");
	psqlconn := fmt.Sprintf("host= %s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",psqlconn)
	if err!=nil{
		fmt.Println(err)	
	}
	defer db.Close()
	insertDynStmt := `insert into test (name, empid) values('1', '2')`
	_, e := db.Exec(insertDynStmt)
	if e!=nil{
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, gin.H{})
}
