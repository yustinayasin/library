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

func AddBooksStocks(bookStocksClient proto.BooksStocksServiceClient, bookClient proto.BookServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BooksStocksRequest

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

		// check if book exist
		bookReq := &proto.BookIdRequest{
			Id: req.BookId,
		}
		bookRes, err := bookClient.GetBookExist(context.Background(), bookReq)
		if err != nil || !bookRes.Exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		res, err := bookStocksClient.AddBooksStocks(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Add book failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Booksstocks})
	}
}

func EditBooksStocks(bookStocksClient proto.BooksStocksServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BooksStocksRequest

		bookStocksID, err := strconv.Atoi(c.Param("bookStockId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.EditBooksStocks(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Edit failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Booksstocks})
	}
}

func DeleteBooksStocks(bookStocksClient proto.BooksStocksServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BooksStocksIdRequest

		bookStocksID, err := strconv.Atoi(c.Param("bookStockId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.DeleteBooksStocks(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	}
}

func GetBooksStocks(bookStocksClient proto.BooksStocksServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BooksStocksIdRequest

		bookStocksID, err := strconv.Atoi(c.Param("bookStockId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.GetBooksStocks(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Retrieve user failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res.Booksstocks})
	}
}
