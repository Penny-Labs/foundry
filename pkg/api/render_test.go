package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RenderJson", func() {
	It("should write valid JSON and status code", func() {
		recorder := httptest.NewRecorder()
		data := map[string]string{"foo": "bar"}
		RenderJson(recorder, data, http.StatusCreated)

		Expect(recorder.Code).To(Equal(http.StatusCreated))
		Expect(recorder.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))

		var resp map[string]string
		Expect(json.Unmarshal(recorder.Body.Bytes(), &resp)).To(Succeed())
		Expect(resp).To(Equal(data))
	})

	It("should handle JSON marshal errors", func() {
		recorder := httptest.NewRecorder()
		ch := make(chan int) // channels can't be marshaled to JSON
		RenderJson(recorder, ch, http.StatusOK)

		Expect(recorder.Code).To(Equal(http.StatusInternalServerError))
		Expect(recorder.Body.String()).To(ContainSubstring("json: unsupported type: chan int"))
	})
})

var _ = Describe("JsonError", func() {
	It("should write valid JSON error message and status code", func() {
		recorder := httptest.NewRecorder()
		message := "An error occurred"
		JsonError(recorder, message, http.StatusBadRequest)

		Expect(recorder.Code).To(Equal(http.StatusBadRequest))
		Expect(recorder.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))

		var resp map[string]string
		Expect(json.Unmarshal(recorder.Body.Bytes(), &resp)).To(Succeed())
		Expect(resp).To(Equal(map[string]string{"error": message}))
	})
})
