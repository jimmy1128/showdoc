import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import UserLogin from '@/components/user/Login'
import UserSetting from '@/components/user/Setting'
import UserRegister from '@/components/user/Register'
import UserResetPassword from '@/components/user/ResetPassword'
import ResetPasswordByUrl from '@/components/user/ResetPasswordByUrl'
import Notice from '@/components/notice/Index'
import ItemIndex from '@/components/item/Index'
import ItemAdd from '@/components/item/add/Index'
import Catalog from '@/components/catalog/Index'
import ItemPassword from '@/components/item/Password'
import PageIndex from '@/components/page/Index'
import ItemShow from '@/components/item/show/Index'
import PageEdit from '@/components/page/edit/Index'
import Team from '@/components/team/Index'
import TeamMember from '@/components/team/Member'
import TeamItem from '@/components/team/Item'
import ItemSetting from '@/components/item/setting/Index'
import ItemExport from '@/components/item/export/Index'
import Attachment from '@/components/attachment/Index'
import Admin from '@/components/admin/Index'
import PageDiff from '@/components/page/Diff'

Vue.use(Router)

const routes = [
  {
    path: '/',
    name: 'Index',
    component: Index
  },
  {
    path: '/admin/index',
    name: 'Admin',
    component: Admin
  },
  {
    path: '/user/login',
    name: 'UserLogin',
    component: UserLogin
  },
  {
    path: '/user/setting',
    name: 'UserSetting',
    component: UserSetting
  },
  {
    path: '/team/index',
    name: 'Team',
    component: Team
  },
  {
    path: '/user/register',
    name: 'UserRegister',
    component: UserRegister
  },
  {
    path: '/user/resetPassword',
    name: 'UserResetPassword',
    component: UserResetPassword
  },
  {
    path: '/user/ResetPasswordByUrl',
    name: 'ResetPasswordByUrl',
    component: ResetPasswordByUrl
  },
  {
    path: '/notice/index',
    name: 'Notice',
    component: Notice
  },
  {
    path: '/item/index',
    name: 'ItemIndex',
    component: ItemIndex
  },
  {
    path: '/item/add',
    name: 'ItemAdd',
    component: ItemAdd
  },
  {
    path: '/catalog/:item_id',
    name: 'Catalog',
    component: Catalog
  },
  {
    path: '/item/password/:item_id',
    name: 'ItemPassword',
    component: ItemPassword
  },
  {
    path: '/page/:page_id',
    name: 'PageIndex',
    component: PageIndex
  },
  {
    path: '/:item_id',
    name: 'ItemShow',
    component: ItemShow
  },
  {
    path: '/page/edit/:item_id/:page_id',
    name: 'PageEdit',
    component: PageEdit
  },
  {
    path: '/team/member/:team_id',
    name: 'TeamMember',
    component: TeamMember
  },
  {
    path: '/team/item/:team_id',
    name: 'TeamItem',
    component: TeamItem
  },
  {
    path: '/item/setting/:item_id',
    name: 'ItemSetting',
    component: ItemSetting
  },
  {
    path: '/item/export/:item_id',
    name: 'ItemExport',
    component: ItemExport
  },
  {
    path: '/attachment/index',
    name: 'Attachment',
    component: Attachment
  },
  {
    path: '/page/diff/:page_id/:page_history_id',
    name: 'PageDiff',
    component: PageDiff
  }
]
const router = new Router({
  routes
})

// router.beforeEach((to, from, next) => {
//   if (to.meta.title) {
//     document.title = to.meta.title
//   }
//   next()

//   const userToken = window.sessionStorage.getItem('token')
//   if (to.path === '/admin') {
//     if (!userToken) {
//       next('/login')
//     } else {
//       next()
//     }
//     next()
//   }
// })
export default router
