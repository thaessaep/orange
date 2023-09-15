package operation_bid

import (
	"bytes"
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

func (n *news) PushBidBuy(bid models.Bid) (bool, error) {
	url := rpc.LimitPriceBuyUrl

	body, err := json.Marshal(bid)
	if err != nil {
		return false, errors.New(fmt.Sprintf("cannot marshal bid: %s", err))
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return false, err
	}
	request.Header.Add("token", config.Token)

	client := http.Client{Timeout: 1 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		return false, err
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	n.log.Println(string(answer))

	return true, nil
}
