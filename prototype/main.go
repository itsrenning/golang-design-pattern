package main

import "fmt"

func main() {
	e1 := Employee{
		Name: "e1",
		Job: &Job{
			Title:     "Backend Engineer",
			YearOfExp: 3,
		},
		Colleague: []string{"c1", "c2"},
	}
	fmt.Println("Create e1 employee info")
	e1.Print()
	fmt.Println()

	fmt.Println("Create e2 employee info with shallow clone of e1")
	e2 := e1.ShallowClone()
	e2.Print()
	fmt.Println()

	fmt.Println("Update e2 employee info and then compared with e1")
	e2.Name = "e2"
	e2.Job.Title = "DBA"
	e2.Job.YearOfExp = 1
	e2.Colleague = append(e2.Colleague, "c3")
	e1.Print()
	e2.Print()
	fmt.Println()

	fmt.Println("Create e3 employee info with deep clone of e1")
	e3 := e1.DeepClone()
	e3.Print()
	fmt.Println()

	fmt.Println("Update e3 employee info and then compared with e1")
	e3.Name = "e3"
	e3.Job.Title = "DevOps"
	e3.Job.YearOfExp = 2
	e3.Colleague = append(e3.Colleague, "c3")
	e1.Print()
	e3.Print()
	fmt.Println()

	fmt.Println("Create e4 employee info with better deep clone of e1")
	e4 := e1.DeepClone()
	e4.Print()
	fmt.Println()

	fmt.Println("Update e4 employee info and then compared with e1")
	e4.Name = "e4"
	e4.Job.Title = "Frontend Engineer"
	e4.Job.YearOfExp = 3
	e4.Colleague = append(e3.Colleague, "c4")
	e1.Print()
	e4.Print()
	fmt.Println()
}
