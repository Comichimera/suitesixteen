# SuiteSixteen Data Logger

SuiteSixteen Data Logger is a command-line application for recording and uploading Joy Ride Turbo leaderboard data to the GitHub page of JRTDB.com. It allows users to enter data interactively, save it to a CSV file, and upload it to a GitHub repository with a custom commit message.

## Table of Contents

- [Installation](#installation)
  - [Arch Linux](#arch-linux)
  - [Windows](#windows)
- [Usage](#usage)
  - [Commands](#commands)
- [Workflow](#workflow)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Arch Linux

1. **Install Go**: Make sure you have Go installed. You can install it from the Arch User Repository (AUR) or using the following command:

    ```bash
    sudo pacman -S go
    ```

2. **Clone the Repository**:

    ```bash
    git clone https://github.com/yourusername/suitesixteen.git
    cd suitesixteen
    ```

3. **Set Up GitHub Personal Access Token**: Before running the program, set the environment variable for your GitHub Personal Access Token (PAT). This token is used to authenticate when uploading files to GitHub.

    ```bash
    export GITHUB_TOKEN=your_github_token_here
    ```

4. **Build the Application**:

    ```bash
    go build -o suitesixteen ./cmd/main
    ```

5. **Run the Application**:

    ```bash
    ./suitesixteen
    ```

### Windows

1. **Install Go**: Download and install Go from [the official Go website](https://golang.org/dl/).

2. **Clone the Repository**: Open a command prompt and run:

    ```cmd
    git clone https://github.com/yourusername/suitesixteen.git
    cd suitesixteen
    ```

3. **Set Up GitHub Personal Access Token**: Before running the program, set the environment variable for your GitHub Personal Access Token (PAT). This token is used to authenticate when uploading files to GitHub. Run the following command in the Command Prompt:

    ```cmd
    setx GITHUB_TOKEN your_github_token_here
    ```

4. **Build the Application**:

    ```cmd
    go build -o suitesixteen.exe ./cmd/main
    ```

5. **Run the Application**:

    ```cmd
    suitesixteen.exe
    ```

## Usage

Once the application is running, follow the on-screen prompts to enter and manage your data.

### Commands

- **`commit`**: Save the current record to the list of committed records.
- **`rollback`**: Remove the last committed record from the list.
- **`track`**: Change the current track you are documenting.
- **`mode`**: Change the current game mode you are documenting.
- **`push`**: Save all committed records to the CSV file and upload it to GitHub.

## Workflow

1. **Start the Program**: Run the `suitesixteen` executable from your terminal or command prompt.

2. **Select Game Mode**: You will be prompted to input a number for the game mode:
   - Type `1` for Time Trial and press Enter.
   - Type `2` for Pro Race and press Enter.
   - Type `3` for Battle Race and press Enter.

3. **Select Track**: After selecting the game mode, choose a track by typing the corresponding number and pressing Enter.

4. **Prepare to Enter Data**: Make sure you are reading from the correct leaderboard in your game. Keep the input format consistent.

5. **Enter Data**:
   - Type the rank of the run on the leaderboard and press Enter.
   - Type the gamertag and press Enter.
   - Type the time in the format `m:ss.xxx` (e.g., `1:23.456`) and press Enter.

6. **Commit the Record**: After entering the data, type `commit` and press Enter to save the record.

7. **Change Game Mode or Track (if needed)**:
   - Type `mode` to change the game mode and follow the prompts.
   - Type `track` to change the track and follow the prompts.

8. **Push Data to GitHub**: When you are done inputting data, type `push` to save all committed records to the CSV file and upload it to GitHub. You will be prompted to enter a commit message before the upload.
9. **Uncommitted data**: Any data that has not been committed will not be saved if another command is used, such as `push`, `rollback`, `track` or `mode`.
