package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"strings"
)

type IEmployeeClone interface {
	ShallowClone() Employee
	DeepClone() *Employee
	BetterDeepClone() Employee
}

type Employee struct {
	Name      string
	Job       *Job
	Colleague []string
}

type Job struct {
	Title     string
	YearOfExp int
}

func (e *Employee) ShallowClone() Employee {
	clone := e
	return *clone
}

func (e *Employee) DeepClone() *Employee {
	clone := *e
	clone.Job = &Job{
		Title:     e.Job.Title,
		YearOfExp: e.Job.YearOfExp,
	}
	copy(clone.Colleague, e.Colleague)
	return &clone
}

func (e *Employee) BetterDeepClone() Employee {
	buffer := bytes.Buffer{}
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(e)
	if err != nil {
		panic(err)
	}
	decoder := gob.NewDecoder(&buffer)
	result := Employee{}
	err1 := decoder.Decode(&result)
	if err1 != nil {
		panic(err)
	}
	return result
}

func (e *Employee) Print() {
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("Employee Info:\n")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Job Title: %s\n", e.Job.Title)
	fmt.Printf("Year of Exp: %d\n", e.Job.YearOfExp)
	fmt.Printf("Colleague: %s\n", e.Colleague)
	fmt.Println(strings.Repeat("=", 40))
}
