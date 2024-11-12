package handler

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	proto "shared/proto/users"

	"github.com/gin-gonic/gin"
)

func AddBorrowRecords(bookStocksClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsRequest

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

		res, err := bookStocksClient.AddBorrowRecords(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Add book failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Borrowrecords})
	}
}

func EditBorrowRecords(bookStocksClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsRequest

		bookStocksID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.EditBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Edit failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Borrowrecords})
	}
}

func DeleteBorrowRecords(bookStocksClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsIdRequest

		bookStocksID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.DeleteBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	}
}

func GetBorrowRecords(bookStocksClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsIdRequest

		bookStocksID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if bookStocksID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bookStocksID is required"})
			return
		}

		req.Id = int32(bookStocksID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := bookStocksClient.GetBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Retrieve user failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res.Borrowrecords})
	}
}
