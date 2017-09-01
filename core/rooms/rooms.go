package rooms

import (
	"net/http"
	"strconv"

	"encoding/json"

	"errors"

	"github.com/bohdanlisovskyi/hotel-managing/core/database"
	"github.com/bohdanlisovskyi/hotel-managing/core/loger"
	"github.com/gorilla/mux"
)

type People struct {
	VisitorName string `json:"visitor_name"`
}

type Room struct {
	RoomNumber string `gorm:"primary_key"`
	Places     int
	Status     int
}

type BusyRoomList struct {
	RoomNumber  string `json:"room_number"`
	VisitorName string `json:"visitor_name"`
}

// add new room to hotel
func AddNewRoom(r *http.Request) error {
	room_number := r.FormValue("room_number")

	if room_number == "" {
		return errors.New("Empty room number parameter")
	}

	places := r.FormValue("places")

	if places == "" {
		return errors.New("Empty places parameter")
	}

	placesInt, err := strconv.Atoi(places)
	status := r.FormValue("status")

	if status == "" {
		return errors.New("Empty status parameter")
	}

	statusInt, err := strconv.Atoi(status)

	if err != nil {
		return err
	}

	return database.InsertNewRoom(room_number, placesInt, statusInt)
}

// return room number from url path variable
func roomNumber(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["room_number"]
}

// add visitor to room (new inhabitants)
func AddPeopleToRoom(r *http.Request) error {
	db, err := database.GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
		return err
	}

	room_number := roomNumber(r)
	room, err := database.CheckIfRoomIsBusy(db, room_number)

	if err != nil {
		return err
	}

	if room == true {
		return errors.New("The room is busy")
	}

	customers := r.FormValue("customers")

	if customers == "" {
		return errors.New("Empty customers parameter")
	}

	visitors := []People{}
	json.Unmarshal([]byte(customers), &visitors)
	if err != nil {
		loger.Log.Errorf("Failed to open db %s", err)
		return err
	}

	for _, visitor := range visitors {

		err := database.AddVisitorToRoom(db, visitor.VisitorName, room_number)
		if err != nil {
			return err
		}
	}

	return database.UpdateRoomStatus(db, 0, room_number)
}

// remove visitors from room (f.e when inhabitants leave)
func RemovePeopleFromRoom(room_number string) error {
	db, err := database.GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
		return err
	}

	err = database.RemoveVisitor(room_number)

	err = database.UpdateRoomStatus(db, 1, room_number)

	if err != nil {
		return err
	}
	return nil
}

// move Visitor from living room to another room
func UpdatePeopleInRoom(room_number string, move_to string) error {
	db, err := database.GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
		return err
	}

	stmt, err := db.Prepare("update customers set room_number=? where room_number=?")

	if err != nil {
		loger.Log.Errorf("Failed to update customers in room: %s", err)
		return err
	}

	_, err = stmt.Exec(move_to, room_number)

	if err != nil {
		loger.Log.Errorf("Failed to update customers in room: %s", err)
		return err
	}

	err = database.UpdateRoomStatus(db, 1, room_number)

	if err != nil {
		return err
	}

	err = database.UpdateRoomStatus(db, 0, move_to)

	if err != nil {
		return err
	}
	return nil
}

// return free rooms list
func GetFreeRooms() (list []Room, err error) {

	db, err := database.GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed get storage %s", err)
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM rooms WHERE status=1")

	if err != nil {
		loger.Log.Errorf("Failed to select rooms %s", err)
		return nil, err
	}

	for rows.Next() {
		var uid int
		var room_number string
		var places int
		var status int
		err = rows.Scan(&uid, &room_number, &places, &status)
		if err != nil {
			loger.Log.Errorf("Failed to scan room %s", err)
		}
		list = append(list, Room{
			RoomNumber: room_number,
			Places:     places,
			Status:     status,
		})
	}
	return list, err
}

// return busy room list with visitors
func GetBusyRooms() (list []BusyRoomList, err error) {
	db, err := database.GetStorage()

	if err != nil {
		loger.Log.Errorf("Failed to insert room %s", err)
		return nil, err
	}

	rows, err := db.Query("SELECT rooms.room_number,customers.visitor_name FROM rooms LEFT JOIN customers ON customers.room_number = rooms.room_number WHERE status=0")

	for rows.Next() {
		var room_number string
		var visitor_name string
		err = rows.Scan(&room_number, &visitor_name)

		if err != nil {
			loger.Log.Errorf("Failed to insert room %s", err)
		}

		list = append(list, BusyRoomList{
			RoomNumber:  room_number,
			VisitorName: visitor_name,
		})
	}
	return list, err
}
