package models

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	mysql.NullTime
}

//Reserva is an exportable type
type Reserva struct {
	ID         int      `json:"reserva_id"`
	Usuario    int      `json:"usuario"`
	Spot       int      `json:"spot"`
	HoraInicio NullTime `json:"hora_inicio"`
	HoraFim    NullTime `json:"hora_fim"`
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}
