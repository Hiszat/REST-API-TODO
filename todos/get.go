package todos

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type TodoResponse struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	Done        bool   `json: "done"`
}

func GetTodosHandler(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		rows, err := db.Query("SELECT title, description, done FROM todos")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}

		var res []TodoResponse

		for rows.Next() {
			var title string
			var description string
			var done int
			err := rows.Scan(&title, &description, &done)

			if err != nil {
				return ctx.String(http.StatusInternalServerError, "Internal Error")
			}

			var todo TodoResponse
			todo.Title = title
			todo.Description = description

			if done == 1 {
				todo.Done = true
			} else {
				todo.Done = false
			}

			res = append(res, todo)
		}
		return ctx.JSON(http.StatusOK, res)
	}
}
