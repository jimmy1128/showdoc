<template>
  <div class="hello"
   @keydown.ctrl.83.prevent="save"
    @keydown.meta.83.prevent="save">
    <Header> </Header>

    <el-container class="container-narrow">

      <el-row class="masthead">
        <el-form :inline="true"   class="demo-form-inline" size="small">
          <el-form-item :label="$t('title')+' : '">
            <el-input  placeholder="" v-model="title"></el-input>
          </el-form-item>
          <el-form-item :label="$t('lang_setting')+' : '">
                <el-select  v-model="cid" :placeholder="$t('lang_choose')" class="lang" @change="selectCatalog">
                <el-option
                v-for="itemlang in belong_to_lang"
                :key="itemlang.id"
                :label="itemlang.name"
                :value="itemlang.id">
               </el-option>
              </el-select>
            </el-form-item>
          <el-form-item :label="$t('catalog') + ' : '">
            <el-select
              :placeholder="$t('optional')"
              class="cat"
              v-model="cat_id"
              v-if="belong_to_catalogs"
            >
             <el-option
                v-for="cat in belong_to_catalogs"
                :key="cat.cat_name"
                :label="cat.cat_name"
                :value="cat.cat_id"
              ></el-option>
            </el-select>
          </el-form-item>
            <el-form-item label>
            <el-button type="text" @click="ShowSortPage">{{
              $t('sort_pages')
            }}</el-button>
          </el-form-item>
          <el-form-item label>
            <el-button type="text" @click="ShowHistoryVersion">{{
              $t('history_version')
            }}</el-button>
          </el-form-item>
          <el-form-item class="pull-right">
              <el-dropdown  @command="dropdown_callback" split-button type="primary" size="medium" trigger="click" @click="save" title="Ctrl + S">
                <span id="save-page">{{ $t('save') }}</span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item :command="save_to_template">{{$t('save_to_templ')}}</el-dropdown-item>
                  <el-tooltip
                  class="item"
                  effect="dark"
                  :content="$t('lock_edit_tips')"
                  placement="left"
                >
                  <el-dropdown-item v-if="!isLock" :command="setLock">{{
                    $t('lock_edit')
                  }}</el-dropdown-item>
                </el-tooltip>
                <el-dropdown-item v-if="isLock" :command="unlock">{{
                  $t('cacel_lock')
                }}</el-dropdown-item>
                  <!-- <el-dropdown-item>保存前添加注释</el-dropdown-item> -->
                </el-dropdown-menu>
              </el-dropdown>
            <el-button type="" size="medium" @click="goback">{{$t('cancel')}}</el-button>
          </el-form-item>
        </el-form>

        <el-row class="fun-btn-group">
          <el-button type="" size="medium" @click="insert_api_template">{{$t('insert_apidoc_template')}}</el-button>
          <el-button type="" size="medium" @click="insert_database_template">{{$t('insert_database_doc_template')}}</el-button>
          <el-button type="" size="medium" @click.native="ShowTemplateList">{{$t('more_templ')}}</el-button>
            <el-dropdown split-button type="" style="margin-left:100px;" size="medium" trigger="hover" >
              {{$t('json_tools')}}
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item @click.native="ShowJsonToTable">{{$t('json_to_table')}}</el-dropdown-item>
                <el-dropdown-item @click.native="ShowJsonBeautify">{{$t('beautify_json')}}</el-dropdown-item>
                <el-dropdown-item @click.native="ShowPasteTable">{{$t('paste_insert_table')}}</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          <el-button type="" size="medium" @click="ShowRunApi">{{$t('http_test_api')}}</el-button>
          <el-badge :value="attachment_count" class="item">
            <el-button type size="medium" @click="ShowAttachment">{{$t('attachment')}}</el-button>
          </el-badge>
        </el-row>

        <Editormd v-bind:content="content" v-if="content" ref="Editormd"  type="editor" ></Editormd>

      </el-row>

        <!-- 更多模板 -->
        <TemplateList :callback="insertValue" ref="TemplateList"></TemplateList>

        <!-- 历史版本 -->
        <HistoryVersion :callback="insertValue" ref="HistoryVersion"></HistoryVersion>

        <!-- Json转表格 组件 -->
        <JsonToTable   :callback="insertValue" ref="JsonToTable" ></JsonToTable>

        <!-- Json格式化 -->
        <JsonBeautify :callback="insertValue" ref="JsonBeautify"></JsonBeautify>

        <!-- 附件列表 -->
      <AttachmentList
        :callback="insertValue"
        :item_id="item_id"
        :manage="true"
        :page_id="page_id"
        ref="AttachmentList"
      ></AttachmentList>
       <!-- 粘贴插入表格 -->
      <PasteTable
        :callback="insertValue"
        :item_id="item_id"
        :manage="true"
        :page_id="page_id"
        ref="PasteTable"
      ></PasteTable>

      <!-- 页面排序 -->
      <SortPage
        :callback="insertValue"
        :belong_to_catalogs="belong_to_catalogs"
        :item_id="item_id"
        :page_id="page_id"
        :cat_id="cat_id"
        ref="SortPage"
      ></SortPage>
      </el-container>
    <Footer> </Footer>
    <div class=""></div>
<!-- 模板存放的地方 -->
<div id="api-doc-templ"  ref="api_doc_templ" style="display:none">

**简要描述：**

- 用户注册接口

**请求URL：**
- ` http://xx.com/api/user/register `

**请求方式：**
- POST

**参数：**

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|username |是  |string |用户名   |
|password |是  |string | 密码    |
|name     |否  |string | 昵称    |

 **返回示例**

```
  {
    "status": 0,
    "data": {
      "uid": "1",
      "username": "12154545",
      "name": "吴系挂",
      "groupid": 2 ,
      "reg_time": "1436864169",
      "last_login_time": "0",
    }
  }
```

 **返回参数说明**

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|groupid |int   |用户组id，1：超级管理员；2：普通用户  |

 **备注**

- 更多返回错误代码请看首页的错误代码描述

</div>
<div id="database-doc-templ" ref="database_doc_templ" style="display:none">

-  用户表，储存用户信息

|字段|类型|空|默认|注释|
|:----    |:-------    |:--- |-- -|------      |
|uid    |int(10)     |否 |  |             |
|username |varchar(20) |否 |    |   用户名  |
|password |varchar(50) |否   |    |   密码    |
|name     |varchar(15) |是   |    |    昵称     |
|reg_time |int(11)     |否   | 0  |   注册时间  |

- 备注：无

</div>
  </div>
</template>

<style scoped>
  .container-narrow{
    margin: 0 auto;
    max-width: 90%;
  }
  .masthead{
    width: 100%;
    margin-top: 5px;
  }
  .cat{
    width: 180px;
  }
  .num{
    width: 60px;
  }
  .lang{
    width: 120px;
  }
  .fun-btn-group{
    margin-top: 15px;
    margin-bottom: 15px;
  }
</style>

<script>
import Editormd from '@/components/common/Editormd'
import JsonToTable from '@/components/common/JsonToTable'
import JsonBeautify from '@/components/common/JsonBeautify'
import TemplateList from '@/components/page/edit/TemplateList'
import HistoryVersion from '@/components/page/edit/HistoryVersion'
import AttachmentList from '@/components/page/edit/AttachmentList'
import PasteTable from '@/components/page/edit/PasteTable'
import SortPage from '@/components/page/edit/SortPage'
import { rederPageContent } from '@/models/page'
import {
  apiTemplateZh,
  databaseTemplateZh,
  apiTemplateEn,
  databaseTemplateEn
} from '@/models/template'
// var $s = require('scriptjs')
var $ = require('jquery')

export default {

  data () {
    return {
      currentDate: new Date(),
      itemList: {},
      content: '',
      title: '',
      item_id: 0,
      s_number: '',
      page_id: '',
      cid: '',
      name: '',
      cat_id: '',
      copy_page_id: '',
      attachment_count: '',
      catalogs: [],
      isLock: 0,
      intervalId: 0,
      saving: false,
      lang: '',
      langs: [],
      lang_array:[],
      infoForm: []
    }
  },
  computed: {
    // 新建/编辑目录时供用户选择的上级目录列表
    belong_to_catalogs: function () {
      if (!this.catalogs || this.catalogs.length <= 0) {
        return []
      }
      var Info = this.catalogs.slice(0)
      var cat_array = []
      // 这个函数将递归
      var rename = function (catalog, p_cat_name) {
        if (catalog.length > 0) {
          for (var j = 0; j < catalog.length; j++) {
            var cat_name = p_cat_name + ' / ' + catalog[j].cat_name
            cat_array.push({
              cat_id: catalog[j].cat_id,
              cat_name: cat_name
            })
            if (catalog[j].sub && catalog[j].sub.length > 0) {
              rename(catalog[j].sub, cat_name)
            }
          }
        }
      }
      for (var i = 0; i < Info.length; i++) {
        cat_array.push(Info[i])
        if (Info[i].sub != null) {
          rename(Info[i].sub, Info[i].cat_name)
        }
      }
      var no_cat = { cat_id: 0, cat_name: this.$t('none') }
      cat_array.push(no_cat)
      return cat_array
    },
    belong_to_lang: function () {
      var that = this
      if (that.infoForm.lang_list === undefined){
              return []
            }
            var langInfo =that.infoForm.lang_list.split(',')
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
  components: {
    Editormd,
    JsonToTable,
    JsonBeautify,
    TemplateList,
    AttachmentList,
    PasteTable,
    SortPage,
    HistoryVersion
  },
  methods: {
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
    get_item_info () {
      var that = this
      var url = DocConfig.server + '/item/detail'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info1 = response.data.data
            that.infoForm = Info1
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    // 获取页面内容
    get_page_content (page_id) {
      if (!page_id) {
        page_id = this.page_id
      }
      var that = this
      var url = DocConfig.server + '/admin/page'
      var params = new URLSearchParams()
      params.append('page_id', page_id)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            // that.$message.success("加载成功");
            that.content = rederPageContent(response.data.data.pagecontent)
            setTimeout(function () {
              that.insertValue(that.content, 1)
              document.body.scrollTop = document.documentElement.scrollTop = 0 // 回到顶部
            }, 500)
            setTimeout(function () {
              // 如果长度大于3000,则关闭预览
              if (that.content.length > 3000) {
                that.editor_unwatch()
              } else {
                that.editor_watch()
              }
              that.draft()
            }, 1000)
            that.title = response.data.data.pagetitle
            that.item_id = response.data.data.itemid
            that.s_number = response.data.data.snumber
            that.cid = response.data.data.cid
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    get_catalog (item_id, value) {
      var that = this
      var url = DocConfig.server + '/catListGroup'
      var params = new URLSearchParams()
      params.append('item_id', item_id)
      params.append('cid', value)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            that.catalogs = Info

            that.get_default_cat()
          } else {
            that.$message(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    // 获取默认该选中的目录
    get_default_cat () {
      var that = this
      var url = DocConfig.server + '/getDefaultCat'
      var params = new URLSearchParams()
      params.append('page_id', this.$route.params.page_id)
      params.append('item_id', that.$route.params.item_id)
      params.append('copy_page_id', this.copy_page_id)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            // that.$message.success("加载成功")
            var json = response.data.data
            that.cat_id = json
          } else {
            that.$alert(response.data.message)
          }
        })
    },
    // 根据语言选择目录
    selectCatalog () {
      var that = this
      that.get_catalog(this.$route.params.item_id, this.cid)
    },
    // 插入数据到编辑器中。插入到光标处。如果参数is_cover为真，则清空后再插入(即覆盖)。
    insertValue (value, is_cover) {
      if (value) {
        const childRef = this.$refs.Editormd // 获取子组件
        if (is_cover) {
          // 清空
          childRef.clear()
        };
        childRef.insertValue(value) // 调用子组件的方法
      };
    },
    // 插入api模板
    insert_api_template () {
      var val
      if (this.$cookies.get('lng') === 'ZH_CN') {
        val = apiTemplateZh
      } else {
        val = apiTemplateEn
      }
      this.insertValue(val)
    },
    // 插入数据字典模板
    insert_database_template () {
      var val
      if (this.$cookies.get('lng') === 'ZH_CN') {
        val = databaseTemplateZh
      } else {
        val = databaseTemplateEn
      }
      this.insertValue(val)
    },
    editor_unwatch () {
      const childRef = this.$refs.Editormd // 获取子组件
      childRef.editor_unwatch()
      if (sessionStorage.getItem('page_id_unwatch_' + this.page_id)) {
      } else {
        this.$message(this.$t('long_page_tips'))
        sessionStorage.setItem('page_id_unwatch_' + this.page_id, 1)
      }
    },
    editor_watch () {
      const childRef = this.$refs.Editormd // 获取子组件
      childRef.editor_watch()
    },
    // json转参数表格
    ShowJsonToTable () {
      const childRef = this.$refs.JsonToTable // 获取子组件
      childRef.dialogFormVisible = true
    },
    // json格式化
    ShowJsonBeautify () {
      const childRef = this.$refs.JsonBeautify // 获取子组件
      childRef.dialogFormVisible = true
    },
    ShowRunApi () {
      window.open('http://runapi.doc.cc/')
    },
    // 更多模板、模板列表
    ShowTemplateList () {
      const childRef = this.$refs.TemplateList // 获取子组件
      childRef.show()
    },
    ShowPasteTable () {
      const childRef = this.$refs.PasteTable // 获取子组件
      childRef.dialogFormVisible = true
    },
    // 展示历史版本
    ShowHistoryVersion () {
      const childRef = this.$refs.HistoryVersion // 获取子组件
      childRef.show()
    },
    // 展示页面排序
    ShowSortPage () {
      this.save(() => {
        const childRef = this.$refs.SortPage // 获取子组件
        childRef.show()
      })
    },
    save (callback) {
      var that = this
      if (this.saving) {
        return false
      }
      this.saving = true
      var loading = that.$loading()
      const childRef = this.$refs.Editormd
      var content = childRef.getMarkdown()
      var cat_id = this.cat_id
      var item_id = that.$route.params.item_id
      var page_id = that.$route.params.page_id
      var url = DocConfig.server + '/page/save'
      var params = new URLSearchParams()
      params.append('page_id', page_id)
      params.append('item_id', item_id)
      params.append('s_number', that.s_number)
      params.append('page_title', that.title)
      params.append('cat_id', cat_id)
      params.append('cid', this.cid)
      params.append('page_content', encodeURIComponent(content))
      params.append('is_urlencode', 1)
      that.$http.post(url, params)
        .then(function (response) {
          loading.close()
          that.saving = false
          if (response.data.status === 200) {
            if (typeof callback === 'function') {
              callback()
            } else {
              // that.$message.success("加载成功")
              // localStorage.removeItem('page_content')
              // that.$router.push({ path: '/' + item_id, query: { page_id: response.data.data.page_id } })
              that.$message({
                showClose: true,
                message: that.$t('save_success'),
                type: 'success'
              })
            }
            that.deleteDraft()
            if (page_id <= 0) {
              that.$router.push({
                path: '/page/edit/' + item_id + '/' + response.data.data.ID
              })
              that.page_id = response.data.data.ID
            }
          } else {
            that.$alert(response.data.message)
          }
        })
      // 设置一个最长关闭时间
      setTimeout(() => {
        loading.close()
        that.saving = false
      }, 20000)
    },
    goback () {
      var url = '/' + this.$route.params.item_id
      this.$router.push({ path: url, query: { page_id: this.$route.params.page_id } })
    },
    dropdown_callback (data) {
      if (data) {
        data()
      }
    },
    // 另存为模板
    save_to_template () {
      var that = this
      const childRef = this.$refs.Editormd
      var content = childRef.getMarkdown()
      this.$prompt(that.$t('save_templ_title'), ' ', {
      }).then(function (data) {
        var url = DocConfig.server + '/template/save'
        var params = new URLSearchParams()
        params.append('template_title', data.value)
        params.append('template_content', content)
        that.$http.post(url, params)
          .then(function (response) {
            if (response.data.status === 200) {
              that.$alert(that.$t('save_templ_text'))
            } else {
              that.$alert(response.data.message)
            }
          })
      })
    },
    ShowAttachment () {
      const childRef = this.$refs.AttachmentList // 获取子组件
      childRef.show()
    },
    /** 粘贴上传图片 **/
    on_paste () {
      var that = this
      var url = DocConfig.server + '/upload'
      document.addEventListener('paste', function (e) {
        var clipboard = e.clipboardData
        for (var i = 0, len = clipboard.items.length; i < len; i++) {
          if (clipboard.items[i].kind === 'file' || clipboard.items[i].type.indexOf('image') > -1) {
            var imageFile = clipboard.items[i].getAsFile()
            var form = new FormData()
            form.append('t', 'ajax-uploadpic')
            form.append('editormd-image-file', imageFile)
            var loading = ''
            var callback = function (type, data) {
              type = type || 'before'
              switch (type) {
                // 开始上传
                case 'before':
                  loading = that.$loading()
                  break
                  // 服务器返回错误
                case 'error':
                  loading.close()
                  that.$alert('图片上传失败')
                  break
                  // 上传成功
                case 'success':
                  loading.close()
                  if (data.success === 1) {
                    var value = '![](/' + data.url + ')'
                    that.insertValue(value)
                  } else {
                    that.$alert('失败')
                  }
                  break
              }
              console.log(type)
            }
            $.ajax({
              url: url,
              type: 'POST',
              dataType: 'json',
              data: form,
              processData: false,
              contentType: false,
              beforeSend: function () {
                // eslint-disable-next-line standard/no-callback-literal
                callback('before')
              },
              error: function () {
                // eslint-disable-next-line standard/no-callback-literal
                callback('error')
              },
              success: function (data) {
                // eslint-disable-next-line standard/no-callback-literal
                callback('success', data)
              }
            })
            e.preventDefault()
          }
        }
      })
    },
    // 草稿
    draft () {
      var that = this
      var pkey = 'page_content_' + this.page_id
      const childRef = this.$refs.Editormd
      // 定时保存文本内容到localStorage
      setInterval(() => {
        var content = childRef.getMarkdown()
        localStorage.setItem(pkey, content)
      }, 30 * 1000)
      // 检测是否有定时保存的内容
      var page_content = localStorage.getItem(pkey)
      if (
        page_content &&
        page_content.length > 0 &&
        page_content !== childRef.getMarkdown() &&
        childRef.getMarkdown() &&
        childRef.getMarkdown().length > 10
      ) {
        localStorage.removeItem(pkey)
        that
          .$confirm(that.$t('draft_tips'), '', {
            showClose: false
          })
          .then(() => {
            that.insertValue(page_content, true)
            localStorage.removeItem(pkey)
          })
          .catch(() => {
            localStorage.removeItem(pkey)
          })
      }
    },
    // 遍历删除草稿
    deleteDraft () {
      for (var i = 0; i < localStorage.length; i++) {
        var name = localStorage.key(i)
        if (name.indexOf('page_content_') > -1) {
          localStorage.removeItem(name)
        }
      }
    },
    // 锁定
    async setLock () {
      var url = DocConfig.server + '/page/setLock'
      var formdata = new FormData()
      formdata.append('page_id', this.page_id)
      formdata.append('item_id', this.item_id)
      if (this.page_id > 0) {
        const { data: res } = await this.$http.post(url, formdata)
        if (res.status === 200) {
          this.isLock = 1
        }
      }
    },
    // 解除锁定
    async unlock () {
      if (!this.isLock) {
        return // 本来处于未锁定中的话，不发起请求
      }
      var url = DocConfig.server + '/page/setLock'
      var formdata = new FormData()
      formdata.append('page_id', this.page_id)
      formdata.append('item_id', this.item_id)
      formdata.append('lock_to', 1000)
      const { data: res } = await this.$http.post(url, formdata)
      if (res.status === 200) {
        this.isLock = 0
      }
    },
    // 如果用户处于锁定状态的话，用心跳保持锁定
    heartBeatLock () {
      this.intervalId = setInterval(() => {
        if (this.isLock) {
          this.setLock()
        }
      }, 3 * 60 * 1000)
    },
    // 判断页面是否被锁定编辑
    async remoteIsLock () {
      var url = DocConfig.server + '/page/isLock'
      var formdata = new FormData()
      formdata.append('page_id', this.page_id)
      const { data: res } = await this.$http.post(url, formdata)
      // 判断已经锁定了不
      if (res.data.lock > 0) {
        if (res.data.isCurUser > 0) {
          this.isLock = 1
        } else {
          this.$alert(this.$t('locking') + res.data.lockUsername)
          this.goback()
        }
      } else {
        this.setLock() // 如果没有被别人锁定，则进编辑页面后自己锁定。
      }
    },
    // 由于页面关闭事件无法直接发起异步的ajax请求，所以用浏览器的navigator.sendBeacon来实现
    unLockOnClose () {
      let user_token = ''
      const userinfostr = localStorage.getItem('userinfo')
      if (userinfostr) {
        const userinfo = JSON.parse(userinfostr)
        if (userinfo && userinfo.user_token) {
          user_token = userinfo.user_token
        }
      }
      const analyticsData = new URLSearchParams({
        page_id: this.page_id,
        item_id: this.item_id,
        lock_to: 1000,
      })
      const url = DocConfig.server + '/page/setLock'
      if ('sendBeacon' in navigator) {
        navigator.sendBeacon(url, analyticsData)
      } else {
        var client = new XMLHttpRequest()
        client.open('POST', url, false)
        client.send(analyticsData)
      }
    }
  },
  mounted () {
    var that = this
    this.get_item_info()
    this.item_id = this.$route.params.item_id
    this.page_id = this.$route.params.page_id
    this.copy_page_id = this.$route.query.copy_page_id ? this.$route.query.copy_page_id : ''
    if (this.copy_page_id > 0) {
      this.get_page_content(this.copy_page_id)
    } else if (this.page_id > 0) {
      this.get_page_content(this.page_id)
    } else {
      this.item_id = this.$route.params.item_id
      this.content = this.$t('welcome_use_doc')
    }
    this.cid = this.itemList.cid
    this.get_catalog(this.$route.params.item_id)
    this.get_lang()
    this.heartBeatLock()
    this.remoteIsLock()
    that.on_paste()
    document.onkeydown = function (e) { // 对整个页面文档监听 其键盘快捷键
      var keyNum = window.event ? e.keyCode : e.which // 获取被按下的键值
      if (keyNum === 83 && e.ctrlKey) { // Ctrl +S 为保存
        that.save()
        e.preventDefault()
      }
    }
    // document.addEventListener('paste', this.on_paste)
    window.addEventListener('beforeunload', this.unLockOnClose)
  },
  beforeDestroy () {
  //   // 解除对粘贴上传图片的监听
  //   document.removeEventListener('paste', this.on_paste)
  this.$message.closeAll()
  clearInterval(this.intervalId)
   this.unlock()
   window.removeEventListener('beforeunload', this.unLockOnClose)
  }
}
</script>
