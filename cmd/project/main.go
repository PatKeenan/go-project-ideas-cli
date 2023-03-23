package main

import (
	"flag"
	"fmt"
	"os"

	project "github.com/Patkeenan/go-project-ideas-cli"
)

var fileName = "./projects.json"

func main() {
	proj := flag.String("pro", "", "Name of project")
	pri := flag.String("pri", "", "Priority of project")
	cat := flag.String("cat", "", "Priority of project")
	list := flag.Bool("list", false, "List of projects")

	flag.Parse()

	l := &project.ProjectList{}

	if err := l.Get(fileName); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch {
	case *list:
		for index, p := range *l {
			fmt.Printf("\nProject %d: %s, Category: %s, Priority: %s\n", index+1, p.Name, p.Priority, p.Category)
		}
	case *proj != "":

		l.Add(*proj, *pri, *cat)

		if err := l.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stderr, "No command specified")
		os.Exit(1)
	}

}
