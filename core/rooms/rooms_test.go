package rooms

import (
	"bytes"
	"net/http"
	"testing"

	"encoding/json"
	"io/ioutil"
	"net/url"
	"strconv"
)

const (
	host = "http://localhost"
	port = ":3001"
)

type ResponseStatus struct {
	Status  string `json:"Status"`
	Message string `json:"Message"`
}

func TestAddNewRoom(t *testing.T) {

	roomTable := []struct {
		room_number string
		places      string
		status      string
		result      string
	}{
		{"998", "1", "1", "true"},
		{"998", "1", "1", "false"},
		{"999", "2", "1", "true"},
	}
	for _, row := range roomTable {
		data := url.Values{}
		data.Set("room_number", row.room_number)
		data.Add("places", row.places)
		data.Add("status", row.status)

		response := doRequest(t, "POST", "/room", data.Encode())
		result := ResponseStatus{}
		json.Unmarshal(response, &result)
		if result.Status != row.result {
			t.Error("Add room error")
		}
	}
}

func TestAddPeopleToRoom(t *testing.T) {

	roomTable := []struct {
		room_number string
		customers   string
		result      string
	}{
		{"998", `[{"visitor_name":"Ivan Petrovich"}]`, "true"},
		{"998", `[{"visitor_name":"Ivan Petrovich"}]`, "false"},
		{"999", `[{"visitor_name":"Milana Petrovich"}]`, "true"},
	}
	for _, row := range roomTable {
		data := url.Values{}
		data.Set("customers", row.customers)

		response := doRequest(t, "POST", "/room/"+row.room_number, data.Encode())
		result := ResponseStatus{}
		json.Unmarshal(response, &result)
		if result.Status != row.result {
			t.Error("Add visitor to room error")
		}
	}
}

func TestRemovePeopleFromRoom(t *testing.T) {

	room := struct {
		room_number string
		customers   string
		result      string
	}{
		"998", `[{"visitor_name":"Ivan Petrovich"}]`, "true",
	}
	data := url.Values{}
	data.Set("customers", room.customers)

	response := doRequest(t, "DELETE", "/room/"+room.room_number, data.Encode())
	result := ResponseStatus{}
	json.Unmarshal(response, &result)
	if result.Status != room.result {
		t.Error("Remove visitor from room error")
	}
}

func TestUpdatePeopleInRoom(t *testing.T) {

	room := struct {
		room_number string
		move_to     string
		result      string
	}{
		"999", "998", "true",
	}
	data := url.Values{}
	data.Set("move_to", room.move_to)

	response := doRequest(t, "PUT", "/room/"+room.room_number, data.Encode())
	result := ResponseStatus{}
	json.Unmarshal(response, &result)
	if result.Status != room.result {
		t.Error("Move visitor to room error")
	}
}

func TestGetFreeRooms(t *testing.T) {
	response, err := GetFreeRooms()

	if err != nil {
		t.Error(err)
	}

	if len(response) == 0 {
		t.Error("Something went wrong with getting free rooms")
	}
}

func TestGetBusyRooms(t *testing.T) {
	response, err := GetBusyRooms()

	if err != nil {
		t.Error(err)
	}

	if len(response) == 0 {
		t.Error("Something went wrong with getting busy rooms")
	}
}

func doRequest(t *testing.T, method string, path string, params string) []byte {

	req, err := http.NewRequest(method, host+port+path, bytes.NewBufferString(params))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(params)))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("Do request Error: ", err.Error())

	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Read Body Error: ", err.Error())
	}

	if resp.StatusCode != 200 {
		t.Errorf("Status code not 200")
	}

	return body
}
