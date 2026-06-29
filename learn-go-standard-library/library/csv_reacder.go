package library

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func CSVReader() {
	csvstring := "Eko, urniawan, jhanedy\n" +
		"budim bagus, setiawan\n" +
		"john, depal, sontoloyo"
	reader := csv.NewReader(strings.NewReader(csvstring))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"eko", "Kuniawan", "Khanaedy"})
	_ = writer.Write([]string{"budi", "setiawan", "joko"})
	_ = writer.Write([]string{"agus", "salim", "wahid"})
	writer.Flush()
}
