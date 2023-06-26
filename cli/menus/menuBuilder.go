package menus

import (
	"strings"
)

func GetMenu(themeType int32) []Menu {
	switch themeType {
	case 101:
		return SpecialMenuStrategy()
	default:
		return DefaultMenuStrategy()
	}
}

func DefaultMenuStrategy() []Menu {
	// =============== CreateReservation ================
	reservationMenu := &CliMenu{
		MenuId:            "1",
		Level:             1,
		Text:              "Create Reservation",
		sharedInputValues: make(map[string]string),
	}

	summaryContent := "Summary of Your Reservation : \n" +
		"Location : {" + InputLocationKey + "}\n" +
		"Room Number : {" + InputRoomNumberKey + "}\n" +
		"Reservation Date : {" + InputReservationDateKey + "}\n" +
		"Do you want to confirm this reservation (yes/no) : "
	createReservationLocationInputContent := &InputContent{"Please Input The Location : ", InputLocationKey, Content{reservationMenu.setContentValue}}
	createReservationRoomInputContent := &InputContent{"Please Input The Room Number : ", InputRoomNumberKey, Content{reservationMenu.setContentValue}}
	createReservationDateInputContent := &InputContent{"Please Input Reservation Date : ", InputReservationDateKey, Content{reservationMenu.setContentValue}}
	createReservationConfirmationInputContent := &InputContent{summaryContent, InputReservationConfirmationKey, Content{reservationMenu.setContentValue}}
	reservationConfirmationFunction := func(valueMap map[string]string) bool {
		return strings.Compare("yes", valueMap[InputReservationConfirmationKey]) == 0
	}
	createReservationActionContent := &ActionContent{
		confirmationValidation: reservationConfirmationFunction,
		buildInputBoundary:     BuildReservationAction,
		Content:                Content{reservationMenu.setContentValue}}
	reservationMenu.contents = []MenuContent{createReservationLocationInputContent, createReservationRoomInputContent, createReservationDateInputContent, createReservationConfirmationInputContent, createReservationActionContent}

	// =============== Reservation Report ================
	reservationReportsMenu := &CliMenu{
		MenuId:            "2",
		Level:             1,
		Text:              "Reservation Reports",
		sharedInputValues: make(map[string]string),
	}

	// =============== Available Room Report ================
	availableRoomReportMenu := &CliMenu{
		MenuId:            "RPT1",
		Level:             2,
		Text:              "Available Room Report",
		sharedInputValues: make(map[string]string),
	}
	// =============== All Room Report ================
	allRoomReportMenu := &CliMenu{
		MenuId:            "RPT2",
		Level:             2,
		Text:              "All Room Report",
		sharedInputValues: make(map[string]string),
	}
	// =============== Set Room Data Menu ================
	setRoomDataMenu := &CliMenu{
		MenuId:            "99",
		Level:             1,
		Text:              "Set Room Data",
		sharedInputValues: make(map[string]string),
	}
	setRoomLocationDataMenuContent := &InputContent{"Please Input The Location : ", InputLocationKey, Content{setRoomDataMenu.setContentValue}}
	setRoomNumberDataMenuContent := &InputContent{"Please Input The Room Number : ", InputRoomNumberKey, Content{setRoomDataMenu.setContentValue}}
	setRoomActionContent := &ActionContent{
		buildInputBoundary: BuildSetRoomDataAction,
		Content:            Content{setRoomDataMenu.setContentValue}}
	setRoomDataMenu.contents = []MenuContent{setRoomLocationDataMenuContent, setRoomNumberDataMenuContent, setRoomActionContent}

	mainMenus := []Menu{reservationMenu, reservationReportsMenu, setRoomDataMenu}
	reportMenus := []Menu{availableRoomReportMenu, allRoomReportMenu}
	reservationMenu.nextMenus = mainMenus
	setRoomDataMenu.nextMenus = mainMenus
	reservationReportsMenu.nextMenus = reportMenus
	availableRoomReportMenu.nextMenus = mainMenus
	allRoomReportMenu.nextMenus = mainMenus

	return mainMenus
}

func SpecialMenuStrategy() []Menu {
	return nil
}
