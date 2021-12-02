package process_transaction

import (
	"testing"

	mock_entity "github.com/Fred-Reis/first-go/entity/mock"
	"github.com/golang/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestProcessTransactionWhenItsValid(t *testing.T) {
	input := TransactionDTOInput{
		ID:        "1",
		AccountID: "1",
		Amount:    200,
	}

	expectedOutput := TransactionDTOOutput{
		ID:           "1",
		Status:       "approved",
		ErrorMessage: "",
	}

	crtl := gomock.NewController(t)
	defer crtl.Finish()

	repositoryMock := mock_entity.NewMockTransactionRepository(crtl)
	repositoryMock.EXPECT().Insert(input.ID, input.AccountID, input.Amount, "approved", "").Return(nil)

	usecase := NewProcessTransaction(repositoryMock)
	output, err := usecase.Execute(input)
	assert.Nil(t, err)
	assert.Equal(t, expectedOutput, output)

}
