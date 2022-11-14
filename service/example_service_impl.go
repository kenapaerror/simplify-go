package service

import (
	"base-project/exception"
	"base-project/helper"
	"base-project/model/entity"
	"base-project/model/web"
	"base-project/repository"
	"base-project/utils"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type ExampleServiceImpl struct {
	ExampleRepository repository.ExampleRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewExampleServiceImpl(exampleRepository repository.ExampleRepository, DB *sql.DB, validate *validator.Validate) ExampleService {
	return &ExampleServiceImpl{
		ExampleRepository: exampleRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *ExampleServiceImpl) Create(ctx context.Context, request web.ExampleCreateRequest) web.ExampleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	example := entity.Example{}

	example = service.ExampleRepository.Create(ctx, tx, example)

	return utils.ToExampleResponse(example)
}

func (service *ExampleServiceImpl) Update(ctx context.Context, request web.ExampleUpdateRequest) web.ExampleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	example, err := service.ExampleRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	example.Name = "example"
	example.UpdatedAt = utils.CurrentMillis()

	example = service.ExampleRepository.Update(ctx, tx, example)

	return utils.ToExampleResponse(example)
}

func (service *ExampleServiceImpl) Delete(ctx context.Context, exampleId string) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	example, err := service.ExampleRepository.FindById(ctx, tx, exampleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.ExampleRepository.Delete(ctx, tx, example)
}

func (service *ExampleServiceImpl) FindById(ctx context.Context, exampleId string) web.ExampleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	example, err := service.ExampleRepository.FindById(ctx, tx, exampleId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return utils.ToExampleResponse(example)
}

func (service *ExampleServiceImpl) FindAll(ctx context.Context) []web.ExampleResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var examples []web.ExampleResponse
	for _, example := range examples {
		examples = append(examples, example)
	}

	return examples
}
