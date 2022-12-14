package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

var cfg = mysql.Config{
	User:      "user",
	Passwd:    "password",
	Net:       "tcp",
	Addr:      "localhost:3306",
	DBName:    "db",
	ParseTime: true,
}

func main() {
	db, _ = sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()

	for {
		fmt.Println("Assigning.....")
		trips, err := getUnassignedTrips()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, element := range trips {
			err := assignTrips(element.Trip_Id)
			if err != nil {
				fmt.Println("Error assigning trip_id", element.Trip_Id)
			}
		}
		time.Sleep(8 * time.Second)
	}
}

func getUnassignedTrips() ([]Trip, error) {
	tList := make([]Trip, 0)
	var rows *sql.Rows
	var err error

	rows, err = db.Query("WITH success_assignment AS ( SELECT ta.trip_id FROM trip_assignment ta WHERE ta.status IN ('riding', 'pending', 'accepted', 'riding', 'done') ) SELECT t.* FROM trip t WHERE t.trip_id NOT IN (SELECT * FROM success_assignment)")

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Trip
		if err := rows.Scan(&t.Trip_Id, &t.Passanger_Id, &t.Pick_Up, &t.Drop_Off, &t.Start, &t.End); err != nil {
			return nil, err
		}
		tList = append(tList, t)
	}
	return tList, nil
}

func assignTrips(tripId int) error {
	var err error

	_, err = db.Query("INSERT INTO trip_assignment(trip_id, driver_id, status) WITH reject_assignment AS ( SELECT ta.driver_id FROM trip_assignment ta WHERE ta.status = 'rejected' AND ta.trip_id = ? ), busy_drivers AS ( SELECT DISTINCT(ta.driver_id) FROM trip_assignment ta WHERE ta.status IN ('ACCEPTED','DRIVING','PENDING') ), random_available_driver AS ( SELECT * FROM driver d WHERE d.driver_id NOT IN (SELECT * FROM reject_assignment) AND d.driver_id NOT IN (SELECT * FROM busy_drivers) AND is_available = TRUE ORDER BY RAND() LIMIT 1 ) SELECT ?, (SELECT driver_id FROM random_available_driver LIMIT 1), ?;", tripId, tripId, TripStatus.PENDING)
	return err
}
