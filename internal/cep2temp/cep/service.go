package cep

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mateusmatinato/goexpert-cep2temp/internal/platform/errors"
)

const (
	NotFoundCEP      = "ERR_NOT_FOUND_CEP"
	FailedGetInfo    = "ERR_FAILED_GET_CEP"
	FailedUnmarshall = "ERR_FAILED_UNMARSHALL_CEP"
)

type Service interface {
	GetInfo(ctx context.Context, request Request) (Response, error)
}

type service struct {
	client    http.Client
	apiConfig APIConfig
}

func (s service) GetInfo(_ context.Context, request Request) (Response, error) {
	res, err := s.client.Get(fmt.Sprintf(s.apiConfig.URL, request.Cep))
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return Response{}, errors.NewNotFoundError(NotFoundCEP, err)
		}
		return Response{}, errors.NewApplicationError(FailedGetInfo, err)
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, errors.NewApplicationError(FailedUnmarshall, err)
	}

	if resp.City == "" {
		return Response{}, errors.NewNotFoundError(NotFoundCEP, err)
	}

	return resp, nil
}

func NewService(client http.Client, config APIConfig) Service {
	return &service{
		client:    client,
		apiConfig: config,
	}
}
