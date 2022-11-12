package part7

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCsvFiles(t *testing.T) {
	csvFile, e := os.Open("data.csv")
	require.NoError(t, e)

	csvReader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		record, e := csvReader.Read()
		if e == io.EOF {
			break
		} /*else if e != nil {
			log.Fatal(e.Error())
			return
		}*/
		require.NoError(t, e)

		f0, e := strconv.ParseFloat(strings.ReplaceAll(record[0], " ", ""), 64)
		require.NoError(t, e)

		f1, e := strconv.ParseFloat(strings.ReplaceAll(record[1], " ", ""), 64)
		require.NoError(t, e)

		tc := TestCase{f0, f1}
		r := Calc(tc)
		require.InDelta(t, f1, r, 0.0001)
	}
}
