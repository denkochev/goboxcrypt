package goboxcrypt

import (
	"encoding/csv"
	"os"
	"testing"
)

type TestCases struct {
	plainText string
	hexKey    string
}

func Test_Sbox_encryption(t *testing.T) {
	tests, _ := readCSV("./test_assets/lorem_hex_dataset.csv")

	for _, test := range tests {
		plainText := test.plainText
		key := test.hexKey

		cipher := Encrypt_Sbox(plainText, key)

		result := Decrypt_Sbox(cipher, key)

		if result != test.plainText {
			t.Errorf("Expected:--------------------- %s,  ----------- but got --------- %s", test.plainText, result)
		}
	}
}

func Test_Pbox_encryption(t *testing.T) {
	tests, _ := readCSV("./test_assets/lorem_hex_dataset.csv")

	for _, test := range tests {
		plainText := test.plainText
		key := test.hexKey

		cipher := Encrypt_Pbox(plainText, key)

		result := Decrypt_Pbox(cipher, key)

		if result != test.plainText {
			t.Errorf("Expected:--------------------- %s,  ----------- but got --------- %s", test.plainText, result)
		}
	}
}

func readCSV(filename string) ([]TestCases, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var records []TestCases
	for _, line := range lines[1:] { // Skip header
		records = append(records, TestCases{
			plainText: line[0],
			hexKey:    line[1],
		})
	}
	return records, nil
}
