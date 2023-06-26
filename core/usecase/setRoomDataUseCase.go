package usecase

import (
	"reservation/persistence/repositories"
)

type SetRoomDataRequest struct {
	Location   string
	RoomNumber string
}

type SetRoomDataUseCase struct {
	request        SetRoomDataRequest
	roomRepository *repositories.RoomRepository
}

func (cru *SetRoomDataUseCase) Execute() {
	location := cru.request.Location
	roomNo := cru.request.RoomNumber
	room := repositories.Room{location, roomNo}
	cru.roomRepository.Save(room)
}
