package cmd

import (
	"context"
	"os"
	"strings"
	"time"

	textprocessing "github.com/Kseniya-cha/text-fixer/pkg/methods"
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
	txt, err := textprocessing.ReadFix(file)
	if err != nil {
		a.log.Fatal(err.Error())
	}
	a.log.Info("file read without extra spaces", zap.String("file", file))

	newTxt, err := textprocessing.ReplaceTabs(txt, a.cfg.Mode, a.cfg.Object)
	if err != nil {
		a.log.Fatal(err.Error())
	}
	a.log.Info(a.cfg.Object+a.cfg.Mode, zap.String("file", file))

	newFile := "./" + strings.Split(a.cfg.Name, ".")[0] + "_" + a.cfg.Mode + ".doc"
	err = textprocessing.SaveFile(file, newFile, newTxt)
	if err != nil {
		a.log.Fatal(err.Error())
	}
	a.log.Info("file saved", zap.String("file", newFile))

	a.sigChan <- os.Interrupt
}
