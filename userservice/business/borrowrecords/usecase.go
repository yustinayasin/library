package borrowrecords

import (
	"errors"
	"time"
)

type BorrowRecordUseCase struct {
	Repo BorrowRecordRepoInterface
}

func NewUseCase(borrowRecordRepo BorrowRecordRepoInterface) BorrowRecordUseCaseInterface {
	return &BorrowRecordUseCase{
		Repo: borrowRecordRepo,
	}
}

func (borrowRecordUseCase *BorrowRecordUseCase) AddBorrowRecord(borrowRecord BorrowRecord) (BorrowRecord, error) {
	if borrowRecord.UserId == 0 {
		return BorrowRecord{}, errors.New("name cannot be empty")
	}

	if borrowRecord.BookId == 0 {
		return BorrowRecord{}, errors.New("email cannot be empty")
	}

	if borrowRecord.BorrowDate == (time.Time{}) {
		return BorrowRecord{}, errors.New("borrow date cannot be empty")
	}

	borrowRecord.IsReturn = false

	borrowRecordRepo, err := borrowRecordUseCase.Repo.AddBorrowRecord(borrowRecord)

	if err != nil {
		return BorrowRecord{}, err
	}

	return borrowRecordRepo, nil
}

func (borrowRecordUseCase *BorrowRecordUseCase) EditBorrowRecord(borrowRecord BorrowRecord, id int) (BorrowRecord, error) {
	if id == 0 {
		return BorrowRecord{}, errors.New("borrowRecord ID cannot be empty")
	}

	if borrowRecord.UserId == 0 {
		return BorrowRecord{}, errors.New("name cannot be empty")
	}

	if borrowRecord.BookId == 0 {
		return BorrowRecord{}, errors.New("email cannot be empty")
	}

	borrowRecordRepo, err := borrowRecordUseCase.Repo.EditBorrowRecord(borrowRecord, id)

	if err != nil {
		return BorrowRecord{}, err
	}

	return borrowRecordRepo, nil
}

func (borrowRecordUseCase *BorrowRecordUseCase) DeleteBorrowRecord(id int) (BorrowRecord, error) {
	if id == 0 {
		return BorrowRecord{}, errors.New("borrowRecord ID cannot be empty")
	}

	borrowRecordRepo, err := borrowRecordUseCase.Repo.DeleteBorrowRecord(id)

	if err != nil {
		return BorrowRecord{}, err
	}

	return borrowRecordRepo, nil
}

func (borrowRecordUseCase *BorrowRecordUseCase) GetBorrowRecord(id int) (BorrowRecord, error) {
	if id == 0 {
		return BorrowRecord{}, errors.New("borrowRecord ID cannot be empty")
	}

	borrowRecordRepo, err := borrowRecordUseCase.Repo.GetBorrowRecord(id)

	if err != nil {
		return BorrowRecord{}, err
	}

	return borrowRecordRepo, nil
}
