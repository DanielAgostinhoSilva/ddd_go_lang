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
		Id:   "3221e70d-2368-49c9-b833-0ff4f4149507",
		Name: "Test Test",
		Cpf:  "48844027070",
	}
}

func (suite *CustomerEntityTestSuit) Test_deve_inicializar_um_customer() {
	customer, err := NewCustomer(suite.customerProps)
	suite.Nil(err)
	suite.NotNil(customer)
	suite.Equal(suite.customerProps.Id, customer.Id().value.String())
	suite.Equal(suite.customerProps.Name, customer.Name().Value())
	suite.Equal(suite.customerProps.Cpf, customer.Cfp().Value())
}

func (suite *CustomerEntityTestSuit) Test_deve_converter_as_propriedades_em_json_string() {
	expectedJson := "{\"value\":\"3221e70d-2368-49c9-b833-0ff4f4149507\",\"cpf\":\"48844027070\",\"name\":\"Test Test\"}"
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
