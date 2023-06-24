package usecase

type InputBoundary interface {
	Execute()
}

func GetCreateReservationUseCase(request CreateReservationUseCaseRequest) InputBoundary {
	return &CreateReservationUseCase{request}
}

func GetSetRoomDataUseCase(request SetRoomDataRequest) InputBoundary {
	return &SetRoomDataUseCase{request}
}
