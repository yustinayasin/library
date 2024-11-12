package authors

import (
	"errors"
	proto "shared/proto/authors"
	"time"
)

type AuthorUseCase struct {
	Repo AuthorRepoInterface
}

func NewUseCase(authorRepo AuthorRepoInterface) AuthorUseCaseInterface {
	return &AuthorUseCase{
		Repo: authorRepo,
	}
}

func (authorUseCase *AuthorUseCase) AddAuthor(author Author) (Author, error) {
	if author.Name == "" {
		return Author{}, errors.New("Name cannot be empty")
	}

	if author.Bio == "" {
		return Author{}, errors.New("Bio cannot be empty")
	}

	if author.Birthdate == (time.Time{}) {
		return Author{}, errors.New("Birthdate year cannot be empty")
	}

	authorRepo, err := authorUseCase.Repo.AddAuthor(author)

	if err != nil {
		return Author{}, err
	}

	return authorRepo, nil
}

func (authorUseCase *AuthorUseCase) EditAuthor(author Author, id int) (Author, error) {
	if id == 0 {
		return Author{}, errors.New("author ID cannot be empty")
	}

	if author.Name == "" {
		return Author{}, errors.New("Name cannot be empty")
	}

	if author.Bio == "" {
		return Author{}, errors.New("Bio cannot be empty")
	}

	if author.Birthdate == (time.Time{}) {
		return Author{}, errors.New("Birthdate year cannot be empty")
	}

	authorRepo, err := authorUseCase.Repo.EditAuthor(author, id)

	if err != nil {
		return Author{}, err
	}

	return authorRepo, nil
}

func (authorUseCase *AuthorUseCase) DeleteAuthor(id int) (Author, error) {
	if id == 0 {
		return Author{}, errors.New("Author ID cannot be empty")
	}

	authorRepo, err := authorUseCase.Repo.DeleteAuthor(id)

	if err != nil {
		return Author{}, err
	}

	return authorRepo, nil
}

func (authorUseCase *AuthorUseCase) GetAuthor(id int) (Author, error) {
	if id == 0 {
		return Author{}, errors.New("Author ID cannot be empty")
	}

	authorRepo, err := authorUseCase.Repo.GetAuthor(id)

	if err != nil {
		return Author{}, err
	}

	return authorRepo, nil
}

func (authorUseCase *AuthorUseCase) GetAuthorExist(id int) (proto.AuthorResponseExist, error) {
	if id == 0 {
		return proto.AuthorResponseExist{Exists: false}, errors.New("Author ID cannot be empty")
	}

	_, err := authorUseCase.Repo.GetAuthor(id)

	if err != nil {
		return proto.AuthorResponseExist{Exists: false}, err
	}

	return proto.AuthorResponseExist{Exists: true}, nil
}
