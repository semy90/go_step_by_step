package main

/*Реализуйте функцию:

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error)

Каждый элемент из inputs должен быть обработан функцией fn в отдельной горутине (не более workers одновременно).
Функция fn принимает элемент из inputs и возвращает результат обработки этого элемента.
Если контекст ctx завершён (например, из-за таймаута), функция должна немедленно вернуть ошибку ctx.Err().
Результат — слайс, где на i-м месте результат обработки inputs[i] через функцию fn.
Реализация должна использовать пул горутин (workers). Последовательные решения не принимаются.
Используйте sync.WaitGroup, каналы и context.*/
import (
	"context"
	"sync"
)

type In struct {
	Id     int
	InData int
}

type Out struct {
	Id  int
	Res int
}

func ParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) ([]int, error) {
	ch, err := TmpParallelMapCtx(ctx, inputs, fn, workers)
	output := make([]int, len(inputs))
	if err != nil {
		return nil, err
	}
	for el := range ch {
		output[el.Id] = el.Res
	}
	return output, nil
}

func TmpParallelMapCtx(ctx context.Context, inputs []int, fn func(int) int, workers int) (chan Out, error) {
	jobs := make(chan In, len(inputs))
	res := make(chan Out, len(inputs))
	done := make(chan any)
	go func() {
		wg := new(sync.WaitGroup)
		for i := 0; i < workers; i++ {
			wg.Add(1)
			go func(id int, ctx context.Context, jobs chan In, res chan Out) {
				defer wg.Done()
				for job := range jobs {
					ans := fn(job.InData)
					// fmt.Println(i, job.InData, ans)
					res <- Out{Id: job.Id, Res: ans}
				}
			}(i, ctx, jobs, res)
		}
		wg.Wait()
		close(done)
	}()

	for i := 0; i < len(inputs); i++ {
		jobs <- In{Id: i, InData: inputs[i]}
	}
	close(jobs)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-done:
		close(res)
		return res, nil
	}

}
