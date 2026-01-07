package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployeeCode(t *testing.T) {
	g := NewGomegaWithT(t)
	e := Employees{
		Name:         "Negative Salary",
		Salary:       17000,
		EmployeeCode: "A-1234",
	}
	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by '-' and 4 digits (0-9)"))
}
