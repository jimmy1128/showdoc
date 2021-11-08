<template>
  <div class="hello">
    <Header></Header>
    <!-- 展示常规项目 -->
    <ShowRegularItem
      :item_info="item_info"
      :search_item="search_item"
      :keyword="keyword"
      v-on:itemlangId2="item_langId2"
      v-if="item_info && (item_info.itemtype == 1 || item_info.itemtype == 3 || item_info.itemtype === '0' || item_info.itemtype === 0 )"
    ></ShowRegularItem>

    <!-- 展示单页项目 -->
    <ShowSinglePageItem :item_info="item_info" v-if="item_info && item_info.itemtype == 2 "></ShowSinglePageItem>

    <!-- 展示表格项目 -->
    <ShowTableItem :item_info="item_info" v-if="item_info && item_info.itemtype == 4 "></ShowTableItem>

    <Footer></Footer>
  </div>
</template>

<script>
import ShowRegularItem from './show_regular_item/Index'
import ShowSinglePageItem from './show_single_page_item/Index'
import ShowTableItem from './show_table_item/Index'
export default {
  data () {
    return {
      item_info: '',
      keyword: '',
      itemlangId2: '',
      itemlangId3: '',
      access : ''
      // loading: false
    }
  },
  components: {
    ShowRegularItem,
    ShowSinglePageItem,
    ShowTableItem
  },
  methods: {
    item_langId2 (value) {
      this.itemlangId2 = value
      if (value !== undefined) {
        this.item_info = ''
        this.get_item_menu(this.keyword)
      }
    },
    // 获取菜单
    async get_item_menu (keyword) {
      if (!keyword) {
        keyword = ''
      }
      // var loading = this.$loading()
      var item_id = this.$route.params.item_id ? this.$route.params.item_id : 0
      var page_id = this.$route.query.page_id ? this.$route.query.page_id : 0
      var url = DocConfig.server + '/item/info'
      var formdata = new FormData()
      formdata.append('id', item_id)
      formdata.append('keyword', keyword)
      formdata.append('lang_id', this.itemlangId2)
      if (!this.keyword) {
        formdata.append('defaultpageid', page_id)
      }
      const { data: res } = await this.$http.post(url, formdata)
      if (res.status === 200) {
        // loading.close()
        var json = res.data
        if (json.defaultpageid <= 0 || json.defaultpageid === undefined) {
          if (json.menu.pages[0]) {
            json.defaultpageid = String(json.menu.pages[0].ID)
          }
        }
        this.item_info = json
        this.itemtitle = json.itemtitle
        this.access = res.access
        if (json.unread_count > 0) {
          this.$message({
            showClose: true,
            duration: 10000,
            dangerouslyUseHTMLString: true,
            message: '<a href="#/notice/index">你有新的未读消息，点击查看</a>'
          })
        }
        document.title =this.itemtitle +'--Doc'
      } else if (res.status === 10307) {
        // 需要输入密码
        this.$router.replace({
          path: '/item/password/' + item_id,
          query: { page_id: page_id, redirect: this.$router.currentRoute.fullPath }
        })
      } else {
        this.$alert(res.message)
      }
      // 设置一个最长关闭时间
      setTimeout(() => {
        // loading.close()
      }, 10000)
    },
    search_item (keyword) {
      this.item_info = ''
      this.$store.dispatch('changeItemInfo', '')
      this.keyword = keyword
      this.get_item_menu(keyword)
    },
    loadConfig () {
      var item_id = this.$route.params.item_id ? this.$route.params.item_id : 0
      var that = this
      var url = DocConfig.server + '/adminSetting/loadLangConfig'
      var params = new URLSearchParams()
      params.append('id', item_id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.length === 0) {
            return
          }
          that.itemlangId2 = 0
        } else {
          this.$alert(response.data.message)
        }
      })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  mounted () {
    // defaultlang 是全局语言
    this.itemlangId2 = localStorage.getItem("defaultlang")
    this.loadConfig()
    this.get_item_menu(this.keyword)
    this.$store.dispatch('changeOpenCatId', 0)
    this.$cookies.remove('selected')
    localStorage.removeItem('defaultlang')
  },
  beforeDestroy () {
    this.$message.closeAll()
    document.title ='Doc'
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
