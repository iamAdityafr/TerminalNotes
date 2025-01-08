package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	titleView
	bodyView
	confirmDeleteView
	searchView
)

type model struct {
	store       *Store
	state       uint
	textarea    textarea.Model
	textinput   textinput.Model
	activeNote  Note
	notes       []Note
	list_index  int
	deleteIndex int
}

func initialModel(store *Store) model {
	notes, err := store.GetNotes()
	if err != nil {
		log.Fatalf("unable to get notes: %v", err)
	}
	//
	searchInput := textinput.New()
	searchInput.Placeholder = "Search notes..."
	searchInput.Focus()

	return model{
		store:     store,
		state:     listView,
		textarea:  textarea.New(),
		textinput: textinput.New(),
		notes:     notes,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds    []tea.Cmd
		command tea.Cmd
	)

	m.textarea, command = m.textarea.Update(msg)
	cmds = append(cmds, command)

	m.textinput, command = m.textinput.Update(msg)
	cmds = append(cmds, command)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit

			case "n":
				m.textinput.SetValue("")
				m.textinput.Focus()
				m.activeNote = Note{}
				m.state = titleView
			case "up", "k":
				if m.list_index > 0 {
					m.list_index--
				}
			case "down", "j":
				if m.list_index < len(m.notes)-1 {
					m.list_index++
				}
			case "enter":
				m.activeNote = m.notes[m.list_index]
				m.state = bodyView
				m.textarea.SetValue(m.activeNote.Body)
				m.textarea.Focus()
				m.textarea.CursorEnd()
			//delete note
			case "d":
				m.deleteIndex = m.list_index
				m.state = confirmDeleteView
			}

		case confirmDeleteView:
			switch key {
			case "y":
				noteToDelete := m.notes[m.deleteIndex]
				err := m.store.DeleteNote(noteToDelete.ID)
				if err != nil {
					log.Printf("failed to delete note: %v", err)
				} else {
					m.notes = append(m.notes[:m.deleteIndex], m.notes[m.deleteIndex+1:]...)
					if m.list_index >= len(m.notes) && m.list_index > 0 {
						m.list_index--
					}
				}
				m.state = listView
			case "n":
				m.state = listView
			}

		case titleView:
			switch key {
			case "enter":
				title := m.textinput.Value()
				if title != "" {
					m.activeNote.Title = title

					m.state = bodyView
					m.textarea.SetValue("")
					m.textarea.Focus()
					m.textarea.CursorEnd()
				}
			case "esc":
				m.state = listView
			}

		case bodyView:
			switch key {
			case "ctrl+s":
				m.activeNote.Body = m.textarea.Value()

				var err error
				if err = m.store.SaveNote(m.activeNote); err != nil {
					return m, tea.Quit
				}

				m.notes, err = m.store.GetNotes()
				if err != nil {
					return m, tea.Quit
				}

				m.state = listView
			case "esc":
				m.state = listView
			}
		}
	}

	return m, tea.Batch(cmds...)
}
