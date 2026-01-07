package unit

import (
	"testing"

	"github.com/SamuHo9/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)
func Testnegativeemployees(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("positive", func(t *testing.T) {
		employees := entity.Employees{
			Name:          "supa dupaluv",
			Salary:        20000,
			EmployeesCode: "",
		}
		ok, err := govalidator.ValidateStruct(employees)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeCode must be 2 uppercase English letters (A-Z) followed by - and 4 digits (0-9)"))
	})
}
