package simplesql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn) error {

	sqlQuery := /* sql */ `
	UPDATE tasks
	SET description = ':)'
	WHERE completed = FALSE;
	`

	_, err := conn.Exec(ctx, sqlQuery)
	return err
}

func UpdateTask(
	ctx context.Context,
	conn *pgx.Conn,
	task TaskModel,
) error {
	sqlQuery := /* sql */ `
	UPDATE tasks
	SET title=$1, description=$2, completed=$3, created_at=$3, completed_at=$4 
	WHERE id =$6;
	`
	_, err := conn.Exec(
		ctx,
		sqlQuery,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.CompletedAt,
		task.ID,
	)

	return err
}
