package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/Fastiraz/comet/breaking-change"
	"github.com/Fastiraz/comet/input"
	"github.com/Fastiraz/comet/menu-type"
	"github.com/Fastiraz/comet/textarea"
	"github.com/charmbracelet/glamour"
)

func main() {
	items := []list.Item{
		menu.Item{TitleStr: "feat", Desc: "Commits that add or remove a new feature to the API or UI"},
		menu.Item{TitleStr: "fix", Desc: "Commits that fix an API or UI bug of a preceded feat commit"},
		menu.Item{TitleStr: "refactor", Desc: "Commits that rewrite/restructure code without changing behavior"},
		menu.Item{TitleStr: "perf", Desc: "Commits that improve performance"},
		menu.Item{TitleStr: "style", Desc: "Commits that do not affect the meaning (formatting, etc)"},
		menu.Item{TitleStr: "test", Desc: "Commits that add or correct tests"},
		menu.Item{TitleStr: "docs", Desc: "Commits that affect documentation only"},
		menu.Item{TitleStr: "build", Desc: "Commits that affect build components"},
		menu.Item{TitleStr: "ops", Desc: "Commits that affect operational components"},
		menu.Item{TitleStr: "chore", Desc: "Miscellaneous commits like modifying .gitignore"},
		menu.Item{TitleStr: "ci", Desc: "Commits related to continuous integration"},
		menu.Item{TitleStr: "revert", Desc: "Commits that revert previous changes"},
	}

	m := menu.NewMenu(items, "Choose a commit type.")
	p := tea.NewProgram(m, tea.WithAltScreen())

	result, err := p.Run()
	if err != nil {
		fmt.Println("Error running program:", err)
		log.Fatal(err)
		os.Exit(1)
	}

	// if finalModel, ok := result.(menu.Model); ok && // finalModel.ItemChosen {
	// 	fmt.Printf(
	// 		"You selected: %s - %s\n",
	// 		finalModel.Selected.TitleStr,
	// 		finalModel.Selected.Desc,
	// 	)
	// } else {
	// 	fmt.Println("No item selected.")
	// }
	finalModel, ok := result.(menu.Model)
	if !ok || !finalModel.ItemChosen {
		fmt.Println("No item selected.")
		return
	}

	commitScope := scope.Input("scope")
	if err != nil {
		fmt.Println("Error running program:", err)
		log.Fatal(err)
		os.Exit(1)
	}

	isBreaking := breakchange.IsBreakingChange()

	subject := scope.Input("subject")
	if err != nil {
		fmt.Println("Error running program:", err)
		log.Fatal(err)
		os.Exit(1)
	}

	body := textarea.TextArea("body")
	footer := textarea.TextArea("footer")

	command := BuildCommand(
		finalModel.Selected.TitleStr,
		commitScope,
		isBreaking,
		subject,
		body,
		footer,
	)

	in := fmt.Sprintf("\n\nGit command: \n\n```bash\n%s\n```\n", command)

	out, err := glamour.Render(in, "dark")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Print(out)
	}
}

func BuildCommand(
	commitType string,
	scope string,
	isBreaking bool,
	subject string,
	body string,
	footer string) string {

	command := fmt.Sprintf("git commit -m \"%s", commitType)

	if scope != "" {
		command = fmt.Sprintf("%s(%s)", command, scope)
	}

	if isBreaking {
		command = fmt.Sprintf("%s!: ", command)
	} else {
		command = fmt.Sprintf("%s: ", command)
	}

	command = fmt.Sprintf("%s%s", command, subject)

	if body != "" {
		command = fmt.Sprintf("%s\n\n%s", command, body)
	}

	if footer != "" {
		command = fmt.Sprintf("%s\n\n%s", command, footer)
	}

	command = fmt.Sprintf("%s\"", command)
	err := CopyToClipboard(command)
	if err != nil {
		fmt.Println("Failed to copy:", err)
	} else {
		fmt.Println("Git command copied to clipboard.")
	}

	return command
}

func CopyToClipboard(text string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("pbcopy")
	case "linux":
		cmd = exec.Command("xclip", "-selection", "clipboard")
	case "windows":
		cmd = exec.Command("cmd", "/c", "clip")
	default:
		return fmt.Errorf("unsupported platform")
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		_, _ = stdin.Write([]byte(text))
	}()

	return cmd.Run()
}
