package menus

import (
	"reservation/core/usecase"
)

type CreateReservationRequestBuilderImpl struct {
	inputContents []*InputContent
}

type SetRoomDataUseCaseRequestBuilderImpl struct {
	inputContents []*InputContent
}

func BuildReservationAction(inputContents []*InputContent) usecase.InputBoundary {
	requestBuilder := &CreateReservationRequestBuilderImpl{inputContents}
	return usecase.GetCreateReservationUseCase(requestBuilder)
}

func (crrb *CreateReservationRequestBuilderImpl) Build() usecase.CreateReservationUseCaseRequest {
	valueMap := convertInputContentIntoMap(crrb.inputContents)
	return usecase.CreateReservationUseCaseRequest{valueMap["seatNumber"]}
}

func BuildSetRoomDataAction(inputContents []*InputContent) usecase.InputBoundary {
	requestBuilder := &SetRoomDataUseCaseRequestBuilderImpl{inputContents}
	return usecase.GetSetRoomDataUseCase(requestBuilder)
}

func (crrb *SetRoomDataUseCaseRequestBuilderImpl) Build() usecase.SetRoomDataRequest {
	valueMap := convertInputContentIntoMap(crrb.inputContents)
	return usecase.SetRoomDataRequest{valueMap[InputLocationKey], valueMap[InputRoomNumberKey]}
}

func convertInputContentIntoMap(inputContents []*InputContent) map[string]string {
	valueMap := make(map[string]string)
	for _, inputContent := range inputContents {
		valueMap[inputContent.Key] = inputContent.Value
	}
	return valueMap
}
