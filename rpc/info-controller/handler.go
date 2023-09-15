package info_controller

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

const (
	latestNews     = "LatestNews"
	latest5MinNews = "LatestNews5Minutes"
)

type news struct {
	log *log.Logger
}

func New(logger *log.Logger) news {
	return news{
		log: logger,
	}
}

func (n *news) LatestNews() (models.News, error) {
	url := fmt.Sprintf("%s%s", rpc.UrlNews, latestNews)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return models.News{}, err
	}
	request.Header.Add("token", config.Token)

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		return models.News{}, errors.New("cannot send request")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.News{}, errors.New("cannot read body")
	}

	n.log.Println(string(body))

	var result models.News
	err = json.Unmarshal(body, &result)
	if err != nil {
		return models.News{}, err
	}

	return result, nil
}

func (n *news) Last5MinutesNews() error {
	url := rpc.UrlNews + latest5MinNews

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	request.Header.Add("token", config.Token)

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	n.log.Println(string(result))

	return nil
}

func (n *news) BillingInfo() (string, error) {
	url := rpc.UrlInfo

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("token", config.Token)

	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
