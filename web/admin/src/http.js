import axios from 'axios'
import Vue from 'vue'
import router from '@/router/index'

const Url = 'http://103.118.82.117:8080/api/v1'

// axios 配置
axios.defaults.timeout = 20000
axios.defaults.baseURL = Url
axios.defaults.withCredentials = true
axios.defaults.headers['Content-Type'] = 'application/x-www-form-urlencoded'
// http request 拦截器
axios.interceptors.request.use(
  config => {
    if (sessionStorage.getItem('token')) {
    // config.headers.Authorization = `token ${store.state.token}`;
    // }
      config.headers.Authorization = `Bearer ${window.sessionStorage.getItem('token')}`
    }
    return config
  },
  err => {
    return Promise.reject(err)
  })
// http response 拦截器
axios.interceptors.response.use(
  response => {
    //   不跳转到登录
    if (response.data.status === 1008 || response.data.code === 1004) {
      router.replace({
        path: '/user/login',
        query: { redirect: router.currentRoute.fullPath }
      })
    }
    return response
  },
  error => {
    // console.log(JSON.stringify(error)) // console : Error: Request failed with status code 402
    return Promise.reject(error.response.data)
  })
Vue.prototype.$http = axios
export { Url }
