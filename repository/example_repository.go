package repository

import (
	"base-project/model/entity"
	"context"
	"database/sql"
)

type ExampleRepository interface {
	Create(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example
	Update(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example
	Delete(ctx context.Context, tx *sql.Tx, example entity.Example)
	FindById(ctx context.Context, tx *sql.Tx, exampleId string) (entity.Example, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Example
}
