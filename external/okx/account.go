package okx

import "fmt"

// AccountBalance 账户余额
type AccountBalance struct {
	AdjEq       string            `json:"adjEq"`       // 美元层面有效保证金
	Details     []BalanceDetail   `json:"details"`     // 各币种资产详细信息
	Imr         string            `json:"imr"`         // 美元层面占用保证金
	IsoEq       string            `json:"isoEq"`       // 美元层面逐仓仓位权益
	MgnRatio    string            `json:"mgnRatio"`    // 美元层面保证金率
	Mmr         string            `json:"mmr"`         // 美元层面维持保证金
	NotionalUsd string            `json:"notionalUsd"` // 以美元价值为单位的持仓数量
	OrdFroz     string            `json:"ordFroz"`     // 美元层面全仓挂单占用保证金
	TotalEq     string            `json:"totalEq"`     // 美元层面权益
	UTime       string            `json:"uTime"`       // 账户信息的更新时间，Unix时间戳的毫秒数格式
}

// BalanceDetail 余额详情
type BalanceDetail struct {
	AvailBal      string `json:"availBal"`      // 可用余额
	AvailEq       string `json:"availEq"`       // 可用保证金
	CashBal       string `json:"cashBal"`       // 币种余额
	Ccy           string `json:"ccy"`           // 币种
	CrossLiab     string `json:"crossLiab"`     // 币种全仓负债额
	DiskEq        string `json:"disEq"`         // 美元层面币种折算权益
	Eq            string `json:"eq"`            // 币种权益
	EqUsd         string `json:"eqUsd"`         // 币种权益美元价值
	FrozenBal     string `json:"frozenBal"`     // 币种占用金额
	Interest      string `json:"interest"`      // 计息
	IsoEq         string `json:"isoEq"`         // 币种逐仓仓位权益
	IsoLiab       string `json:"isoLiab"`       // 币种逐仓负债额
	IsoUpl        string `json:"isoUpl"`        // 逐仓未实现盈亏
	Liab          string `json:"liab"`          // 币种负债额
	MaxLoan       string `json:"maxLoan"`       // 币种最大可借
	MgnRatio      string `json:"mgnRatio"`      // 保证金率
	NotionalLever string `json:"notionalLever"` // 杠杆倍数
	OrdFrozen     string `json:"ordFrozen"`     // 挂单冻结数量
	Twap          string `json:"twap"`          // 当前负债币种触发系统自动换币的风险
	UTime         string `json:"uTime"`         // 更新时间，Unix时间戳的毫秒数格式
	Upl           string `json:"upl"`           // 未实现盈亏
	UplLiab       string `json:"uplLiab"`       // 由于仓位未实现亏损导致的负债
	StgyEq        string `json:"stgyEq"`        // 策略权益
}

// GetAccountBalance 获取账户余额
func (c *Client) GetAccountBalance(ccy ...string) ([]AccountBalance, error) {
	path := "/api/v5/account/balance"
	if len(ccy) > 0 && ccy[0] != "" {
		path = fmt.Sprintf("%s?ccy=%s", path, ccy[0])
	}
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var balances []AccountBalance
	if err := ParseResponse(resp, &balances); err != nil {
		return nil, err
	}
	
	return balances, nil
}

// AccountPosition 持仓信息
type AccountPosition struct {
	InstType    string `json:"instType"`    // 产品类型
	MgnMode     string `json:"mgnMode"`     // 保证金模式：cross 全仓 isolated 逐仓
	PosId       string `json:"posId"`       // 持仓ID
	PosSide     string `json:"posSide"`     // 持仓方向：long short net
	Pos         string `json:"pos"`         // 持仓数量
	AvailPos    string `json:"availPos"`    // 可平仓数量
	AvgPx       string `json:"avgPx"`       // 开仓平均价
	Upl         string `json:"upl"`         // 未实现收益
	UplRatio    string `json:"uplRatio"`    // 未实现收益率
	InstId      string `json:"instId"`      // 产品ID
	Lever       string `json:"lever"`       // 杠杆倍数
	LiqPx       string `json:"liqPx"`       // 预估强平价
	MarkPx      string `json:"markPx"`      // 最新标记价格
	Imr         string `json:"imr"`         // 占用保证金
	Margin      string `json:"margin"`      // 保证金余额
	MgnRatio    string `json:"mgnRatio"`    // 保证金率
	Mmr         string `json:"mmr"`         // 维持保证金
	Liab        string `json:"liab"`        // 负债额
	LiabCcy     string `json:"liabCcy"`     // 负债币种
	Interest    string `json:"interest"`    // 利息
	TradeId     string `json:"tradeId"`     // 最新成交ID
	NotionalUsd string `json:"notionalUsd"` // 以美元价值为单位的持仓数量
	Adl         string `json:"adl"`         // 信号区
	Ccy         string `json:"ccy"`         // 占用保证金的币种
	Last        string `json:"last"`        // 最新成交价
	UsdPx       string `json:"usdPx"`       // 美元价格（仅适用于期权）
	BePx        string `json:"bePx"`        // 盈亏平衡价
	DeltaBS     string `json:"deltaBS"`     // 美元本位持仓仓位delta
	DeltaPA     string `json:"deltaPA"`     // 币本位持仓仓位delta
	GammaBS     string `json:"gammaBS"`     // 美元本位持仓仓位gamma
	GammaPA     string `json:"gammaPA"`     // 币本位持仓仓位gamma
	ThetaBS     string `json:"thetaBS"`     // 美元本位持仓仓位theta
	ThetaPA     string `json:"thetaPA"`     // 币本位持仓仓位theta
	VegaBS      string `json:"vegaBS"`      // 美元本位持仓仓位vega
	VegaPA      string `json:"vegaPA"`      // 币本位持仓仓位vega
	CTime       string `json:"cTime"`       // 持仓创建时间
	UTime       string `json:"uTime"`       // 最近一次持仓更新时间
	PTime       string `json:"pTime"`       // 持仓信息更新时间
}

// GetAccountPositions 获取持仓信息
func (c *Client) GetAccountPositions(instType, instId string) ([]AccountPosition, error) {
	path := "/api/v5/account/positions"
	
	// 构建查询参数
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
	
	var positions []AccountPosition
	if err := ParseResponse(resp, &positions); err != nil {
		return nil, err
	}
	
	return positions, nil
}

// AccountConfig 账户配置
type AccountConfig struct {
	AcctLv      string `json:"acctLv"`      // 账户层级：1 简单交易模式 2 单币种保证金模式 3 跨币种保证金模式 4 组合保证金模式
	AutoLoan    bool   `json:"autoLoan"`    // 是否自动借币：true false
	CtIsoMode   string `json:"ctIsoMode"`   // 合约逐仓保证金模式：automatic 自动转入 autonomy 自主转入
	GreeksType  string `json:"greeksType"`  // 希腊字母展示方式：PA 币本位 BS 美元本位
	Level       string `json:"level"`       // 当前带单合约的倍数
	LevelTmp    string `json:"levelTmp"`    // 临时带单合约的倍数
	MgnIsoMode  string `json:"mgnIsoMode"`  // 现货对冲逐仓保证金模式：automatic 自动转入 autonomy 自主转入
	PosMode     string `json:"posMode"`     // 持仓方式：long_short_mode 双向持仓 net_mode 单向持仓
	Uid         string `json:"uid"`         // 用户标识
	Label       string `json:"label"`       // APIKey的备注名
	RoleType    string `json:"roleType"`    // 当前APIKey的权限：0 普通用户 1 资管户用户 2 资管户子账户
	TraderInsts []struct {
		InstId string `json:"instId"` // 产品ID
	} `json:"traderInsts"` // 当前带单合约
	KycLv   string `json:"kycLv"`   // 身份认证等级
	IP      string `json:"ip"`      // API绑定ip地址
	Perm    string `json:"perm"`    // API权限
	MainUid string `json:"mainUid"` // 母账户UID
}

// GetAccountConfig 获取账户配置
func (c *Client) GetAccountConfig() (*AccountConfig, error) {
	path := "/api/v5/account/config"
	
	resp, err := c.Get(path)
	if err != nil {
		return nil, err
	}
	
	var configs []AccountConfig
	if err := ParseResponse(resp, &configs); err != nil {
		return nil, err
	}
	
	if len(configs) == 0 {
		return nil, fmt.Errorf("no account config found")
	}
	
	return &configs[0], nil
}

// SetLeverage 设置杠杆倍数
type SetLeverageRequest struct {
	InstId  string `json:"instId"`            // 产品ID
	Ccy     string `json:"ccy,omitempty"`     // 保证金币种，仅适用于跨币种保证金模式下的全仓杠杆
	Lever   string `json:"lever"`             // 杠杆倍数
	MgnMode string `json:"mgnMode"`           // 保证金模式：isolated 逐仓 cross 全仓
	PosSide string `json:"posSide,omitempty"` // 持仓方向：long short，仅适用于逐仓模式
}

// SetLeverageResponse 设置杠杆响应
type SetLeverageResponse struct {
	InstId  string `json:"instId"`  // 产品ID
	Lever   string `json:"lever"`   // 杠杆倍数
	MgnMode string `json:"mgnMode"` // 保证金模式
	PosSide string `json:"posSide"` // 持仓方向
}

// SetLeverage 设置杠杆倍数
func (c *Client) SetLeverage(req SetLeverageRequest) (*SetLeverageResponse, error) {
	path := "/api/v5/account/set-leverage"
	
	resp, err := c.Post(path, req)
	if err != nil {
		return nil, err
	}
	
	var results []SetLeverageResponse
	if err := ParseResponse(resp, &results); err != nil {
		return nil, err
	}
	
	if len(results) == 0 {
		return nil, fmt.Errorf("no response data")
	}
	
	return &results[0], nil
}

