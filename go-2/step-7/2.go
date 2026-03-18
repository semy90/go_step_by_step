package main

/*Напишите программу, имитирующую портал викторин с несколькими вопросами и правильными ответами на них. Вопросы будут приходить один за другим. Время на ответ — 1 секунда. Если время истекло, переходите к следующему вопросу. В конце покажите количество правильных ответов. Если вы прислали ответ не в том регистре (например "GO" вместо "go") ответ все равно будет защитан. Сигнатура проверяемой функции:

func QuizRunner(questions, answers []string, answerCh chan string) int
questions — список вопросов answers — список ответов answerCh — канал с пользовательским вводом

Результат — число правильных ответов.

Пример 1

Список вопросов:= "Q1", "Q2", "Q3"
Список ответов:= "A1", "A2", "A3"
В канал будем писать ответы:
answerCh <- "A1"
answerCh <- "A2"
answerCh <- "A3"

получаем 3 правильных ответа
Пример 2

Список вопросов := "Q1", "Q2", "Q3", "Q4"
Список ответов: "A1", "A2", "A3", "A4"
В канал будем писать ответы:
answerCh <- "A1". 1 правильный ответ
ждем 1.2 секунды. Не успели ответить на 2-ой вопрос
answerCh <- "A3". 2 правильных ответа
ждем 0.5 секунды. Ответить на второй вопрос успеваем
answerCh <- "A4". 3 правильных ответа

получаем 3 правильных ответа*/
import (
	"strings"
	"time"
)

func Check(s1, s2 string) bool {
	return strings.ToLower(s1) == strings.ToLower(s2)
}

func QuizRunner(questions, answers []string, answerCh chan string) int {
	curr := 0
	count := 0
	for range questions {
		select {
		case answ := <-answerCh:

			if Check(strings.TrimSpace(answ), answers[curr]) {
				count++
			}
			curr++
		case <-time.After(time.Second):
			curr++
		}
	}
	return count
}
