package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	notesheaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("15")).
				Background(lipgloss.Color("0")).
				Padding(1, 2).
				Border(lipgloss.ThickBorder()).
				BorderForeground(lipgloss.Color("15")).
				Align(lipgloss.Center)

	instructionStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("248")).
				Align(lipgloss.Center)

	//selected note style
	activeNoteStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("15")).Background(lipgloss.Color("141")).Bold(true)

	noteTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("14")).
			Bold(true)

	faint          = lipgloss.NewStyle().Foreground(lipgloss.Color("245")).Faint(true)
	itemLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(4)
	contentStyle   = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Italic(true)
)

func (m model) View() string {
	s := notesheaderStyle.Render("N O T E S") + "\n\n"

	// Render the body of the app
	if m.state == bodyView {
		s += noteTitleStyle.Render("Note:") + "\n\n"
		s += m.textarea.View() + "\n\n"
		s += instructionStyle.Render("ctrl+s - save • esc - discard")
		return s
	}

	//title input view
	if m.state == titleView {
		s += noteTitleStyle.Render("Note title:") + "\n\n"
		s += m.textinput.View() + "\n\n"
		s += instructionStyle.Render("enter - save • esc - discard")
		return s
	}

	//list of notes
	if m.state == listView {
		for i, n := range m.notes {
			selectionIndic := " "
			if i == m.list_index {
				selectionIndic = ">"
			}

			shortBody := strings.ReplaceAll(n.Body, "\n", " ")
			if len(shortBody) > 40 {
				shortBody = shortBody[:40] + "..."
			}

			if i == m.list_index {
				s += activeNoteStyle.Render(selectionIndic+n.Title+" | "+contentStyle.Render(shortBody)) + "\n\n"
			} else {
				s += itemLabelStyle.Render(selectionIndic+n.Title+" | "+faint.Render(shortBody)) + "\n\n"
			}
		}
		s += instructionStyle.Render("n - new note • q - quit")
	}

	return s
}
