package middleware_test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Middleware/Authentication", func() {
	var (
		db gorm.DB
	)
	BeforeEach(func() {
		db, _ = gorm.Open("postgres", "user=postgres dbname=postgres")
	})

	Describe("public resources", func() {
		It("doesn't restrict public resources", func() {
			Expect(db).To(Equal(3))
		})
	})

	Describe("session validity", func() {
		It("lets autenticated sessions make requests", func() {

		})

		It("redirects authenticated requests when session missing", func() {

		})

		It("redirects authenticated requests when session expired", func() {

		})

		It("redirects authenticated requests when session is bogus", func() {

		})
	})

	Describe("user roles", func() {
		It("allows actions where users have the correct role", func() {})
		It("defaults to allowing all users", func() {})
		It("returns permission denied when user lacks role", func() {})
	})
})
