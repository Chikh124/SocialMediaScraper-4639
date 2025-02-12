Ось приклад простого коду на Go для обробки даних, цей код використовує структури для зберігання даних про студентів і функції для обробки цих даних:

```go
package main

import (
	"fmt"
	"sort"
	"strings"
)

// Структура для зберігання студентів
type Student struct {
	name string
	age  int
	gpa  float64
}

// Функція для створення нового студента
func NewStudent(name string, age int, gpa float64) Student {
	return Student{name, age, gpa}
}

// Функція для виводу інформації про студента
func (s Student) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, GPA: %.2f", s.name, s.age, s.gpa)
}

// Функція для порівняння двох студентів за ім'ям
type ByName []Student

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].name < n[j].name }

// Функція для порівняння двох студентів за GPA
type ByGPA []Student

func (g ByGPA) Len() int           { return len(g) }
func (g ByGPA) Swap(i, j int)      { g[i], g[j] = g[j], g[i] }
func (g ByGPA) Less(i, j int) bool { return g[i].gpa < g[j].gpa }

func main() {
	// Список студентів
	students := []Student{
		NewStudent("John", 20, 3.5),
		NewStudent("Alice", 19, 3.8),
		NewStudent("Bob", 21, 3.2),
	}

	// Виведення інформації про студентів
	fmt.Println("Original List:")
	for _, s := range students {
		fmt.Println(s)
	}

	// Сортування студентів за ім'ям
	sort.Sort(ByName(students))
	fmt.Println("\nSorted by name:")
	for _, s := range students {
		fmt.Println(s)
	}

	// Сортування студентів за GPA
	sort.Sort(ByGPA(students))
	fmt.Println("\nSorted by GPA:")
	for _, s := range students {
		fmt.Println(s)
	}

  // Пошук студентів за іменем
  name := "Alice"
  fmt.Println("\nSearching for student by name:")
  for _, s := range students {
    if strings.Compare(s.name, name) == 0 {
      fmt.Println(s)
    }
  }
}
```

Цей код створює структуру `Student` для зберігання даних про студентів, а також функції для створення нових студентів, виводу інформації про студентів, сортування студентів за ім'ям або GPA, і пошуку студентів за іменем. Програма створює список студентів, сортує цей список і виводить результат.