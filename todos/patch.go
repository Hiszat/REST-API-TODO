package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UpdateRequest struct {
	Title       string `json: "title"`
	Description string `json: "description"`
}

func TodosUpdateHandler(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request UpdateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		res, err := db.Exec("UPDATE todos SET title = ?, description = ? WHERE id = ?", request.Title, request.Description, id)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}
		if rowsAffected == 0 {
			return ctx.String(http.StatusNotFound, "Data not found")
		}

		return ctx.String(http.StatusOK, "OK")
	}
}
