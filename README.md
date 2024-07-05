# 🚀 nixi

## What is nixi? 🤔

`nixi` is a tool designed to simplify software installation on NixOS. Instead of manually modifying the `configuration.nix` file, `nixi` automates the process for you.

## Why does nixi exist? 💡

When I started using NixOS, I found it challenging to install software with a single command. While `nix-env` offers some functionality, it requires ongoing maintenance. Therefore, I created `nixi` to streamline software management on NixOS.

## Tools used and why 🛠️

`nixi` is programmed in Go. I chose Go because I wanted to learn the language, and it allows me to quickly compile binaries.

## How to install 🧩

1. Download the `nixi` executable file from the project repository.
2. Move the `nixi` file to `/usr/local/bin/` (create the directory if it doesn't exist):

    ```sh
    sudo mkdir -p /usr/local/bin
    sudo mv nixi /usr/local/bin/
    ```

3. Add `/usr/local/bin` to your `PATH` permanently:

   For `bash` users, create and edit `~/.bashrc`:

    ```sh
    touch ~/.bashrc
    echo 'export PATH=$PATH:/usr/local/bin' >> ~/.bashrc
    source ~/.bashrc
    ```

   For `zsh` users, create and edit `~/.zshrc`:

    ```sh
    touch ~/.zshrc
    echo 'export PATH=$PATH:/usr/local/bin' >> ~/.zshrc
    source ~/.zshrc
    ```

4. Verify the installation by running `nixi`. If it doesn't work, ensure your `PATH` is correctly set by checking the contents of `~/.bashrc` or `~/.zshrc`.

## How to use nixi 💻

`nixi` provides two main commands for installing and removing software:

- To install software:

    ```sh
    nixi install steam
    nixi install steam,qbittorrent,librewolf
    ```

- To remove software:

    ```sh
    nixi remove steam
    nixi remove steam,qbittorrent,librewolf
    ```

You can install or remove multiple tools by separating them with commas.

## Issues and contributions 🛠️

If you encounter any issues or have suggestions for improvements, please open a GitHub issue. Your feedback is appreciated!

---

I hope you find this tool useful. Thank you for using `nixi`! 🙌
