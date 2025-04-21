package scope

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	textInput textinput.Model
	err       error
	input     string
	mode      string
}

func initialModel(mode string) model {
	ti := textinput.New()
	if mode == "scope" {
		ti.Placeholder = "Enter your scope"
	} else if mode == "subject" {
		ti.Placeholder = "Enter your subject"
	}
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err:       nil,
		mode:      mode,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			m.input = m.textInput.Value()
			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var message string
	if m.mode == "scope" {
		message = "What’s the commit scope? (let empty if none)"
	} else if m.mode == "subject" {
		message = "What’s the commit subject?"
	}

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		message,
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}

func Input(mode string) string {
	p := tea.NewProgram(
		initialModel(mode),
		tea.WithAltScreen(),
	)
	result, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	if finalModel, ok := result.(model); ok {
		return finalModel.input
	}

	return ""
}
