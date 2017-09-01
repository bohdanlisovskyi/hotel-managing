package database

import (
	"github.com/bohdanlisovskyi/hotel-managing/core/loger"
	"database/sql"
)

// insert room to db
func InsertNewRoom(room_number string, places int, status int) error {
	db, err := GetStorage()

	if err != nil {
		loger.Log.Panicf("Failed to open db %s", err)
	}

	stmt, err := db.Prepare("INSERT INTO rooms(room_number, places, status) values(?,?,?)")

	_, err = stmt.Exec(room_number, places, status)

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
	}
	return err
}

// check if room busy or not
func CheckIfRoomIsBusy(db *sql.DB, room_number string) (bool, error) {
	rows, err := db.Query("SELECT status FROM rooms WHERE room_number=$1", room_number)
	busy := false

	if err != nil {
		return busy, err
	}

	for rows.Next() {
		var status int
		err = rows.Scan(&status)

		if err != nil {
			loger.Log.Errorf("Failed to select room %s", err)
			return busy, err
		}

		if status == 0 {
			busy = true
		}
	}
	return busy, nil
}

// add visitor to room
func AddVisitorToRoom(db *sql.DB, visitor_name string, room_number string) error {
	stmt, err := db.Prepare("INSERT INTO customers(visitor_name, room_number) values(?,?)")

	if err != nil {
		loger.Log.Errorf("Insert visitor error: %s", err)
		return err
	}

	_, err = stmt.Exec(visitor_name, room_number)

	if err != nil {
		loger.Log.Errorf("Failed to execute %s", err)
		return err
	}
	return nil
}

// update room status to 0 if busy or 1 if free
func UpdateRoomStatus(db *sql.DB, status int, room_number string) error {
	stmt, err := db.Prepare("update rooms set status=? where room_number=?")

	if err != nil {
		loger.Log.Errorf("Update room status error: %s", err)
		return err
	}

	_, err = stmt.Exec(status, room_number)

	if err != nil {
		loger.Log.Errorf("Update room status error: %s", err)
		return err
	}
	return nil
}

// remove visitor from room
func RemoveVisitor(room_number string) error {
	db, err := GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
		return err
	}

	stmt, err := db.Prepare("delete from customers where room_number=?")

	if err != nil {
		loger.Log.Errorf("Failed to delete from room %s", err)
		return err
	}

	_, err = stmt.Exec(room_number)

	if err != nil {
		loger.Log.Errorf("Failed to remove from room %s", err)
		return err
	}
	return nil
}
