package usecase

import "fmt"

type CreateReservationUseCaseRequestBuilder interface {
	Build() CreateReservationUseCaseRequest
}

type CreateReservationUseCaseRequest struct {
	SeatNumber string
}

type CreateReservationUseCase struct {
	RequestBuilder CreateReservationUseCaseRequestBuilder
}

func (cru *CreateReservationUseCase) Execute() {
	createReservationRequest := cru.RequestBuilder.Build()
	seatNumber := createReservationRequest.SeatNumber
	fmt.Println("you ordered seatNumber : ", seatNumber)
}
