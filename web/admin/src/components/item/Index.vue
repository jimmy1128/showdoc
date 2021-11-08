<template>
  <div class="hello" >
    <Header> </Header>

    <el-container class="container-narrow" >

      <el-row class="masthead" >
          <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
          <div class="logo-title" >
              <h2 class="muted">Doc</h2>
          </div>
          <div class="header-btn-group pull-right" style="display: flex; align-items: center;">
          <el-tooltip class="item" effect="dark" :content="$t('language')" placement="top">
            <router-link to>
              <i class="material-icons md-20" style="display: flex; align-items: baseline;" :key="lang?'ZH_CN':'EN_US'" @click="changeLang()" >language</i>
            </router-link>
          </el-tooltip>
          <el-tooltip
            class="item"
            effect="dark"
            content="接口开发调试工具"
            placement="top"
          >
            <a target="_blank" href="">
              <i class="el-icon-connection"></i>
            </a>
          </el-tooltip>

          <el-tooltip class="item" effect="dark" :content="$t('team_mamage')" placement="top">
            <router-link to="/team/index">
              <i class="el-icon-s-flag"></i>
            </router-link>
          </el-tooltip>

          <el-tooltip
            v-if="isAdmin"
            class="item"
            effect="dark"
            :content="$t('background')"
            placement="top"
          >
            <router-link to="/admin/index">
              <i class="el-icon-s-tools"></i>
            </router-link>
          </el-tooltip>
        &nbsp;&nbsp;
          <el-tooltip class="item" effect="dark" :content="$t('more')" placement="top">
            <el-dropdown @command="dropdownCallback" trigger="click">
              <span class="el-dropdown-link">
                <i class="el-icon-caret-bottom el-icon--right header-btn-group" ></i>
              </span>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item>
                  <router-link to="/user/setting" :disabled="true">{{$t("Logged")}}:{{username}}</router-link>
                </el-dropdown-item>
                <el-tooltip class="item" effect="dark" :content="$t('under_develop')" placement="top">
                <el-dropdown-item >
                  <!-- <router-link to="/attachment/index" >{{$t("my_attachment")}}</router-link> -->
                  <router-link to="" >{{$t("my_attachment")}}</router-link>
                </el-dropdown-item>
                </el-tooltip>
                <el-dropdown-item :command="logout">{{$t("logout")}}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </el-tooltip>
        </div>
      </el-row>

      </el-container>
      <el-container class="container-narrow">
        <div class="container-thumbnails">
        <div class="search-box-div" v-if="itemList.length > 9">
          <div class="search-box el-input el-input--prefix">
            <input
              autocomplete="off"
              type="text"
              rows="2"
              validateevent="true"
              class="el-input__inner"
              v-model="keyword"
            />
            <span class="el-input__prefix">
              <i class="el-input__icon el-icon-search"></i>
            </span>
          </div>
        </div>

        <ul class="thumbnails" id="item-list" v-if="itemListByKeyword">
          <draggable
            v-model="itemListByKeyword"
            tag="span"
            group="item"
            @end="endMove"
            ghostClass="sortable-chosen"
          >
            <li
              class="text-center"
              v-for="item in itemListByKeyword"
              v-dragging="{ item: item, list: itemListByKeyword, group: 'item' }"
              :key="item.itemId"
            >
              <router-link
                class="thumbnail item-thumbnail"
                :to="'/' +  (item.item_domain ? item.item_domain:item.ID)"
                title
              >
                <!-- 自己创建的话显示项目设置按钮 -->
                <span
                  class="item-setting"
                  @click.prevent="clickItemSetting(item.ID)"
                  :title="$t('item_setting')"
                  v-if="item.creator"
                >
                  <i class="el-icon-setting"></i>
                </span>
                <!-- 如果是加入的项目的话，这里显示退出按钮 -->
                <span
                  class="item-exit"
                  @click.prevent="clickItemExit(item.ID)"
                  :title="$t('item_exit')"
                  v-if="! item.creator"
                >
                  <i class="el-icon-close"></i>
                </span>
                <p class="my-item">{{item.titlename}}</p>
                <!-- 如果是加密项目的话，这里显示一个加密图标 -->
                <span class="item-private" v-if="item.is_private">
                  <el-tooltip
                    class="item"
                    effect="dark"
                    :content="$t('private_tips')"
                    placement="right"
                  >
                    <i class="el-icon-lock"></i>
                  </el-tooltip>
                </span>
              </router-link>
            </li>
          </draggable>
          <li class="text-center">
            <router-link class="thumbnail item-thumbnail" to="/item/add" title>
              <p class="my-item">
                {{$t('new_item')}}
                <i class="el-icon-plus"></i>
              </p>
            </router-link>
          </li>
        </ul>
      </div>

    </el-container>
    <Footer> </Footer>

  </div>
</template>

<style scoped>
.container-narrow {
  margin: 0 auto;
  max-width: 930px;
  color: #fff;
}
.material-icons.md-20 {
  font-size: 20px;
}
.masthead {
  width: 100%;
  margin-top: 30px;
  background: rgb(44, 96, 106);
  color: #fff;
}
.header-btn-group {
  margin-top: -38px;
  font-size: 20px;
  color: #fff;
}
.header-btn-group a {
  margin-left: 25px;
  color: #fff;
}
.el-dropdown {
  font-size: 18px;
}
.el-dropdown-link,
a {
  color:#444;
}
.logo-title {
  margin-left: 10px;
}
.container-thumbnails {
  margin-top: 30px;
  max-width: 1000px;
}
.my-item {
  margin: 40px 5px;
}
.thumbnails li {
  float: left;
  margin-bottom: 20px;
  margin-left: 20px;
}
.thumbnails li a {
  color: #444;
  font-weight: bold;
  height: 100px;
  width: 180px;
}
.thumbnails li a:hover,
.thumbnails li a:focus {
  border-color: #f2f5e9;
  -webkit-box-shadow: none;
  box-shadow: none;
  text-decoration: none;
  background-color: #F4F7F9;
}
.thumbnail {
  display: block;
  padding: 4px;
  line-height: 20px;
  border: 1px solid #ddd;
  -webkit-border-radius: 4px;
  -moz-border-radius: 4px;
  border-radius: 4px;
  -webkit-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.055);
  -moz-box-shadow: 0 1px 3px rgba(0, 0, 0, 0.055);
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.055);
  -webkit-transition: all 0.2s ease-in-out;
  -moz-transition: all 0.2s ease-in-out;
  -o-transition: all 0.2s ease-in-out;
  transition: all 0.2s ease-in-out;
  list-style: none;
  background-color: #ffffff;
}
.item-setting {
  float: right;
  margin-right: 15px;
  margin-top: 5px;
  display: none;
}
.item-exit {
  float: right;
  margin-right: 5px;
  margin-top: 5px;
  display: none;
}
.item-private {
  float: right;
  margin-right: 15px;
  margin-top: -20px;
  display: none;
  cursor: default;
}
.thumbnails li a i {
  color: #777;
  font-weight: bold;
  margin-left: 5px;
}
.item-thumbnail:hover .item-setting {
  display: block;
}
.item-thumbnail:hover .item-exit {
  display: block;
}
.item-thumbnail:hover .item-private {
  display: block;
}
.search-box-div {
  width: 190px;
  margin-left: 60px;
}
.sortable-chosen .item-thumbnail {
  color: #ffffff;
  background-color: #ffffff;
}
</style>

<script>
import draggable from 'vuedraggable'
var $s = require('scriptjs')
var $ = require('jquery')

export default {
  components: {
    draggable
  },
  data () {
    return {
      currentDate: new Date(),
      itemList: [],
      isAdmin: false,
      keyword: '',
      lang: false,
      username: '',
      locale: 'ZH_CN'
    }
  },
  computed: {
    itemListByKeyword: function () {
      if (!this.keyword) {
        return this.itemList
      }
      const itemListByKeyword = []
      for (var i = 0; i < this.itemList.length; i++) {
        if (this.itemList[i].titlename.indexOf(this.keyword) > -1) {
          itemListByKeyword.push(this.itemList[i])
        }
      }
      return itemListByKeyword
    }
  },
  methods: {
    changeLang () {
      // 增加传入语言
      if (this.lang === false) {
        this.lang = true
        this.locale = 'EN_US'
      } else {
        this.lang = false
        this.locale = 'ZH_CN'
      }
      this.$cookies.set('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, 50)
      window.location.reload() // 进行刷新改变cookie里的值
    },
    getItemList () {
      var that = this
      var url = DocConfig.server + '/admin/list'
      var params = new URLSearchParams()
      that.$http.get(url, params).then(function (response) {
        if (response.data.status === 200) {
          var json = response.data.data
          that.itemList = json
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    feedback () {
      if (this.lang === true) {
        window.open('https://')
      } else {
        var msg =
          '你正在使用免费开源版doc，如有问题或者建议，请到github提issue：'
        msg +=
          "<a href='https://' target='_blank'></a><br>"
        msg +=
          '如果你觉得doc好用，不妨给开源项目点一个star。良好的关注度和参与度有助于开源项目的长远发展。'
        this.$alert(msg, {
          dangerouslyUseHTMLString: true
        })
      }
    },
    itemTopClass (top) {
      if (top) {
        return 'el-icon-arrow-down'
      };
      return 'el-icon-arrow-up'
    },
    bindItemEven () {
      // 这里偷个懒，直接用jquery来操作DOM。因为老版本的代码就是基于jquery的，所以复制过来稍微改下
      $s(['static/jquery.min.js'], () => {
      // 当鼠标放在项目上时将浮现设置和置顶图标
        $('.item-thumbnail').mouseover(function () {
          $(this).find('.item-setting').show()
          // $(this).find('.item-top').show()
          // $(this).find('.item-down').show()
        })
        // 当鼠标离开项目上时将隐藏设置和置顶图标
        $('.item-thumbnail').mouseout(function () {
          $(this).find('.item-setting').hide()
          $(this).find('.item-top').hide()
          $(this).find('.item-down').hide()
        })
      })
    },
    // 进入项目设置页
    clickItemSetting (itemId) {
      this.$router.push({ path: '/item/setting/' + itemId })
    },
    clickItemTop (itemId, oldTop) {
      if (oldTop) {
        var action = 'cancel'
      } else {
        action = 'top'
      }
      var that = this
      var url = DocConfig.server + '/item/top'
      var params = new URLSearchParams()
      params.append('action', action)
      params.append('item_id', itemId)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.getItemList()
          } else {
            that.$alert(response.data)
          }
        })
    },
    clickItemExit (itemId) {
      var that = this
      this.$confirm(that.$t('confirm_exit_item'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var url = DocConfig.server + '/item/exitItem'
        var params = new URLSearchParams()
        params.append('item_id', itemId)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            window.location.reload()
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    },
    async logout () {
      var url = DocConfig.server + '/exit'
      const { data: res } = await this.$http.get(url)
      window.sessionStorage.clear()
      this.$router.push('/')
      if (res.status === 200) return this.$message.success('成功退出')
    },
    user_info () {
      var that = this
      this.get_user_info(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.role === 1) {
            that.isAdmin = true
          }
        }
      })
    },
    sort_item (data) {
      var that = this
      var url = DocConfig.server + '/item/sort'
      var params = new URLSearchParams()
      params.append('data', JSON.stringify(data))
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          // window.location.reload();
        } else {
          that.$alert(response.data.message, '', {
            callback: function () {
              window.location.reload()
            }
          })
        }
      })
    },
    exchangeArray (data, oldIndex, newIndex) {
      const tmp = data[oldIndex]
      data.splice(oldIndex, 1)
      data.splice(newIndex, 0, tmp)
      return data
    },
    endMove (evt) {
      const data = {}
      const list = this.exchangeArray(
        this.itemList,
        evt.oldIndex,
        evt.newIndex
      )
      this.itemList = []
      this.$nextTick(() => {
        this.itemList = list
      })
      for (var i = 0; i < list.length; i++) {
        const key = list[i].ID
        data[key] = i + 1
      }
      this.sort_item(data)
    },
    dropdownCallback (data) {
      if (data) {
        data()
      }
    }
  },
  mounted () {
    this.getItemList()
    this.user_info()
    this.get_user_info(response => {
      if (response.data.status === 200) {
        this.username = response.data.data.username
      }
    })
    if (this.$cookies.get('lng') === 'EN_US') {
      this.locale = 'EN_US'
      this.lang = true
    } else {
      this.locale = 'ZH_CN'
      this.lang = false
    }

    this.$cookies.set('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, '30d')
  },
  watch: {
    locale (val) {
      this.$i18n.locale = val
    }
  },
  beforeDestroy () {
    this.$message.closeAll()
    // 去掉添加的背景色*/
    document.body.removeAttribute('class', 'grey-bg')
  }
}
</script>
