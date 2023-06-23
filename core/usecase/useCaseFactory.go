package usecase

type InputBoundary interface {
	Execute()
}

func GetCreateReservationUseCase(requestBuilder CreateReservationUseCaseRequestBuilder) InputBoundary {
	return &CreateReservationUseCase{requestBuilder}
}
