package value_objects

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type IdVoTestSuite struct {
	suite.Suite
}

func (suite *IdVoTestSuite) Test_deve_iniciar_um_id_valido() {
	id, err := NewID("e1f9b24e-c934-4831-b63a-5c884a544a8c")
	suite.Nil(err)
	suite.NotNil(id)
	suite.Equal("e1f9b24e-c934-4831-b63a-5c884a544a8c", id.value.String())

	id, err = NewID("")
	suite.Nil(err)
	suite.NotNil(id)
	suite.NotNil(id.Value())
}

func (suite *IdVoTestSuite) Test_deve_comparar_se_os_id_sao_iguais_ou_nao() {
	idA, err := NewID("e1f9b24e-c934-4831-b63a-5c884a544a8c")
	suite.Nil(err)
	suite.NotNil(idA)
	suite.Equal("e1f9b24e-c934-4831-b63a-5c884a544a8c", idA.value.String())

	idB, err := NewID("e1f9b24e-c934-4831-b63a-5c884a544a8c")
	suite.Nil(err)
	suite.NotNil(idB)
	suite.Equal("e1f9b24e-c934-4831-b63a-5c884a544a8c", idB.value.String())

	idC, err := NewID("9c904f91-4f80-4e45-88f5-19701d379597")
	suite.Nil(err)
	suite.NotNil(idC)
	suite.Equal("9c904f91-4f80-4e45-88f5-19701d379597", idC.value.String())

	suite.True(idA.Equals(idB))
	suite.False(idA.Equals(idC))
}

func (suite *IdVoTestSuite) Test_deve_retornar_um_erro_se_o_id_estiver_invalido() {
	id, err := NewID("e1f9b24e-c934-4831-b63a-5c884a544")
	suite.Nil(id)
	suite.NotNil(err)
	suite.ErrorIs(ErrInvalidId, err)

}

func TestIdVoSuite(t *testing.T) {
	suite.Run(t, new(IdVoTestSuite))
}
