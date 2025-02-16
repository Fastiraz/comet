package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/Fastiraz/conventional-commits-maker/breaking-change"
	"github.com/Fastiraz/conventional-commits-maker/input"
	"github.com/Fastiraz/conventional-commits-maker/menu-type"
	"github.com/Fastiraz/conventional-commits-maker/textarea"
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

	if finalModel, ok := result.(menu.Model); ok && finalModel.ItemChosen {
		fmt.Printf(
			"You selected: %s - %s\n",
			finalModel.Selected.TitleStr,
			finalModel.Selected.Desc,
		)
	} else {
		fmt.Println("No item selected.")
	}

	commitScope := scope.ScopeInput()
	if err != nil {
		fmt.Println("Error running program:", err)
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Printf(
		"You select scope: %s\n",
		commitScope,
	)

	isBreaking := breakchange.IsBreakingChange()
	fmt.Printf("Is breaking change: %v\n", isBreaking)

	body := textarea.TextArea("body")
	fmt.Printf("Body content: %s\n", body)

	footer := textarea.TextArea("footer")
	fmt.Printf("Footer content: %s\n", footer)
}
