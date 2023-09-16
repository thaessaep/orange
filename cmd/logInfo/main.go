package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/thaessaep/rpc/buyer_by_news"
	info_controller "github.com/thaessaep/rpc/info-controller"
	"github.com/thaessaep/rpc/operation_bid"
	stock_info "github.com/thaessaep/rpc/stock-info"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	logger := log.New(file, "INFO\t", log.Ldate|log.Ltime)

	bidFile, err := os.OpenFile("bid_operation.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	bidLogger := log.New(bidFile, "INFO\t", log.Ldate|log.Ltime)

	//stockFile, err := os.OpenFile("stock_info.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//stockInfoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	newsController := info_controller.New(logger)
	stockInfoController := stock_info.New(nil)
	oparationBidController := operation_bid.New(bidLogger)

	operationBid := buyer_by_news.New(&newsController, &stockInfoController, &oparationBidController)

	timer := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timer.C:
			_, err := newsController.BillingInfo()
			if err != nil {
				panic(err.Error())
			}

			//news, err := controller.LatestNews()
			//if err != nil {
			//	panic(err.Error())
			//}
			bids, err := operationBid.PushBidByNews()
			if err != nil {
				fmt.Println(err.Error())
			}
			bidLogger.Println("buy bids: ", bids)
		}
	}
}
