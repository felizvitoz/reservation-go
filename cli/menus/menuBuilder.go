package menus

func GetMenu(clientId int32) []Menu {
	switch clientId {
	case 10001:
		return SpecialMenuStrategy()
	default:
		return DefaultMenuStrategy()
	}
}

func DefaultMenuStrategy() []Menu {
	reservationMenu := &CliMenu{
		MenuId: "1",
		Level:  1,
		Order:  1,
		Text:   "Create Reservation",
	}

	createReservationInputContent := &InputContent{"", "Please Input Seat Number : ", "seatNumber", Content{}}
	createReservationAction := BuildReservationAction([]*InputContent{createReservationInputContent})
	createReservationActionContent := &ActionContent{createReservationAction, Content{}}
	reservationMenu.contents = []MenuContent{createReservationInputContent, createReservationActionContent}
	reservationMenu.nextMenus = []Menu{reservationMenu}

	return []Menu{reservationMenu}
}

func SpecialMenuStrategy() []Menu {
	return nil
}
