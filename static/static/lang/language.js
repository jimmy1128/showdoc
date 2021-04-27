class Language {
  currentLang = 'ZH_CNS'
  // currentLangPackage = zhcn;
  langKeys = ['ZH_CN', 'EN_US']

  init () {
    if (
      localStorage.getItem('locale') === null ||
      !localStorage.getItem('locale')
    ) {
      this.currentLang = 'ZH_CN'
    } else {
      this.currentLang = localStorage.getItem('locale')
    }
  }

  // 特权方法
  getLang () {
    return this.currentLang
  }

  setLang (lang) {
    if (!this.langKeys.includes(lang)) {
      throw new Error('PLEASE SET THE RIGHT KEY')
    }
    localStorage.setItem('locale', lang)
  }

  clearLang (lang) {
    if (!this.langKeys.includes(lang)) {
      throw new Error('PLEASE SET THE RIGHT KEY')
    }
    localStorage.removeItem('locale')
  }
}
const LANG = new Language()
export default LANG
