package menus

import "fmt"

type Menu interface {
	Show()
	Enter()
	Move() []Menu
	GetId() string
}

type CliMenu struct {
	MenuId            string
	Level             int32
	Text              string
	sharedInputValues map[string]string
	contents          []MenuContent
	nextMenus         []Menu
}

func (mn *CliMenu) Show() {
	fmt.Println(mn.MenuId, ".", mn.Text)
}

func (mn *CliMenu) Enter() {
	for _, content := range mn.contents {
		content.Execute(mn.sharedInputValues)
	}
}

func (mn *CliMenu) Move() []Menu {
	return mn.nextMenus
}

func (mn *CliMenu) GetId() string {
	return mn.MenuId
}

func (mn *CliMenu) setContentValue(key string, value string) {
	mn.sharedInputValues[key] = value
}
