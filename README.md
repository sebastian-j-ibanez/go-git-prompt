# go-git-prompt

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sebastian-j-ibanez/go-git-prompt?logo=go&logoColor=white&color=00ADD8)
![GitHub Release](https://img.shields.io/github/v/release/sebastian-j-ibanez/go-git-prompt?include_prereleases&logo=github&color=00ADD8)

A cross-platform Git status prompt written in Go.

## Usage

1. Download the latest release for your system.
2. Make sure `go-git-prompt` is in your `$PATH`.
3. Add `go-git-prompt prompt` to your shell prompt.
4. Open a git repo in your terminal and you will see the Git prompt. Enjoy!

<img src="https://github.com/user-attachments/assets/d1b8a4a6-2621-4e11-a9fd-7b08cadb6c70" alt="Adding go-git-prompt to .bashrc" width="500"/>
<img src="https://github.com/user-attachments/assets/7b5387c1-0353-4ccf-af23-ec673152558e" alt="Example status prompts" width="500"/>

## FAQ

### What does it do?

`go-git-prompt` shows information about a Git repository in the terminal. It is meant to be added to your shell prompt.

### Why another shell prompt project?

Projects like [oh-my-posh](https://ohmyposh.dev/) and [starship](https://starship.rs/) are fantastic but overkill for a simple git status prompt. I created `go-git-prompt` because I wanted a solution like [bash-git-prompt](https://github.com/magicmonty/bash-git-prompt) on **both** Linux and Windows.

### What is the project status?

The project is in early development. `go-git-prompt` will be ready for public use in release 1.0.

## License

`go-git-prompt` is released under the MIT license. Please see [LICENSE.md](LICENSE.md) for more details.
