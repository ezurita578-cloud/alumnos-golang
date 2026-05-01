package main

import "fmt"

// Materia
type Subject struct {
	Name  string
	Grade float64
}

// Alumno
type Student struct {
	Name     string
	Subjects map[string]Subject
}

// Agregar materia
func (s *Student) AddSubject(name string, grade float64) {
	if s.Subjects == nil {
		s.Subjects = make(map[string]Subject)
	}
	s.Subjects[name] = Subject{
		Name:  name,
		Grade: grade,
	}
}

// Promedio
func (s Student) Average() float64 {
	total := 0.0
	for _, sub := range s.Subjects {
		total += sub.Grade
	}
	if len(s.Subjects) == 0 {
		return 0
	}
	return total / float64(len(s.Subjects))
}

// Mostrar info
func (s Student) Print() {
	fmt.Println("Alumno:", s.Name)
	for _, sub := range s.Subjects {
		fmt.Println("-", sub.Name, ":", sub.Grade)
	}
	fmt.Println("Promedio:", s.Average())
}

// Aprobado
func (s Student) IsApproved() bool {
	return s.Average() >= 70
}

// Buscar materia
func (s Student) FindSubject(name string) {
	sub, ok := s.Subjects[name]
	if ok {
		fmt.Println("Materia encontrada:", sub.Name, "| Calificacion:", sub.Grade)
	} else {
		fmt.Println("Materia no encontrada")
	}
}

// Eliminar materia
func (s *Student) RemoveSubject(name string) {
	_, ok := s.Subjects[name]
	if ok {
		delete(s.Subjects, name)
		fmt.Println("Materia eliminada:", name)
	} else {
		fmt.Println("Materia no encontrada")
	}
}

// Actualizar calificacion
func (s *Student) UpdateGrade(name string, grade float64) {
	sub, ok := s.Subjects[name]
	if ok {
		sub.Grade = grade
		s.Subjects[name] = sub
		fmt.Println("Calificacion actualizada:", name, "->", grade)
	} else {
		fmt.Println("Materia no encontrada")
	}
}

// Materia con mayor calificacion
func (s Student) TopSubject() {
	best := Subject{}
	for _, sub := range s.Subjects {
		if sub.Grade > best.Grade {
			best = sub
		}
	}
	fmt.Println("Mejor materia:", best.Name, "| Calificacion:", best.Grade)
}

func main() {
	student := Student{Name: "Carlos"}
	student.AddSubject("Matematicas", 90)
	student.AddSubject("Programacion", 95)
	student.AddSubject("Fisica", 80)

	student.Print()

	if student.IsApproved() {
		fmt.Println("Alumno aprobado")
	} else {
		fmt.Println("Alumno reprobado")
	}

	student.FindSubject("Fisica")
	student.RemoveSubject("Fisica")
	student.UpdateGrade("Matematicas", 85)
	student.TopSubject()
}
