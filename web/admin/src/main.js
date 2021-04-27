import Vue from 'vue'
import App from './App.vue'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import Header from '@/components/common/Header'
import Footer from '@/components/common/Footer'
import util from '@/util.js'
import VueI18n from 'vue-i18n'
import VueDND from 'awe-dnd'
// import enLocale from 'element-ui/lib/locale/lang/en'
// import zhLocale from 'element-ui/lib/locale/lang/zh-CN'
import 'url-search-params-polyfill'
import 'babel-polyfill'
import VueClipboard from 'vue-clipboard2'
// import { Url } from './http'
import { EN_US } from '../static/lang/en'
import { ZH_CN } from '../static/lang/zh-CN'
import store from './store/'
import VueCookies from 'vue-cookies'
// import LANG from '../static/lang/language'

Vue.use(util)
Vue.config.productionTip = false
Vue.component('Header', Header)
Vue.component('Footer', Footer)
Vue.use(ElementUI)
Vue.use(VueI18n)
Vue.use(VueClipboard)
Vue.use(VueCookies)
Vue.use(VueDND)
// Vue.config.lang = this.DocConfig.lang

const i18n = new VueI18n({
  fallbackLocale: 'ZH_CN',
  messages: {
    ZH_CN,
    EN_US
  }
})
export default i18n
// 多语言相关
// 将axios挂载到prototype上，在组件中可以直接使用this.axios访问

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  i18n,
  store,
  template: '<App/>',
  render: h => h(App),
  components: { App }
}).$mount('#app')
