import { isValidUrl } from "./index"

describe("isValidUrl", () => {
  const validUrls = [
    "google.com",
    "http://www.domain1-test-here.com",
    "http://domain1-test-here.com",
    "https://www.domain1-test-here.com",
    "https://domain1-test-here.comwww.domain1-test-here.com",
    "domain-here.com",
  ]
  validUrls.forEach((url) => {
    it(`should return true for ${url}`, () => {
      expect(isValidUrl(url)).toBe(true)
    })
  })
  const invalidUrls = ["", "asdasdasdasdasd", "canvas", "1231231231239"]
  invalidUrls.forEach((url) => {
    it(`should return true for ${url}`, () => {
      expect(isValidUrl(url)).toBe(false)
    })
  })
})
