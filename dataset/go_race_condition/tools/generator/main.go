package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	sampleFd, err := os.Create("../../sample.csv")
	if err != nil {
		panic(err)
	}
	defer sampleFd.Close()

	sampleCsvWriter := csv.NewWriter(sampleFd)

	tplFd, err := os.Open("../../template.csv")
	if err != nil {
		panic(err)
	}
	defer tplFd.Close()

	tplCsvReader := csv.NewReader(tplFd)

	for {
		row, err := tplCsvReader.Read()
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		if errors.Is(err, io.EOF) || len(row) == 0 {
			break
		}

		if row[0] == "code" {
			if err := sampleCsvWriter.Write(row); err != nil {
				panic(err)
			}
			continue
		}

		for i := 0; i < 10000; i++ {
			newCode := row[0]
			newCode = strings.Replace(newCode, "{{placeholder_counter}}", gofakeit.Regex(`^[a-zA-Z][a-zA-Z0-9]+$`), -1)
			newCode = strings.Replace(newCode, "{{placeholder_wait_group}}", gofakeit.Regex(`^[a-zA-Z][a-zA-Z0-9]+$`), -1)
			newCode = strings.Replace(newCode, "{{placeholder_incr_func}}", gofakeit.Regex(`^[a-zA-Z][a-zA-Z0-9]+$`), -1)
			newCode = strings.Replace(newCode, "{{placeholder_decr_func}}", gofakeit.Regex(`^[a-zA-Z][a-zA-Z0-9]+$`), -1)
			newCode = strings.Replace(newCode, "{{placeholder_process_num_func}}", gofakeit.Regex(`^[a-zA-Z][a-zA-Z0-9]+$`), -1)
			newCode = strings.Replace(newCode, "{{placeholder_recycle_init}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)
			newCode = strings.Replace(newCode, "{{placeholder_recycle}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)
			newCode = strings.Replace(newCode, "{{placeholder_process_num1}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)
			newCode = strings.Replace(newCode, "{{placeholder_process_num2}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)
			newCode = strings.Replace(newCode, "{{placeholder_append_data1}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)
			newCode = strings.Replace(newCode, "{{placeholder_append_data2}}", fmt.Sprintf("%d", gofakeit.Int64()), -1)

			if err := sampleCsvWriter.Write([]string{newCode, row[1]}); err != nil {
				panic(err)
			}
		}
	}

	sampleCsvWriter.Flush()
}
