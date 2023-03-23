# Go Project Idea Tracker CLI

### Overview

This is a simple for fun project to help me learn go. It's a command line tool that allows someone to easily add a project, assign a  category, and add a level of priority. The project will be saved in `project.json` file located in the `cmd/project` directory.

### Usage

Navigate to the project directory located in `cmd/project`.

1. Add a project

```zsh

go run main.go -pro "Write a cool command line tool" -pri "High" -cat "GO"

```

2. List project ideas

```zsh

go run main.go -list

```


