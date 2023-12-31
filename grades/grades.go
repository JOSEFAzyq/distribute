package grades

import (
	"fmt"
	"sync"
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

func (ss Student) Average() float32 {
	var result float32
	for _, grade := range ss.Grades {
		result += grade.Score
	}
	return result / float32(len(ss.Grades))
}

type Students []Student

var (
	students      Students
	studentsMutex sync.Mutex
)

func (ss Students) GetById(id int) (*Student, error) {
	for i := range ss {
		if ss[i].ID == id {
			return &ss[i], nil
		}
	}
	return nil, fmt.Errorf("student with ID %d not found", id)
}

const (
	GradeQuiz = GradeType("Quiz")
	GradeTest = GradeType("Test")
	GradeExam = GradeType("Exam")
)

type GradeType string
type Grade struct {
	Title string
	Type  GradeType
	Score float32
}
