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
	ID         int       `json:"reserva_id"`
	Nome       string    `json:"usuario_nome"`
	Sobrenome  string    `json:"usuario_sobrenome"`
	Unidade    string    `json:"unidade"`
	Tipo       string    `json:"tipo"`
	HoraInicio NullTime2 `json:"hora_inicio"`
}

// Marshal2JSON for NullTime
func (nt *NullTime2) Marshal2JSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}
