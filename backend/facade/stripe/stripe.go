package stripe

import (
	"fmt"
	"log"

	"github.com/aogen-fiber/backend/facade/middleware"
	"github.com/aogen-fiber/backend/infrastructure/stripe"
	"github.com/gofiber/fiber/v2"
)

// Handler Stripe处理器
type Handler struct {
	stripeService stripe.Service
}

// NewHandler 创建新的Stripe处理器
func NewHandler(stripeService stripe.Service) *Handler {
	return &Handler{
		stripeService: stripeService,
	}
}

// CreateStripeSessionRequest 表示创建 Stripe Session 的请求
type CreateStripeSessionRequest struct {
	SuccessURL string             `json:"success_url"`
	CancelURL  string             `json:"cancel_url"`
	Mode       string             `json:"mode"`
	LineItems  []*stripe.LineItem `json:"line_items"`
}

// CreateStripeSessionResponse 表示创建 Stripe Session 的响应
type CreateStripeSessionResponse struct {
	Success bool   `json:"success"`
	URL     string `json:"url,omitempty"`
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
}

// SubscriptionResponse 表示订阅信息的响应
type SubscriptionResponse struct {
	Success       bool                `json:"success"`
	Subscriptions []*SubscriptionInfo `json:"subscriptions,omitempty"`
	Error         string              `json:"error,omitempty"`
}

// SubscriptionInfo 表示订阅信息
type SubscriptionInfo struct {
	ID                 string `json:"id"`
	Status             string `json:"status"`
	CurrentPeriodStart int64  `json:"current_period_start"`
	CurrentPeriodEnd   int64  `json:"current_period_end"`
	CancelAtPeriodEnd  bool   `json:"cancel_at_period_end"`
	PriceID            string `json:"price_id"`
}

// CancelSubscriptionResponse 表示取消订阅的响应
type CancelSubscriptionResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

// CreateSession 处理创建 Stripe Session 的请求
func (h *Handler) CreateSession(c *fiber.Ctx) error {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		log.Printf("[Stripe] CreateSession failed: unauthorized access attempt")
		return c.Status(fiber.StatusUnauthorized).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "Authentication required",
			Code:    "AUTH_REQUIRED",
		})
	}

	log.Printf("[Stripe] CreateSession: processing request for user %s", user.Email)

	// 检查用户是否已有有效订阅
	subscriptions, err := h.stripeService.ListActiveSubscriptions(user.Email)
	if err != nil {
		log.Printf("[Stripe] CreateSession failed: error checking existing subscriptions - %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to check existing subscriptions: %v", err),
			Code:    "SUBSCRIPTION_CHECK_FAILED",
		})
	}

	if len(subscriptions) > 0 {
		log.Printf("[Stripe] CreateSession failed: user %s already has active subscription", user.Email)
		return c.Status(fiber.StatusBadRequest).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "You already have an active subscription",
			Code:    "SUBSCRIPTION_EXISTS",
		})
	}

	var req CreateStripeSessionRequest
	if err := c.BodyParser(&req); err != nil {
		log.Printf("[Stripe] CreateSession failed: invalid request data - %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "Invalid request data",
			Code:    "INVALID_REQUEST",
		})
	}

	// 验证请求数据
	if req.SuccessURL == "" || req.CancelURL == "" {
		return c.Status(fiber.StatusBadRequest).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "Success URL and Cancel URL are required",
			Code:    "MISSING_URLS",
		})
	}

	if req.Mode == "" {
		return c.Status(fiber.StatusBadRequest).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "Mode is required",
			Code:    "MISSING_MODE",
		})
	}

	if len(req.LineItems) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   "At least one line item is required",
			Code:    "MISSING_LINE_ITEMS",
		})
	}

	// 创建 Stripe Session
	session, err := h.stripeService.CreateCheckoutSession(&stripe.CreateCheckoutSessionParams{
		Email:      user.Email,
		SuccessURL: req.SuccessURL,
		CancelURL:  req.CancelURL,
		Mode:       req.Mode,
		LineItems:  req.LineItems,
	})

	if err != nil {
		log.Printf("[Stripe] CreateSession failed for user %s: %v", user.Email, err)
		return c.Status(fiber.StatusInternalServerError).JSON(CreateStripeSessionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to create checkout session: %v", err),
			Code:    "SESSION_CREATE_FAILED",
		})
	}

	log.Printf("[Stripe] CreateSession successful for user %s: session URL created", user.Email)
	return c.JSON(CreateStripeSessionResponse{
		Success: true,
		URL:     session.URL,
	})
}

// GetActiveSubscriptions 获取当前用户的有效订阅
func (h *Handler) GetActiveSubscriptions(c *fiber.Ctx) error {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		log.Printf("[Stripe] GetActiveSubscriptions failed: unauthorized access attempt")
		return c.Status(fiber.StatusUnauthorized).JSON(SubscriptionResponse{
			Success: false,
			Error:   "Authentication required",
		})
	}

	log.Printf("[Stripe] GetActiveSubscriptions: fetching subscriptions for user %s", user.Email)

	subscriptions, err := h.stripeService.ListActiveSubscriptions(user.Email)
	if err != nil {
		log.Printf("[Stripe] GetActiveSubscriptions failed for user %s: %v", user.Email, err)
		return c.Status(fiber.StatusInternalServerError).JSON(SubscriptionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to list subscriptions: %v", err),
		})
	}

	log.Printf("[Stripe] GetActiveSubscriptions successful for user %s: found %d subscriptions", user.Email, len(subscriptions))

	// 转换为响应格式
	var subInfos []*SubscriptionInfo
	for _, sub := range subscriptions {
		// 获取订阅的第一个项目的价格ID
		var priceID string
		if len(sub.Items.Data) > 0 {
			priceID = sub.Items.Data[0].Price.ID
		}

		subInfos = append(subInfos, &SubscriptionInfo{
			ID:                 sub.ID,
			Status:             string(sub.Status),
			CurrentPeriodStart: sub.CurrentPeriodStart,
			CurrentPeriodEnd:   sub.CurrentPeriodEnd,
			CancelAtPeriodEnd:  sub.CancelAtPeriodEnd,
			PriceID:            priceID,
		})
	}

	return c.JSON(SubscriptionResponse{
		Success:       true,
		Subscriptions: subInfos,
	})
}

// CancelSubscription 处理取消订阅的请求
func (h *Handler) CancelSubscription(c *fiber.Ctx) error {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		log.Printf("[Stripe] CancelSubscription failed: unauthorized access attempt")
		return c.Status(fiber.StatusUnauthorized).JSON(CancelSubscriptionResponse{
			Success: false,
			Error:   "Authentication required",
		})
	}

	subscriptionID := c.Params("id")
	log.Printf("[Stripe] CancelSubscription: processing cancellation request for subscription %s by user %s", subscriptionID, user.Email)

	if subscriptionID == "" {
		log.Printf("[Stripe] CancelSubscription failed: missing subscription ID")
		return c.Status(fiber.StatusBadRequest).JSON(CancelSubscriptionResponse{
			Success: false,
			Error:   "Subscription ID is required",
		})
	}

	// 获取用户的有效订阅
	subscriptions, err := h.stripeService.ListActiveSubscriptions(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(CancelSubscriptionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to verify subscription: %v", err),
		})
	}

	// 验证订阅是否属于当前用户
	var found bool
	for _, sub := range subscriptions {
		if sub.ID == subscriptionID {
			found = true
			break
		}
	}

	if !found {
		log.Printf("[Stripe] CancelSubscription failed: subscription %s not found or not owned by user %s", subscriptionID, user.Email)
		return c.Status(fiber.StatusForbidden).JSON(CancelSubscriptionResponse{
			Success: false,
			Error:   "Subscription not found or not owned by current user",
		})
	}

	// 取消订阅
	_, err = h.stripeService.CancelSubscription(subscriptionID, &stripe.CancelSubscriptionParams{
		InvoiceNow: false,
		Prorate:    false,
	})

	if err != nil {
		log.Printf("[Stripe] CancelSubscription failed for subscription %s: %v", subscriptionID, err)
		return c.Status(fiber.StatusInternalServerError).JSON(CancelSubscriptionResponse{
			Success: false,
			Error:   fmt.Sprintf("Failed to cancel subscription: %v", err),
		})
	}

	log.Printf("[Stripe] CancelSubscription successful for subscription %s by user %s", subscriptionID, user.Email)
	return c.JSON(CancelSubscriptionResponse{
		Success: true,
	})
}
