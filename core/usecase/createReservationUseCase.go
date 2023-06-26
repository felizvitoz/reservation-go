package usecase

import "fmt"

type CreateReservationUseCaseRequest struct {
	Location        string
	RoomNo          string
	ReservationDate string
}

type CreateReservationUseCase struct {
	request CreateReservationUseCaseRequest
}

func (cru *CreateReservationUseCase) Execute() {
	fmt.Println("you ordered seatNumber : ", cru.request.Location, cru.request.RoomNo)
}
