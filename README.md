# todo-cli  <!-- exclude-toc -->
Manage TODO List at CLI

[![Go][go-test-image]][go-test-url]

[go-test-image]: https://github.com/dondakeshimo/todo-cli/workflows/Go/badge.svg
[go-test-url]: https://github.com/dondakeshimo/todo-cli/actions?query=workflow%3AGo

## Table of Contents  <!-- exclude-toc -->
  * [Why todo-cli](#sec0-1-0-0)
  * [Install](#sec0-2-0-0)
      * [go installation](#sec0-2-0-1)
      * [download binary](#sec0-2-0-2)
  * [Usage](#sec0-3-0-0)
    * [Reminder](#sec0-3-1-0)
      * [macos](#sec0-3-1-1)
      * [slack](#sec0-3-1-2)
    * [Completion](#sec0-3-2-0)
  * [Uninstall](#sec0-4-0-0)
  * [Contributing](#sec0-5-0-0)
  * [License](#sec0-6-0-0)

*created by [tocot](https://pypi.org/project/tocot/)*

<a id="sec0-1-0-0"></a>
## Why todo-cli
- simple and light
- remind the task with MacOS popup or slack
- only show tasks you want to check

<a id="sec0-2-0-0"></a>
## Install
<a id="sec0-2-0-1"></a>
#### go installation
This is a simple way, but require [golang](https://golang.org/) .

Make sure that you have already added binary path to your PATH.

```
$ export PATH=$PATH:$(go env GOPATH)/bin
```

##### Go version \< 1.16
```
$ go get -u github.com/dondakeshimo/todo-cli/cmd/todo
```

##### Go 1.16+
```
$ go install github.com/dondakeshimo/todo-cli/cmd/todo@latest
```

<a id="sec0-2-0-2"></a>
#### download binary
You can download binary from our repository.
Bellow example is for MaxOS.

```
$ TODO_VERSION=1.0.0

$ curl -O https://github.com/dondakeshimo/todo-cli/releases/download/v${TODO_VERSION}/todo-${TODO_VERSION}.macos-10.15.tar.gz

$ tar -xvf todo-${TODO_VERSION}.macos-10.15.tar.gz

$ mv todo path/to/your/$PATH
```

<a id="sec0-3-0-0"></a>
## Usage

```
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

```
$ todo list
+----+--------------------------------+----------------+-------+----------+----------+
| ID |              Task              |   RemindTime   | Group | Reminder | Priority |
+----+--------------------------------+----------------+-------+----------+----------+
|  1 | deleting or modifying this     | 2099/1/1 00:00 |       |          |        0 |
|    | task is your first TODO        |                |       |          |          |
+----+--------------------------------+----------------+-------+----------+----------+

$ todo close -i=1

$ todo l
+----+------+------------+-------+----------+----------+
| ID | Task | RemindTime | Group | Reminder | Priority |
+----+------+------------+-------+----------+----------+
+----+------+------------+-------+----------+----------+

$ todo add "must task" -d="2021/03/03 12:00" -g="project x" -r=slack -p=0

$ todo l
+----+-----------+----------------+-----------+----------+----------+
| ID |   Task    |   RemindTime   |   Group   | Reminder | Priority |
+----+-----------+----------------+-----------+----------+----------+
|  1 | must task | 2021/3/3 12:00 | project x | slack    |        0 |
+----+-----------+----------------+-----------+----------+----------+

$ todo a "boring task" -g="project x"

$ todo a "important task" -d="2022/01/01" -p=50

$ todo l
+----+----------------+----------------+-----------+----------+----------+
| ID |      Task      |   RemindTime   |   Group   | Reminder | Priority |
+----+----------------+----------------+-----------+----------+----------+
|  1 | must task      | 2021/3/3 12:00 | project x | slack    |        0 |
|  2 | important task | 2022/1/1 00:00 |           |          |       50 |
|  3 | boaring task   |                | project x |          |      100 |
+----+----------------+----------------+-----------+----------+----------+

$ todo m -i=1 -t="should task" -p=10

$ todo l -g="project x"
+----+--------------+----------------+-----------+----------+----------+
| ID |     Task     |   RemindTime   |   Group   | Reminder | Priority |
+----+--------------+----------------+-----------+----------+----------+
|  1 | should task  | 2021/3/3 12:00 | project x | slack    |       10 |
|  3 | boaring task |                | project x |          |      100 |
+----+--------------+----------------+-----------+----------+----------+

$ todo conf --hide_reminder=true --show_config
columnwidth: 30
hidegroup: false
hidepriority: false
hidereminder: true
slackmentionto: XXXXXXXXXX
slackwebhookurl: https://hooks.slack.com/services/XXXXXXXXXX/XXXXXXXXXX/XXXXXXXXXX
taskfilepath: /home/dondakeshimo/.local/share/todo/todo.json

$ todo l
+----+----------------+----------------+-----------+----------+
| ID |      Task      |   RemindTime   |   Group   | Priority |
+----+----------------+----------------+-----------+----------+
|  1 | should task    | 2021/3/3 12:00 | project x |       10 |
|  2 | important task | 2022/1/1 00:00 |           |       50 |
|  3 | boaring task   |                | project x |      100 |
+----+----------------+----------------+-----------+----------+
```

<a id="sec0-3-1-0"></a>
### Reminder
You can choose reminder from macos or slack.

If you use linux OS, reminder feature needs cron daemon.

<a id="sec0-3-1-1"></a>
#### macos
You don't have to any configuration.
You just add a task with option `-r=macos`.

```
$ todo a "remind me this task" -r=macos -d=+1m

$ todo l
+----+---------------------+-----------------+-------+----------+----------+
| ID |        Task         |   RemindTime    | Group | Reminder | Priority |
+----+---------------------+-----------------+-------+----------+----------+
|  1 | remind me this task | 2021/3/3 01:01  |       | macos    |      100 |
+----+---------------------+-----------------+-------+----------+----------+
```

After 1 minute, you will get a message like bellow.

![notification](https://user-images.githubusercontent.com/23194960/126190791-be2dae4a-5e56-4e59-8151-a6d88e48f0e9.png)

If you push `done` button, the task will be closed.


<a id="sec0-3-1-2"></a>
#### slack
You have to configure Slack App.
Install Incomming Webhook to your workspace from [here](https://slack.com/apps) and configure information for todo-cli.
Member ID of slack shown by profile is set to `--slack_mention_to` option but it is not required.

```
$ todo conf --slack_webhook_url="https://hooks.slack.com/services/XXXXXXXX/XXXXXXXX" --slack_mention_to=XXXXXXXXXX
```

Then, add a task with `-r=slack` option.

```
$ todo a "remind me this task in slack\!" -r=slack -d=+1m

$ todo l
+----+-------------------------------+-----------------+-------+----------+----------+
| ID |             Task              |   RemindTime    | Group | Reminder | Priority |
+----+-------------------------------+-----------------+-------+----------+----------+
|  1 | remind me this task in slack! | 2021/7/20 01:11 |       | slack    |      100 |
+----+-------------------------------+-----------------+-------+----------+----------+
```

you will get a message in slack.

![notification in slack](https://user-images.githubusercontent.com/23194960/126192217-cee8469b-b917-4770-ab76-f604556bd3e2.png)


<a id="sec0-3-2-0"></a>
### Completion
todo-cli is made with cobra, then cobra supports generating completion files.
look at [cobra document](https://github.com/spf13/cobra/blob/master/shell_completions.md).

<a id="sec0-4-0-0"></a>
## Uninstall
```
$ make uninstall
```

If you installed todo-cli from downloading binary, please remove the binary by your hand.

```
$ rm path/to/your/todo/binary
```

<a id="sec0-5-0-0"></a>
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

<a id="sec0-6-0-0"></a>
## License
[MIT](https://choosealicense.com/licenses/mit/)
