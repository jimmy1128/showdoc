<template>
  <div class="hello">
    <link href="https://fonts.googleapis.com/icon?family=Material+Icons" rel="stylesheet">
    <div v-if="show_menu_btn" id="header-right-btn">
      <el-dropdown trigger="click" @command="handleCommand">
        <span class="el-dropdown-link">
          <i class="el-icon-caret-bottom el-icon--right"></i>
        </span>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item command="goback">{{$t('goback')}}</el-dropdown-item>
          <el-dropdown-item command="share">{{$t('share')}}</el-dropdown-item>
          <el-dropdown-item v-if="item_info.itempermn" command="new_page">{{$t('new_page')}}</el-dropdown-item>
          <el-dropdown-item v-if="item_info.itempermn" command="new_catalog">{{$t('new_catalog')}}</el-dropdown-item>
          <el-dropdown-item v-if="item_info.itempermn" command="edit_page">{{$t('edit_page')}}</el-dropdown-item>
          <el-dropdown-item v-if="item_info.itempermn" command="copy">{{$t('copy')}}</el-dropdown-item>
          <el-dropdown-item
            v-if="item_info.itempermn"
            command="ShowHistoryVersion"
          >{{$t('history_version')}}</el-dropdown-item>
          <el-dropdown-item v-if="item_info.itempermn" command="export">{{$t('export')}}</el-dropdown-item>
          <el-dropdown-item
            v-if="item_info.itempermn"
            command="delete_page"
          >{{$t('delete_interface')}}</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
    <div class="op-bar" v-if="show_op_bar">
      <!-- <div >
     <el-tooltip class="item" effect="dark" :content="$t('language1')" placement="top">
     <el-dropdown trigger="click" class='el-select-lang' @command="LangChange" style="padding-left:0px" placement="top-start" v-if="selected != null">
      <span class="el-dropdown-link" >
        <svg class="icons" aria-hidden="true" style="float: left; color: #8492a6; font-size: 10px" v-html="selected">
            </svg>
      </span>
      <el-dropdown-menu slot="dropdown">
        <el-dropdown-item
        v-for="itemlang in belong_to_lang"
        :key="itemlang.id"
        :label="itemlang.name"
        :value="itemlang"
        :command="itemlang"> {{ itemlang.name }}
        <svg class="icon" aria-hidden="true" style="float: left; color: #8492a6; font-size: 13px" v-html="itemlang.icon" >
            </svg>
        </el-dropdown-item>
      </el-dropdown-menu>
    </el-dropdown>
    </el-tooltip>
     </div> -->
      <span v-if="item_info.is_login">
        <el-tooltip class="item" effect="dark" :content="$t('goback')" placement="left">
          <router-link to="/item/index">
            <i class="el-icon-back"></i>
          </router-link>
        </el-tooltip>

        <el-tooltip class="item" effect="dark" :content="$t('share')" placement="top">
          <i class="el-icon-share" @click="share_page"></i>
        </el-tooltip>
        <el-tooltip
          v-if="! item_info.itempermn"
          class="item"
          effect="dark"
          :content="$t('detail')"
          placement="top"
        >
          <i class="el-icon-info" @click="show_page_info"></i>
        </el-tooltip>
      </span>

      <span v-if="item_info.itempermn">
        <el-tooltip class="item" effect="dark" :content="$t('new_page')" placement="top">
          <i class="el-icon-plus" @click="new_page"></i>
        </el-tooltip>
        <el-tooltip class="item" effect="dark" :content="$t('new_catalog')" placement="left">
          <i class="el-icon-folder" @click="mamage_catalog"></i>
        </el-tooltip>
        <el-tooltip class="item" effect="dark" :content="$t('edit_page')" placement="top">
          <i class="el-icon-edit" @click="edit_page"></i>
        </el-tooltip>
        <el-tooltip class="item" effect="dark" :content="$t('language')" placement="top" v-if="link !== '/0'">
            <router-link :to="link" title :key="$route.params.ID" v-on:click.native="$router.go(0)">
              <i class="material-icons md-20" :key="lang? 'ZH_CN':'EN_US'" @click="changeLang()">language</i>
            </router-link>
          </el-tooltip>

        <el-tooltip
          v-show="!showMore"
          class="item"
          effect="dark"
          :content="$t('more')"
          placement="top"
        >
          <i class="el-icon-caret-top" @click="showMoreAction"></i>
        </el-tooltip>
        <el-tooltip
          v-show="showMore"
          class="item"
          effect="dark"
          :content="$t('more')"
          placement="top"
        >
          <i class="el-icon-caret-bottom" @click="hideMoreAction"></i>
        </el-tooltip>

        <span v-show="showMore">
          <el-tooltip class="item" effect="dark" :content="$t('copy')" placement="left">
            <router-link :to="'/page/edit/'+item_id+'/0?copy_page_id='+page_id">
              <i class="el-icon-document"></i>
            </router-link>
          </el-tooltip>
          <el-tooltip class="item" effect="dark" :content="$t('history_version')" placement="top">
            <i class="el-icon-goods" @click="ShowHistoryVersion"></i>
          </el-tooltip>
          <el-tooltip class="item" effect="dark" :content="$t('detail')" placement="top">
            <i class="el-icon-info" @click="show_page_info"></i>
          </el-tooltip>
          <el-tooltip class="item" effect="dark" :content="$t('export')" placement="left">
            <router-link :to="'/item/export/'+item_info.id" v-if="item_info.itempermn">
              <i class="el-icon-download"></i>
            </router-link>
          </el-tooltip>
          <el-tooltip class="item" effect="dark" :content="$t('delete_interface')" placement="top">
            <i class="el-icon-delete" @click="delete_page"></i>
          </el-tooltip>

          <span v-if="item_info.itemcreator">
            <el-tooltip class="item" effect="dark" :content="$t('item_setting')" placement="top">
              <router-link :to="'/item/setting/'+item_info.id" v-if="item_info.itemcreator">
                <i class="el-icon-setting"></i>
              </router-link>
            </el-tooltip>
          </span>
        </span>
      </span>
    </div>

    <el-dialog
      :title="$t('share_page')"
      :visible.sync="dialogVisible"
      width="600px"
      :close-on-click-modal="false"
    >
      <p>
        {{$t('item_page_address')}} :
        <code>{{share_page_link}}</code>
        <i
          class="el-icon-document-copy"
          v-clipboard:copy="share_page_link"
          v-clipboard:success="onCopy"
        ></i>
      </p>
      <p v-if="false" style="border-bottom: 1px solid #eee;">
        <img id="qr-page-link" style="width:114px;height:114px;" :src="qr_page_link" />
      </p>

      <div v-show="item_info.itempermn">
        <el-checkbox
          v-model="isCreateSiglePage"
          @change="checkCreateSiglePage"
        >{{$t('create_sigle_page')}}</el-checkbox>

        <p v-if="isCreateSiglePage">
          {{$t('single_page_address')}} :
          <code>{{share_single_link}}</code>
          <i
            class="el-icon-document-copy"
            v-clipboard:copy="share_single_link"
            v-clipboard:success="onCopy"
          ></i>
        </p>
        <p></p>
        <p>{{$t('create_sigle_page_tips')}}</p>
      </div>

      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogVisible = false">{{$t('confirm')}}</el-button>
      </span>
    </el-dialog>

    <!-- 历史版本 -->
    <HistoryVersion
      :page_id="page_id"
      :is_show_recover_btn="false"
      :is_modal="false"
      callback="insertValue"
      ref="HistoryVersion"
    ></HistoryVersion>
  </div>
</template>

<style scoped>
.op-bar {
  color: #333;
  position: fixed;
  top: 100px;
  margin-left: 840px;
  max-width: 250px;
}
.op-bar i {
  cursor: pointer;
  font-size: 16px;
  margin-right: 55px;
  margin-bottom: 30px;
}
.icon-folder {
  width: 15px;
  height: 12px;
  cursor: pointer;
  margin-right: 55px;
}
a {
  color: #333;
}
#header-right-btn {
  font-size: 20px;
  top: 15px;
  right: 5%;
  cursor: pointer;
  position: fixed;
}
.el-dropdown-link {
  color: #000;
  font-size: 18px;
  font-weight: bolder;
}
.el-icon-document-copy {
  cursor: pointer;
}
.icon {
  width: 1.9em;
  height: 1.9em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
  padding-top: 5px;
}
.el-select-lang{
    border: none;
    border-radius: 0px;
}
.icons {
  width: 5em;
  height: 5em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}

</style>

<script>
import HistoryVersion from '@/components/page/edit/HistoryVersion'
import '../../../../../static/fonts/iconfont.js'
export default {
  props: {
    item_id: {},
    item_domain: {},
    page_id: {},
    item_info: {},
    page_info: []

  },
  data () {
    return {
      menu: [],
      dialogVisible: false,
      qr_page_link: '#',
      qr_single_link: '#',
      share_page_link: '',
      share_single_link: '',
      copyText1: '',
      copyText2: '',
      isCreateSiglePage: false,
      showMore: false,
      lang: '',
      show_menu_btn: false,
      show_op_bar: true,
      link: {},
      langs: [],
      locale: '',
      selected: '',
      itemlang: [],
      icon: '',
      value: {},
      lang_array:[]
    }
  },
  components: {
    HistoryVersion
  },
  computed:{
    belong_to_lang: function () {
      var that = this
      if (that.item_info.lang_list === undefined){
              return []
            }
            var langInfo =that.item_info.lang_list.split(',')
            var lang_array = []

            if (that.langs.length > 0){
              for (var k=0; k < langInfo.length; k++){
                if (langInfo[k] === ""){
                  return that.langs
                }
              for (var j=0; j < that.langs.length; j++){
                if (that.langs[j].id === parseInt(langInfo[k])){
                  lang_array.push({
                    id: that.langs[j].id,
                    name: that.langs[j].name,
                    icon: that.langs[j].icon
                  })
                }
              }
              }
            }
            return lang_array

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
      //this.$cookies.set('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, '30d')
      // this.$cookies.config('30d')
      // window.location.reload() // 进行刷新改变cookie里的值
    },
    get_lang () {
      var that = this
      var url = DocConfig.server + '/lang'
      var params = new URLSearchParams()
      that.$http.get(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.langs = response.data.data
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },

    LangChange (value) {
      this.$emit('itemlangId', value.id)
      this.selected = value
      // if (this.$cookies.get('selected') === null) {
        this.$cookies.set('selected', value.icon)
      // }
      if (value.name === 'English' || value.icon === '#icon-world-flag_-GBR--UnitedKingdom') {
        this.locale = 'EN_US'
        this.lang = true
        // this.$cookies.set('lng', this.locale === 'EN_US' ? this.locale : this.locale, '1d')
      } else {
        this.lang = false
        this.locale = 'ZH_CN'
        // this.$cookies.set('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, '30d')
      }

    },
    edit_page () {
      var page_id = this.page_id > 0 ? this.page_id : 0
      var url = '/page/edit/' + this.item_id + '/' + page_id
      this.$router.push({ path: url })
    },
    share_page () {
      var page_id = this.page_id > 0 ? this.page_id : 0
      const path = this.item_domain ? this.item_domain : this.item_id
      this.share_page_link =
        this.getRootPath() + '#/' + path + '?page_id=' + page_id
      // this.share_single_link= this.getRootPath()+"/page/"+page_id ;
      this.qr_page_link =
        DocConfig.server +
        '/api/common/qrcode&size=3&url=' +
        encodeURIComponent(this.share_page_link)
      // this.qr_single_link = DocConfig.server +'/api/common/qrcode&size=3&url='+encodeURIComponent(this.share_single_link);
      this.dialogVisible = true
      this.copyText1 =
        this.item_info.item_name +
        ' - ' +
        this.page_info.pagetitle +
        '\r\n' +
        this.share_page_link
      this.copyText2 =
        this.page_info.pagetitle + '\r\n' + this.share_single_link
    },
    dropdown_callback (data) {
      if (data) {
        data()
      }
    },
    show_page_info () {
      var html =
        '本页面由 ' +
        this.page_info.authoruid +
        ' 于 ' +
        this.page_info.CreatedAt +
        ' 更新'
      this.$alert(html)
    },
    // 展示历史版本
    ShowHistoryVersion () {
      const childRef = this.$refs.HistoryVersion // 获取子组件
      childRef.show()
    },
    delete_page () {
      var page_id = this.page_id > 0 ? this.page_id : 0
      var that = this
      var url = DocConfig.server + '/page/delete'
      this.$confirm(that.$t('comfirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('page_id', page_id)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            window.location.reload()
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    },
    onCopy () {
      this.$message(this.$t('copy_success'))
    },
    checkCreateSiglePage (newvalue) {
      if (newvalue) {
        this.CreateSiglePage()
      } else {
        this.$confirm(this.$t('cancelSingle'), ' ', {
          confirmButtonText: this.$t('cancelSingleYes'),
          cancelButtonText: this.$t('cancelSingleNo'),
          type: 'warning'
        }).then(() => {
          this.CreateSiglePage()
        }, () => {
          this.isCreateSiglePage = true
        })
      }
    },
    CreateSiglePage () {
      var page_id = this.page_id > 0 ? this.page_id : 0
      var that = this
      var url = DocConfig.server + '/page/createSinglePage'
      var params = new URLSearchParams()
      params.append('page_id', page_id)
      params.append('isCreateSiglePage', this.isCreateSiglePage)
      that.axios.post(url, params).then(function (response) {
        if (response.data.error_code === 0) {
          var unique_key = response.data.data.unique_key
          if (unique_key) {
            that.share_single_link = that.getRootPath() + '#/p/' + unique_key
          } else {
            that.share_single_link = ''
          }
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    new_page () {
      var url = '/page/edit/' + this.item_info.id + '/0'
      this.$router.push({ path: url })
    },
    mamage_catalog () {
      var url = '/catalog/' + this.item_info.id
      this.$router.push({ path: url })
    },
    showMoreAction () {
      this.showMore = true
      var element = document
        .getElementById('page_md_content')
        .getElementsByClassName('open-list')
      element[0].style.top = '330px'
    },
    hideMoreAction () {
      this.showMore = false
      var element = document
        .getElementById('page_md_content')
        .getElementsByClassName('open-list')
      element[0].style.top = '230px'
    },
    handleCommand (command) {
      switch (command) {
        case 'goback':
          this.$router.push({ path: '/item/index' })
          break
        case 'share':
          this.share_page()
          break
        case 'new_page':
          this.new_page()
          break
        case 'new_catalog':
          this.mamage_catalog()
          break
        case 'edit_page':
          this.edit_page()
          break
        case 'ShowHistoryVersion':
          this.ShowHistoryVersion()
          break
        case 'copy':
          this.$router.push({
            path:
              '/page/edit/' +
              this.item_info.item_id +
              '/0?copy_page_id=' +
              this.page_id
          })
          break
        case 'export':
          this.$router.push({ path: '/item/export/' + this.item_info.item_id })
          break
        case 'delete_page':
          this.delete_page()
          break
      }
    },
    get_item_info () {
      var that = this
      var url = DocConfig.server + '/item/detail'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            if (Info.password) {
              that.isOpenItem = false
            }
            if (Info.lang === 1) {
              Info.lang = false
            } else if (Info.lang === 2) {
              Info.lang = true
            }
            that.MyForm = Info
            that.link = '/' + that.MyForm.link
            that.lang = that.MyForm.lang
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    loadConfig () {
      var item_id = this.$route.params.item_id ? this.$route.params.item_id : 0
      var that = this
      var url = DocConfig.server + '/adminSetting/loadLangConfig'
      var params = new URLSearchParams()
      params.append('id', item_id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return
          }
          that.selected = '<use xlink:href="' + response.data.data.icon + '"></use>'
          that.$cookies.set('selected', that.selected)
          that.value = response.data.data
          that.value.icon = '<use xlink:href="' + response.data.data.icon + '"></use>'

          if (response.data.access === 1) {
            // defaultlang 是全局语言
            localStorage.setItem('defaultlang', that.value.id)
            that.LangChange(that.value)
          }
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    publicloadConfig (value) {
      var that = this
      var url = DocConfig.server + '/public/lang'
      var params = new URLSearchParams()
      params.append('lang', value)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return
          }
          that.value = response.data.data
          that.value.icon = '<use xlink:href="' + response.data.data.icon + '"></use>'
          if (response.data.access === 1) {
             that.LangChange(that.value)
             this.selected = '<use xlink:href="' + response.data.data.icon + '"></use>'
             this.$cookies.set('selected', that.selected)
          }
        } else {
          this.$alert(response.data.message)
        }
      })
    },
    get_lang1 (value) {
      var that = this
      var url = DocConfig.server + '/lang/info'
      var params = new URLSearchParams()
      params.append('icon', value)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return
          }
          that.selected = '<use xlink:href="' + response.data.data.icon + '"></use>'
          that.$cookies.set('selected', that.selected)
          that.value = response.data.data
          that.value.icon = '<use xlink:href="' + response.data.data.icon + '"></use>'
        } else {
          that.$alert(response.data.message)
        }
      })
    }
  },

  created(){
  },
  mounted () {
    var that = this
    var page_id = this.$route.query.page_id ? this.$route.query.page_id : 0

    this.get_item_info()
    //this.get_lang()
    if (this.$cookies.get('lng') === null ) {
      this.locale = DocConfig.lang
    this.$cookies.set('lng', this.locale === 'EN_US' ? this.locale : this.locale, '1d')
    }
    //  if (this.$cookies.get('selected') === null) {
    //   this.loadConfig()
    // }
    // if (this.$cookies.get('selected') === null) {
    //   this.publicloadConfig(page_id)
    // }
    if (this.page_info.unique_key) {
      this.isCreateSiglePage = true
      this.share_single_link =
        this.getRootPath() + '#/p/' + this.page_info.unique_key
    }
    if (this.$cookies.get('lng') === 'EN_US') {
      this.locale = 'EN_US'
      this.lang = true
    } else {
      this.locale = 'ZH_CN'
      this.lang = false
    }
    // this.selected = this.$cookies.get('selected')
    // if (this.selected !== null) {
    //   this.get_lang1(this.selected)
    //   //this.$cookies.remove('selected')
    // }
    this.lang = this.$cookies.get('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, 50)
    document.onkeydown = function (e) {
      // 对整个页面文档监听 其键盘快捷键
      var keyNum = window.event ? e.keyCode : e.which // 获取被按下的键值
      if (keyNum === 69 && e.ctrlKey) {
        // Ctrl +e 为编辑
        that.edit_page()
        e.preventDefault()
      }
    }
    if
    (
      this.isMobile() ||
      (window.innerWidth < 1300 && !this.item_info.is_login)
    ) {
      this.show_menu_btn = false
      this.show_op_bar = false
    }
    if
    (
      this.isMobile() ||
      (window.innerWidth < 1300 && this.item_info.is_login)
    ) {
      this.show_menu_btn = true
      this.show_op_bar = false
    }

  },
  watch: {
    page_info: function () {
      if (this.page_info.unique_key) {
        this.isCreateSiglePage = true
        this.share_single_link =
          this.getRootPath() + '#/p/' + this.page_info.unique_key
      } else {
        this.isCreateSiglePage = false
        this.share_single_link = ''
      }
    },
    locale (val) {
      this.$i18n.locale = val
    }
  },
  destroyed () {
    document.onkeydown = undefined

  }
}
</script>
