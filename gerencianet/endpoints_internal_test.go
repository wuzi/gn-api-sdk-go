package gerencianet

import (
	"testing"
)

type MockRequester struct {}

func (m *MockRequester) request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error) {
	return "", nil
}

func Test_endpoints(t *testing.T) {
	requester := &MockRequester{}
	endpoints := endpoints{requester}
	endpoints.CreateCharge(nil)
	endpoints.DetailCharge(1)
	endpoints.UpdateChargeMetadata(1, nil)
	endpoints.UpdateBillet(1, nil)
	endpoints.PayCharge(1, nil)
	endpoints.CancelCharge(1)
	endpoints.CreateCarnet(nil)
	endpoints.DetailCarnet(1)
	endpoints.UpdateParcel(1, 1, nil)
	endpoints.UpdateCarnetMetadata(1, nil)
	endpoints.GetNotification("")
	endpoints.GetPlans(1, 1)
	endpoints.CreatePlan(nil)
	endpoints.DeletePlan(1)
	endpoints.CreateSubscription(1, nil)
	endpoints.DetailSubscription(1)
	endpoints.PaySubscription(1, nil)
	endpoints.CancelSubscription(1)
	endpoints.UpdateSubscriptionMetadata(1, nil)
	endpoints.GetInstallments(20000, "visa")
	endpoints.ResendBillet(1, nil)
	endpoints.CreateChargeHistory(1, nil)
	endpoints.ResendCarnet(1, nil)
	endpoints.ResendParcel(1, 1, nil)
	endpoints.CreateCarnetHistory(1, nil)
	endpoints.CancelCarnet(1)
	endpoints.CancelParcel(1, 1)
	endpoints.ChargeLink(1, nil)
	endpoints.UpdateChargeLink(1, nil)
	endpoints.UpdatePlan(1, nil)
	endpoints.CreateSubscriptionHistory(1, nil)
	endpoints.CreateChargeBalanceSheet(1, nil)
	t.Skip("skipping endpoints tests")
}