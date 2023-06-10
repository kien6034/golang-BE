package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/kien6034/golang-BE/db/mock"
	db "github.com/kien6034/golang-BE/db/sqlc"
	"github.com/kien6034/golang-BE/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)

	// Build stubs
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil) // Expect the GetAccount to be call 1 times

	// Srart test server and request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest("GET", url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// Check response
	require.Equal(t, http.StatusOK, recorder.Code)
	requireBodyMatchAccount(t, recorder.Body, account)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)

	require.Equal(t, account, gotAccount)
}
