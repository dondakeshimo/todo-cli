# todo-cli
Manage TODO List At CLI

[![Go][go-test-image]][go-test-url]

[go-test-image]: https://github.com/dondakeshimo/todo-cli/workflows/Go/badge.svg
[go-test-url]: https://github.com/dondakeshimo/todo-cli/actions?query=workflow%3AGo

## Why todo-cli
- simple and light
- supply shell completion
- not show information if you not need

## Install
#### go installation
This is a simple way, but require [golang](https://golang.org/) .

Make sure that you have already added binary path to your PATH.

```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

##### Go version \< 1.16
```bash
$ go get -u github.com/dondakeshimo/todo-cli/cmd/todo
```

##### Go 1.16+
```bash
$ go install github.com/dondakeshimo/todo-cli/cmd/todo@latest
```

#### download binary
You can download binary from our repository.
Bellow example is for MaxOS.

```bash
$ TODO_VERSION=0.4.0

$ curl -O https://github.com/dondakeshimo/todo-cli/releases/download/v${TODO_VERSION}/todo-${TODO_VERSION}.macos-10.15.tar.gz

$ tar -xvf todo-${TODO_VERSION}.macos-10.15.tar.gz

$ mv todo path/to/your/$PATH
```

## Usage

```bash
$ todo --help
Manage Your TODO

Usage:
  todo [command]

Available Commands:
  add         Add a task
  close       Close tasks
  completion  generate the autocompletion script for the specified shell
  configure   Configure your todo-cli
  help        Help about any command
  list        List tasks
  modify      Modify a task
  notify      Notify a task (basicaly be used by system)

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.
```

```bash
$ todo list
+----+--------------------------------+----------------+----------+----------+
| ID |              Task              |   RemindTime   | Reminder | Priority |
+----+--------------------------------+----------------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |          |        0 |
|    | task is your first TODO        |                |          |          |
+----+--------------------------------+----------------+----------+----------+

$ todo close -i=1

$ todo l
+----+------+------------+----------+----------+
| ID | Task | RemindTime | Reminder | Priority |
+----+------+------------+----------+----------+
+----+------+------------+----------+----------+

$ todo add "must task" -d="2021/03/03 12:00" -p=0

$ todo l
+----+-----------+----------------+----------+----------+
| ID |   Task    |   RemindTime   | Reminder | Priority |
+----+-----------+----------------+----------+----------+
|  1 | must task | 2021/3/3 12:00 |          |        0 |
+----+-----------+----------------+----------+----------+

$ todo a "boring task"

$ todo a "important task" -d="2022/01/01" -p=50

$ todo l
+----+----------------+----------------+----------+----------+
| ID |      Task      |   RemindTime   | Reminder | Priority |
+----+----------------+----------------+----------+----------+
|  1 | must task      | 2021/3/3 12:00 |          |        0 |
|  2 | important task | 2022/1/1 00:00 |          |       50 |
|  3 | boring task    |                |          |      100 |
+----+----------------+----------------+----------+----------+

$ todo m -i=1 -t="should task" -p=10

$ todo l
+----+----------------+----------------+----------+----------+
| ID |      Task      |   RemindTime   | Reminder | Priority |
+----+----------------+----------------+----------+----------+
|  1 | should task    | 2021/3/3 12:00 |          |       10 |
|  2 | important task | 2022/1/1 00:00 |          |       50 |
|  3 | boring task    |                |          |      100 |
+----+----------------+----------------+----------+----------+

$ todo conf --hide_reminder=true --show_config
taskfilepath: /home/dondakeshimo/.local/share/todo/todo.json
hidepriority: false
hidereminder: true

$ todo l
+----+----------------+----------------+----------+
| ID |      Task      |   RemindTime   | Priority |
+----+----------------+----------------+----------+
|  1 | should task    | 2021/3/3 12:00 |       10 |
|  2 | important task | 2022/1/1 00:00 |       50 |
|  3 | boring task    |                |      100 |
+----+----------------+----------------+----------+
```

### Reminder
:warning: **reminder feature is only supported by macos**

You can choose reminder from macos or slack.

#### macos
You don't have to any configuration.
You just add a task with option `-r=macos`.

```bash
$ todo a "remind me this task" -r=macos -d=+1m

$ todo l
+----+---------------------+-----------------+----------+----------+
| ID |        Task         |   RemindTime    | Reminder | Priority |
+----+---------------------+-----------------+----------+----------+
|  1 | remind me this task | 2021/3/3 01:01  | macos    |      100 |
+----+---------------------+-----------------+----------+----------+
```

After 1 minute, you will get a message like bellow.

![notification](https://user-images.githubusercontent.com/23194960/126190791-be2dae4a-5e56-4e59-8151-a6d88e48f0e9.png)


#### slack
You have to configure Slack App.
Install Incomming Webhook to your workspace from [here](https://slack.com/apps) and configure information for todo-cli.
Member ID of slack shown by profile is set to `--slack_mention_to` option but it is not required.

```bash
$ todo conf --slack_webhook_url="https://hooks.slack.com/services/XXXXXXXX/XXXXXXXX" --slack_mention_to=XXXXXXXXXX
```

Then, add a task with `-r=slack` option.

```bash
$ todo a "remind me this task in slack\!" -r=slack -d=+1m

$ todo l
+----+-------------------------------+-----------------+----------+----------+
| ID |             Task              |   RemindTime    | Reminder | Priority |
+----+-------------------------------+-----------------+----------+----------+
|  1 | remind me this task in slack! | 2021/7/20 01:11 | slack    |      100 |
+----+-------------------------------+-----------------+----------+----------+
```

you will get a message in slack.

![notification in slack](https://user-images.githubusercontent.com/23194960/126192217-cee8469b-b917-4770-ab76-f604556bd3e2.png)


<!--
TODO: rewrite for cobra
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
source $(todo completion zsh)
```
-->

## Uninstall
```bash
$ make uninstall
```

If you installed todo-cli from downloading binary, please remove the binary by your hand.

```bash
$ rm path/to/your/todo/binary
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
