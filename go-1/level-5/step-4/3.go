package main

/*Предприятие, на котором работает программист Арсений, крупное, там много сотрудников. У каждого из них своя должность, зарплата и стаж работы.

Напишите программу, в которой тип Company реализует такой интерфейс:

type CompanyInterface interface{
    AddWorkerInfo(name, position string, salary, experience uint) error
    SortWorkers() ([]string, error)
}
AddWorkerInfo — метод добавления информации о новых сотрудниках, где name — имя сотрудника, position — должность, salary — зарплата, experience — стаж работы (месяцев).

SortWorkers — метод, который сортирует список сотрудников по доходу за время работы на предприятии (по убыванию), должности (от высокой до низкой) и возвращает слайс формата: *имя* — *доход* — *должность*. Допустимые должности в порядке убывания: «директор», «зам. директора», «начальник цеха», «мастер», «рабочий».*/
import (
	"fmt"
	"sort"
)

var pos = []string{"директор", "зам. директора", "начальник цеха", "мастер", "рабочий"}

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Worker struct {
	Name       string
	Position   string
	Salary     uint
	Experience uint
}

type Company struct {
	workers []Worker
}

func (cmpn *Company) AddWorkerInfo(name, position string, salary, experience uint) error {

	f := false
	for _, pos1 := range pos {
		if pos1 == position {
			f = true
		}
	}
	if !f {
		return fmt.Errorf("не существующая позиция")
	}

	cmpn.workers = append(cmpn.workers, Worker{
		Name:       name,
		Position:   position,
		Salary:     salary,
		Experience: experience,
	})

	return nil
}

func linear_search(name string) int {
	for i := range pos {
		if pos[i] == name {
			return i
		}
	}
	return -1 // не используется если че
}

func (cmpn *Company) SortWorkers() ([]string, error) {
	sort.Slice(cmpn.workers, func(i, j int) bool {
		if cmpn.workers[i].Salary*cmpn.workers[i].Experience > cmpn.workers[j].Salary*cmpn.workers[j].Experience {
			return true
		}
		if cmpn.workers[i].Salary*cmpn.workers[i].Experience < cmpn.workers[j].Salary*cmpn.workers[j].Experience {
			return false
		}
		if cmpn.workers[i].Salary*cmpn.workers[i].Experience == cmpn.workers[j].Salary*cmpn.workers[j].Experience && linear_search(cmpn.workers[i].Name) < linear_search(cmpn.workers[j].Name) {
			return true
		}
		return false
	})
	ans := []string{}
	for _, el := range cmpn.workers {
		ans = append(ans, fmt.Sprintf("%s — %d — %s", el.Name, el.Salary*el.Experience, el.Position))
	}
	return ans, nil

}
