package common

import (
	"github.com/cwntr/crypto-sdk/pkg/coins/grph"
	"github.com/cwntr/crypto-sdk/pkg/coins/mc"
	"github.com/dnaeon/go-vcr/recorder"
)

var (
	clients map[string]interface{}
)


func InitClients(rec *recorder.Recorder) {
	clients = make(map[string]interface{}, 2)
	clients[NameGRPH] = grph.InitApi(grph.APIUrl, rec)
	clients[NameMC] = mc.InitApi(mc.APIUrl, rec)
}

func GetClientByName(name string) interface{} {
	return clients[name]
}

func GetClients() map[string]interface{} {
	return clients
}

func GetGrphClient() *grph.API {
	grphClient := clients[NameGRPH]
	switch v := grphClient.(type) {
	case *grph.API:
		return v
	}
	return nil
}

func GetMcClient() *mc.API {
	mcClient := clients[NameMC]
	switch v := mcClient.(type) {
	case *mc.API:
		return v
	}
	return nil
}
