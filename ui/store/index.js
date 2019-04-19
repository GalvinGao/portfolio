export const state = () => ({
  locales: ['en', 'zh'],
  locale: 'zh',
  extendedTitle: null,
  contentList: []
})

export const mutations = {
  setLang(state, locale) {
    if (state.locales.indexOf(locale) !== -1) {
      state.locale = locale
    }
  },
  setExtendedTitle(state, name) {
    state.extendedTitle = name
  },
  setContentList(state, list) {
    state.contentList = list
  }
}
