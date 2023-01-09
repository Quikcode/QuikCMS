import React from 'react'
import ReactDOM from 'react-dom/client'
import App from "@/app";
import {I18nextProvider} from "react-i18next";
import i18next from "i18next";
import language from "@/lib/language.js";
import getTypeLocale from "@/utils/getTypeLocale.js";
import locales from "@/locales/index.js";

import "normalize.css"
import "@/styles/global.scss"

if (!language.Get()) {
    language.Set(getTypeLocale(navigator.language))
}

i18next.init({
    interpolation: {escapeValue: false},
    lng: language.Get(),
    resources: locales()
})

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
      <I18nextProvider i18n={i18next}>
          <App/>
      </I18nextProvider>
  </React.StrictMode>,
)
