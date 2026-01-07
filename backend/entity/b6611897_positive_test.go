package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestPositive(t *testing.T) {
	g := NewGomegaWithT(t)
	e := Employees{
		Name:         "Positive Name",
		Salary:       18000,
		EmployeeCode: "HR-1024",
	}
	ok, err := govalidator.ValidateStruct(e)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}
