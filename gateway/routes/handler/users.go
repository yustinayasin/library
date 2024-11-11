package handler

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"shared/proto"

	"github.com/gin-gonic/gin"
)

func SignupHandler(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UserEditRequest

		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Reading body failed"})
			return
		}

		// Reassign the body to the context (so it can be parsed again)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Binding failed"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer cancel()

		res, err := userClient.SignUp(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Signup failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.User})
	}
}

func LoginHandler(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UserLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Binding failed"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := userClient.Login(ctx, &req)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.User})
	}
}

func EditUserHandler(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UserEditRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Binding failed"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := userClient.EditUser(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Edit failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.User})
	}
}

func DeleteUserHandler(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UserIdRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Binding failed"})
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := userClient.DeleteUser(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	}
}

func GetUserHandler(userClient proto.UserServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.UserIdRequest

		userID, err := strconv.Atoi(c.Param("userId"))

		if userID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "userID is required"})
			return
		}

		req.Id = int32(userID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := userClient.GetUser(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Retrieve user failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res.User})
	}
}
