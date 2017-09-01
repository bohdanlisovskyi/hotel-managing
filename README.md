# hotel-management

This repo have hotel management code!

### Set Up Instructions
**1.Clone Repository**

```bash
$   git clone https://github.com/bohdanlisovskyi/hotel-managing.git
```

**2.Run to install glide and gometalinter**

```bash
$   make install-helpers
```

**3.Run to install application dependencies**

```bash
$   make dependencies
```    

### REST API

#### 1. Add new room

* Method `POST`

* Path `/room`
    
Request Parameters:

```
room_number={room number},
places={how many place in room},
status={status 1 -free, 0 -busy}
```
    
Response:

```
{
    "Status":"true",
    "Message":"Add new room"
}
```
```
{
    "Status":"false",
    "Message":"Error message"
}
```
    
#### 2. Add Visitor to room

* Method `POST`

* Path `/room/{room_number}`

Request Parameters:

```
customer=
[
    {
        "visitor_name":"Ivan Petrovich"
    },
    {
        "visitor_name":"Milana Petrovich"
    }
]
```    
Response:

```
{
    "Status":"true",
    "Message":"Add people to room"
}
```
```
{
    "Status":"false",
    "Message":"Error message"
}
```
#### 3. Remove Visitor from room

* Method `DELETE`

* Path `/room/{room_number}`
Response:

```
{
    "Status":"true",
    "Message":"Remove people from room: 005"
}
```
```
{
    "Status":"false",
    "Message":"Error message"
}
```        
#### 4. Move visitor from room to another room

* Method `PUT`

* Path `/room/{room_number}`

Request Parameters:

```
move_to={new room number}
```    
Response:

```
{
    "Status":"true",
    "Message":"Move people from room 888to room 005"
}
```
```
{
    "Status":"false",
    "Message":"Error message"
}
```    
#### 5. Get free rooms list

* Method `GET`

* Path `/rooms/free`

Response:

```
[
    {
        "RoomNumber":"001",
        "Places":2,"Status":1
    },
    {
        "RoomNumber":"003",
        "Places":1,"Status":1
    },
    ...
]

```
```
{
    "Status":"false",
    "Message":"Error message"
}
```   
    
#### 6. Get busy rooms list

* Method `GET`

* Path `/rooms/busy`

Response:

```
[
    {
        "room_number":"001",
        "visitor_name":"Ivan Petrovich"
    },
    {
        "room_number":"001",
        "visitor_name":"Milana Petrovich
    },
    ...
]

```
```
{
    "Status":"false",
    "Message":"Error message"
}
```   
    
    