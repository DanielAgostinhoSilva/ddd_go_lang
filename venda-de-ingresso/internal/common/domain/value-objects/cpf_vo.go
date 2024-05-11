package value_objects

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
)

var (
	ErrInvalidCpf = errors.New("Invalid Cpf")
)

type Cpf struct {
	value string
}

func (c *Cpf) Equals(cpf *Cpf) bool {
	return reflect.DeepEqual(c, cpf)
}

func (c *Cpf) Value() string {
	return c.value
}

func NewCpf(cpf string) (*Cpf, error) {
	re := regexp.MustCompile(`\D`)
	cpf = re.ReplaceAllString(cpf, "")

	if len(cpf) != 11 {
		return nil, ErrInvalidCpf
	}

	digitos := make([]int, 11)
	for i := 0; i < 11; i++ {
		digitos[i], _ = strconv.Atoi(string(cpf[i]))
	}

	soma := 0
	for i := 0; i < 9; i++ {
		soma += digitos[i] * (10 - i)
	}

	resto := soma % 11
	digitoVerificador1 := 11 - resto
	if digitoVerificador1 >= 10 {
		digitoVerificador1 = 0
	}

	if digitoVerificador1 != digitos[9] {
		return nil, ErrInvalidCpf
	}

	soma = 0
	for i := 0; i < 10; i++ {
		soma += digitos[i] * (11 - i)
	}

	resto = soma % 11
	digitoVerificador2 := 11 - resto
	if digitoVerificador2 >= 10 {
		digitoVerificador2 = 0
	}

	if digitoVerificador2 != digitos[10] {
		return nil, ErrInvalidCpf
	}

	return &Cpf{value: cpf}, nil
}
