# tri-nit-2024
In this project we develop an interactive online platform that connects language learners with
teachers for personalized online lessons by offering a platform where learners can choose a
tutor based on their target language, fluency and budget.

To get started, clone this repositry in your local machine.

## Pre-requisite
You need the following to run the web application. <br>
Postgresql <br>
Go 

## Setting Up back-end
Inside the `../backend` directory, the following commands need to be run in terminal to download required Go packages.
Install Chi router
```
go get -u github.com/go-chi/chi/v5
```
Install `goose` to run migrations into database.
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
Import a postgresql driver
```
go get github.com/lib/pq
```
Configure postgres as follows:
username: postgres <br>
password: dummypass <br>
host: localhost <br>
port: 5432 <br>
database: lingua <br>
`cd` into `/sql/schema` and run the following code in terminal to set-up the database.
```
goose postgres postgres://postgres:dummypass@localhost:5432/lingua up
```
Bacl in `../backend` directory, run the below command in terminal.
```
go build && ./backend
```
If no error is shown, the backend server is up and running in port: 5000
