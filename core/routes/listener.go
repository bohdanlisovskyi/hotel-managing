package router

import "github.com/bohdanlisovskyi/hotel-managing/core/handlers"

// write which route you'd like to listen
var routes = Routes{
	Route{
		"Free rooms",
		"GET",
		"/rooms/free",
		handlers.FreeRooms,
	},
	Route{
		"Busy rooms",
		"GET",
		"/rooms/busy",
		handlers.BusyRooms,
	},
	Route{
		"Add new room",
		"POST",
		"/room",
		handlers.AddRoom,
	},
	Route{
		"Add people to room",
		"POST",
		"/room/{room_number}",
		handlers.AddToRoom,
	},
	Route{
		"Update people in room",
		"PUT",
		"/room/{room_number}",
		handlers.UpdatePeopleInRoom,
	},
	Route{
		"Remove people from room",
		"DELETE",
		"/room/{room_number}",
		handlers.RemovePeopleFromRoom,
	},
}
