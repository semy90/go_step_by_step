package main

/*
рсений решил помочь Яндекс Лицею автоматизировать сбор статистики по ученикам, ведь их много, и вручную всех не пересчитаешь. Нужно понять, кто набрал достаточно баллов для перехода на следующий модуль, а кто ещё нет. Помогите Арсению написать такую программу.

Создайте структуру Student с полями name (string), solvedProblems — количество решённых задач (int), scoreForOneTask — количество баллов за одну задачу (float64) и passingScore — проходной балл в следующий модуль (float64). Будем считать, что все задачи дают одинаковые баллы.

Создайте метод IsExcellentStudent для этой структуры, который возвращает true, если ученик проходит по баллам в следующий модуль, и false в ином случае.
*/
type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() bool {
	return s.scoreForOneTask*float64(s.solvedProblems) >= s.passingScore
}
