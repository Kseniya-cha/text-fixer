package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Kseniya-cha/text-fixer/pkg/config"
	"github.com/Kseniya-cha/text-fixer/pkg/logger"
	"go.uber.org/zap"
)

type app struct {
	log     *zap.Logger
	cfg     *config.Config
	ctx     context.Context
	sigChan chan os.Signal
}

func NewApp(ctx context.Context, cfg *config.Config) *app {

	ctx, _ = context.WithCancel(ctx)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	return &app{
		log:     logger.NewLogger(cfg),
		cfg:     cfg,
		ctx:     ctx,
		sigChan: sigChan,
	}
}
