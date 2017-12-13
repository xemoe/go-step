package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/sclevine/agouti/matchers"
)

var _ = Describe("HelloXemoe Page", func() {

	It("should see json response", func() {

		page := getPage()
		defer page.Destroy()

		By("go to page", func() {

			Expect(page.Navigate("http://localhost:3000/hi")).To(Succeed())
			Expect(page).To(HaveURL("http://localhost:3000/hi"))

		})
	})
})
