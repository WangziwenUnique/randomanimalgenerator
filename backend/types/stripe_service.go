package types

import "github.com/stripe/stripe-go/v81"

// StripeService 定义Stripe服务接口
type StripeService interface {
	ListActiveSubscriptions(email string) ([]*stripe.Subscription, error)
}
