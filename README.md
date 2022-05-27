# GOBLIN

The Backend repo for the PlayStation developer club website. Written in go.

## HOW TO RUN

Clone this repo from Github (obviously) and you should be at the project's root directory. You can navigate to the
project's root directory if you're not there.

Add the `config.env` under the `config` folder in the project.

Automatically, your go modules should be installed, if not, you can manually install it by running `go mod download`
or `go get`.

Now, from the root directory, run `go run main.go`. Your API service should be running on port `5000`.