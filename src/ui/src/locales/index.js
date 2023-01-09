import en from "./en"
import es from "./es"

export const list = [
  {
    "name": "English",
    "label": "en",
    "module": en
  },
  {
    "name": "Español",
    "label": "es",
    "module": es
  }
]

export default () => {
  let newObject = {}
  list.map((locale) => {
    newObject[locale.label] = locale.module
  })
  return newObject
}