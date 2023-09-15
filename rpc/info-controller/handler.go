package info_controller

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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

func (n *news) LatestNews(time time.Time) error {
	url := fmt.Sprintf("%s%s", rpc.UrlNews, latestNews)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	n.log.Println(time, string(result))

	return nil
}

func (n *news) Last5MinutesNews() error {
	url := fmt.Sprintf("%s%s", rpc.UrlNews, latest5MinNews)

	resp, err := http.Get(url)
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

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(result), nil
}
