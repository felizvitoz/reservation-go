package usecase

import "fmt"

type SetRoomDataUseCaseRequestBuilder interface {
	Build() SetRoomDataRequest
}

type SetRoomDataRequest struct {
	Location   string
	RoomNumber string
}

type SetRoomDataUseCase struct {
	RequestBuilder SetRoomDataUseCaseRequestBuilder
}

func (cru *SetRoomDataUseCase) Execute() {
	setRoomDataRequest := cru.RequestBuilder.Build()
	location := setRoomDataRequest.Location
	roomNo := setRoomDataRequest.RoomNumber
	fmt.Println("Room Data : ", location, roomNo)
}
