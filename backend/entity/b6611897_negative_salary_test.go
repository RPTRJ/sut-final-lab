package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestSalary(t *testing.T) {
	g := NewGomegaWithT(t)
	e := Employees{
		Name:         "Negative Salary",
		Salary:       13000,
		EmployeeCode: "HR-2048",
	}
	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(BeFalse())
	g.Expect(err).To(Not(BeNil()))
	g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
}
