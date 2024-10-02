package main

import (
	// net "net/http"
	// "os/user"

	handlers "go-rest-api/pkg/handlers"
	// model "go-rest-api/pkg/model"

	gin "github.com/gin-gonic/gin"
)

//  Write an REST server in Golang that has endpoints to do created, updated, deleted operations. For example,
// GET /v1/user --> Gets all user
// POST /v1/user --> Creates an user
// GET /v1/user/id --> Get an user with ID
// DELETE /v1/user/id --> Deleted an user with ID
// PUT /v1/user/id --> Updates an user with ID

// So you can have a simple user schema as:
// {
// "name": "<some-name>"
// "id": "<some-id>"
// "location": "some-city"
// }
// And then have mysql store the data via the golang rest server
 
 func main(){
	path := gin.Default()
	
	// Main page
	path.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome to the homepage",
		})
	})
	// Fetch all User
	path.GET("/v1/user", handlers.GetUsers)
	// path.GET("/v1/user", func(ctx *gin.Context) {
	// 	// user := ctx.Param("user")
	// 	users := handlers.GetUsers
	// 	// ctx.JSON(200, gin.H{
	// 	// 	"message": "You fetched all users",
	// 	// 	"users": users,
	// 	// })
	// })
	// Add a user
	path.POST("/v1/:user", func(ctx *gin.Context) {
		user := ctx.Param("user")
		ctx.JSON(200, gin.H{
			"message": "new user " + user + " created",
		})
	})
	// Delete a User 
	path.DELETE("/v1/:user", func(ctx *gin.Context) {
		user := ctx.Param("user")
		ctx.JSON(200, gin.H{
			"message": "User " + user + " deleted",
		})
	})

	path.Run(":8081") 

 }