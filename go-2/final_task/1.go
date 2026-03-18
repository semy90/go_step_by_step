package main

/*Помните, как вы пришли проджект-менеджером в команду, работавшую над обречённым на провал проектом без чёткого списка задач? Поздравляем: ваш скрипт сработал и вы смогли наладить процесс, собрать бэклог и найти ответственных. Пока вы писали предыдущий код, вам очень понравилось программировать. К тому же вы молоды, амбициозны и прошли курс по Golang.

Теперь вам обратился друг, попавший в такую же ситуацию; он помнил, что вы говорили о готовом решении, и попросил помочь. Вы соглашаетесь, но, будучи конкретным человеком, просите поставить вам задачу в систему и уходите на выходные. В понедельник вы приходите на работу, наливаете себе чашечку кофе, открываете задачу и читаете:

TASK_BETS_PROJECT_UPLOADER-1 Перепишите функцию из вступительного задания следующим образом

type Ticket struct {
    Ticket string
    User   string
    Status string
    Date   time.Time
}

func GetTasks(
    ctx context.Context,
    r io.Reader,
    w io.Writer,
    user *string,
    status *string,
    timeout time.Duration,
) error
Раньше вы получали переменную text типа string в качестве входных данных, теперь данные будут передаваться в функцию через r io.Reader. Ответ будем писать в w io.Writer в формате json. Добавим возможность остановить выполнение функции прерыванием контекста ctx context.Context или по истечению таймаута timeout time.Duration. Текст ошибки можете определить на свое усмотрение.*/
import (
	"context"
	"encoding/json"
	"io"

	"strings"
	"time"
)

type Ticket struct {
	Ticket string
	User   string
	Status string
	Date   time.Time
}

func ParseFromR(ctx context.Context, user *string, status *string, r io.Reader) ([]Ticket, error) {
	bytetext, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	raw := strings.Split(string(bytetext), "\n")
	var answ []Ticket

	for _, el := range raw {
		if el == "" {
			continue
		}

		tmp := strings.Split(el, "_")
		if len(tmp) != 4 || !strings.HasPrefix(tmp[0], "TICKET-") {
			continue
		}
		time1, err := time.Parse(time.DateOnly, tmp[3])
		if err != nil {
			continue
		}
		if tmp[2] != "В работе" && tmp[2] != "Готово" && tmp[2] != "Не будет сделано" {
			continue
		}

		if user == nil {
			if status == nil {
				answ = append(answ, Ticket{
					Ticket: tmp[0],
					User:   tmp[1],
					Status: tmp[2],
					Date:   time1,
				})
			} else if *status == tmp[2] {
				answ = append(answ, Ticket{
					Ticket: tmp[0],
					User:   tmp[1],
					Status: tmp[2],
					Date:   time1,
				})
			}
		} else if *user == tmp[1] {
			if status == nil {
				answ = append(answ, Ticket{
					Ticket: tmp[0],
					User:   tmp[1],
					Status: tmp[2],
					Date:   time1,
				})
			} else if *status == tmp[2] {
				answ = append(answ, Ticket{
					Ticket: tmp[0],
					User:   tmp[1],
					Status: tmp[2],
					Date:   time1,
				})
			}
		}

	}
	return answ, nil
}

func SubGetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	done chan any,
	Err chan error,
) {

	tickets, err := ParseFromR(ctx, user, status, r)
	if err != nil {
		Err <- err
		return
	}

	var jsons []interface{}

	for _, el := range tickets {
		tmpjson := make(map[string]string)
		tmpjson["user"] = el.User
		tmpjson["status"] = el.Status
		tmpjson["ticket"] = el.Ticket
		tmpjson["date"] = el.Date.Format(time.RFC3339)
		jsons = append(jsons, tmpjson)
	}

	jsondyte, err := json.Marshal(jsons)
	if err != nil {
		Err <- err
		return
	}
	_, err = w.Write(jsondyte)
	if err != nil {
		Err <- err
		return
	}
	close(done)
}

func GetTasks(
	ctx context.Context,
	r io.Reader,
	w io.Writer,
	user *string,
	status *string,
	timeout time.Duration,
) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	done := make(chan any)
	Err := make(chan error)
	go SubGetTasks(ctx, r, w, user, status, done, Err)
	select {
	case <-done:
		return nil
	case err := <-Err:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
