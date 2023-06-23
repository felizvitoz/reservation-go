package menus

import "fmt"

type Menu interface {
	Show()
	Enter()
	Move() []Menu
	GetId() string
}

type CliMenu struct {
	MenuId    string
	Level     int32
	Order     int32
	Text      string
	contents  []MenuContent
	nextMenus []Menu
}

func (mn CliMenu) Show() {
	fmt.Println(mn.MenuId, ".", mn.Text)
}

func (mn CliMenu) Enter() {
	for _, content := range mn.contents {
		content.Execute()
	}
}

func (mn CliMenu) Move() []Menu {
	return mn.nextMenus
}

func (mn CliMenu) GetId() string {
	return mn.MenuId
}
