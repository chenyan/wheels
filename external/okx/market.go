package okx

import "fmt"

// Ticker 行情数据
type Ticker struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Last      string `json:"last"`      // 最新成交价
	LastSz    string `json:"lastSz"`    // 最新成交的数量
	AskPx     string `json:"askPx"`     // 卖一价
	AskSz     string `json:"askSz"`     // 卖一数量
	BidPx     string `json:"bidPx"`     // 买一价
	BidSz     string `json:"bidSz"`     // 买一数量
	Open24h   string `json:"open24h"`   // 24小时开盘价
	High24h   string `json:"high24h"`   // 24小时最高价
	Low24h    string `json:"low24h"`    // 24小时最低价
	VolCcy24h string `json:"volCcy24h"` // 24小时成交量，以币为单位
	Vol24h    string `json:"vol24h"`    // 24小时成交量，以张为单位
	SodUtc0   string `json:"sodUtc0"`   // UTC 0 时开盘价
	SodUtc8   string `json:"sodUtc8"`   // UTC+8 时开盘价
	Ts        string `json:"ts"`        // ticker数据产生时间
}

// GetTicker 获取单个产品行情信息
func (c *Client) GetTicker(instId string) (*Ticker, error) {
	path := fmt.Sprintf("/api/v5/market/ticker?instId=%s", instId)
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var tickers []Ticker
	if err := ParseResponse(resp, &tickers); err != nil {
		return nil, err
	}
	
	if len(tickers) == 0 {
		return nil, fmt.Errorf("no ticker data found")
	}
	
	return &tickers[0], nil
}

// GetTickers 获取所有产品行情信息
func (c *Client) GetTickers(instType string) ([]Ticker, error) {
	path := fmt.Sprintf("/api/v5/market/tickers?instType=%s", instType)
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var tickers []Ticker
	if err := ParseResponse(resp, &tickers); err != nil {
		return nil, err
	}
	
	return tickers, nil
}

// Orderbook 深度数据
type Orderbook struct {
	Asks [][]string `json:"asks"` // 卖方深度，[价格, 数量, 已废弃, 订单数量]
	Bids [][]string `json:"bids"` // 买方深度，[价格, 数量, 已废弃, 订单数量]
	Ts   string     `json:"ts"`   // 深度产生的时间
}

// GetOrderbook 获取产品深度
func (c *Client) GetOrderbook(instId string, sz int) (*Orderbook, error) {
	path := fmt.Sprintf("/api/v5/market/books?instId=%s", instId)
	if sz > 0 {
		path += fmt.Sprintf("&sz=%d", sz)
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var orderbooks []Orderbook
	if err := ParseResponse(resp, &orderbooks); err != nil {
		return nil, err
	}
	
	if len(orderbooks) == 0 {
		return nil, fmt.Errorf("no orderbook data found")
	}
	
	return &orderbooks[0], nil
}

// Candle K线数据
type Candle struct {
	Ts       string `json:"ts"`       // 开始时间
	O        string `json:"o"`        // 开盘价格
	H        string `json:"h"`        // 最高价格
	L        string `json:"l"`        // 最低价格
	C        string `json:"c"`        // 收盘价格
	Vol      string `json:"vol"`      // 交易量（张）
	VolCcy   string `json:"volCcy"`   // 交易量（币）
	VolCcyQt string `json:"volCcyQt"` // 交易量（计价货币）
	Confirm  string `json:"confirm"`  // K线状态：0 K线未完结，1 K线已完结
}

// GetCandles 获取K线数据
// bar: 时间粒度，默认值1m
// 如 [1m/3m/5m/15m/30m/1H/2H/4H]
// 香港时间开盘价k线：[6H/12H/1D/2D/3D/1W/1M/3M]
// UTC时间开盘价k线：[/6Hutc/12Hutc/1Dutc/2Dutc/3Dutc/1Wutc/1Mutc/3Mutc]
func (c *Client) GetCandles(instId, bar string, limit int) ([]Candle, error) {
	path := fmt.Sprintf("/api/v5/market/candles?instId=%s", instId)
	if bar != "" {
		path += fmt.Sprintf("&bar=%s", bar)
	}
	if limit > 0 {
		path += fmt.Sprintf("&limit=%d", limit)
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var candles []Candle
	if err := ParseResponse(resp, &candles); err != nil {
		return nil, err
	}
	
	return candles, nil
}

// Instrument 产品信息
type Instrument struct {
	InstType  string `json:"instType"`  // 产品类型
	InstId    string `json:"instId"`    // 产品ID
	Uly       string `json:"uly"`       // 标的指数
	Category  string `json:"category"`  // 手续费类别
	BaseCcy   string `json:"baseCcy"`   // 交易货币币种
	QuoteCcy  string `json:"quoteCcy"`  // 计价货币币种
	SettleCcy string `json:"settleCcy"` // 盈亏结算和保证金币种
	CtVal     string `json:"ctVal"`     // 合约面值
	CtMult    string `json:"ctMult"`    // 合约乘数
	CtValCcy  string `json:"ctValCcy"`  // 合约面值币种
	OptType   string `json:"optType"`   // 期权类型：C看涨期权 P看跌期权
	Stk       string `json:"stk"`       // 行权价格
	ListTime  string `json:"listTime"`  // 上线日期
	ExpTime   string `json:"expTime"`   // 产品下线时间
	Lever     string `json:"lever"`     // 该instId交易时可用的最高杠杆倍数
	TickSz    string `json:"tickSz"`    // 下单价格精度
	LotSz     string `json:"lotSz"`     // 下单数量精度
	MinSz     string `json:"minSz"`     // 最小下单数量
	CtType    string `json:"ctType"`    // 合约类型：linear 正向合约 inverse 反向合约
	Alias     string `json:"alias"`     // 合约日期别名
	State     string `json:"state"`     // 产品状态
}

// GetInstruments 获取交易产品基础信息
func (c *Client) GetInstruments(instType, instId string) ([]Instrument, error) {
	path := fmt.Sprintf("/api/v5/public/instruments?instType=%s", instType)
	if instId != "" {
		path += fmt.Sprintf("&instId=%s", instId)
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var instruments []Instrument
	if err := ParseResponse(resp, &instruments); err != nil {
		return nil, err
	}
	
	return instruments, nil
}

