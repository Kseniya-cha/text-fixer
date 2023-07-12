package main

import (
	"context"
	"fmt"

	"github.com/Kseniya-cha/text-fixer/cmd"
	"github.com/Kseniya-cha/text-fixer/pkg/config"
)

// 1. Заменить "\n\n" на "\n\n<tab>"
// 1. Заменить "\n\n<tab>" на "\n\n"

func main() {
	// Чтение конфигурационного файла
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := cmd.NewApp(ctx, cfg)

	// Запуск алгоритма в отдельной горутине
	go app.Run(ctx)

	// Ожидание прерывающего сигнала
	app.GracefulShutdown(cancel)
}
