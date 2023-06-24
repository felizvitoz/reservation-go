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
		MenuId:            "1",
		Level:             1,
		Text:              "Create Reservation",
		sharedInputValues: make(map[string]string),
	}
	createReservationLocationInputContent := &InputContent{"Please Input The Location : ", InputLocationKey, Content{reservationMenu.setContent}}
	createReservationRoomInputContent := &InputContent{"Please Input The Room Number : ", InputRoomNumberKey, Content{reservationMenu.setContent}}
	createReservationActionContent := &ActionContent{BuildReservationAction, Content{reservationMenu.setContent}}
	reservationMenu.contents = []MenuContent{createReservationLocationInputContent, createReservationRoomInputContent, createReservationActionContent}

	setRoomDataMenu := &CliMenu{
		MenuId:            "99",
		Level:             1,
		Text:              "Set Room Data",
		sharedInputValues: make(map[string]string),
	}
	setRoomLocationDataMenuContent := &InputContent{"Please Input The Location : ", InputLocationKey, Content{setRoomDataMenu.setContent}}
	setRoomNumberDataMenuContent := &InputContent{"Please Input The Room Number : ", InputRoomNumberKey, Content{setRoomDataMenu.setContent}}
	setRoomActionContent := &ActionContent{BuildSetRoomDataAction, Content{setRoomDataMenu.setContent}}
	setRoomDataMenu.contents = []MenuContent{setRoomLocationDataMenuContent, setRoomNumberDataMenuContent, setRoomActionContent}

	mainMenus = []Menu{reservationMenu, setRoomDataMenu}
	reservationMenu.nextMenus = mainMenus
	setRoomDataMenu.nextMenus = mainMenus

	return mainMenus
}

func SpecialMenuStrategy() []Menu {
	return nil
}
