package main

/*
Считается, что много повторений одних и тех же слов в одном тексте — плохо. К сожалению, программист Арсений нередко этим грешит.
Помогите ему и напишите программу, которая будет принимать на вход число слов в тексте (без знаков препинания и специальных символов), слово, повторения которого нужно подсчитать, и сам текст, а выводить количество повторений.
*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scanner.Text()

	scanner.Scan()
	sub := strings.ToLower(scanner.Text())

	scanner.Scan()
	str := strings.ToLower(scanner.Text())

	count := strings.Count(str, sub)
	fmt.Println(count)
}
