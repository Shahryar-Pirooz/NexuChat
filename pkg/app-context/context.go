package appcontext

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

var defaultLogger *zap.Logger

type AppContextOption func(*appContext) *appContext

func init() {
	var err error
	defaultLogger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func WithDatabase(db *gorm.DB) AppContextOption {
	return func(ctx *appContext) *appContext {
		ctx.db = db
		return ctx
	}
}

func WithLogger(logger *zap.Logger) AppContextOption {
	return func(ctx *appContext) *appContext {
		ctx.logger = logger
		return ctx
	}
}

func SetDatabase(ctx context.Context, db *gorm.DB) {
	ac, ok := ctx.(*appContext)
	if !ok {
		return
	}
	ac.db = db
}

func GetDatabase(ctx context.Context) *gorm.DB {
	ac, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return ac.db
}

func SetLogger(ctx context.Context, logger *zap.Logger) {
	ac, ok := ctx.(*appContext)
	if !ok {
		return
	}
	ac.logger = logger
}

func GetLogger(ctx context.Context) *zap.Logger {
	ac, ok := ctx.(*appContext)
	if !ok {
		return nil
	}
	return ac.logger
}


func NewAppContext(parent context.Context , opts ...AppContextOption) context.Context{
	ctx := &appContext{Context: parent}

	for _,opt := range opts{
		ctx = opt(ctx)
	}
	return ctx
}
