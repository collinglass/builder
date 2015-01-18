package builder

import (
	"errors"
)

// HouseBuilder interface
//
// implements GetHouse which
// returns pointer to house associated with house builder
//
// implements BuildHouse which
// builds house associated with house builder
//
// implements BuildRoom which
// builds room with new roomId on house
//
// implements BuildDoor which
// builds door between two rooms
type HouseBuilder interface {
	GetHouse() *House
	BuildHouse()
	BuildRoom(roomId int)
	BuildDoor(doorId, roomId, roomId2 int) error
}

// Room interface
//
// implements SetRoomId which
// sets roomId to new int
//
// implements GetRoomId which
// returns current roomid
type Room interface {
	SetRoomId(roomId int)
	GetRoomId() int
}

// Door interface
//
// implements IsOpen which
// returns isOpen boolean of door
//
// implements SetOpen which
// sets isOpen boolean of door
//
// implements GetDoorId which
// returns doorId
type Door interface {
	IsOpen() bool
	SetOpen(isOpen bool)
	GetDoorId() int
}

// House struct
//
// contains rooms which
// is slice of Rooms in house
//
// contains doors which
// is slice of Doors in house
type House struct {
	rooms *[]Room
	doors *[]Door
}

// NewHouse
//
// returns new House
func NewHouse() *House {
	return &House{rooms: new([]Room), doors: new([]Door)}
}

// AddRoom
//
// appends Room to rooms slice
func (this *House) AddRoom(room Room) {
	*this.rooms = append(*this.rooms, room)
}

// GetRoom
//
// returns Room with roomId and nil error
//
// or
//
// returns nil Room with error "room is not here"
func (this *House) GetRoom(roomId int) (Room, error) {
	if len(*this.rooms) == 0 {
		return nil, errors.New("room is not here")
	}

	for _, room := range *this.rooms {
		if room != nil && room.GetRoomId() == roomId {
			return room, nil
		}
	}
	return nil, errors.New("room is not here")
}

// AddDoor
//
// appends Door to doors slice
func (this *House) AddDoor(door Door) {
	*this.doors = append(*this.doors, door)
}

// GetDoor
//
// returns Door with doorId and nil error
//
// or
//
// returns nil Door with error "door is not here"
func (this *House) GetDoor(doorId int) (Door, error) {
	for _, door := range *this.doors {
		if door.GetDoorId() == doorId {
			return door, nil
		}
	}
	return nil, errors.New("door is not here")
}

// BuildHouse
//
// builds house with concrete HouseBuild struct
// returns pointer to newly built House
func BuildHouse(builder HouseBuilder) *House {
	builder.BuildHouse()
	builder.BuildRoom(1)
	builder.BuildRoom(2)
	builder.BuildDoor(1, 1, 2)
	return builder.GetHouse()
}
