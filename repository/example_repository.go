package repository

import (
	"context"
	"database/sql"
	"simplify-go/model/entity"
)

type ExampleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example
	Update(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example
	Delete(ctx context.Context, tx *sql.Tx, example entity.Example)
	FindById(ctx context.Context, tx *sql.Tx, exampleId string) (entity.Example, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Example
}
