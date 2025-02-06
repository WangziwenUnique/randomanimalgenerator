package stripe

import (
	"fmt"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"github.com/stripe/stripe-go/v81/customer"
	"github.com/stripe/stripe-go/v81/subscription"
)

// Service Stripe服务接口
type Service interface {
	CreateCheckoutSession(params *CreateCheckoutSessionParams) (*stripe.CheckoutSession, error)
	CancelSubscription(subscriptionID string, params *CancelSubscriptionParams) (*stripe.Subscription, error)
	CreateCustomer(params *CreateCustomerParams) (*stripe.Customer, error)
	FindOrCreateCustomer(email string) (*stripe.Customer, error)
	ListActiveSubscriptions(email string) ([]*stripe.Subscription, error)
}

// CreateCheckoutSessionParams 创建支付会话的参数
type CreateCheckoutSessionParams struct {
	Email      string
	SuccessURL string
	CancelURL  string
	LineItems  []*LineItem
	Mode       string // payment, subscription, or setup
}

// CancelSubscriptionParams 取消订阅的参数
type CancelSubscriptionParams struct {
	InvoiceNow bool // 是否立即生成发票
	Prorate    bool // 是否按比例退款
}

// CreateCustomerParams 创建客户的参数
type CreateCustomerParams struct {
	Email       string
	Name        string
	Description string
	Phone       string
	Metadata    map[string]string
}

// LineItem 商品项
type LineItem struct {
	PriceID  string
	Quantity int64
}

type service struct {
	secretKey string
}

// NewService 创建新的Stripe服务实例
func NewService(secretKey string) Service {
	stripe.Key = secretKey
	return &service{
		secretKey: secretKey,
	}
}

// CreateCheckoutSession 创建支付会话
func (s *service) CreateCheckoutSession(params *CreateCheckoutSessionParams) (*stripe.CheckoutSession, error) {
	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, item := range params.LineItems {
		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceID),
			Quantity: stripe.Int64(item.Quantity),
		})
	}

	// 查找或创建客户
	customer, err := s.FindOrCreateCustomer(params.Email)
	if err != nil {
		return nil, err
	}

	sessionParams := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String(params.SuccessURL),
		CancelURL:  stripe.String(params.CancelURL),
		Mode:       stripe.String(params.Mode),
		LineItems:  lineItems,
		Customer:   stripe.String(customer.ID),
	}

	return session.New(sessionParams)
}

// CancelSubscription 取消订阅
func (s *service) CancelSubscription(subscriptionID string, params *CancelSubscriptionParams) (*stripe.Subscription, error) {
	cancelParams := &stripe.SubscriptionCancelParams{
		InvoiceNow: stripe.Bool(false),
		Prorate:    stripe.Bool(false),
	}

	return subscription.Cancel(subscriptionID, cancelParams)
}

// CreateCustomer 创建客户
func (s *service) CreateCustomer(params *CreateCustomerParams) (*stripe.Customer, error) {
	customerParams := &stripe.CustomerParams{
		Email:       stripe.String(params.Email),
		Name:        stripe.String(params.Name),
		Description: stripe.String(params.Description),
		Phone:       stripe.String(params.Phone),
	}

	// 添加元数据
	if params.Metadata != nil {
		for key, value := range params.Metadata {
			customerParams.AddMetadata(key, value)
		}
	}

	return customer.New(customerParams)
}

// FindOrCreateCustomer 根据邮箱查找客户，如果不存在则创建
func (s *service) FindOrCreateCustomer(email string) (*stripe.Customer, error) {
	// 根据邮箱查询客户
	customers := customer.List(&stripe.CustomerListParams{
		Email: stripe.String(email),
	})

	// 获取第一个匹配的客户
	if customers.Next() {
		return customers.Customer(), nil
	}

	// 如果客户不存在，创建新客户
	return s.CreateCustomer(&CreateCustomerParams{
		Email: email,
	})
}

// ListActiveSubscriptions 获取客户的有效订阅列表
func (s *service) ListActiveSubscriptions(email string) ([]*stripe.Subscription, error) {
	// 先查找或创建客户
	customer, err := s.FindOrCreateCustomer(email)
	if err != nil {
		return nil, fmt.Errorf("failed to find or create customer: %v", err)
	}

	params := &stripe.SubscriptionListParams{
		Customer: stripe.String(customer.ID),
		Status:   stripe.String("active"), // 只获取有效的订阅
	}

	var subscriptions []*stripe.Subscription
	i := subscription.List(params)
	for i.Next() {
		subscriptions = append(subscriptions, i.Subscription())
	}

	return subscriptions, i.Err()
}
