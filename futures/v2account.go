package futures

import (
	"context"
	"encoding/json"
)

// GetBalanceService get account balance
type GetBalanceServiceV2 struct {
	c *Client
}

// Do send request
func (s *GetBalanceServiceV2) Do(ctx context.Context, opts ...RequestOption) (res []*BalanceV2, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v2/balance",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*BalanceV2{}, err
	}
	res = make([]*BalanceV2, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*BalanceV2{}, err
	}
	return res, nil
}

// Balance define user balance of your account
type BalanceV2 struct {
	AccountAlias       string `json:"accountAlias"`
	Asset              string `json:"asset"`
	Balance            string `json:"balance"`
	CrossWalletBalance string `json:"crossWalletBalance"`
	CrossUnPnl         string `json:"crossUnPnl"`
	availableBalance   string `json:"availableBalance"`
	MaxWithdrawAmount  string `json:"maxWithdrawAmount"`
}

// GetBalanceService get account balance
type GetPositionServiceV2 struct {
	c *Client
	symbol string
}

// Symbol set symbol
func (s *GetPositionServiceV2) Symbol(symbol string) *GetPositionServiceV2 {
	s.symbol = symbol
	return s
}

// Do send request
func (s *GetPositionServiceV2) Do(ctx context.Context, opts ...RequestOption) (res []*PositionV2, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/fapi/v2/positionRisk",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*PositionV2{}, err
	}
	res = make([]*PositionV2, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*PositionV2{}, err
	}
	return res, nil
}

// Position define user balance of your account
type PositionV2 struct {
	EntryPrice       string `json:"entryPrice"`
	MarginType       string `json:"marginType"`
	IsAutoAddMargin  string `json:"isAutoAddMargin"`
	IsolatedMargin   string `json:"isolatedMargin"`
	Leverage         string `json:"leverage"`
	LiquidationPrice string `json:"liquidationPrice"`
	MarkPrice        string `json:"markPrice"`
	MaxNotionalValue string `json:"maxNotionalValue"`
	PositionAmt      string `json:"positionAmt"`
	Symbol           string `json:"symbol"`
	UnRealizedProfit string `json:"unRealizedProfit"`
	PositionSide     string `json:"positionSide"`
}
