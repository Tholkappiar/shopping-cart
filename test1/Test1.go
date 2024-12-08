package test1

import (
	"fmt"
	"gin-test/inits"
	"gin-test/migrate"
	"net/http"

	"github.com/gin-gonic/gin"
)


func Test1(c *gin.Context) {
	fmt.Println("Connecting to DB")
	inits.ConnectToDB()
	c.JSON(http.StatusOK, gin.H{"message": "DB connected, from test 1"})
	migrate.Migrate()
	fmt.Println("migrated")
}
