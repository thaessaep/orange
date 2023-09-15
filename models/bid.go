package models

// заявка на продажу/покупку (post-запрос)
type Bid struct {
	SymbolId int64 `json:"symbolId"`
	Price    int64 `json:"price"`
	Quantity int32 `json:"quantity"`
}
