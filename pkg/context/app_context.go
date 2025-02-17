package context

import (
	"context"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type appContext struct {
	context.Context
	db     *gorm.DB
	logger *zap.Logger
}

type AppContextOpt func(*appContext) *appContext

var defaultLogger *zap.Logger

func init() {
	var err error
	defaultLogger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func WithDB(db *gorm.DB) AppContextOpt {
	return func(ctx *appContext) *appContext {
		ctx.db = db
		return ctx
	}
}

func WithLogger(logger *zap.Logger) AppContextOpt {
	return func(ctx *appContext) *appContext {
		ctx.logger = logger
		return ctx
	}
}

func SetDB(ctx context.Context, db *gorm.DB) {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return
	}
	appCtx.db = db
}

func SetLogger(ctx context.Context, logger *zap.Logger) {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return
	}
	appCtx.logger = logger
}

func GetDB(ctx context.Context) *gorm.DB {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return appCtx.db
}

func GetLogger(ctx context.Context) *zap.Logger {
	appCtx, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return appCtx.logger
}

func NewAppContext(parrent context.Context, opts ...AppContextOpt) context.Context {
	ctx := &appContext{
		Context: parrent,
	}
	for _, opt := range opts {
		ctx = opt(ctx)
	}
	return ctx
}
