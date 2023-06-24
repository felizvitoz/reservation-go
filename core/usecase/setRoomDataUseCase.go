package usecase

import "fmt"

type SetRoomDataRequest struct {
	Location   string
	RoomNumber string
}

type SetRoomDataUseCase struct {
	request SetRoomDataRequest
}

func (cru *SetRoomDataUseCase) Execute() {
	location := cru.request.Location
	roomNo := cru.request.RoomNumber
	fmt.Println("Room Data : ", location, roomNo)
}
