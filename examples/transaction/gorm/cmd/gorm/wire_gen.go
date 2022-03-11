// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SeeMusic/kratos/examples/transaction/gorm/internal/biz"
	"github.com/SeeMusic/kratos/examples/transaction/gorm/internal/conf"
	"github.com/SeeMusic/kratos/examples/transaction/gorm/internal/data"
	"github.com/SeeMusic/kratos/examples/transaction/gorm/internal/server"
	"github.com/SeeMusic/kratos/examples/transaction/gorm/internal/service"
	"github.com/SeeMusic/kratos/v2"
	"github.com/SeeMusic/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	cardRepo := data.NewCardRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	userUsecase := biz.NewUserUsecase(userRepo, cardRepo, transaction, logger)
	transactionService := service.NewTransactionService(userUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, transactionService)
	grpcServer := server.NewGRPCServer(confServer, logger, transactionService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
