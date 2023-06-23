package menus

func GetMenu(themeType int32) []Menu {
	switch themeType {
	case 101:
		return SpecialMenuStrategy()
	default:
		return DefaultMenuStrategy()
	}
}

func DefaultMenuStrategy() []Menu {
	var mainMenus []Menu
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

	setRoomDataMenu := &CliMenu{
		MenuId: "99",
		Level:  1,
		Order:  99,
		Text:   "Set Room Data",
	}
	setRoomLocationDataMenuContent := &InputContent{"", "Please Input The Location : ", InputLocationKey, Content{}}
	setRoomNumberDataMenuContent := &InputContent{"", "Please Input The Room Number : ", InputRoomNumberKey, Content{}}
	SetRoomDataAction := BuildSetRoomDataAction([]*InputContent{setRoomLocationDataMenuContent, setRoomNumberDataMenuContent})
	setRoomActionContent := &ActionContent{SetRoomDataAction, Content{}}
	setRoomDataMenu.contents = []MenuContent{setRoomLocationDataMenuContent, setRoomNumberDataMenuContent, setRoomActionContent}

	mainMenus = []Menu{reservationMenu, setRoomDataMenu}
	reservationMenu.nextMenus = mainMenus
	setRoomDataMenu.nextMenus = mainMenus

	return mainMenus
}

func SpecialMenuStrategy() []Menu {
	return nil
}
