package models

// Заявки на продажу (get запрос)
type StockInfoSell struct {
	Id     int64     `json:"id"`
	Ticker string    `json:"ticker"`
	Bids   []BidStat `json:",inline"`
}

type BidStat struct {
	Price    int64 `json:"price"`
	Quantity int32 `json:"quantity"`
}
