package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Загрузчик файлов, который может загружать файлы с нескольких URL-
// адресов. Есть таймауты, лимиты на воркеров
// Шаблоны: worker pool, fan-in, timeout

const workerNum int = 5

const timeout time.Duration = time.Second * 2

func main() {
	apiUrls := []string{
		"https://jsonplaceholder.typicode.com/posts/1",
		"https://jsonplaceholder.typicode.com/posts/2",
		"https://jsonplaceholder.typicode.com/posts/3",
		"https://jsonplaceholder.typicode.com/posts/4",
		"https://jsonplaceholder.typicode.com/posts/5",
		"https://jsonplaceholder.typicode.com/posts/6",
	}

	input := make(chan string, len(apiUrls))
	go func() {
		for _, apiUrl := range apiUrls {
			input <- apiUrl
		}
		close(input)
	}()

	result := make(chan Result, len(apiUrls))
	fn := func(apiUrl string) Result {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiUrl, nil)
		if err != nil {
			return Result{Error: fmt.Errorf("http.NewRequestWithContext error: %w", err)}
		}

		// как я понимаю timeout здесь работает как и context выше
		cl := http.Client{
			Timeout: timeout,
		}

		resp, err := cl.Do(req)
		if err != nil {
			// тут в ошибках еще желательно отслеживать при каком запросе произошла ошибка,
			// но я упущу это в этом программе)))
			// например, через уникальный идентификатор
			return Result{Error: fmt.Errorf("cl.Do error: %w", err)}
		}
		defer resp.Body.Close()

		data := struct {
			Title string `json:"title"`
		}{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return Result{Error: fmt.Errorf("json.NewDecoder.Decode error: %w", err)}
		}

		return Result{Result: data.Title}
	}
	wp := NewWorkerPool(uint64(workerNum), input)
	wp.Run(result, fn)

	for res := range result {
		if res.Error != nil {
			// тут тоже, что и выше
			fmt.Println(res.Error)
			continue
		}
		fmt.Println(res.Result)
	}
}
