package vowels

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Delete(outputFile string, resultFileName string, vowelsType string) error {
	records, err := ReadCsvFile(outputFile)
	if err != nil {
		return err
	}
	var d []string
	for _, p := range records {
		d = append(d, DeleteCharacter(vowelsType, p))
	}
	if err := WriteFileCsv(d, resultFileName); err != nil {
		return err
	}
	return nil

}
func DeleteCharacter(vowels string, text string) string {
	var builder strings.Builder
	for i := range text {
		if !strings.Contains(vowels, string(text[i])) {
			builder.WriteString(string(text[i]))
		}
	}
	return builder.String()
}

// ReadCsvFile Read file
func ReadCsvFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var vList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vList = append(vList, line)
	}
	return vList, nil
}

// WriteFileCsv Write csv file
func WriteFileCsv(text []string, resultFileName string) error {
	file, err := os.OpenFile(resultFileName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	dataWriter := bufio.NewWriter(file)
	for _, h := range text {
		s := fmt.Sprintf("%s\n", h)
		_, err = dataWriter.WriteString(s)
		if err != nil {
			return err
		}
	}

	dataWriter.Flush()
	return nil
}
