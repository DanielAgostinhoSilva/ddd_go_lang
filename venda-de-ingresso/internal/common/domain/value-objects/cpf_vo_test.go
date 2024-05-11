package value_objects

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CpfVoSuiteTest struct {
	suite.Suite
}

func (suite *CpfVoSuiteTest) Test_deve_criar_um_cpf_valido() {
	cpf, err := NewCpf("488.440.270-70")
	suite.Nil(err)
	suite.Equal("48844027070", cpf.value)
	cpf, err = NewCpf("48844027070")
	suite.Nil(err)
	suite.Equal("48844027070", cpf.value)
}

func (suite *CpfVoSuiteTest) Test_deve_comparar_os_cpf_se_sao_iguais_ou_nao() {
	cpfA, err := NewCpf("488.440.270-70")
	suite.Nil(err)
	suite.Equal("48844027070", cpfA.value)

	cpfB, err := NewCpf("488.440.270-70")
	suite.Nil(err)
	suite.Equal("48844027070", cpfA.value)

	cpfC, err := NewCpf("767.796.550-40")
	suite.Nil(err)
	suite.Equal("48844027070", cpfA.value)

	suite.True(cpfA.Equals(cpfB))
	suite.False(cpfA.Equals(cpfC))
}

func (suite *CpfVoSuiteTest) Test_deve_lancar_um_erro_quando_o_cpf_for_invalido() {
	cpf, err := NewCpf("")
	suite.Nil(cpf)
	suite.ErrorIs(ErrInvalidCpf, err)

	cpf, err = NewCpf("488.440.270")
	suite.Nil(cpf)
	suite.ErrorIs(ErrInvalidCpf, err)

	cpf, err = NewCpf("488.440.270-71")
	suite.Nil(cpf)
	suite.ErrorIs(ErrInvalidCpf, err)

	cpf, err = NewCpf("488.440.270-10")
	suite.Nil(cpf)
	suite.ErrorIs(ErrInvalidCpf, err)
}

func TestCpfVoSuite(t *testing.T) {
	suite.Run(t, new(CpfVoSuiteTest))
}
