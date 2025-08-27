package hash

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hash Functions", func() {
	Describe("Sha512Hash", func() {
		It("should return a valid SHA-512 hash", func() {
			key := "testkey"
			hash := Sha512Hash(key)
			Expect(hash).To(HaveLen(128))
		})
	})

	Describe("Sha256Hash", func() {
		It("should return a valid SHA-256 hash", func() {
			key := "testkey"
			hash := Sha256Hash(key)
			Expect(hash).To(HaveLen(64))
		})
	})

	Describe("BcryptHashKey and BcryptCompareHashAndKey", func() {
		It("should hash and compare a key successfully", func() {
			key := "testkey"
			hashed, err := BcryptHashKey(key, BcryptCost)
			Expect(err).To(BeNil())
			Expect(hashed).NotTo(BeEmpty())

			err = BcryptCompareHashAndKey(hashed, key)
			Expect(err).To(BeNil())
		})

		It("should fail to compare with wrong key", func() {
			key := "testkey"
			wrongKey := "wrongkey"
			hashed, err := BcryptHashKey(key, BcryptCost)
			Expect(err).To(BeNil())

			err = BcryptCompareHashAndKey(hashed, wrongKey)
			Expect(err).NotTo(BeNil())
		})
	})
})
