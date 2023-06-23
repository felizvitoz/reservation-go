package menus

import (
	"reservation/core/usecase"
)

func BuildReservationAction(inputContents []*InputContent) usecase.InputBoundary {
	requestBuilder := &CreateReservationRequestBuilderImpl{inputContents}
	return usecase.GetCreateReservationUseCase(requestBuilder)
}

type CreateReservationRequestBuilderImpl struct {
	inputContents []*InputContent
}

func (crrb *CreateReservationRequestBuilderImpl) Build() usecase.CreateReservationUseCaseRequest {
	valueMap := make(map[string]string)
	for _, inputContent := range crrb.inputContents {
		valueMap[inputContent.Key] = inputContent.Value
	}
	return usecase.CreateReservationUseCaseRequest{valueMap["seatNumber"]}
}
