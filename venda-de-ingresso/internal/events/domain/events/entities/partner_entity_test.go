package entities

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PartnerEntityTestsSuite struct {
	suite.Suite
}

func (suite *PartnerEntityTestsSuite) Test_deve_criar_um_partner_valido() {
	command := &CreatePartnerCommand{
		Name: "Fulano da silva",
	}
	partner, err := CreatePartner(command)
	suite.Nil(err)
	suite.NotNil(partner)
	suite.NotNil(partner.Id())
	suite.Equal(command.Name, partner.Name())
}

func (suite *PartnerEntityTestsSuite) Test_deve_inicializar_um_parter() {

	partner, err := NewPartner(&PartnerProps{
		id:   "5e1ab23c-af07-4a16-8755-4c7a5affdc03",
		name: "Fulana da Silva",
	})
	suite.Nil(err)
	suite.NotNil(partner)
	suite.NotNil(partner.Id())
	suite.Equal("Fulana da Silva", partner.Name())
}

func (suite *PartnerEntityTestsSuite) Test_deve_comparar_os_partner_se_sao_iguais_ou_nao() {

	partnerA, err := NewPartner(&PartnerProps{
		id:   "5e1ab23c-af07-4a16-8755-4c7a5affdc03",
		name: "Fulana da Silva",
	})
	suite.Nil(err)
	suite.NotNil(partnerA)

	partnerB, err := NewPartner(&PartnerProps{
		id:   "5e1ab23c-af07-4a16-8755-4c7a5affdc03",
		name: "Teste A",
	})
	suite.Nil(err)
	suite.NotNil(partnerB)

	partnerC, err := NewPartner(&PartnerProps{
		id:   "f547086c-b5e4-41cd-914d-3cf5e949166c",
		name: "Teste C",
	})
	suite.Nil(err)
	suite.NotNil(partnerC)

	suite.True(partnerA.Equals(partnerB))
	suite.False(partnerA.Equals(partnerC))
}

func TestPartnerEntitySuite(t *testing.T) {
	suite.Run(t, new(PartnerEntityTestsSuite))
}
