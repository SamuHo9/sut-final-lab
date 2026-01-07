package unit

import (
	"testing"

	"github.com/SamuHo9/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)
func Testnegative(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("positive", func(t *testing.T) {
		employees := entity.Employees{
			Name:          "supa dupaluv",
			Salary:        200,
			EmployeesCode: "HR-1877",
		}
		ok, err := govalidator.ValidateStruct(employees)

		g.Expect(ok).To(BeFalse())
		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 15000 and 200000"))
	})
}
