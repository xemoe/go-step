package main_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"

	"github.com/onsi/gomega/gexec"
	"os"
	"os/exec"
)

func TestWebHelloworldApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WebHelloworldApp Suite")
}

var agoutiDriver *agouti.WebDriver
var websiteSession *gexec.Session

var _ = BeforeSuite(func() {

	// agoutiDriver = agouti.PhantomJS()
	// agoutiDriver = agouti.Selenium()
	agoutiDriver = agouti.ChromeDriver()

	Expect(agoutiDriver.Start()).To(Succeed())

	//
	// start server.go
	//
	startWebsite()
})

var _ = AfterSuite(func() {
	Expect(agoutiDriver.Stop()).To(Succeed())
	websiteSession.Kill()
})

//
//
//

func getPage() *agouti.Page {
	var page *agouti.Page
	var err error
	if os.Getenv("TEST_ENV") == "CI" {
		page, err = agouti.NewPage(agoutiDriver.URL(),
			agouti.Desired(agouti.Capabilities{
				"chromeOptions": map[string][]string{
					"args": []string{
						// There is no GPU on our Ubuntu box!
						"disable-gpu",

						// Sandbox requires namespace permissions that we don't have on a container
						"no-sandbox",
					},
				},
			}),
		)
	} else {
		// Local machines can start chrome normally
		page, err = agoutiDriver.NewPage(agouti.Browser("chrome"))
	}
	Expect(err).NotTo(HaveOccurred())

	return page
}

func startWebsite() {
	command := exec.Command("go", "run", "server.go")
	Eventually(func() error {
		var err error
		websiteSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
		return err
	}).Should(Succeed())
}
