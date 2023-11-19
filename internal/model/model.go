package model

type CompetentItem struct {
	ID                     int64 `json:"id" reindex:"id,hash,pk"`
	UID                    int64 `json:"uid" reindex:"uid,hash"`
	PaymentIssue           int32 `json:"payment_issue" reindex:"payment_issue,hash"`
	CreateAccount          int32 `json:"create_account" reindex:"create_account,hash"`
	ContactCustomerService int32 `json:"contact_customer_service" reindex:"contact_customer_service,hash"`
	GetInvoice             int32 `json:"get_invoice" reindex:"get_invoice,hash"`
	TrackOrder             int32 `json:"track_order" reindex:"track_order,hash"`
	GetRefund              int32 `json:"get_refund" reindex:"get_refund,hash"`
	ContactHumanObject     int32 `json:"contact_human_object" reindex:"contact_human_object,hash"`
	RecoverPassword        int32 `json:"recover_password" reindex:"recover_password,hash"`
	ChangeOrder            int32 `json:"change_order" reindex:"change_order,hash"`
	DeleteAccount          int32 `json:"delete_account" reindex:"delete_account,hash"`
	Complaint              int32 `json:"complaint" reindex:"complaint,hash"`
	CheckInvoices          int32 `json:"check_invoices" reindex:"check_invoices,hash"`
	Review                 int32 `json:"review" reindex:"review,hash"`
	CheckRefundPolicy      int32 `json:"check_refund_policy" reindex:"check_refund_policy,hash"`
	DeliveryOptions        int32 `json:"delivery_options" reindex:"delivery_options,hash"`
	CheckCancellationFee   int32 `json:"check_cancellation_fee" reindex:"check_cancellation_fee,hash"`
	TrackRefund            int32 `json:"track_refund" reindex:"track_refund,hash"`
	CheckPaymentMethods    int32 `json:"check_payment_methods" reindex:"check_payment_methods,hash"`
	SwitchAccount          int32 `json:"switch_account" reindex:"switch_account,hash"`
	NewsletterSubscription int32 `json:"newsletter_subscription" reindex:"newsletter_subscription,hash"`
	DeliveryPeriod         int32 `json:"delivery_period" reindex:"delivery_period,hash"`
	EditAccount            int32 `json:"edit_account" reindex:"edit_account,hash"`
	RegistrationProblems   int32 `json:"registration_problems" reindex:"registration_problems,hash"`
	ChangeShippingAddress  int32 `json:"change_shipping_address" reindex:"change_shipping_address,hash"`
	SetUpShippingAddress   int32 `json:"set_up_shipping_address" reindex:"set_up_shipping_address,hash"`
	PlaceOrder             int32 `json:"place_order" reindex:"place_order,hash"`
	CancelOrder            int32 `json:"cancel_order" reindex:"cancel_order,hash"`
	CheckInvoice           int32 `json:"check_invoice" reindex:"check_invoice,hash"`
}

type CompetentRes struct {
	ID                     int64 `json:"id"`
	UID                    int64 `json:"uid"`
	PaymentIssue           int32 `json:"payment_issue"`
	CreateAccount          int32 `json:"create_account"`
	ContactCustomerService int32 `json:"contact_customer_service"`
	GetInvoice             int32 `json:"get_invoice"`
	TrackOrder             int32 `json:"track_order"`
	GetRefund              int32 `json:"get_refund"`
	ContactHumanObject     int32 `json:"contact_human_object"`
	RecoverPassword        int32 `json:"recover_password"`
	ChangeOrder            int32 `json:"change_order"`
	DeleteAccount          int32 `json:"delete_account"`
	Complaint              int32 `json:"complaint"`
	CheckInvoices          int32 `json:"check_invoices"`
	Review                 int32 `json:"review"`
	CheckRefundPolicy      int32 `json:"check_refund_policy"`
	DeliveryOptions        int32 `json:"delivery_options"`
	CheckCancellationFee   int32 `json:"check_cancellation_fee"`
	TrackRefund            int32 `json:"track_refund"`
	CheckPaymentMethods    int32 `json:"check_payment_methods"`
	SwitchAccount          int32 `json:"switch_account"`
	NewsletterSubscription int32 `json:"newsletter_subscription"`
	DeliveryPeriod         int32 `json:"delivery_period"`
	EditAccount            int32 `json:"edit_account"`
	RegistrationProblems   int32 `json:"registration_problems"`
	ChangeShippingAddress  int32 `json:"change_shipping_address"`
	SetUpShippingAddress   int32 `json:"set_up_shipping_address"`
	PlaceOrder             int32 `json:"place_order"`
	CancelOrder            int32 `json:"cancel_order"`
	CheckInvoice           int32 `json:"check_invoice"`
}

type AddCompetentReq struct {
	Uid int64 `json:"uid"`
}

type AddCompetentRes struct {
	Uid    int64  `json:"uid"`
	Status string `json:"status"`
}
