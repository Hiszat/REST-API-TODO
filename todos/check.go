package todos

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckRequest struct {
	Done bool `json: "done"`
}

func CheckUncheckHandler(db *sql.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request CheckRequest
		json.NewDecoder(ctx.Request().Body).Decode(&request)

		var doneInt int

		if request.Done {
			doneInt = 1
		}

		res, err := db.Exec("UPDATE todos SET done = ? WHERE id = ?", doneInt, id)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}

		rowsAffected, err := res.RowsAffected()
		if err != nil {
			return ctx.String(http.StatusInternalServerError, "Internal Error")
		}

		if rowsAffected == 0 {
			return ctx.String(http.StatusInternalServerError, "Not Found")
		}

		return ctx.String(http.StatusOK, "OK")
	}
}
