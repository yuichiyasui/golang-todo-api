package repository

import (
	"api/model"
	"context"
	"database/sql"
)

func GetTasks(ctx context.Context, db *sql.DB) (model.TaskSlice, error) {
	tasks, err := model.Tasks().All(ctx, db)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
