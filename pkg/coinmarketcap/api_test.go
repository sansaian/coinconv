package coinmarketcap

import (
	"github.com/sansaian/coinconv/config"
	"github.com/sansaian/coinconv/internal/entities"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func Test_priceConvertor_GetConvertingPrice(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		_, err := rw.Write([]byte(successResponse))
		require.NoError(t, err)

	}))
	defer server.Close()

	cfg := config.CoinMarket{
		Url:     server.URL,
		Token:   "someToken",
		Timeout: 5 * time.Second,
	}

	type fields struct {
		client *http.Client
		cfg    *config.CoinMarket
	}
	type args struct {
		data *entities.InputData
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.ConvertingResult
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				client: server.Client(),
				cfg:    &cfg,
			},
			args: args{
				data: &entities.InputData{
					Amount: 123.45,
					From:   "USD",
					To:     "BTC",
				},
			},
			want: &entities.ConvertingResult{
				Result: 7154,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &priceConvertor{
				client: tt.fields.client,
				cfg:    tt.fields.cfg,
			}
			got, err := pc.GetConvertingPrice(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConvertingPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConvertingPrice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

const successResponse = `{
    "status": {
        "timestamp": "2022-08-16T12:17:41.112Z",
        "error_code": 0,
        "error_message": null,
        "elapsed": 1,
        "credit_count": 1,
        "notice": null
    },
    "data": {
        "USD": {
            "symbol": "USD",
            "id": "ucofp8a0nhr",
            "name": "up3mgmklcf",
            "amount": 123.45,
            "last_updated": "2022-08-16T12:17:41.112Z",
            "quote": {
                "BTC": {
                    "price": 7154,
                    "last_updated": "2022-08-16T12:17:41.112Z"
                }
            }
        }
    }
}`
