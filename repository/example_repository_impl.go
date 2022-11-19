package repository

import (
	"context"
	"database/sql"
	"errors"
	"simplify-go/helper"
	"simplify-go/model/entity"
)

type ExampleRepositoryImpl struct{}

func NewExampleRepositoryImpl() ExampleRepository {
	return &ExampleRepositoryImpl{}
}

func (repository *ExampleRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example {
	SQL := "INSERT INTO example (id, name, email, created_at, updated_at) VALUES (?,?,?,?,?)"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		example.Id,
		example.Name,
		example.Email,
		example.CreatedAt,
		example.UpdatedAt,
	)
	helper.PanicIfError(err)

	return example
}

func (repository *ExampleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, example entity.Example) entity.Example {
	SQL := "UPDATE example SET name=?, updated_at=? WHERE id=?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		example.Name,
		example.UpdatedAt,
		example.Id,
	)
	helper.PanicIfError(err)

	return example
}

func (repository *ExampleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, example entity.Example) {
	SQL := "DELETE FROM example WHERE id = ?"

	_, err := tx.ExecContext(
		ctx,
		SQL,
		example.Id,
	)
	helper.PanicIfError(err)
}

func (repository *ExampleRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, exampleId string) (entity.Example, error) {
	SQL := "SELECT id, name, email, created_at, updated_at FROM example WHERE id=?"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
		exampleId,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	example := entity.Example{}
	if rows.Next() {
		err := rows.Scan(
			&example.Id,
			&example.Name,
			&example.Email,
			&example.CreatedAt,
			&example.UpdatedAt,
		)
		helper.PanicIfError(err)

		return example, nil
	}

	return example, errors.New("example not found")
}

func (repository *ExampleRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.Example {
	SQL := "SELECT id, name, email, created_at, updated_at FROM example"

	rows, err := tx.QueryContext(
		ctx,
		SQL,
	)
	helper.PanicIfError(err)
	defer rows.Close()

	examples := []entity.Example{}
	for rows.Next() {
		example := entity.Example{}
		err := rows.Scan(
			&example.Id,
			&example.Name,
			&example.Email,
			&example.CreatedAt,
			&example.UpdatedAt,
		)
		helper.PanicIfError(err)
		examples = append(examples, example)
	}

	return examples
}
