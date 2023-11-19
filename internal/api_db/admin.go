package api_db

import (
	"fmt"
	"github.com/b0gochort/admin_panel/internal/model"
	"github.com/restream/reindexer/v3"
)

type AdminAPIImpl struct {
	db *reindexer.Reindexer
}

func NewChatApi(db *reindexer.Reindexer) *AdminAPIImpl {
	return &AdminAPIImpl{
		db: db,
	}
}

func (a *AdminAPIImpl) TakeCompetent() ([]model.CompetentItem, error) {
	err := a.db.OpenNamespace("competent", reindexer.DefaultNamespaceOptions(), model.CompetentItem{})
	if err != nil {
		return nil, fmt.Errorf("AdminAPIImpl.TakeModerators.OpenNamespace: %v", err)
	}
	elem := a.db.Query("competent")

	var response []model.CompetentItem

	iter := elem.Exec()
	if iter.Error() != nil {
		return nil, fmt.Errorf("AdminAPIImpl.TakeModerators.Exec: %v", iter.Error())
	}

	for iter.Next() {
		elem := iter.Object().(*model.CompetentItem)
		response = append(response, model.CompetentItem{
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
		})
	}

	return response, nil
}

func (a *AdminAPIImpl) ChangeCompetent(item model.CompetentItem) (model.CompetentItem, error) {
	err := a.db.OpenNamespace("competent", reindexer.DefaultNamespaceOptions(), model.CompetentItem{})
	if err != nil {
		return model.CompetentItem{}, fmt.Errorf("AdminAPIImpl.ChangeCompetent.OpenNamespace: %v", err)
	}

	err = a.db.Upsert("competent", item)
	if err != nil {
		return model.CompetentItem{}, fmt.Errorf("AdminAPIImpl.ChangeCompetent.Upsert: %v", err)
	}

	return item, nil
}

func (a *AdminAPIImpl) AddCompetent(userId int64) error {
	err := a.db.OpenNamespace("general", reindexer.DefaultNamespaceOptions(), model.CompetentItem{})
	if err != nil {
		return fmt.Errorf("AdminAPIImpl.AddCompetents.OpenNamespace: %v", err)
	}

	elem := a.db.Query("general").Set("sub", "competents").Where("uid", reindexer.EQ, userId).Update()
	if elem.Error() != nil {
		return fmt.Errorf("AdminAPIImpl.ChangeCompetent.Upsert: %v", elem.Error())
	}

	return nil
}
