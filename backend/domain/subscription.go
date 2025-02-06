package domain

// SubscriptionPlan 定义订阅计划的配置
type SubscriptionPlan struct {
	MonthlyPriceID  string      `json:"monthly_price_id"`
	MonthlyPrice    string      `json:"monthly_price"`
	YearlyPriceID   string      `json:"yearly_price_id"`
	YearlyPrice     string      `json:"yearly_price"`
	PreProcessWords interface{} `json:"pre_process_words"` // 可以是数字或 "unlimited"
	Words           int         `json:"words"`
}

// SubscriptionConfig 定义所有订阅计划的配置
type SubscriptionConfig struct {
	Lite     SubscriptionPlan `json:"lite"`
	Standard SubscriptionPlan `json:"standard"`
	Pro      SubscriptionPlan `json:"pro"`
}

// DefaultSubscriptionConfig 返回默认的订阅配置
func DefaultSubscriptionConfig() *SubscriptionConfig {
	return &SubscriptionConfig{
		Lite: SubscriptionPlan{
			MonthlyPriceID:  "price_1QjJuGGRvSbloEkVm7ePymM5",
			MonthlyPrice:    "19",
			YearlyPriceID:   "price_1QjJuGGRvSbloEkVmXzcnP4F",
			YearlyPrice:     "114",
			PreProcessWords: 500,
			Words:           20000,
		},
		Standard: SubscriptionPlan{
			MonthlyPriceID:  "price_1QjJuDGRvSbloEkV3Y8FHD0l",
			MonthlyPrice:    "29",
			YearlyPriceID:   "price_1QjJuDGRvSbloEkVCXiSam8F",
			YearlyPrice:     "174",
			PreProcessWords: "unlimited",
			Words:           50000,
		},
		Pro: SubscriptionPlan{
			MonthlyPriceID:  "price_1QjJu9GRvSbloEkVopg5abca",
			MonthlyPrice:    "79",
			YearlyPriceID:   "price_1QjJu9GRvSbloEkV6LGTDnnd",
			YearlyPrice:     "474",
			PreProcessWords: "unlimited",
			Words:           150000,
		},
	}
}
