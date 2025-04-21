package textarea

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

func TextArea(mode string) string {
	p := tea.NewProgram(initialModel(mode))

	result, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	if m, ok := result.(model); ok {
		return m.textarea.Value()
	}

	return ""
}

type errMsg error

type model struct {
	textarea textarea.Model
	err      error
}

func initialModel(mode string) model {
	ti := textarea.New()
	if mode == "body" {
		ti.Placeholder = "Tell me the body of the commit..."
	} else if mode == "footer" {
		ti.Placeholder = "Tell me the footer of the commit..."
	} else {
		ti.Placeholder = "Write something..."
	}

ti.ShowLineNumbers = true

ti.Focus()

	return model{
		textarea: ti,
		err:      nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if msg.Alt { // Alt+Enter for newline
				m.textarea.SetValue(m.textarea.Value() + "\n")
			} else {
				return m, tea.Quit
			}
		case tea.KeyEsc:
			if m.textarea.Focused() {
				m.textarea.Blur()
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.textarea.Focused() {
				cmd = m.textarea.Focus()
				cmds = append(cmds, cmd)
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		commitMessage(m.textarea.Placeholder),
		m.textarea.View(),
		"(enter to submit, alt+enter for new line, ctrl+c to quit)",
	) + "\n\n"
}

func commitMessage(placeholder string) string {
	if placeholder == "Tell me the body of the commit..." {
		return "Body commit"
	} else if placeholder == "Tell me the footer of the commit..." {
		return "Footer commit"
	}
	return "Write something..."
}
