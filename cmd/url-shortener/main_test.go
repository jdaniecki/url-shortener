package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Main Suite")
}

var serverURL string

var _ = BeforeSuite(func() {
	//slog.SetLogLoggerLevel(slog.LevelDebug)

	host := "localhost:8081"
	serverURL = "http://" + host

	// Start the server in a goroutine
	go func() {
		err := startServer(host)
		Expect(err).NotTo(HaveOccurred())
	}()

	// Give the server a moment to start
	Eventually(func() bool {
		resp, err := http.Get(serverURL)
		if err != nil {
			return false
		}
		resp.Body.Close()
		return true
	}, "5s", "100ms").Should(BeTrue())
})

var _ = AfterSuite(func() {
	// Implement server shutdown logic if needed
})

var _ = Describe("Main", func() {

	Context("when sending a request to the server", func() {

		It("should return 404 for the root path", func() {
			resp, err := http.Get(serverURL)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

		It("should return 404 for a non-existent short URL", func() {
			resp, err := http.Get(serverURL + "/non-existent")
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
		})

		It("should return 200 for valid short URL", func() {
			// create a short URL for example.com domain
			body := `{"url":"http://example.com"}`
			resp, err := http.Post(serverURL+"/shorten", "application/json", strings.NewReader(body))
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()

			// validate response
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			var result map[string]string
			err = json.NewDecoder(resp.Body).Decode(&result)
			Expect(err).NotTo(HaveOccurred())
			shortURL, ok := result["shortUrl"]
			Expect(ok).To(BeTrue())
			Expect(shortURL).NotTo(BeEmpty())

			// retrive the short URL
			resp, err = http.Get(serverURL + "/" + shortURL)
			Expect(err).NotTo(HaveOccurred())
			defer resp.Body.Close()
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
		})

	})
})
