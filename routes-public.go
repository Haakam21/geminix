package gemini

import (
	"encoding/json"
	"strconv"
)

// Symbols
func (g *Api) Symbols() ([]string, error) {

	url := g.url + SYMBOLS_URI

	var symbols []string

	body, err := g.request("GET", url, nil, nil)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &symbols)

	return symbols, nil
}

// Ticker
func (g *Api) Ticker(symbol string) (Ticker, error) {

	url := g.url + TICKER_URI + symbol

	var ticker Ticker

	body, err := g.request("GET", url, nil, nil)
	if err != nil {
		return ticker, err
	}

	json.Unmarshal(body, &ticker)

	return ticker, nil
}

// Order Book
func (g *Api) OrderBook(symbol string, limitBids, limitAsks int) (Book, error) {

	url := g.url + BOOK_URI + symbol
	params := map[string]interface{}{
		"limit_bids": strconv.Itoa(limitBids),
		"limit_asks": strconv.Itoa(limitAsks),
	}

	var book Book

	body, err := g.request("GET", url, nil, params)
	if err != nil {
		return book, err
	}

	json.Unmarshal(body, &book)

	return book, nil
}

// Trades
func (g *Api) Trades(symbol string, since int64, limitTrades int, includeBreaks bool) ([]Trade, error) {

	url := g.url + TRADES_URI + symbol
	params := map[string]interface{}{
		"since":          strconv.Itoa(int(since)),
		"limit_trades":   strconv.Itoa(limitTrades),
		"include_breaks": strconv.FormatBool(includeBreaks),
	}

	var res []Trade

	body, err := g.request("GET", url, nil, params)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &res)

	return res, nil
}

// Current Auction
func (g *Api) CurrentAuction(symbol string) (CurrentAuction, error) {

	url := g.url + AUCTION_URI + symbol

	var auction CurrentAuction

	body, err := g.request("GET", url, nil, nil)
	if err != nil {
		return auction, err
	}

	json.Unmarshal(body, &auction)

	return auction, nil
}

// Auction History
func (g *Api) AuctionHistory(symbol string, since int64, limit int, includeIndicative bool) ([]Auction, error) {

	url := g.url + AUCTION_URI + symbol + "/history"
	params := map[string]interface{}{
		"since":                 strconv.Itoa(int(since)),
		"limit_auction_results": strconv.Itoa(limit),
		"include_indicative":    strconv.FormatBool(includeIndicative),
	}

	var auctions []Auction

	body, err := g.request("GET", url, nil, params)
	if err != nil {
		return auctions, err
	}

	json.Unmarshal(body, &auctions)

	return auctions, nil
}
