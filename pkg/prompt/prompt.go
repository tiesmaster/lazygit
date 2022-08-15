package prompt

import (
	"os"

	"github.com/chzyer/readline"
	"github.com/manifoldco/promptui"
)

// This is for creating CLI menus and prompts when we're not inside the full lazygit UI.
// Useful for asking questions before we properly start lazygit.

// TODO: use https://github.com/AlecAivazis/survey

func Menu(label string, items []string) (string, error) {
	prompt := promptui.Select{
		Label:  label,
		Items:  items,
		Stdin:  os.Stdin,
		Stdout: NoBellStdout,
	}

	_, result, err := prompt.Run()
	return result, err
}

func Prompt(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: func(string) error { return nil },
		Stdin:    os.Stdin,
		Stdout:   NoBellStdout,
	}

	return prompt.Run()
}

// see https://github.com/manifoldco/promptui/issues/49#issuecomment-1012640880
// This is a hack to get around the bell noise that mac users hear when selecting items
type noBellStdout struct{}

func (n *noBellStdout) Write(p []byte) (int, error) {
	if len(p) == 1 && p[0] == readline.CharBell {
		return 0, nil
	}
	return readline.Stdout.Write(p)
}

func (n *noBellStdout) Close() error {
	return readline.Stdout.Close()
}

var NoBellStdout = &noBellStdout{}
