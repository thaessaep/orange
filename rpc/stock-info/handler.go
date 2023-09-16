package stock_info

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/thaessaep/config"
	"github.com/thaessaep/models"
	"github.com/thaessaep/rpc"
)

type news struct {
	log *log.Logger
}

func New(logger *log.Logger) news {
	return news{
		log: logger,
	}
}

func (n *news) SalesRequests() ([]models.StockInfoSell, error) {
	url := rpc.SellStockUrl

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("token", config.Token)

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot take StockInfoSell: %s", err))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//n.log.Println(string(body))

	var result []models.StockInfoSell

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("cannot unmarshal StockInfoSell: %s", err))
	}

	return result, nil
}

//
//func (n *news) SalesRequests(time time.Time) error {
//	url := rpc.SellStockUrl
//
//	resp, err := http.Get(url)
//	if err != nil {
//		return err
//	}
//
//	result, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	n.log.Println(time, string(result))
//
//	return nil
//}
