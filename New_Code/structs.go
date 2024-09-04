package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	ID     int
	Name   string
	Email  string
	Rollno int
}

var students []Student
var nextID = 1

// Student -> Create a New Student
func CreateStudent(name, email string, rollno int) Student {

	newStudent := Student{
		ID:     nextID,
		Name:   name,
		Email:  email,
		Rollno: rollno,
	}

	nextID++
	students = append(students, newStudent)

	return newStudent
}

// Student -> Display all the student in the slice
func ReadStudents() []Student {

	return students
}

// Student -> Update the student using their Student id
func UpdateStudent(id int, newName, newEmail string, newRollNo int) bool {

	for i := range students {
        if students[i].ID == id {
            s := &students[i]  // Take a pointer to the slice element
            s.Email = newEmail
            s.Name = newName
            s.Rollno = newRollNo

			fmt.Println(s)

            return true
        }
    }

	return false

}

// Student -> Delete each student using their Student id
func DeleteStudent(id int) bool {
	for i, s := range students {
		if s.ID == id {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}

	return false
}

func PrintChoices() {

	fmt.Println("Index")
	fmt.Println("1. Create Student")
	fmt.Println("2. Display all students")
	fmt.Println("3. Update a student")
	fmt.Println("4. Delete a student")
	fmt.Println("5. Exit")

}


func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Student Application")
	fmt.Println("Let's do some operations in our application")
	PrintChoices()
	
	for {
		fmt.Print("\nEnter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter Roll no: ")
			tid, _ := reader.ReadString('\n')
			tid = strings.TrimSpace(tid)
			rollno, _ := strconv.ParseInt(tid, 10, 64)

			fmt.Print("Enter Student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter Student email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			CreateStudent(name, email, int(rollno))
			fmt.Println("Student created successfully!")

			PrintChoices()

		case "2":
			students := ReadStudents()

			if len(students) == 0 {
				fmt.Println("No students were added yet !!")
				continue
			}

			fmt.Println("All Students:")
			for _, student := range students {
				fmt.Printf("ID: %d, Name: %s, Email: %s, Roll No: %d\n", student.ID, student.Name, student.Email, student.Rollno)
			}

			PrintChoices()

		case "3":
			fmt.Print("Enter Student id: ")
			sid, _ := reader.ReadString('\n')
			sid = strings.TrimSpace(sid)
			id, _ := strconv.Atoi(sid)

			fmt.Print("Enter new Roll no: ")
			tid, _ := reader.ReadString('\n')
			tid = strings.TrimSpace(tid)
			rollno, _ := strconv.ParseInt(tid, 10, 64)

			fmt.Print("Enter new Student name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Enter new Student email: ")
			email, _ := reader.ReadString('\n')
			email = strings.TrimSpace(email)

			if UpdateStudent(id, name, email, int(rollno)) {
				fmt.Println("Student updated successfully!")
			} else {
				fmt.Println("Student not found.")
			}

			PrintChoices()

		case "4":
			fmt.Print("Enter Student id: ")
			sid, _ := reader.ReadString('\n')
			sid = strings.TrimSpace(sid)
			id, _ := strconv.Atoi(sid)

			if DeleteStudent(id) {
				fmt.Println("Student deleted successfully!")
			} else {
				fmt.Println("Student not found.")
			}

			PrintChoices()

		case "5":
			fmt.Println("Thank you for using our application...!!")
			return // Exits the loop and the program

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
