package menus

import (
	"reservation/core/usecase"
)

func BuildReservationAction(inputMap map[string]string) usecase.InputBoundary {
	request := usecase.CreateReservationUseCaseRequest{
		inputMap[InputLocationKey], inputMap[InputRoomNumberKey]}
	return usecase.GetCreateReservationUseCase(request)
}

func BuildSetRoomDataAction(inputMap map[string]string) usecase.InputBoundary {
	request := usecase.SetRoomDataRequest{
		inputMap[InputLocationKey], inputMap[InputRoomNumberKey]}
	return usecase.GetSetRoomDataUseCase(request)
}
