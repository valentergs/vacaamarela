package models

import (
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullTime2 is an alias for mysql.NullTime data type
type NullTime2 struct {
	mysql.NullTime
}

//ReservaJoin is an exportable type
type ReservaJoin struct {
	ID         int       `json:"reserva"`
	Nome       string    `json:"nome"`
	Sobrenome  string    `json:"sobrenome"`
	Unidade    string    `json:"unidade"`
	Spot       int       `json:"spot"`
	Tipo       string    `json:"tipo"`
	HoraInicio NullTime2 `json:"hora_inicio"`
	HoraFim    NullTime2 `json:"hora_fim"`
}

// MarshalJSON for NullTime
func (nt *NullTime2) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}
