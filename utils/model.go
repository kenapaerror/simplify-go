package utils

import (
	"simplify-go/model/entity"
	"simplify-go/model/web"
)

func ToExampleResponse(example entity.Example) web.ExampleResponse {
	return web.ExampleResponse{
		Id:        example.Id,
		Name:      example.Name,
		Email:     example.Email,
		CreatedAt: example.CreatedAt,
		UpdatedAt: example.UpdatedAt,
	}
}
