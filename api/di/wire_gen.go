// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/andrewhowdencom/x40.link/api"
	"github.com/andrewhowdencom/x40.link/storage/di"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func WireGRPCServer() (*grpc.Server, error) {
	storer, err := di.WireStorage()
	if err != nil {
		return nil, err
	}
	v, err := OptsFromViper()
	if err != nil {
		return nil, err
	}
	server := api.NewGRPCMux(storer, v...)
	return server, nil
}
