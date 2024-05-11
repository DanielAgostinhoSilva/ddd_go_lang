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
		Cpf:  "488.440.270-70",
	}
}

func (suite *CustomerEntityTestSuit) Test_deve_inicializar_um_customer() {
	customer, err := NewCustomer(suite.customerProps)
	suite.Nil(err)
	suite.NotNil(customer)
	suite.Equal("123456", customer.Id())
	suite.Equal("Test Test", customer.Name().Value())
	suite.Equal("48844027070", customer.Cfp().Value())
}

func (suite *CustomerEntityTestSuit) Test_deve_converter_as_propriedades_em_json_string() {
	expectedJson := "{\"id\":\"123456\",\"cpf\":\"48844027070\",\"name\":\"Test Test\"}"
	customer, err := NewCustomer(suite.customerProps)
	suite.Nil(err)
	suite.NotNil(customer)
	jsonData, err := customer.ToJson()
	suite.Nil(err)
	suite.Equal(expectedJson, jsonData)
}

func TestCustomerEntity(t *testing.T) {
	suite.Run(t, new(CustomerEntityTestSuit))
}
