package entities

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CustomerEntityTestSuit struct {
	suite.Suite
	customerProps *CustomerConstructorProps
}

func (suite *CustomerEntityTestSuit) SetupTest() {
	suite.customerProps = &CustomerConstructorProps{
		Id:   "123456",
		Name: "Test Test",
		Cpf:  "12345678901",
	}
}

func (suite *CustomerEntityTestSuit) Test_deve_inicializar_um_customer() {
	customer := NewCustomer(suite.customerProps)
	suite.NotNil(customer)
	suite.Equal("123456", customer.Id())
	suite.Equal("Test Test", customer.Name())
	suite.Equal("12345678901", customer.Cfp())
}

func (suite *CustomerEntityTestSuit) Test_deve_converter_as_propriedades_em_json_string() {
	expectedJson := "{\"id\":\"123456\",\"cpf\":\"12345678901\",\"name\":\"Test Test\"}"
	customer := NewCustomer(suite.customerProps)
	suite.NotNil(customer)
	jsonData, err := customer.ToJson()
	suite.Nil(err)
	suite.Equal(expectedJson, jsonData)
}

func TestCustomerEntity(t *testing.T) {
	suite.Run(t, new(CustomerEntityTestSuit))
}
