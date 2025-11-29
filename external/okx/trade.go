package okx

import "fmt"

// OrderRequest 下单请求
type OrderRequest struct {
	InstId     string `json:"instId"`               // 产品ID
	TdMode     string `json:"tdMode"`               // 交易模式：isolated 逐仓 cross 全仓 cash 非保证金
	Side       string `json:"side"`                 // 订单方向：buy sell
	OrdType    string `json:"ordType"`              // 订单类型：market 市价单 limit 限价单 post_only 只做maker fok 全部成交或立即取消 ioc 立即成交并取消剩余
	Sz         string `json:"sz"`                   // 委托数量
	Px         string `json:"px,omitempty"`         // 委托价格，仅限价单有效
	PosSide    string `json:"posSide,omitempty"`    // 持仓方向：long short net，默认net
	Ccy        string `json:"ccy,omitempty"`        // 保证金币种，仅适用于跨币种保证金模式下的全仓杠杆订单
	ClOrdId    string `json:"clOrdId,omitempty"`    // 客户自定义订单ID
	Tag        string `json:"tag,omitempty"`        // 订单标签
	ReduceOnly bool   `json:"reduceOnly,omitempty"` // 是否只减仓，true 或 false，默认false
	TgtCcy     string `json:"tgtCcy,omitempty"`     // 委托数量的类型：base_ccy 交易货币 quote_ccy 计价货币，仅适用于币币市价单
}

// OrderResponse 下单响应
type OrderResponse struct {
	ClOrdId string `json:"clOrdId"` // 客户自定义订单ID
	OrdId   string `json:"ordId"`   // 订单ID
	Tag     string `json:"tag"`     // 订单标签
	SCode   string `json:"sCode"`   // 事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    // 事件执行失败时的msg
}

// PlaceOrder 下单
func (c *Client) PlaceOrder(req OrderRequest) (*OrderResponse, error) {
	path := "/api/v5/trade/order"
	
	resp, err := c.Post(path, req)
	if err != nil {
		return nil, err
	}
	
	var results []OrderResponse
	if err := ParseResponse(resp, &results); err != nil {
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, nil
	}
	
	return &results[0], nil
}

// CancelOrderRequest 撤单请求
type CancelOrderRequest struct {
	InstId  string `json:"instId"`            // 产品ID
	OrdId   string `json:"ordId,omitempty"`   // 订单ID，ordId和clOrdId必须传一个
	ClOrdId string `json:"clOrdId,omitempty"` // 客户自定义订单ID
}

// CancelOrderResponse 撤单响应
type CancelOrderResponse struct {
	ClOrdId string `json:"clOrdId"` // 客户自定义订单ID
	OrdId   string `json:"ordId"`   // 订单ID
	SCode   string `json:"sCode"`   // 事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    // 事件执行失败时的msg
}

// CancelOrder 撤单
func (c *Client) CancelOrder(req CancelOrderRequest) (*CancelOrderResponse, error) {
	path := "/api/v5/trade/cancel-order"
	
	resp, err := c.Post(path, req)
	if err != nil {
		return nil, err
	}
	
	var results []CancelOrderResponse
	if err := ParseResponse(resp, &results); err != nil {
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, nil
	}
	
	return &results[0], nil
}

// Order 订单信息
type Order struct {
	InstType        string `json:"instType"`        // 产品类型
	InstId          string `json:"instId"`          // 产品ID
	TgtCcy          string `json:"tgtCcy"`          // 委托数量的类型
	Ccy             string `json:"ccy"`             // 保证金币种
	OrdId           string `json:"ordId"`           // 订单ID
	ClOrdId         string `json:"clOrdId"`         // 客户自定义订单ID
	Tag             string `json:"tag"`             // 订单标签
	Px              string `json:"px"`              // 委托价格
	Sz              string `json:"sz"`              // 委托数量
	Pnl             string `json:"pnl"`             // 收益
	OrdType         string `json:"ordType"`         // 订单类型
	Side            string `json:"side"`            // 订单方向
	PosSide         string `json:"posSide"`         // 持仓方向
	TdMode          string `json:"tdMode"`          // 交易模式
	AccFillSz       string `json:"accFillSz"`       // 累计成交数量
	FillPx          string `json:"fillPx"`          // 最新成交价格
	TradeId         string `json:"tradeId"`         // 最新成交ID
	FillSz          string `json:"fillSz"`          // 最新成交数量
	FillTime        string `json:"fillTime"`        // 最新成交时间
	AvgPx           string `json:"avgPx"`           // 成交均价
	State           string `json:"state"`           // 订单状态：canceled 撤单成功 live 等待成交 partially_filled 部分成交 filled 完全成交
	Lever           string `json:"lever"`           // 杠杆倍数
	TpTriggerPx     string `json:"tpTriggerPx"`     // 止盈触发价
	TpTriggerPxType string `json:"tpTriggerPxType"` // 止盈触发价类型
	TpOrdPx         string `json:"tpOrdPx"`         // 止盈委托价
	SlTriggerPx     string `json:"slTriggerPx"`     // 止损触发价
	SlTriggerPxType string `json:"slTriggerPxType"` // 止损触发价类型
	SlOrdPx         string `json:"slOrdPx"`         // 止损委托价
	FeeCcy          string `json:"feeCcy"`          // 交易手续费币种
	Fee             string `json:"fee"`             // 订单交易手续费
	RebateCcy       string `json:"rebateCcy"`       // 返佣金币种
	Rebate          string `json:"rebate"`          // 返佣金额
	Category        string `json:"category"`        // 订单种类
	UTime           string `json:"uTime"`           // 订单状态更新时间
	CTime           string `json:"cTime"`           // 订单创建时间
}

// GetOrder 获取订单信息
func (c *Client) GetOrder(instId, ordId, clOrdId string) (*Order, error) {
	path := "/api/v5/trade/order?instId=" + instId
	if ordId != "" {
		path += "&ordId=" + ordId
	}
	if clOrdId != "" {
		path += "&clOrdId=" + clOrdId
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var orders []Order
	if err := ParseResponse(resp, &orders); err != nil {
		return nil, err
	}
	
	if len(orders) == 0 {
		return nil, nil
	}
	
	return &orders[0], nil
}

// GetOrderList 获取未成交订单列表
func (c *Client) GetOrderList(instType, instId string) ([]Order, error) {
	path := "/api/v5/trade/orders-pending"
	
	params := ""
	if instType != "" {
		params += "?instType=" + instType
	}
	if instId != "" {
		if params == "" {
			params += "?instId=" + instId
		} else {
			params += "&instId=" + instId
		}
	}
	path += params
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var orders []Order
	if err := ParseResponse(resp, &orders); err != nil {
		return nil, err
	}
	
	return orders, nil
}

// GetOrderHistory 获取历史订单记录（近7天）
func (c *Client) GetOrderHistory(instType string, limit int) ([]Order, error) {
	path := "/api/v5/trade/orders-history?instType=" + instType
	if limit > 0 {
		path += fmt.Sprintf("&limit=%d", limit)
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var orders []Order
	if err := ParseResponse(resp, &orders); err != nil {
		return nil, err
	}
	
	return orders, nil
}

