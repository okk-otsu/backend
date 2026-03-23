package simplesql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/k0kubun/pp/v3"
)

func SelectRows(
	ctx context.Context,
	conn *pgx.Conn,
) ([]TaskModel, error) {

	sqlQuery := /* sql */ `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return []TaskModel{}, err
	}
	defer rows.Close()

	tasks := make([]TaskModel, 0)
	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)

		if err != nil {
			return []TaskModel{}, err
		}

		tasks = append(tasks, task)
		printTask(task)
	}

	return tasks, nil
}

func printTask(tasks TaskModel) {
	fmt.Println("----------------------------------")
	pp.Println(tasks)

}
