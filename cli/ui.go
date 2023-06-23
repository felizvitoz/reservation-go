package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reservation/cli/menus"
	"strings"
)

func main() {
	currentMenus := menus.GetMenu(1)
	exit := false
	for !exit {
		showAllMenus(currentMenus)

		input := showInput()
		if strings.Compare("exit", input) == 0 {
			exit = true
		}

		selectedMenu, errorOutput := findSelectedMenu(input, currentMenus)
		if errorOutput != nil {
			continue
		}
		selectedMenu.Enter()
		currentMenus = selectedMenu.Move()
	}
}

func showAllMenus(menus []menus.Menu) {
	for _, menu := range menus {
		menu.Show()
	}
}

func showInput() string {
	fmt.Print("select one menu...")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.Trim(input, "\n")
}

func findSelectedMenu(input string, currentMenus []menus.Menu) (menus.Menu, error) {
	var errorOutput error
	var selectedMenu menus.Menu
	for _, menu := range currentMenus {
		if strings.Compare(menu.GetId(), input) == 0 {
			selectedMenu = menu
		}
	}

	if selectedMenu == nil {
		errorOutput = errors.New("There is no matched menu")
	}
	return selectedMenu, errorOutput
}
