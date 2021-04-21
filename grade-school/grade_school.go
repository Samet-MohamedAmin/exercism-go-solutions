package school

import (
	"sort"
)

// Grade holds grades
type Grade struct {
	number   int
	students []string
}

// School holds enrollments
type School struct {
	grades map[int]*Grade
}

// New returns new school
func New() *School {
	return &School{map[int]*Grade{}}
}

// Add adds name
func (s *School) Add(student string, n int) {
	if _, ok := s.grades[n]; !ok {
		s.grades[n] = &Grade{n, []string{student}}
	} else {
		grade := s.grades[n]
		grade.students = append(grade.students, student)
		sort.Strings(grade.students)
	}
}

// Grade return array of strings
func (s *School) Grade(n int) []string {

	if _, ok := s.grades[n]; !ok {
		return []string{}
	}
	return s.grades[n].students
}

// Enrollment return array of grades
func (s *School) Enrollment() []Grade {
	grades := []Grade{}
	for _, g := range s.grades {
		grades = append(grades, *g)
	}
	sort.SliceStable(grades, func(i, j int) bool { return grades[i].number < grades[j].number })
	return grades
}
