package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/mateusmatinato/goexpert-cep2temp/internal/platform/errors"
)

const (
	FailedGetInfo    = "ERR_FAILED_GET_WEATHER"
	FailedUnmarshall = "ERR_FAILED_UNMARSHALL_WEATHER"
)

type Service interface {
	GetInfo(ctx context.Context, request Request) (Response, error)
}

type service struct {
	client    http.Client
	apiConfig APIConfig
}

func (s service) GetInfo(_ context.Context, request Request) (Response, error) {
	weatherURL := fmt.Sprintf(s.apiConfig.URL, s.apiConfig.APIKey, url.QueryEscape(request.Query))
	res, err := s.client.Get(weatherURL)
	if err != nil {
		return Response{}, errors.NewApplicationError(FailedGetInfo, err)
	}

	var resp Response
	err = json.NewDecoder(res.Body).Decode(&resp)
	if err != nil {
		return Response{}, errors.NewApplicationError(FailedUnmarshall, err)
	}

	return resp, nil
}

func NewService(client http.Client, apiConfig APIConfig) Service {
	return &service{client: client, apiConfig: apiConfig}
}
