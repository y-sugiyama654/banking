package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	appErr := request.Validate()

	// Assert
	if appErr.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing transaction type.")
	}
	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction type.")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	request := TransactionRequest{
		TransactionType: DEPOSIT,
		Amount:          -100,
	}

	appErr := request.Validate()

	if appErr.Message != "Amount cannot be less than zero." {
		t.Error("Invalid message while validating amount.")
	}
	if appErr.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while validating amount.")
	}
}
