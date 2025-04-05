# ASCII Game Menu

A colorful TUI (Text User Interface) menu system for launching ASCII-based games.

## Features

- Beautiful terminal-based interface using tview
- Game descriptions displayed in a side panel
- Easy configuration via JSON
- Support for launching external game executables
- Keyboard navigation (arrow keys, Enter, Space)
- Self-contained executable (no external dependencies needed)
- Support for both embedded and external configurations

## Installation

1. Clone the repository
2. Install Go dependencies:
   ```bash
   go mod download
   ```
3. Build the self-contained executable:
   ```bash
   # For macOS (Intel)
   CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-w -s" -o ascii-menu
   
   # For macOS (Apple Silicon)
   CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-w -s" -o ascii-menu
   
   # For Linux
   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ascii-menu
   ```

The resulting binary can be copied anywhere and run directly.

## Interface

The menu interface consists of two main panels:

1. **Left Panel**: Game List
   - Shows all available games
   - Currently selected game is highlighted in green
   - Games are navigated using arrow keys

2. **Right Panel**: Game Description
   - Shows detailed description of the currently selected game
   - Updates automatically as you navigate through games
   - Yellow text for better readability

## Controls

- **↑ (Up Arrow)**: Move selection up in the game list
- **↓ (Down Arrow)**: Move selection down in the game list
- **Enter**: Launch the selected game
- **Space**: Return to menu (if in a game)
- **q**: Quit the application

## Game Launching

When you select a game and press Enter:
1. The menu will temporarily suspend
2. The selected game will launch in the same terminal
3. When the game exits, the menu will automatically return
4. If the game fails to launch, an error message will be displayed

## Configuration

The menu system supports two types of configuration:

1. **Embedded Configuration** (built into the binary)
   - Located in `config/games/config.json`
   - Contains default games
   - Always available

2. **External Configuration** (optional)
   - Located at `/etc/ascii-menu/config.json`
   - Can be used to add additional games
   - Games from external config are added to the embedded games
   - Format is the same as embedded config

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

To add your own games:
1. Create the directory: `sudo mkdir -p /etc/ascii-menu`
2. Create the config file: `sudo touch /etc/ascii-menu/config.json`
3. Add your games following the example format
4. Set appropriate permissions: `sudo chmod 644 /etc/ascii-menu/config.json`

## Usage

Run the application:
```bash
./ascii-menu
```

## Development

The project is organized into:
- `main.go`: Main application logic
- `config/`: Configuration handling
- `config/games/`: Embedded game configuration files 