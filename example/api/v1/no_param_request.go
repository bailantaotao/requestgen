package api

import "github.com/c9s/requestgen"

type NoParamResponse struct {
	ID string `json:"id"`
}

//go:generate /usr/local/go/bin/go run ../../../cmd/requestgen -method GET -url /v1/bullet -debug -type NoParamRequest -responseType NoParamResponse
type NoParamRequest struct {
	client requestgen.APIClient
}
