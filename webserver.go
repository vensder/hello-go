package main

import (
	"fmt"
	"log"
	"net/http"
)

const jenkinsHelp = "Use commands: \n/jenkins help (this help)\n/jenkins help deploy"

const jenkinsHelpDeploy = "> Deploy combinations of containers with custom *branches* or *tags* on `staging` or `demo`:\n" +
	"```repos: api, middleware, frontend```\n" +
	"`@jenkins deploy staging <repo1> <branch1> <repo2> <branch2>`\n" +
	"`@jenkins deploy staging api bugfix/test1 check-in bugfix/test2`\n\n" +

	"> Deploy combinations of containers with release (tag) numbers on `production` or `testing`:\n" +
	"```containers: api, middleware, frontend```\n" +
	"`@jenkins deploy production <container> <release>`\n" +
	"`@jenkins deploy production api 1.0.4 [frontend 1.0.5 middleware 0.1]`\n" +
	"`@jenkins deploy testing api 1.1 [middleware 0.2]`\n\n" +

	"> <_here next help commands or messages from slack slash command_>"

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go!")
}

func handlerSlack(w http.ResponseWriter, r *http.Request) {

	//Read the Request Parameter "command"
	command := r.FormValue("command")
	text := r.FormValue("text")

	//Ideally do other checks for Slack tokens/username/etc

	if command == "/jenkins" {
		switch text {
		case "help":
			fmt.Fprint(w, jenkinsHelp)
		case "help deploy":
			fmt.Fprint(w, jenkinsHelpDeploy)
		default:
			fmt.Fprint(w, jenkinsHelp)
		}
	} else {
		fmt.Fprint(w, "I do not understand your command.")
	}
}

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", handler)
	http.HandleFunc("/slack", handlerSlack)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
