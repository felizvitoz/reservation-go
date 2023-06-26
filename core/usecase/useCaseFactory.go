package usecase

import "reservation/persistence/repositories"

type InputBoundary interface {
	Execute()
}

var roomRepository = &repositories.RoomRepository{make([]repositories.Room, 0)}

func GetCreateReservationUseCase(request CreateReservationUseCaseRequest) InputBoundary {
	return &CreateReservationUseCase{request}
}

func GetSetRoomDataUseCase(request SetRoomDataRequest) InputBoundary {
	return &SetRoomDataUseCase{request, roomRepository}
}
