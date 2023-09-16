package buyer_by_news

import (
	"errors"
	"fmt"

	"github.com/thaessaep/models"
)

const PriceForBuy = 500

var lastNews models.News

type latestNews interface {
	LatestNews() (models.News, error)
}

type bidForSell interface {
	SalesRequests() ([]models.StockInfoSell, error)
}

type postOperationsBid interface {
	PushBidBuy(bid models.Bid) (bool, error)
}

type buyerByNews struct {
	lNews             latestNews
	bidForSell        bidForSell
	postOperationsBid postOperationsBid
}

func New(lNews latestNews, bidForSell bidForSell, postOperationsBid postOperationsBid) buyerByNews {
	return buyerByNews{
		lNews:             lNews,
		bidForSell:        bidForSell,
		postOperationsBid: postOperationsBid,
	}
}

func (b *buyerByNews) PushBidByNews() ([]models.Bid, error) {
	// LatestNews - получаем последние новости
	// если новость старая - скипаем
	news, err := b.lNews.LatestNews()
	if err != nil {
		return nil, err
	}

	if news.Date == lastNews.Date {
		return nil, errors.New("old date")
	}

	// получаем заявки на продажу
	stockInfoSell, err := b.bidForSell.SalesRequests()
	if err != nil {
		return nil, err
	}

	// если в заявке на продажу rate > 0 -> смотрим цену и если цена нам подходит -> выставляем заявку на покупку
	if news.Rate > 0 {
		res, err := b.pushBidForBuy(stockInfoSell, news.CompaniesAffected)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	// если в заявке на продажу rate < 0 -> смотрим есть ли у нас эта акция и если есть -> выставляем заявку на продажу

	return nil, nil
}

func (b *buyerByNews) pushBidForBuy(stockInfoSell []models.StockInfoSell, companiesAffected []string) ([]models.Bid, error) {
	var result []models.Bid
	for _, stock := range stockInfoSell {
		bid := minBid(stock)
		fmt.Println("check bid", bid)
		if bid != nil && inCompaniesAffected(companiesAffected, stock.Ticker) {
			bidToPush := models.Bid{
				SymbolId: stock.Id,
				Price:    bid.Price,
				Quantity: 1,
			}

			// пушим заявку на покупку
			pushed, err := b.postOperationsBid.PushBidBuy(bidToPush)
			if err != nil {
				return nil, err
			}
			if pushed {
				result = append(result, bidToPush)
			}

		}
	}

	return result, nil
}

// смотрим подходящая ли цена
func minBid(stock models.StockInfoSell) *models.BidStat {
	minPrice := int64(PriceForBuy)
	var result *models.BidStat

	fmt.Println("check stock by bid price", stock)
	for _, bid := range stock.Bids {
		fmt.Println("check bid price", bid.Price)
		if bid.Price < PriceForBuy && bid.Price < minPrice {
			result = &bid
		}
	}

	return result
}

func inCompaniesAffected(companiesAffected []string, ticker string) bool {
	for _, company := range companiesAffected {

		// отделяем часть Oranges/ от тикера
		if company == ticker[8:] {
			return true
		}
	}

	return false
}
