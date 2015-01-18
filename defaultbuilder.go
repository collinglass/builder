package builder

//
//
//
// Default Room
//
//
//
type DefaultRoom struct {
	roomId int
}

func NewDefaultRoom(roomId int) *DefaultRoom {
	return &DefaultRoom{roomId: roomId}
}

func (this *DefaultRoom) SetRoomId(roomId int) {
	this.roomId = roomId
}

func (this *DefaultRoom) GetRoomId() int {
	return this.roomId
}

//
//
//
// Default Door
//
//
//

type DefaultDoor struct {
	doorId int
	room1  Room
	room2  Room
	isOpen bool
}

func NewDefaultDoor(doorId int, room1, room2 Room) *DefaultDoor {
	return &DefaultDoor{doorId: doorId, room1: room1, room2: room2}
}

func (this *DefaultDoor) GetDoorId() int {
	return this.doorId
}

func (this *DefaultDoor) IsOpen() bool {
	return this.isOpen
}

func (this *DefaultDoor) SetOpen(isOpen bool) {
	this.isOpen = isOpen
}

//
//
//
// Default House Builder
//
//
//
type DefaultHouseBuilder struct {
	house *House
}

func NewDefaultHouseBuilder() *DefaultHouseBuilder {
	return &DefaultHouseBuilder{}
}

func (this *DefaultHouseBuilder) GetHouse() *House {
	return this.house
}

func (this *DefaultHouseBuilder) BuildHouse() {
	this.house = NewHouse()
}

func (this *DefaultHouseBuilder) BuildRoom(roomId int) {
	room, err := this.house.GetRoom(roomId)
	if err != nil && room == nil {
		room := NewDefaultRoom(roomId)
		this.house.AddRoom(room)
	}
}

func (this *DefaultHouseBuilder) BuildDoor(doorId, roomId1, roomId2 int) error {
	room1, err := this.house.GetRoom(roomId1)
	if err != nil {
		return err
	}

	room2, err := this.house.GetRoom(roomId2)
	if err != nil {
		return err
	}

	door := NewDefaultDoor(doorId, room1, room2)
	this.house.AddDoor(door)

	return nil
}
