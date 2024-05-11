package value_objects

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type NameVoTestSuite struct {
	suite.Suite
}

func (suite *NameVoTestSuite) Test_deve_iniciar_um_nome_valido() {
	nameA, err := NewName("Fulano da Silva")
	suite.Nil(err)
	suite.Equal("Fulano da Silva", nameA.Value())
}

func (suite *NameVoTestSuite) Test_deve_comparar_dois_nomes_se_sao_iguais() {
	nameA, err := NewName("Fulano da Silva")
	suite.Nil(err)
	suite.Equal("Fulano da Silva", nameA.Value())

	nameB, err := NewName("Fulano da Silva")
	suite.Nil(err)
	suite.Equal("Fulano da Silva", nameB.Value())

	nameC, err := NewName("Fulana da Silva")
	suite.Nil(err)
	suite.Equal("Fulana da Silva", nameC.Value())

	suite.True(nameA.Equals(nameB))
	suite.False(nameA.Equals(nameC))
}

func (suite *NameVoTestSuite) Test_deve_retornar_um_error_se_o_tamanho_do_nome_for_menor_que_3() {
	nameA, err := NewName("Fu")
	suite.Nil(nameA)
	suite.ErrorIs(ErrNameIsRequired, err)
}

func TestNameVo(t *testing.T) {
	suite.Run(t, new(NameVoTestSuite))
}
