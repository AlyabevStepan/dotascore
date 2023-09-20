package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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



func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.PUT("/put_game", put_game )
	r.PUT("/put_games", put_games)
	r.GET("/get_players",get_players)
	r.Run(":8080")
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


func TestPingRoute(t *testing.T) {
	router := SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}