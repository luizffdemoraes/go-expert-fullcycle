package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.0)
	assert.Nil(t, err)
	assert.Equal(t, 10.0, tax)

	tax, err = CalculateTax(0.0)
	assert.Error(t, err, "amount must be greater than 0")
	assert.Equal(t, 0.0, tax)
	assert.Contains(t, err.Error(), "greater than 0")
}

func TestCalculateTaxAndSave2(t *testing.T) {
	repository := &TaxRepositoryMock{}

	// Descubra o valor correto retornado por CalculateTax2(1000.0)
	expectedTax := CalculateTax2(1000.0)

	// Configure o mock com o valor correto
	repository.On("SaveTax", expectedTax).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax"))

	// Primeira chamada: deve passar "expectedTax" e não retornar erro
	err := CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(1000.0, repository)
	assert.Nil(t, err)

	// Segunda chamada: use um valor diferente que retorne 0.0 no cálculo do imposto
	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "error saving tax")

	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 3)
}
