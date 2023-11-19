package service

import (
	"fmt"
	"github.com/b0gochort/admin_panel/internal/api_db"
	"github.com/b0gochort/admin_panel/internal/model"
)

type AdminServiceImpl struct {
	adminAPI api_db.AdminAPI
}

func NewAdminService(adminAPI api_db.AdminAPI) *AdminServiceImpl {
	return &AdminServiceImpl{
		adminAPI: adminAPI,
	}
}

func (s *AdminServiceImpl) GetCompetent() ([]model.CompetentRes, error) {
	competents, err := s.adminAPI.TakeCompetent()
	if err != nil {
		return nil, fmt.Errorf("AdminServiceImpl.GetCompetent: %v", err)
	}

	res := make([]model.CompetentRes, 0)

	for _, elem := range competents {
		competent := model.CompetentRes{
			ID:                     elem.ID,
			UID:                    elem.UID,
			PaymentIssue:           elem.PaymentIssue,
			CreateAccount:          elem.CreateAccount,
			ContactCustomerService: elem.ContactCustomerService,
			GetInvoice:             elem.GetInvoice,
			TrackOrder:             elem.TrackOrder,
			GetRefund:              elem.GetRefund,
			ContactHumanObject:     elem.ContactHumanObject,
			RecoverPassword:        elem.RecoverPassword,
			ChangeOrder:            elem.ChangeOrder,
			DeleteAccount:          elem.DeleteAccount,
			Complaint:              elem.Complaint,
			CheckInvoices:          elem.CheckInvoice,
			Review:                 elem.Review,
			CheckRefundPolicy:      elem.CheckRefundPolicy,
			DeliveryOptions:        elem.DeliveryOptions,
			CheckCancellationFee:   elem.CheckCancellationFee,
			TrackRefund:            elem.TrackRefund,
			CheckPaymentMethods:    elem.CheckPaymentMethods,
			SwitchAccount:          elem.SwitchAccount,
			NewsletterSubscription: elem.NewsletterSubscription,
			DeliveryPeriod:         elem.DeliveryPeriod,
			EditAccount:            elem.EditAccount,
			RegistrationProblems:   elem.RegistrationProblems,
			ChangeShippingAddress:  elem.ChangeShippingAddress,
			SetUpShippingAddress:   elem.SetUpShippingAddress,
			PlaceOrder:             elem.PlaceOrder,
			CancelOrder:            elem.CancelOrder,
			CheckInvoice:           elem.CheckInvoices,
		}
		res = append(res, competent)
	}

	return res, nil
}

func (s AdminServiceImpl) ChangeCompetent(elem model.CompetentRes) (model.CompetentRes, error) {
	cmpt := model.CompetentItem{
		ID:                     elem.ID,
		UID:                    elem.UID,
		PaymentIssue:           elem.PaymentIssue,
		CreateAccount:          elem.CreateAccount,
		ContactCustomerService: elem.ContactCustomerService,
		GetInvoice:             elem.GetInvoice,
		TrackOrder:             elem.TrackOrder,
		GetRefund:              elem.GetRefund,
		ContactHumanObject:     elem.ContactHumanObject,
		RecoverPassword:        elem.RecoverPassword,
		ChangeOrder:            elem.ChangeOrder,
		DeleteAccount:          elem.DeleteAccount,
		Complaint:              elem.Complaint,
		CheckInvoices:          elem.CheckInvoice,
		Review:                 elem.Review,
		CheckRefundPolicy:      elem.CheckRefundPolicy,
		DeliveryOptions:        elem.DeliveryOptions,
		CheckCancellationFee:   elem.CheckCancellationFee,
		TrackRefund:            elem.TrackRefund,
		CheckPaymentMethods:    elem.CheckPaymentMethods,
		SwitchAccount:          elem.SwitchAccount,
		NewsletterSubscription: elem.NewsletterSubscription,
		DeliveryPeriod:         elem.DeliveryPeriod,
		EditAccount:            elem.EditAccount,
		RegistrationProblems:   elem.RegistrationProblems,
		ChangeShippingAddress:  elem.ChangeShippingAddress,
		SetUpShippingAddress:   elem.SetUpShippingAddress,
		PlaceOrder:             elem.PlaceOrder,
		CancelOrder:            elem.CancelOrder,
		CheckInvoice:           elem.CheckInvoices,
	}
	_, err := s.adminAPI.ChangeCompetent(cmpt)
	if err != nil {
		return model.CompetentRes{}, fmt.Errorf("AdminServiceImpl.ChangeCompetent: %v", err)
	}

	return elem, nil
}
func (s *AdminServiceImpl) AddCompetent(userId int64) error {
	if err := s.adminAPI.AddCompetent(userId); err != nil {
		return fmt.Errorf("AdminServiceImpl.AddCompetent: %v", err)
	}

	return nil
}
