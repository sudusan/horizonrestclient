package main

import "time"

// Todo defining a Todo structure
type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos defining a Todos structure
type Todos []Todo
