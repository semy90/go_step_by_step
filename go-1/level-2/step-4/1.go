package main

/*
Программист Арсений так увлекается кодингом на Go, что совсем не следит за временем. День от ночи он ещё более-менее отличает, когда изредка выглядывает в окно, но определять дату на глаз у него не получается.
Помогите ему и напишите программу, которая поможет разработчику быстро сориентироваться во времени.
Для этого реализуйте такие функции:
currentDayOfTheWeek() string, которая возвращает текущий день недели (Понедельник, Вторник, Среда, Четверг, Пятница, Суббота, Воскресенье)
dayOrNight() string, которая возвращает, День сейчас или Ночь. День для программы начинается в 10 утра, а заканчивается в 22 вечера включительно (Арсений, как и многие другие программисты, — сова)
nextFriday() int, которая вычисляет, сколько дней до следующей пятницы, включая сегодняшний день (если сегодня пятница, то 0). Арсений с нетерпением ждёт выходных, чтобы наконец выбраться на улицу
CheckCurrentDayOfTheWeek(answer string) bool, которая проверяет, знает ли Арсений, какой сейчас день недели. Она должна принимать строку с предположением Арсения и выводить true, если Арсений угадал, и false в ином случае. Регистр ввода учитывать не нужно: Понедельник, понедельник и ПОНЕДЕЛЬНИК должны возвращать true (если, конечно, сегодня действительно понедельник)
CheckNowDayOrNight(answer string) (bool, error), которая проверяет, знает ли Арсений, день сейчас или ночь. Она должна принимать строку с предположением Арсения и выводить true или false. Регистр также учитывать не нужно. Если на вход подана строка длиной больше или меньше четырёх символов, нужно возвращать пустую строку и ошибку: исправь свой ответ, а лучше ложись поспать
Чтобы получить текущее время, используйте функцию TimeNow(). Она вернёт время в формате time.Time. Саму функцию TimeNow реализовывать не нужно.

Формат ввода
Функции currentDayOfTheWeek() string, dayOrNight() string, nextFriday() int, CheckCurrentDayOfTheWeek(answer string) bool и СheckNowDayOrNight(answer string) (bool, error)
*/
import (
	"errors"
	"strings"
	"unicode/utf8"
)

func currentDayOfTheWeek() string {
	week := TimeNow().Weekday()
	switch week {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	default:
		return ""
	}
}

func dayOrNight() string {
	hour := TimeNow().Hour()
	if hour >= 10 && hour <= 22 {
		return "День"
	} else {
		return "Ночь"
	}
}

func nextFriday() int {
	day := int(TimeNow().Weekday())
	if day <= 5 {
		return 5 - day
	} else {
		return 7 - day + 5
	}
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	curDay := strings.ToLower(currentDayOfTheWeek())
	ans := strings.ToLower(answer)
	return curDay == ans
}

func CheckNowDayOrNight(answer string) (bool, error) {
	curDayOrNight := strings.ToLower(dayOrNight())
	ans := strings.ToLower(answer)
	if utf8.RuneCountInString(ans) == 4 && ans == curDayOrNight {
		return true, nil
	} else if utf8.RuneCountInString(ans) == 4 && ans != curDayOrNight {
		return false, nil
	}
	return false, errors.New("исправь свой ответ, а лучше ложись поспать")

}
