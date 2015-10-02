# minikick

Minikick is a simple cli version of kickstarter made for their coding challenge.

#Prerequisites 

1. Go version 1.5
2. mongoDB

# Instructions to build minikick 

1. `go get github.com/faiq/minikick`
2. `cd $GOPATH/src/github.com/faiq/minikick`
3. `go install`

Assuming you have `go/bin` in your `$PATH` you should now have an application called minikick.

If you want to use `minikick`, run `mongod` and you should be able run the commands.

```
NAME:
   minikick - Create and back some awesome projects

USAGE:
   minikick [global options] command [command options] [arguments...]
   
VERSION:
   0.0.0
   
COMMANDS:
   project, p	Create a new project! Just pass in a project name and target amount. (Dont use $ for the amount)
   back, b	Back a project! The arguments are name, project name, credit card number, and an amount.
   list, l	Display a project the backers and the amount they backed for a project!
   backer, br	Display a list of projects that a backer has backed and the amounts backed
   help, h	Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version
   

```
