package utils

import (
	"base-project/model/entity"
	"base-project/model/web"
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
