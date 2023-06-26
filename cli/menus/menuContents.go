package menus

import (
	"bufio"
	"fmt"
	"os"
	"reservation/core/usecase"
	"strings"
)

type MenuContent interface {
	Execute(map[string]string)
}

type Content struct {
	setValueCallBack func(string, string)
}

type TextContent struct {
	text string
	Content
}

type InputContent struct {
	Text string
	Key  string
	Content
}

type ActionContent struct {
	confirmationValidation func(map[string]string) bool
	buildInputBoundary     func(map[string]string) usecase.InputBoundary
	Content
}

func (tc *TextContent) Execute(sharedInputValues map[string]string) {
	fmt.Println(tc.text)
}

func (ic *InputContent) Execute(sharedInputValues map[string]string) {
	fmt.Print(replaceTextWithActualValue(ic.Text, sharedInputValues))
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	ic.Content.setValueCallBack(ic.Key, strings.Trim(input, "\n"))
}

func (ac *ActionContent) Execute(sharedInputValues map[string]string) {
	confirmed := true
	if ac.confirmationValidation != nil {
		confirmed = ac.confirmationValidation(sharedInputValues)
	}
	if confirmed {
		ac.buildInputBoundary(sharedInputValues).Execute()
	}
}

func replaceTextWithActualValue(input string, sharedInputValues map[string]string) string {
	for key, value := range sharedInputValues {
		input = strings.Replace(input, "{"+key+"}", value, 1)
	}
	return input
}
