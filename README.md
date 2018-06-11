# Golang Moneywatcher

Golang Moneywatcher is an application written in Go (https://golang.org/). With this web application you can monitor your expenses. 

# Installation

1. Install Go on your computer and set your $GOPATH
2. Run the command 'go get github.com/jonashellmann/golang-moneywatcher'
3. Change in the directory $GOPATH/src/github.com/jonashellmann/golang-moneywatcher
4. Open the file 'config.json' and change the values to fit your database. Therefore you need a user and a SQL database.
5. Run the following commands to start the application
	5.1 go build
	5.2 nohup ./moneywatcher > ~/moneywatcher.log 2>&1 &
