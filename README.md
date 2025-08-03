# AstraShell

AstraShell is a CLI tool designed for developers, providing a suite of microservices to streamline your workflow. Inspired by platforms like Vercel, AstraShell brings together multiple shell-based utilities to create a mini CLI "LAN Vercel"-like experience.
---

## Features

- **Microservice Architecture:** Each feature is implemented as a standalone microservice, making it easy to extend and maintain.
- **Git Push Utility:** Push your code to Git with a single command. AstraShell prompts you for a commit message and handles the rest, simplifying your version control workflow.
- **Shell Script Integration:** Many services are implemented in shell scripts for maximum portability and ease of use.
---

## Getting Started

1. Clone the repository:
   ```sh
   git clone https://github.com/Gustavo-DCosta/AstraShell.git
   ```
2. Explore the available microservices in the subdirectories.
---

## Example: Push Code to Git

AstraShell includes a microservice to automate pushing code to Git. Simply run the provided script, enter your commit message when prompted, and your changes will be pushed to the main branch.

- **Bash:**
  ```sh
  ./pushLynx/script/git.sh
  ```
- **PowerShell:**
  ```powershell
  ./pushLynx/script/git.ps1
  ```

## License

MIT

