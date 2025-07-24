package convert

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func ToJSON(p string, saveToFile bool) (*string, error) {
	file, err := os.Open(p)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	json := string(buffer)

	err = Normalize(&json)
	if err != nil {
		return nil, err
	}

	if saveToFile {
		save(p, &json)
	}

	return &json, nil
}

func Normalize(jsonString *string) error {
	repl, err := getReplacement()
	if err != nil {
		return err
	}

	for o, n := range repl {
		*jsonString = strings.ReplaceAll(*jsonString, o, n)
	}

	*jsonString = strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, *jsonString)

	*jsonString = (*jsonString)[1 : len(*jsonString)-1]

	replacer := strings.NewReplacer(
		"{", "{\n",
		"},", "\n\t},\n",
		"\",", "\",\n",
		"}  ]", "\n\t}\n]")
	*jsonString = replacer.Replace(*jsonString)

	return nil
}

func getReplacement() (map[string]string, error) {
	file, err := os.Open("convert/replacement.json")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	buffer, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data = make(map[string]string)
	err = json.Unmarshal(buffer, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func save(fName string, json *string) error {
	nameParts := strings.Split(fName, ".")
	fName = nameParts[0] + "_normalized.json"
	newFile, err := os.Create(fName)
	if err != nil {
		return err
	}

	_, err = newFile.WriteString(*json)
	if err != nil {
		return err
	}

	return nil
}
