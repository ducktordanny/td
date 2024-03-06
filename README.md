# TD

> Just a simple git like TODO list application for the terminal where I can do project based TODOs or global ones if I want to. Later on I want to use it inside NeoVim with telescope or something like that.

## Install

WIP

Currently, it can be built from the repo:

```sh
git clone https://github.com/ducktordanny/td.git
```

then `cd` into the repo's directory and run:

```sh
./setup
```

This script will initialize a default config to store the TODOs in the `~/.config/.tdconfig.json` file. This might change in the future, for now it is good enough for me.

## Commands

| Command             | Alias     | Type   | Description                                                                                                                                         |
| ------------------- | --------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------------- |
| `--add`             | `-a`      |        | Creates a new TODO.                                                                                                                                 |
| `--content`         | `-c`      | string | Takes a content for adding or editing a todo. This is useful if you don't want the default behaviour which opens up an editor and takes your input. |
| `--edit`            | `-e`      | string | Edits a TODO based on the provided ID.                                                                                                              |
| `--global`          | `-g`      |        | Switches to global scope, and handles global TODOs.                                                                                                 |
| `--list`            | `-ls`     |        | Returns the list of all TODOs in the scope.                                                                                                         |
| `--list-resolved`   | `-ls-rs`  |        | Returns the list of resolved TODOs in the scope.                                                                                                    |
| `--list-unresolved` | `-ls-urs` |        | Returns the list of unresolved TODOs in the scope.                                                                                                  |
| `--remove`          | `-rm`     | string | Removes a TODO based on the provided ID.                                                                                                            |
| `--resolve`         | `-rs`     | string | Resolves a TODO based on the provided ID.                                                                                                           |
| `--toggle`          | `-t`      | string | Toggles the resolved status of a TODO based on the provided ID.                                                                                     |
| `--unresolve`       | `-urs`    | string | Unresolves a TODO based on the provided ID.                                                                                                         |

For more info see `td --help`

-a Alias for 'add'.
-c string Alias for 'content'.
-e string Alias for 'edit'.
-g Alias for 'global'.
-ls Alias for 'list'.
-ls-rs Alias for 'list-resolved'.
-ls-urs Alias for 'list-unresolved'.
-rm string Alias for 'remove'.
-rs string Alias for 'resolve'.
-t string Alias for 'toggle'.
-urs string Alias for 'unresolve'.
