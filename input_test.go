package interact_test

import (
	"errors"

	"github.com/deiwin/interact"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Input", func() {
	var message = "Please answer"
	Describe("PromptAndRetry", func() {
		Context("without any checks", func() {
			BeforeEach(func() {
				userInput = " user-input \n"
			})

			It("should return the trimmed input (just as without retry)", func() {
				input, err := actor.PromptAndRetry(message)
				Expect(err).NotTo(HaveOccurred())
				Expect(input).To(Equal("user-input"))
				Eventually(output).Should(gbytes.Say(`Please answer: `))
			})
		})

		Context("with a check that fails the first time", func() {
			var (
				checkErr = errors.New("The first time fails!")
				check    interact.InputCheck
				i        int
			)
			BeforeEach(func() {
				i = 0
				check = func(input string) error {
					if i == 0 {
						Expect(input).To(Equal("user-input1"))
						i++
						return checkErr
					}
					Expect(input).To(Equal("correct-input"))
					return nil
				}
			})

			Context("with user retrying", func() {
				BeforeEach(func() {
					userInput = "user-input1\ny\ncorrect-input\n"
				})

				It("should return the second (correct) input", func() {
					input, err := actor.PromptAndRetry(message, check)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal("correct-input"))
				})

				It("should have correct prompts", func() {
					actor.PromptAndRetry(message, check)
					Eventually(output).Should(gbytes.Say(`Please answer: `))
					Eventually(output).Should(gbytes.Say(`The first time fails!`))
					Eventually(output).Should(gbytes.Say(`Do you want to try again\? \[y/N\]: `))
					Eventually(output).Should(gbytes.Say(`Please answer: `))
				})
			})

			Context("with user retrying, but answering gibberish on first retry", func() {
				BeforeEach(func() {
					userInput = "user-input1\nasdfsadf\ny\ncorrect-input\n"
				})

				It("should return the second (correct) input", func() {
					input, err := actor.PromptAndRetry(message, check)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal("correct-input"))
				})

				It("should have correct prompts", func() {
					actor.PromptAndRetry(message, check)
					Eventually(output).Should(gbytes.Say(`Please answer: `))
					Eventually(output).Should(gbytes.Say(`The first time fails!`))
					Eventually(output).Should(gbytes.Say(`Do you want to try again\? \[y/N\]: `))
					Eventually(output).Should(gbytes.Say(`Please select y/n!`))
					Eventually(output).Should(gbytes.Say(`Do you want to try again\? \[y/N\]: `))
					Eventually(output).Should(gbytes.Say(`Please answer: `))
				})
			})
		})
	})

	Describe("Prompt", func() {
		Context("with user input", func() {
			BeforeEach(func() {
				userInput = " user-input \n"
			})

			It("should return the trimmed input", func() {
				input, err := actor.Prompt(message)
				Expect(err).NotTo(HaveOccurred())
				Expect(input).To(Equal("user-input"))
				Eventually(output).Should(gbytes.Say(`Please answer: `))
			})

			Context("with a check", func() {
				var (
					checkErr error
					check    interact.InputCheck
				)

				JustBeforeEach(func() {
					check = func(input string) error {
						Expect(input).To(Equal("user-input"))
						return checkErr
					}
				})
				Context("with a failing check", func() {
					BeforeEach(func() {
						checkErr = errors.New("Check failed!")
					})

					It("should return the error from the check", func() {
						_, err := actor.Prompt(message, check)
						Expect(err).To(Equal(checkErr))
					})

					Context("with another check after the failed one", func() {
						It("should not call the second check", func() {
							actor.Prompt(message, check, func(input string) error {
								Fail("should not be called")
								return nil
							})
						})
					})
				})

				Context("with a passing check", func() {
					BeforeEach(func() {
						checkErr = nil
					})

					It("should not return an error", func() {
						_, err := actor.Prompt(message, check)
						Expect(err).NotTo(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("PromptOptional", func() {
		var def = "default value"
		Context("without any checks", func() {
			Context("with a simple input", func() {
				BeforeEach(func() {
					userInput = " user-input \n"
				})

				It("should return the trimmed input", func() {
					input, err := actor.PromptOptional(message, def)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal("user-input"))
					Eventually(output).Should(gbytes.Say(`Please answer: `))
				})
			})

			Context("with no input (just enter)", func() {
				BeforeEach(func() {
					userInput = "\n"
				})

				It("should return the default value", func() {
					input, err := actor.PromptOptional(message, def)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal(def))
					Eventually(output).Should(gbytes.Say(`Please answer: \(default value\) `))
				})
			})

			Context("with just whitespace as input", func() {
				BeforeEach(func() {
					userInput = " 	\n"
				})

				It("should return the default value", func() {
					input, err := actor.PromptOptional(message, def)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal(def))
					Eventually(output).Should(gbytes.Say(`Please answer: \(default value\) `))
				})
			})
		})

		Context("with a check", func() {
			var (
				checkErr error
				check    interact.InputCheck
			)
			JustBeforeEach(func() {
				check = func(input string) error {
					Expect(input).To(Equal("user-input"))
					return checkErr
				}
			})

			Context("with no input (just enter), but a failing check", func() {
				BeforeEach(func() {
					userInput = "\n"
					checkErr = errors.New("Check failed!")
				})

				It("should return the default value (not run the check)", func() {
					input, err := actor.PromptOptional(message, def, check)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal(def))
					Eventually(output).Should(gbytes.Say(`Please answer: \(default value\) `))
				})
			})

			Context("with a simple input", func() {

				BeforeEach(func() {
					userInput = "user-input \n"
				})

				Context("with a failing check", func() {
					BeforeEach(func() {
						checkErr = errors.New("Check failed!")
					})

					It("should return the error from the check", func() {
						_, err := actor.PromptOptional(message, def, check)
						Expect(err).To(Equal(checkErr))
					})

					Context("with another check after the failed one", func() {
						It("should not call the second check", func() {
							actor.PromptOptional(message, def, check, func(input string) error {
								Fail("should not be called")
								return nil
							})
						})
					})
				})

				Context("with a passing check", func() {
					BeforeEach(func() {
						checkErr = nil
					})

					It("should not return an error", func() {
						_, err := actor.PromptOptional(message, def, check)
						Expect(err).NotTo(HaveOccurred())
					})
				})
			})
		})
	})

	Describe("PromptOptionalAndRetry", func() {
		var def = "default value"
		Context("without any checks", func() {
			BeforeEach(func() {
				userInput = " user-input \n"
			})

			It("should return the trimmed input (just as without retry)", func() {
				input, err := actor.PromptOptionalAndRetry(message, def)
				Expect(err).NotTo(HaveOccurred())
				Expect(input).To(Equal("user-input"))
				Eventually(output).Should(gbytes.Say(`Please answer: `))
			})
		})

		Context("with a check that fails the first time", func() {
			var (
				checkErr = errors.New("The first time fails!")
				check    interact.InputCheck
				i        int
			)
			BeforeEach(func() {
				i = 0
				check = func(input string) error {
					if i == 0 {
						Expect(input).To(Equal("user-input1"))
						i++
						return checkErr
					}
					Expect(input).To(Equal("correct-input"))
					return nil
				}
			})

			Context("with user retrying", func() {
				BeforeEach(func() {
					userInput = "user-input1\ny\ncorrect-input\n"
				})

				It("should return the second (correct) input", func() {
					input, err := actor.PromptOptionalAndRetry(message, def, check)
					Expect(err).NotTo(HaveOccurred())
					Expect(input).To(Equal("correct-input"))
				})

				It("should have correct prompts", func() {
					actor.PromptOptionalAndRetry(message, def, check)
					Eventually(output).Should(gbytes.Say(`Please answer: \(default value\) `))
					Eventually(output).Should(gbytes.Say(`The first time fails!`))
					Eventually(output).Should(gbytes.Say(`Do you want to try again\? \[y/N\]: `))
					Eventually(output).Should(gbytes.Say(`Please answer: \(default value\) `))
				})
			})
		})
	})
})
