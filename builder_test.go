package builder

import (
	"testing"
)

func TestBuildHouse(t *testing.T) {
	hb := NewDefaultHouseBuilder()

	BuildHouse(hb)
	house := hb.GetHouse()

	if house == nil {
		t.Errorf("TestBuildHouse Failed: %s\n", "house is nil")
	}
}

func TestGetDoor(t *testing.T) {
	hb := NewDefaultHouseBuilder()

	BuildHouse(hb)
	house := hb.GetHouse()

	door, err := house.GetDoor(1)

	if err != nil {
		t.Errorf("TestGetDoor Failed: %s\n", "door is nil")
	}
	if door == nil {
		t.Errorf("TestGetDoor Failed: %s\n", "door is nil")
	}
}

func TestBuildDoor(t *testing.T) {
	hb := NewDefaultHouseBuilder()

	BuildHouse(hb)
	house := hb.GetHouse()

	err := hb.BuildDoor(2, 1, 2)
	door, err := house.GetDoor(2)

	if err != nil {
		t.Errorf("TestBuildDoor Failed: %s\n", "door is nil")
	}
	if door == nil {
		t.Errorf("TestBuildDoor Failed: %s\n", "door is nil")
	}
}

func TestGetRoom(t *testing.T) {
	hb := NewDefaultHouseBuilder()

	BuildHouse(hb)
	house := hb.GetHouse()

	room, err := house.GetRoom(1)

	if err != nil {
		t.Errorf("TestGetRoom Failed: %s\n", "room is nil")
	}
	if room == nil {
		t.Errorf("TestGetRoom Failed: %s\n", "room is nil")
	}
}

func TestBuildRoom(t *testing.T) {
	hb := NewDefaultHouseBuilder()

	BuildHouse(hb)
	house := hb.GetHouse()

	hb.BuildRoom(3)
	room, err := house.GetRoom(1)

	if err != nil {
		t.Errorf("TestBuildRoom Failed: %s\n", "room is nil")
	}
	if room == nil {
		t.Errorf("TestBuildRoom Failed: %s\n", "room is nil")
	}
}
