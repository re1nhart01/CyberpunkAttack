#!/bin/bash

# Define the command you want to run
COMMAND="redis-server"

# Function to open a new terminal and execute the command
run_command_in_new_terminal() {
  case "$OSTYPE" in
    linux*) # For Linux
      if command -v gnome-terminal &>/dev/null; then
        gnome-terminal -- bash -c "$COMMAND; exec bash"
      elif command -v konsole &>/dev/null; then
        konsole -e bash -c "$COMMAND; exec bash"
      elif command -v xterm &>/dev/null; then
        xterm -e "$COMMAND; bash"
      else
        echo "No supported terminal emulator found!"
        exit 1
      fi
      ;;
    darwin*) # For macOS
      osascript -e "tell application \"Terminal\" to do script \"$COMMAND\""
      ;;
    cygwin* | msys* | win32*) # For Windows (Git Bash, MSYS, or Cygwin)
      cmd.exe /c start bash -c "$COMMAND"
      ;;
    *)
      echo "Unsupported OS: $OSTYPE"
      exit 1
      ;;
  esac
}

# Run the function
run_command_in_new_terminal
