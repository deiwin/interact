package interact_test

import (
	"github.com/deiwin/interact"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Confirm", func() {
	var (
		def     interact.ConfirmDefault
		message = "Are you sure?"
	)

	Context("with no default", func() {
		BeforeEach(func() {
			def = interact.ConfirmNoDefault
		})

		It("should ask with yes displayed as default", func() {
			actor.Confirm(message, def)
			Eventually(output).Should(gbytes.Say(`Are you sure\? \[y/n\]: `))
		})

		Context("with user answering yes", func() {
			BeforeEach(func() {
				userInput = "y\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})

		Context("with user answering no", func() {
			BeforeEach(func() {
				userInput = "n\n"
			})

			It("should return false", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeFalse())
			})
		})

		Context("with user answering nothing and then y", func() {
			BeforeEach(func() {
				userInput = "\ny\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})

		Context("with user answering gibberish and then y", func() {
			BeforeEach(func() {
				userInput = "asdfsadfa\ny\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})

			It("should ask twice with yes displayed as default", func() {
				actor.Confirm(message, def)
				Eventually(output).Should(gbytes.Say(`Are you sure\? \[y/n\]: `))
				Eventually(output).Should(gbytes.Say(`Please select y/n`))
				Eventually(output).Should(gbytes.Say(`Are you sure\? \[y/n\]: `))
			})
		})
	})

	Context("with no as default", func() {
		BeforeEach(func() {
			def = interact.ConfirmDefaultToNo
		})

		It("should ask with yes displayed as default", func() {
			actor.Confirm(message, def)
			Eventually(output).Should(gbytes.Say(`Are you sure\? \[y/N\]: `))
		})

		Context("with user answering yes", func() {
			BeforeEach(func() {
				userInput = "y\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})

		Context("with user answering no", func() {
			BeforeEach(func() {
				userInput = "n\n"
			})

			It("should return false", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeFalse())
			})
		})

		Context("with user answering nothing", func() {
			BeforeEach(func() {
				userInput = "\n"
			})

			It("should return false", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeFalse())
			})
		})

		Context("with user answering gibberish and then y", func() {
			BeforeEach(func() {
				userInput = "asdfasdf\ny\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})
	})

	Context("with yes as default", func() {
		BeforeEach(func() {
			def = interact.ConfirmDefaultToYes
		})

		It("should ask with yes displayed as default", func() {
			actor.Confirm(message, def)
			Eventually(output).Should(gbytes.Say(`Are you sure\? \[Y/n\]: `))
		})

		Context("with user answering yes", func() {
			BeforeEach(func() {
				userInput = "y\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})

		Context("with user answering no", func() {
			BeforeEach(func() {
				userInput = "n\n"
			})

			It("should return false", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeFalse())
			})
		})

		Context("with user answering nothing", func() {
			BeforeEach(func() {
				userInput = "\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})

		Context("with user answering gibberish and then y", func() {
			BeforeEach(func() {
				userInput = "sadfasdf\ny\n"
			})

			It("should return true", func() {
				confirmed, err := actor.Confirm(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmed).To(BeTrue())
			})
		})
	})
})
