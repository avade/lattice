package s3_blob_store_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/lattice/ltc/blob_store/s3_blob_store"
	config_package "github.com/cloudfoundry-incubator/lattice/ltc/config"
	"github.com/cloudfoundry-incubator/lattice/ltc/config/persister"
)

var _ = Describe("S3BlobStore", func() {
	var (
		verifier s3_blob_store.Verifier
	)

	BeforeEach(func() {
		verifier = s3_blob_store.Verifier{}

		config := config_package.New(persister.NewMemPersister())
		config.SetS3BlobStore("", "", "", "")
	})

	Describe("Verify", func() {
		It("verifies a blob store with valid credentials", func() {
			//XXX

			authorized, err := verifier.Verify(config)
			Expect(err).NotTo(HaveOccurred())
			Expect(authorized).To(BeTrue())

			Expect(fakeServer.ReceivedRequests()).To(HaveLen(1))
		})

		Context("when the blob store credentials are incorrect", func() {
			It("rejects a blob store with invalid credentials", func() {
				//XXX

				authorized, err := verifier.Verify(config)
				Expect(err).NotTo(HaveOccurred())
				Expect(authorized).To(BeFalse())

				Expect(fakeServer.ReceivedRequests()).To(HaveLen(1))
			})
		})

		Context("when the blob store is inaccessible", func() {
			It("returns an error", func() {
				//XXX

				_, err := verifier.Verify(config)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
