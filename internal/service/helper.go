package service

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

func floatToNumeric(f float64) pgtype.Numeric {
	var n pgtype.Numeric
	n.Scan(fmt.Sprintf("%f", f))
	return n
}