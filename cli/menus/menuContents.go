package menus

import (
	"bufio"
	"fmt"
	"os"
	"reservation/core/usecase"
	"strings"
)

type MenuContent interface {
	Execute()
	GetOrder() int32
}

type Content struct {
	Order int32
}

type TextContent struct {
	text string
	Content
}

type InputContent struct {
	Value string
	Text  string
	Key   string
	Content
}

type ActionContent struct {
	InputBoundary usecase.InputBoundary
	Content
}

func (tc *TextContent) Execute() {
	fmt.Println(tc.text)
}

func (ic *InputContent) Execute() {
	fmt.Print(ic.Text)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	ic.Value = strings.Trim(input, "\n")
}

func (ac *ActionContent) Execute() {
	ac.InputBoundary.Execute()
}

func (ct *Content) GetOrder() int32 {
	return ct.Order
}
