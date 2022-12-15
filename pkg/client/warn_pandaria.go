package client

import (
	"github.com/rancher/apiserver/pkg/types"
	"net/url"
)

type APIWarnings struct {
	ar *types.APIRequest
}

func (am APIWarnings) HandleWarningHeader(code int, agent string, message string) {
	if code != 299 || len(message) == 0 {
		return
	}

	//The front-end js needs to process this header similarly to QueryUnescape to avoid passing unsafe characters and display Chinese characters correctly.
	am.ar.Response.Header().Add("X-API-Warnings", url.QueryEscape(message))
}
