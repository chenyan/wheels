# OKX API Go 客户端

OKX 交易所 API v5 的 Go 语言实现，支持账户查询、交易下单、行情获取等功能。

## 功能特性

- ✅ 完整的鉴权机制（HMAC SHA256 签名）
- ✅ 账户管理（余额查询、持仓查询、账户配置等）
- ✅ 交易功能（下单、撤单、订单查询等）
- ✅ 行情数据（Ticker、深度、K线等）
- ✅ 支持模拟交易环境
- ✅ 使用 `reqs` 模块进行 HTTP 请求

## 安装

```bash
go get github.com/jimosrv/wheels/external/okx
```

## 快速开始

### 1. 创建客户端

```go
import "github.com/jimosrv/wheels/external/okx"

// 创建客户端
client := okx.NewClient("your-api-key", "your-secret-key", "your-passphrase")

// 可选：使用模拟交易环境
client.SetSimulated(true)

// 可选：使用 AWS 服务器
client.SetBaseURL(okx.BaseURLAWS)
```

### 2. 获取账户信息

```go
// 获取账户余额
balances, err := client.GetAccountBalance()
if err != nil {
    log.Fatal(err)
}

for _, balance := range balances {
    fmt.Printf("总权益: %s USD\n", balance.TotalEq)
    for _, detail := range balance.Details {
        fmt.Printf("%s: %s (可用: %s)\n", 
            detail.Ccy, detail.Eq, detail.AvailBal)
    }
}

// 获取持仓信息
positions, err := client.GetAccountPositions("SWAP", "BTC-USDT-SWAP")
if err != nil {
    log.Fatal(err)
}

// 获取账户配置
config, err := client.GetAccountConfig()
if err != nil {
    log.Fatal(err)
}
fmt.Printf("账户模式: %s\n", config.AcctLv)
```

### 3. 交易下单

```go
// 限价买入
orderReq := okx.OrderRequest{
    InstId:  "BTC-USDT",
    TdMode:  okx.TdModeCash,
    Side:    okx.SideBuy,
    OrdType: okx.OrdTypeLimit,
    Px:      "40000",      // 价格
    Sz:      "0.001",      // 数量
}

result, err := client.PlaceOrder(orderReq)
if err != nil {
    log.Fatal(err)
}
fmt.Printf("订单ID: %s\n", result.OrdId)

// 市价卖出
marketOrder := okx.OrderRequest{
    InstId:  "BTC-USDT",
    TdMode:  okx.TdModeCash,
    Side:    okx.SideSell,
    OrdType: okx.OrdTypeMarket,
    Sz:      "0.001",
}

result, err = client.PlaceOrder(marketOrder)
```

### 4. 订单管理

```go
// 查询订单
order, err := client.GetOrder("BTC-USDT", "order-id", "")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("订单状态: %s\n", order.State)

// 获取未成交订单列表
orders, err := client.GetOrderList("SPOT", "")
if err != nil {
    log.Fatal(err)
}

// 撤销订单
cancelReq := okx.CancelOrderRequest{
    InstId: "BTC-USDT",
    OrdId:  "order-id",
}
cancelResult, err := client.CancelOrder(cancelReq)
```

### 5. 行情数据

```go
// 获取单个产品行情
ticker, err := client.GetTicker("BTC-USDT")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("最新价: %s\n", ticker.Last)
fmt.Printf("买一价: %s, 卖一价: %s\n", ticker.BidPx, ticker.AskPx)

// 获取所有现货行情
tickers, err := client.GetTickers("SPOT")

// 获取深度数据
orderbook, err := client.GetOrderbook("BTC-USDT", 20)
if err != nil {
    log.Fatal(err)
}

// 获取K线数据
candles, err := client.GetCandles("BTC-USDT", "1m", 100)
```

### 6. 杠杆设置

```go
// 设置杠杆倍数
req := okx.SetLeverageRequest{
    InstId:  "BTC-USDT-SWAP",
    Lever:   "10",
    MgnMode: okx.TdModeCross,
}

result, err := client.SetLeverage(req)
if err != nil {
    log.Fatal(err)
}
```

## API 常量

```go
// 产品类型
okx.InstTypeSpot     // 币币
okx.InstTypeMargin   // 币币杠杆
okx.InstTypeSwap     // 永续合约
okx.InstTypeFutures  // 交割合约
okx.InstTypeOption   // 期权

// 交易模式
okx.TdModeCash       // 非保证金
okx.TdModeCross      // 全仓
okx.TdModeIsolated   // 逐仓

// 订单方向
okx.SideBuy          // 买入
okx.SideSell         // 卖出

// 订单类型
okx.OrdTypeMarket    // 市价单
okx.OrdTypeLimit     // 限价单
okx.OrdTypePostOnly  // 只做maker
okx.OrdTypeFok       // 全部成交或立即取消
okx.OrdTypeIoc       // 立即成交并取消剩余

// 持仓方向
okx.PosSideLong      // 多头
okx.PosSideShort     // 空头
okx.PosSideNet       // 单向持仓
```

## 获取 API Key

1. 登录 [OKX 官网](https://www.okx.com)
2. 进入用户中心 -> API 管理
3. 创建 API Key，获取：
   - API Key
   - Secret Key
   - Passphrase

**注意**：Secret Key 只显示一次，请妥善保存。

## 安全提示

- 不要将 API Key 硬编码在代码中
- 建议使用环境变量或配置文件存储敏感信息
- 为 API Key 设置合理的权限（如只读、交易等）
- 建议使用 IP 白名单限制 API 访问
- 先在模拟交易环境中测试

## API 文档

详细的 API 文档请参考：
- [OKX API 文档](https://www.okx.com/docs-v5/zh/)
- [API 鉴权说明](https://www.okx.com/docs-v5/zh/#overview-v5-api-key-creation-generating-an-api-key)

## 已实现的接口

### 账户接口
- ✅ 获取账户余额 `GetAccountBalance`
- ✅ 获取持仓信息 `GetAccountPositions`
- ✅ 获取账户配置 `GetAccountConfig`
- ✅ 设置杠杆倍数 `SetLeverage`

### 交易接口
- ✅ 下单 `PlaceOrder`
- ✅ 撤单 `CancelOrder`
- ✅ 查询订单 `GetOrder`
- ✅ 获取未成交订单列表 `GetOrderList`
- ✅ 获取历史订单 `GetOrderHistory`

### 行情接口
- ✅ 获取单个产品行情 `GetTicker`
- ✅ 获取所有产品行情 `GetTickers`
- ✅ 获取深度数据 `GetOrderbook`
- ✅ 获取K线数据 `GetCandles`
- ✅ 获取产品信息 `GetInstruments`

## 许可证

MIT License

