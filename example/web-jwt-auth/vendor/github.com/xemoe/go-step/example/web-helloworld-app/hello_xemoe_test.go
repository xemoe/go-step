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

			Expect(page.Navigate("http://127.0.0.1:3000/hi")).To(Succeed())
			Eventually(page).Should(HaveURL("http://127.0.0.1:3000/hi"))

			//
			// Expected returned json string
			//
			expected := `{"message":"Hello world (fallback)"}`
			pre := page.Find("pre")
			Eventually(pre).Should(BeVisible())
			Eventually(pre).Should(HaveText(expected))

		})
	})
})
