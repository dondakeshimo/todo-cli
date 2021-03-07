[![Go][go-test-image]][go-test-url]
[![golangci-lint][golangci-lint-image]][golangci-lint-url]

[go-test-image]: https://github.com/dondakeshimo/todo-cli/workflows/Go/badge.svg
[go-test-url]: https://github.com/dondakeshimo/todo-cli/actions?query=workflow%3AGo
[golangci-lint-image]: https://github.com/dondakeshimo/todo-cli/workflows/golangci-lint/badge.svg
[golangci-lint-url]: https://github.com/dondakeshimo/todo-cli/actions?query=workflow%3Agolangci-lint


# todo-cli
Manage TODO List At CLI

:warning: only support MacOS now...

![result](https://user-images.githubusercontent.com/23194960/110210135-a49ec580-7ed3-11eb-8edb-7c7509847d8d.gif)

## Installation
Use the `go get` or `make install` after cloning from GitHub

```bash
go get github.com/dondakeshimo/todo-cli/cmd/todo 
```
OR
```bash
make install
```

make sure that you have already added binary path to your PATH

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

## Usage

```bash
$ todo --help
NAME:
   todo - Manage Your TODO

USAGE:
   todo [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   list, l    List tasks
   add, a     Add a task
   close, c   Close tasks
   modify, m  Modify a task
   notify     Notify a task (basicaly be used by system)
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Completion
You can use completion with bash or zsh.

##### :warning: zsh completion
ZSH completion may show inappropriate candidates
if you didn't configure below setting.  
We recommend that you set zsh-completion configuration.

#### Bash
Set `PROG=todo` and load `scripts/bash_autocomplete`.
Adding the following lines to your BASH configuration file (usually `.bash_profile` )
will allow the auto-completion to persist across new shells.

```bash
PROG=todo source path/to/todo-cli/scripts/bash_autocomplete
```

#### Zsh
Set `PROG=todo` and `_CLI_ZSH_AUTOCOMPLETE_HACK=1` , then load `scripts/zsh_autocomplete`.
Adding the following lines to your BASH configuration file (usually `.zshrc` )
will allow the auto-completion to persist across new shells.

```bash
PROG=todo
_CLI_ZSH_AUTOCOMPLETE_HACK=1
source path/to/todo-cli/scripts/zsh_autocomplete
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
