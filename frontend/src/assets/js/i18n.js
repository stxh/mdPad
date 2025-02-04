// i18n.js
import i18next from 'i18next';
import { createI18nStore } from 'svelte-i18next';
import LanguageDetector from 'i18next-browser-languagedetector';

// 导入你的翻译资源
import translationEN from './locales/en.json';
import translationZH from './locales/zh.json';

i18next.use(LanguageDetector) // 使用 LanguageDetector
  .init({
  //lng: 'zh', // 默认语言
  fallbackLng: 'en', // 回退语言
  debug: false, // 开启调试模式 (可选)
  resources: {
    en: {
      translation: translationEN,
    },
    zh: {
      translation: translationZH,
    },
  },
  interpolation: {
    escapeValue: false, // Svelte 默认会转义，所以这里设置为 false
  },
});

const i18n = createI18nStore(i18next);

export default i18n;