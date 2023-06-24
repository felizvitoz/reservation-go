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
	buildInputBoundary func(map[string]string) usecase.InputBoundary
	Content
}

func (tc *TextContent) Execute(sharedInputValues map[string]string) {
	fmt.Println(tc.text)
}

func (ic *InputContent) Execute(sharedInputValues map[string]string) {
	fmt.Print(ic.Text)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	ic.Content.setValueCallBack(ic.Key, strings.Trim(input, "\n"))
}

func (ac *ActionContent) Execute(sharedInputValues map[string]string) {
	ac.buildInputBoundary(sharedInputValues).Execute()
}
