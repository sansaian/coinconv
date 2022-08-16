package coinmarketcap

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/entities"
	"net/http"
	"net/url"
)

type response struct {
	Data map[string]struct {
		Quote map[string]struct {
			Price float64 `json:"price"`
		} `json:"quote"`
	} `json:"data"`
}

type priceConvertor struct {
	client *http.Client
	cfg    *config.CoinMarket
}

func New(cfg *config.CoinMarket, client *http.Client) *priceConvertor {
	return &priceConvertor{
		cfg:    cfg,
		client: client,
	}
}

func (pc *priceConvertor) GetConvertingPrice(data *entities.InputData) (*entities.ConvertingResult, error) {
	req, err := http.NewRequest("GET", pc.cfg.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("can't create http request: %w", err)
	}

	ctx, cancel := context.WithTimeout(req.Context(), pc.cfg.Timeout)
	defer cancel()

	req = req.WithContext(ctx)

	req = pc.withQuery(req, fmt.Sprintf("%f", data.Amount), data.From, data.To)
	req = pc.withHeaders(req)

	resp, err := pc.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request to server: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status=%v", resp.Status)
	}
	var r response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("invalid json in response: %w", err)
	}
	return &entities.ConvertingResult{Result: r.Data[data.From].Quote[data.To].Price}, nil
}

func (pc *priceConvertor) withHeaders(req *http.Request) *http.Request {
	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", pc.cfg.Token)
	return req
}

func (pc *priceConvertor) withQuery(req *http.Request, amount, from, to string) *http.Request {
	q := url.Values{}
	q.Add("amount", amount)
	q.Add("symbol", from)
	q.Add("convert", to)
	req.URL.RawQuery = q.Encode()
	return req
}
