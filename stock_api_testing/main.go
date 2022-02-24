package main

import (
	"context"
	"fmt"
	"log"

	finnhub "github.com/Finnhub-Stock-API/finnhub-go/v2"
)

func main() {
	cfg := finnhub.NewConfiguration()
	cfg.AddDefaultHeader("X-Finnhub-Token", "c84livaad3i9u79hakkg")
	finnhubClient := finnhub.NewAPIClient(cfg).DefaultApi

	//Earnings calendar
	earningsCalendar, _, err := finnhubClient.EarningsCalendar(context.Background()).From("2021-07-01").To("2021-07-25").Execute()
	fmt.Printf("%+v\n", earningsCalendar)
	if err != nil {
		log.Fatal(err)
	}

	// NBBO
	bboData, _, err := finnhubClient.StockNbbo(context.Background()).Symbol("AAPL").Date("2021-07-23").Limit(50).Skip(0).Execute()
	fmt.Printf("%+v\n", bboData)

	//Stock candles
	stockCandles, _, err := finnhubClient.StockCandles(context.Background()).Symbol("AAPL").Resolution("D").From(1590988249).To(1591852249).Execute()
	fmt.Printf("%+v\n", stockCandles)

	fmt.Println(stockCandles.GetH())

}
