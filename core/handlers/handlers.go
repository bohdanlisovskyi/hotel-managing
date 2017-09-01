package handlers

import (
	"net/http"

	"github.com/bohdanlisovskyi/hotel-managing/core/response"
	"github.com/bohdanlisovskyi/hotel-managing/core/rooms"
	"github.com/gorilla/mux"
)

type ResponseStatus struct {
	Status  string
	Message string
}

// response free room list
func FreeRooms(w http.ResponseWriter, r *http.Request) {
	list, err := rooms.GetFreeRooms()

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, list)
}

// response busy room list
func BusyRooms(w http.ResponseWriter, r *http.Request) {
	list, err := rooms.GetBusyRooms()

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, list)
}

// add new room to hotel
func AddRoom(w http.ResponseWriter, r *http.Request) {
	err := rooms.AddNewRoom(r)

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, ResponseStatus{
		Status:  "true",
		Message: "Add new room",
	})
}

// add visitor to room
func AddToRoom(w http.ResponseWriter, r *http.Request) {
	err := rooms.AddPeopleToRoom(r)

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, ResponseStatus{
		Status:  "true",
		Message: "Add people to room ",
	})
}

// move on visitor from room to another room
func UpdatePeopleInRoom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	room_number := vars["room_number"]
	move_to := r.FormValue("move_to")
	err := rooms.UpdatePeopleInRoom(room_number, move_to)

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, ResponseStatus{
		Status:  "true",
		Message: "Move people from room " + room_number + " to room " + move_to,
	})
}

// remove visitor from room
func RemovePeopleFromRoom(w http.ResponseWriter, r *http.Request) {
	room_number := mux.Vars(r)["room_number"]
	err := rooms.RemovePeopleFromRoom(room_number)

	if err != nil {
		response.New(w, ResponseStatus{
			Status:  "false",
			Message: err.Error(),
		})
		return
	}
	response.New(w, ResponseStatus{
		Status:  "true",
		Message: "Remove people from room: " + room_number,
	})
}
