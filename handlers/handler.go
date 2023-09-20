package handler

import (
	"os"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "postgres"
	port_server= "8080"
)



func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}


func put_game(c *gin.Context){
	password := os.Getenv("DB_PASSWORD");
	psqlconn := fmt.Sprintf("host= %s port=%s user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)
	db, err := sql.Open("postgres",psqlconn)
	if err!=nil{
		fmt.Println(err)	
	}
	defer db.Close()
	insertDynStmt := `insert into test (name, empid) values('1', '2')`+`,('44','44');`
	_, e := db.Exec(insertDynStmt)
	if e!=nil{
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, gin.H{})
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
	insertDynStmt := `insert into test (name, empid) values('1', '2');`
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
	insertDynStmt := `insert into test (name, empid) values('1', '2');`
	_, e := db.Exec(insertDynStmt)
	if e!=nil{
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, gin.H{})
}