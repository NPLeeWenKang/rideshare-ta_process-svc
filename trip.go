package main

import (
	"database/sql"
)

type Trip struct {
	Trip_Id      int          `json:"trip_id"`
	Passanger_Id int          `json:"passanger_id"`
	Pick_Up      string       `json:"pick_up"`
	Drop_Off     string       `json:"drop_off"`
	Start        sql.NullTime `json:"start"`
	End          sql.NullTime `json:"end"`
}
