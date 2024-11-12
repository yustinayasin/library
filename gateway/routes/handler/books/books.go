package handler

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	proto "shared/proto/books"

	"github.com/gin-gonic/gin"
)

func AddBook(bookClient proto.BookServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BookRequest

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

		res, err := bookClient.AddBook(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Add book failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Book})
	}
}

func EditBook(bookClient proto.BookServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BookRequest

		bookID, err := strconv.Atoi(c.Param("bookId"))

		if bookID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookID is required"})
			return
		}

		req.Id = int32(bookID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookClient.EditBook(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Edit failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Book})
	}
}

func DeleteBook(bookClient proto.BookServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BookIdRequest

		bookID, err := strconv.Atoi(c.Param("bookId"))

		if bookID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookID is required"})
			return
		}

		req.Id = int32(bookID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookClient.DeleteBook(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	}
}

func GetBook(bookClient proto.BookServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BookIdRequest

		bookID, err := strconv.Atoi(c.Param("bookId"))

		if bookID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookID is required"})
			return
		}

		req.Id = int32(bookID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookClient.GetBook(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Retrieve user failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res.Book})
	}
}
