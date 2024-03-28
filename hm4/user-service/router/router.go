package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-service/db"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:id", getUser)
	r.POST("/users", postUser)
	r.PUT("/users/:id", putUser)
	r.DELETE("/users/:id", deleteUser)
	return r
}

func getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func postUser(ctx *gin.Context) {
	var movie db.User
	err := ctx.Bind(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateUser(&movie)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": res,
	})
}

func putUser(ctx *gin.Context) {
	var updatedUser db.User
	err := ctx.Bind(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbUser, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbUser.UserName = updatedUser.UserName
	dbUser.FirstName = updatedUser.FirstName
	dbUser.LastName = updatedUser.LastName
	dbUser.Phone = updatedUser.Phone

	res, err := db.UpdateUser(dbUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
