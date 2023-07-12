package textprocessing

import (
	"errors"
	"os"
	"strings"

	"github.com/lu4p/cat"
)

func ReadFix(path string) (string, error) {

	txt, err := cat.File(path)
	if err != nil {
		return "", err
	}

	listWrong := []string{",", " ", "-"}
	for _, item := range listWrong {
		txt = strings.Replace(txt, item+item, item, -1)
	}

	return txt, nil
}

func ReplaceTabs(txt string, mode string) (string, error) {

	var newTxt string

	switch mode {
	case "add":
		if string(txt[:5]) == "<tab>" {
			return newTxt, errors.New("<tab> already in text")
		}
		newTxt = "<tab>" + strings.Replace(txt, "\n\n", "\n\n<tab>", -1)

	case "remove":
		if string(txt[:5]) != "<tab>" {
			return newTxt, errors.New("there is no <tab> in text")
		}
		newTxt = string(strings.Replace(txt, "\n\n<tab>", "\n\n", -1)[5:])
	}

	return newTxt, nil
}

func SaveFile(path, newPath, txt string) error {
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}

	_, err = newFile.WriteString(txt)
	if err != nil {
		return err
	}

	return nil
}
