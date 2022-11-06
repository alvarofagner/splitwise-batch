# Splitwise-cli-go
A splitwise API command line client written in go.

# Dev environment

This project is written in Go so it is necessary to follow the instruction for [Download and install Go](https://go.dev/doc/install). Alternativelly, one can use the ready docker environment which uses VS Code devcontainers described as follows.

## VS Code DevContainer environment

### Tools needed

* Visual Studio Code [https://code.visualstudio.com/Download](https://code.visualstudio.com/Download)
* Remote - Containers extension for VS Code: `ms-vscode-remote.remote-containers`

### Open the repo

The complete toolchain and environment needed to develop is defined in a [DevContainer](https://code.visualstudio.com/docs/remote/containers) in the `.devcontainer` folder.

1. Open the repo in VSCode
1. Hit `F1` and type/select `Remote-Containers: Reopen in Container`

# Authentication

Register an aplication [here](https://secure.splitwise.com/apps) in order to obtain an API Key.

Then make a copy of .credentials.template.json and name it as .credentials.json (or some other name but in the commands this is the default file name). Update the file with the created API Key.

# Commands

TODO: Create one main command "splitwise" and add the following commands to it. Also, build the project with mage and call go install.

## Categories

Lists all possible categories.

~~~
go run cmd/categories/main.go list
~~~

## User

### Current

Get data about the current user.

~~~
go run cmd/users/main.go current
~~~

## Friends

List data of all friends of current user.
~~~
go run cmd/friends/main.go list
~~~

## Groups

List data about all current user's groups.

~~~
go run cmd/groups/main.go list
~~~

## Expenses

### By shares

Create a list of expenses by shares. It requires a csv file in the format **date;description;cost**. So, for example, the contents of the following file are valid:

~~~
2022-11-05;Dinner at Restaurant;1,350.00
2022-11-06;Taxi to the bar;600.50
~~~

Example of the command:
~~~
run cmd/expenses/main.go byshares --from-file expenses.csv --group-id 39229928 --currency-code SEK --paid-by 27742729  --user-share 27742729:2,59044772:1
~~~

in this example, the user-share flag means that the user 27742729 will pay for 2 parts, while 59044772 will pay for one. User shares can also be passed like `--user-share 27742729:2 --user-share 59044772:1` and as many as it is required. 