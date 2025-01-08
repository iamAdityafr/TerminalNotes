# Terminal Notes App

A simple terminal-based notes-taking application built using Go, the `Bubbletea` framework, `Go-SQLite3`, and `Lipgloss` for styling.

## Features

- **Create Notes**: Allows you to create new notes directly in the terminal.
- **Edit Notes**: Edit your existing notes from within the app.
- **View Notes**: Display a list of all your notes in the terminal.
- **Delete Notes**: Easily delete notes from the app.
- **SQLite Database**: All notes are stored in a local SQLite database.

## Technologies Used

- **Bubbletea**: A Go framework for building terminal-based applications with a model-update-view architecture.
- **Go-SQLite3**: Used for database management and storing notes in SQLite format.
- **Lipgloss**: A package for styling terminal output.

## Prerequisites

Before installing and running the Terminal Notes App, ensure you have the following installed on your system:

1. **Go**: The application is built using Go. You can check if Go is installed by running:
   ```bash
   go version
2. **SQLite**: The app uses SQLite for database management. Ensure SQLite is installed on your system:
* On Linux, install it using your package manager:

```bash
sudo apt-get install sqlite3
```
* On Windows, download the SQLite binaries from the official SQLite website and follow the installation instructions.

## Installation
Clone the repository to your local machine:

```bash
git clone https://github.com/iamAdityafr/TerminalNotes
cd TerminalNotes/src
```
Install the required Go dependencies:

```bash
go mod download
```

Build the application:

```bash
go build -o terminal_notes
```
Run the application:

```bash
./terminal_notes
```

You should now be able to use the Terminal Notes App to create, edit, view, and delete notes directly from your terminal!

## Images

![Terminal notes1](https://github.com/iamAdityafr/TerminalNotes/blob/main/ig1.png?raw=true)

![Terminal notes2](https://github.com/iamAdityafr/TerminalNotes/blob/main/ig2.png?raw=true)

