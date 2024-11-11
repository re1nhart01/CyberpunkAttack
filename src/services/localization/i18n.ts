import i18n from 'i18next';
import { initReactI18next } from 'react-i18next';
import { Languages, Localization } from './Localization';

i18n
  .use(initReactI18next)
  .init({
    resources: Localization.resources,
    lng: Languages.English,
    fallbackLng: Languages.Ukrainian,
    interpolation: {
      escapeValue: false,
    },
  });

export default i18n;
