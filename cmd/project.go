package cmd

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/lu4p/cat"
	"go.uber.org/zap"
)

func (a *app) GracefulShutdown(cancel context.CancelFunc) {

	defer close(a.sigChan)
	<-a.sigChan

	cancel()
	a.log.Info("exiting...")
	time.Sleep(time.Second * 1)
}

func (a *app) Run(ctx context.Context) {

	file := a.cfg.Path + a.cfg.Name

	txt, err := readFix(file)
	if err != nil {
		a.log.Fatal(err.Error())
	}
	a.log.Info("file readed without extra spaces", zap.String("file", file))

	newTxt, err := replaceTabs(txt, a.cfg.Mode)
	if err != nil {
		a.log.Fatal(err.Error())
	}
	a.log.Info("tabs replaced", zap.String("file", file))

	err = a.saveFile(file, "./"+strings.Split(a.cfg.Name, ".")[0]+"_"+a.cfg.Mode+".txt", newTxt)
	if err != nil {
		a.log.Fatal(err.Error())
	}

	a.sigChan <- os.Interrupt
}

func readFix(path string) (string, error) {

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

func replaceTabs(txt string, mode string) (string, error) {

	var newTxt string

	switch mode {
	case "add":
		if string(txt[:5]) == "<tab>" {
			return newTxt, errors.New("<tab> already in text!")
		}
		newTxt = "<tab>" + strings.Replace(txt, "\n\n", "\n\n<tab>", -1)

	case "remove":
		if string(txt[:5]) != "<tab>" {
			return newTxt, errors.New("there is no <tab> in text!")
		}
		newTxt = string(strings.Replace(txt, "\n\n<tab>", "\n\n", -1)[5:])
	}

	return newTxt, nil
}

func (a *app) saveFile(path, newPath, txt string) error {
	newFile, err := os.Create(newPath)
	if err != nil {
		return err
	}

	_, err = newFile.WriteString(txt)
	if err != nil {
		return err
	}

	a.log.Info("file saved", zap.String("file", newPath))

	return nil
}
