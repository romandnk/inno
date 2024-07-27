package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

// Загрузчик файлов, который может загружать файлы с нескольких URL-
// адресов. Есть таймауты, лимиты на воркеров
// Шаблоны: worker pool, fan-in, timeout

const pathToSave string = "./"

const workerNum int = 5

const timeout time.Duration = time.Second * 2

func main() {
	apiUrls := []string{
		"https://earthly.dev/blog/assets/images/golang-csv-files/jSkrisz.png",
		"https://www.adobe.com/creativecloud/file-types/image/raster/media_1655923d8b0386180ca09391e8654c4e468090bb3.jpeg?width=1200&format=pjpg&optimize=medium",
		"https://itsm365.com/documents_rus/web/Content/Resources/doc/import_empl_csv.csv",
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
			return Result{Error: fmt.Errorf("cl.Do [%s] error: %w", apiUrl, err)}
		}
		defer resp.Body.Close()

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(resp.Body)
		if err != nil {
			return Result{Error: fmt.Errorf("buf.ReadFrom [%s] error: %w", apiUrl, err)}
		}

		return Result{Url: apiUrl, Data: buf.Bytes()}
	}
	wp := NewWorkerPool(uint64(workerNum), input)
	wp.Run(result, fn)

	for res := range result {
		if res.Error != nil {
			fmt.Println(res.Error)
			continue
		}

		fileName, err := makeFileNameFromUrl(res.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = saveFile(fileName, res.Data)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func makeFileNameFromUrl(apiUrl string) (string, error) {
	u, err := url.Parse(apiUrl)
	if err != nil {
		return "", err
	}

	fileName := path.Base(u.Path)

	return fileName, nil
}

func saveFile(fileName string, data []byte) error {
	f, err := createFile(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("write file [%s] error: %w", fileName, err)
	}
	return nil
}

func createFile(fileName string) (*os.File, error) {
	f, err := os.Create(pathToSave + fileName)
	if err != nil {
		return nil, fmt.Errorf("error creating file [%s]: %w", fileName, err)
	}
	return f, nil
}
