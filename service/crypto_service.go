package service

import (
	"context"
	"cryptogo/database"
	"cryptogo/model"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
	"go.mongodb.org/mongo-driver/bson"
)

type cryptoService struct{}

var singleton CryptoService
var once sync.Once

func GetCryptoService() CryptoService {
	if singleton != nil {
		return singleton
	}
	once.Do(func() {
		singleton = &cryptoService{}
	})
	return singleton
}

type CryptoService interface {
	GetListToken(id string, currency string) *model.Token
}

func (s *cryptoService) GetListToken(id string, currency string) *model.Token {
	t := TestDb()
	fmt.Println(t.Name)
	res, err := LatestQuote(id, currency)
	if err != nil {
		log.Fatal(err)
	}

	token := model.Token{
		ID:         id,
		CoinMarket: res,
		// Tokenlist: model.Tokenlist{},
	}
	return &token
}

type Client struct {
	httpClient *http.Client
}

func LatestQuote(ids string, currency string) (interface{}, error) {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})
	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Symbol:  ids,
		Convert: currency,
	})
	if err != nil {
		panic(err)
	}
	var res model.CoinMarketQuote
	for _, quote := range quotes {
		// fmt.Println(quote.Name) // "Bitcoin"
		res.Slug = quote.Slug
		res.Price = quote.Quote[currency].Price
		res.Coin = quote.Name
		res.Currency = currency
		res.PercentChange24H = quote.Quote[currency].PercentChange24H
		// BTC price converted to CHF
		// fmt.Println(quote.Quote["CHF"].Price) // 3464.684951197137
	}
	return &res, nil
}

type Token struct {
	ID   int
	Name string
	Code string
}

func TestDb() Token {
	var token Token
	db := database.GetInstance()
	collection := db.Database("cryptodb").Collection("token_list")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res := collection.FindOne(ctx, bson.M{"name": "test"})
	res.Decode(&token)
	fmt.Println(token)
	return token
}
