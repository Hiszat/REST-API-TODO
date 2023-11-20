package todos

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteTodosHandler(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		res, err := db.Exec("DELETE FROM todos WHERE id= ?", id)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}
		if rowsAffected == 0 {
			return ctx.String(http.StatusNotFound, "Not Found")
		}

		return ctx.String(http.StatusOK, "OK")
	}
}
