package okx

// Package okx 实现了 OKX 交易所 API v5 的 Go 客户端
// 支持账户查询、交易下单、行情获取等功能
//
// 使用方法：
//   client := okx.NewClient("api-key", "secret-key", "passphrase")
//   balance, err := client.GetAccountBalance()

const (
	// 产品类型
	InstTypeSpot    = "SPOT"    // 币币
	InstTypeMargin  = "MARGIN"  // 币币杠杆
	InstTypeSwap    = "SWAP"    // 永续合约
	InstTypeFutures = "FUTURES" // 交割合约
	InstTypeOption  = "OPTION"  // 期权

	// 交易模式
	TdModeCash     = "cash"     // 非保证金
	TdModeCross    = "cross"    // 全仓
	TdModeIsolated = "isolated" // 逐仓

	// 订单方向
	SideBuy  = "buy"  // 买入
	SideSell = "sell" // 卖出

	// 订单类型
	OrdTypeMarket   = "market"    // 市价单
	OrdTypeLimit    = "limit"     // 限价单
	OrdTypePostOnly = "post_only" // 只做maker
	OrdTypeFok      = "fok"       // 全部成交或立即取消
	OrdTypeIoc      = "ioc"       // 立即成交并取消剩余

	// 持仓方向
	PosSideLong  = "long"  // 多头
	PosSideShort = "short" // 空头
	PosSideNet   = "net"   // 单向持仓
)
