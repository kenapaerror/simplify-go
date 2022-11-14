package service

import (
	"context"
	"simplify-go/model/web"
)

type ExampleService interface {
	Create(ctx context.Context, request web.ExampleCreateRequest) web.ExampleResponse
	Update(ctx context.Context, request web.ExampleUpdateRequest) web.ExampleResponse
	Delete(ctx context.Context, exampleId string)
	FindById(ctx context.Context, exampleId string) web.ExampleResponse
	FindAll(ctx context.Context) []web.ExampleResponse
}
