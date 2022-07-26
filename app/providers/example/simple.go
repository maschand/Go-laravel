package provider_test

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		isError,
	}
}

func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("Failed to create Service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}

}
