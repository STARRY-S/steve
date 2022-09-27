package client

import (
	"github.com/rancher/apiserver/pkg/types"
)

type APIWarnings struct {
	ar *types.APIRequest
}

func (am APIWarnings) HandleWarningHeader(code int, agent string, message string) {
	if code != 299 || len(message) == 0 {
		return
	}

	am.ar.Response.Header().Add("X-API-Warnings", message)
}
