# ASCII Game Menu

A colorful TUI (Text User Interface) menu system for launching ASCII-based games.

## Features

- Beautiful terminal-based interface using tview
- Game descriptions displayed in a side panel
- Easy configuration via JSON
- Support for launching external game executables
- Keyboard navigation (arrow keys, Enter, Space)

## Installation

1. Clone the repository
2. Install Go dependencies:
   ```bash
   go mod download
   ```
3. Build the application:
   ```bash
   go build
   ```

## Configuration

Games are configured in `games/config.json`. Each game entry requires:
- `name`: Display name
- `description`: Short description
- `path`: Absolute path to the game executable

Example configuration:
```json
{
    "games": [
        {
            "name": "Snake",
            "description": "Classic snake game",
            "path": "/usr/games/snake"
        }
    ]
}
```

## Usage

Run the application:
```bash
./ascii-menu
```

Controls:
- ↑/↓: Navigate menu
- Enter: Launch selected game
- Space: Return to menu
- q: Quit application

## Development

The project is organized into:
- `main.go`: Main application logic
- `config/`: Configuration handling
- `games/`: Game configuration files 