import i18n from 'i18next';
import { ServiceBuilder } from '../service.d';
import en from './en';
import ua from './ua';

export enum Languages {
    English = 'en',
    Ukrainian = 'ua',
}

export class Localization extends ServiceBuilder {

  constructor() {
    super();
  }
  public static get currentLanguage(): Languages {
    return i18n.language as Languages || Languages.Ukrainian;
  }

  public static get resources() {
    return {
      ua: {
        translation: ua,
      },
      en: {
        translation: en,
      },
    };
  };

  public async init() {
    console.log('STUB!');
  }
}

export const TRANSLATIONS =
    Localization.resources[Localization.currentLanguage].translation;
