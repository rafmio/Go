// https://github.com/sgreben/testing-with-gomock/blob/master/user/user_test.go

package user_test

import (
	"testing"

	"github.com/sgreben/testing-with-gomock/mocks"
	"github.com/sgreben/testing-with-gomock/tree/master/user"
)

func TestUse(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)
	testUser := &user.User{Doer: mockDoer}

	// Expect Do to be called once with 123 and "Hello GoMock" as parameters, and return nil from the mocked call
	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

	testUser.Use()
}
