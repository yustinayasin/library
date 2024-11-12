package borrowrecords

import (
	"errors"
	borrowrecords "userservice/business/borrowrecords"

	"gorm.io/gorm"
)

type BorrowRecordRepository struct {
	Db *gorm.DB
}

func NewBorrowRecordRepository(database *gorm.DB) borrowrecords.BorrowRecordRepoInterface {
	return &BorrowRecordRepository{
		Db: database,
	}
}

func (repo *BorrowRecordRepository) AddBorrowRecord(borrowRecord borrowrecords.BorrowRecord) (borrowrecords.BorrowRecord, error) {
	borrowRecordDB := FromUsecase(borrowRecord)

	result := repo.Db.Create(&borrowRecordDB)

	if result.Error != nil {
		return borrowrecords.BorrowRecord{}, result.Error
	}

	return borrowRecordDB.ToUsecase(), nil
}

func (repo *BorrowRecordRepository) EditBorrowRecord(borrowRecord borrowrecords.BorrowRecord, id int) (borrowrecords.BorrowRecord, error) {
	borrowRecordDb := FromUsecase(borrowRecord)

	var newBorrowRecord BorrowRecord

	result := repo.Db.Preload("Role").First(&newBorrowRecord, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return borrowrecords.BorrowRecord{}, errors.New("BorrowRecord not found")
		}
		return borrowrecords.BorrowRecord{}, errors.New("error in database")
	}

	newBorrowRecord.UserId = borrowRecordDb.UserId
	newBorrowRecord.BookId = borrowRecordDb.BookId
	newBorrowRecord.BorrowDate = borrowRecordDb.BorrowDate
	newBorrowRecord.ReturnDate = borrowRecordDb.ReturnDate
	newBorrowRecord.IsReturn = borrowRecordDb.IsReturn

	repo.Db.Save(&newBorrowRecord)
	return newBorrowRecord.ToUsecase(), nil
}

func (repo *BorrowRecordRepository) DeleteBorrowRecord(id int) (borrowrecords.BorrowRecord, error) {
	var borrowRecordDb BorrowRecord

	resultFind := repo.Db.First(&borrowRecordDb, id)

	if resultFind.Error != nil {
		return borrowrecords.BorrowRecord{}, errors.New("borrowRecord not found")
	}

	result := repo.Db.Delete(&borrowRecordDb, id)

	if result.Error != nil {
		return borrowrecords.BorrowRecord{}, errors.New("Delete failed")
	}

	return borrowRecordDb.ToUsecase(), nil
}

func (repo *BorrowRecordRepository) GetBorrowRecord(id int) (borrowrecords.BorrowRecord, error) {
	var borrowRecordDb BorrowRecord

	resultFind := repo.Db.First(&borrowRecordDb, id)

	if resultFind.Error != nil {
		return borrowrecords.BorrowRecord{}, errors.New("borrowRecord not found")
	}

	return borrowRecordDb.ToUsecase(), nil
}
