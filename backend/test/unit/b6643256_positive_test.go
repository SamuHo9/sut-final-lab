package unit

import (
	"testing"

	"github.com/SamuHo9/sut-final-lab/entity"
	"github.com/asaskevich/govalidator"

	. "github.com/onsi/gomega"
)
func Testpositive(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run("positive", func(t *testing.T) {
		employees := entity.Employees{
			Name:          "dddd",
			Salary:        20000,
			EmployeesCode: "HR-1024",
		}
		ok, err := govalidator.ValidateStruct(employees)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
