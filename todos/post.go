package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateRequest struct {
	Title       string `json: "title"`
	Description string `json: "description"`
}

func CreateTodosHandler(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		_, err := db.Exec(
			"INSERT INTO todos (title, description, done) VALUES (?, ?, 0)",
			request.Title,
			request.Description,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}
		return ctx.String(http.StatusOK, "OK")
	}
}
