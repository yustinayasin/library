package handler

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	protoBook "shared/proto/books"
	proto "shared/proto/users"

	"github.com/gin-gonic/gin"
)

func AddBorrowRecords(borrowRecordClient proto.BorrowRecordsServiceClient, bookClient protoBook.BookServiceClient) gin.HandlerFunc {
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

		// check if book exist
		bookReq := &protoBook.BookIdRequest{
			Id: req.BookId,
		}
		bookRes, err := bookClient.GetBookExist(context.Background(), bookReq)
		if err != nil || !bookRes.Exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		res, err := borrowRecordClient.AddBorrowRecords(ctx, &req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Add book failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Borrowrecords})
	}
}

func EditBorrowRecords(borrowRecordClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsRequest

		borrowRecordID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if borrowRecordID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "borrowRecordID is required"})
			return
		}

		req.Id = int32(borrowRecordID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := borrowRecordClient.EditBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Edit failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message, "user": res.Borrowrecords})
	}
}

func DeleteBorrowRecords(borrowRecordClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsIdRequest

		borrowRecordID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if borrowRecordID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "borrowRecordID is required"})
			return
		}

		req.Id = int32(borrowRecordID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := borrowRecordClient.DeleteBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Delete failed"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": res.Message})
	}
}

func GetBorrowRecords(borrowRecordClient proto.BorrowRecordsServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req proto.BorrowRecordsIdRequest

		borrowRecordID, err := strconv.Atoi(c.Param("borrowRecordId"))

		if borrowRecordID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "borrowRecordID is required"})
			return
		}

		req.Id = int32(borrowRecordID)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		res, err := borrowRecordClient.GetBorrowRecords(ctx, &req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Retrieve user failed"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": res.Borrowrecords})
	}
}
