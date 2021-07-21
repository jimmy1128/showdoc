<template>
  <div class="hello">
    <Header> </Header>

      <el-card class="center-card">
      <el-row>
      <el-button  type="text" class="add-cat" @click="add_cat()">{{$t('add_cat')}}</el-button>
      <el-button type="text" class="goback-btn" @click="goback">{{$t('goback')}}</el-button>
      </el-row>
      <el-tabs v-model="activeName" @tab-click="handleClick">
      <el-tab-pane :label="$t('memu_manage')" name="first">
      <div class="itxst">
      <p class="tips" v-if="treeData.length > 1">{{ $t('cat_tips') }}</p>
       <el-tree
          class="tree-node"
          :data="treeData"
          node-key="id"
          :props="defaultProps"
          default-expand-all
          @node-drag-end="handleDragEnd"
          draggable
        >
          <span class="custom-tree-node" slot-scope="{ node, data }">
            <span>{{ node.label }}</span>
            <span class="right-bar">
              <el-button
                type="text"
                size="mini"
                :title="$t('edit')"
                class="el-icon-edit"
                @click.stop="edit(node, data)"
              ></el-button>
              <el-button
                type="text"
                size="mini"
                class="el-icon-plus"
                :title="$t('add_cat')"
                @click.stop="add_cat(node, data)"
              ></el-button>
              <el-button
                type="text"
                size="mini"
                class="el-icon-document"
                :title="$t('sort_pages')"
                @click.stop="showSortPage(node, data)"
              ></el-button>
              <el-button
                type="text"
                size="mini"
                class="el-icon-copy-document"
                :title="$t('copy_or_mv_cat')"
                @click.stop="copyCat(node, data)"
              ></el-button>
              <el-button
                type="text"
                size="mini"
                class="el-icon-delete"
                @click.stop="delete_cat(node, data)"
              ></el-button>
            </span>
          </span>
        </el-tree>
        </div>
        </el-tab-pane>
        <el-tab-pane :label="$t('page_manage')" name="second">
        <el-container class="container-narrow">
        <el-card class="page-card left">
          <div slot="header" class="clearfix">
            <el-select v-model="id" :placeholder="$t('select')" @change="langlist(id)">
              <el-option
                v-for="item in belong_to_lang"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            </el-select>
        <el-button style="float: right; padding: 3px 0" type="text" @click="endMove" >{{$t('update')}}</el-button>
        </div>
        <div class="itxst">
        <draggable v-model="pages" tag="div" group="page" @end="endMove" :scroll="true" animation="300">
            <div class="page-box" v-for="(page,index) in pages" :key="page.ID">
              <span class="page-name"><i class="fa fa-times close" @click="removeAt(index)"></i>{{ page.pagetitle }}</span>
            </div>
          </draggable>
        </div>
        </el-card>
        <el-card class="page-card right">
          <div slot="header" class="clearfix ">
            <el-select v-model="langid" :placeholder="$t('select')" @change="langlist1(langid)" >
              <el-option
                v-for="item in belong_to_lang"
                :key="item.id"
                :label="item.name"
                :value="item.id">
              </el-option>
            </el-select>
        <el-button style="float: right; padding: 3px 0" type="text" @click="endMove2">{{$t('update')}}</el-button>
        </div>
        <div class="itxst">
        <draggable v-model="pages2" tag="div" group="page" @end="endMove2" :scroll="true" animation="300">
            <div class="page-box" v-for="(page,idx) in pages2" :key="page.ID">
              <span class="page-name"> <i class="fa fa-times close" @click="removeAt2(idx)"></i>{{ page.pagetitle }}</span>

            </div>
          </draggable>
        </div>
        </el-card>
        </el-container>
        </el-tab-pane>
        </el-tabs>
      </el-card>

      <el-dialog :visible.sync="dialogFormVisible"  width="300px" :close-on-click-modal="false" >
        <el-form >
            <el-form-item :label="$t('cat_name')+' : '" >
              <el-input  :placeholder="$t('input_cat_name')" v-model="MyForm.cat_name"></el-input>
            </el-form-item>
            <el-form-item :label="$t('parent_cat_name')+' : '" >
              <el-select v-model="MyForm.parent_cat_id" :placeholder="$t('none')">
                  <el-option
                    v-for="item in belong_to_catalogs"
                    :key="item.cat_id"
                    :label="item.cat_name"
                    :value="item.cat_id">
                  </el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="$t('cat_lang')+' : '" >
              <el-select v-model="MyForm.cid" :placeholder="$t('select')">
                  <el-option
                    v-for="item in belong_to_lang"
                    :key="item.id"
                    :label="item.name"
                    :value="item.id">
                  </el-option>
                </el-select>
            </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">{{$t('cancel')}}</el-button>
          <el-button type="primary" @click="MyFormSubmit" >{{$t('confirm')}}</el-button>
        </div>
      </el-dialog>
    <SortPage
      :belong_to_catalogs="belong_to_catalogs"
      :item_id="item_id"
      :cat_id="curl_cat_id"
      ref="SortPage"
    ></SortPage>

    <Copy
      v-if="copyFormVisible"
      :item_id="item_id"
      :cat_id="curl_cat_id"
      :callback="copyCallback"
    ></Copy>
    <Footer> </Footer>
  </div>
</template>

<script>
import SortPage from '@/components/page/edit/SortPage'
import Copy from './Copy'
import draggable from 'vuedraggable'
export default {
  name: 'Login',
  components: {
    SortPage,
    Copy,
    draggable
  },
  data () {
    return {
      MyForm: {
        cat_id: 0,
        parent_cat_id: '',
        cid: '',
        cat_name: '',
        s_number: '',
      },
      lang: [],
      catalogs: [],
      dialogFormVisible: false,
      copyFormVisible: false,
      treeData: [],
      btnedit: false,
      defaultProps: {
        children: 'children',
        label: 'label'
      },
      activeName: 'first',
      item_id: '',
      curl_cat_id: '',
      pages: [],
      pages2: [],
      id: '',
      langid: '',
      drag: false,
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
            if (that.lang.length > 0){
              for (var k=0; k < langInfo.length; k++){
                if (langInfo[k] === ""){
                  return that.lang
                }
              for (var j=0; j < that.lang.length; j++){
                if (that.lang[j].id === parseInt(langInfo[k])){
                  lang_array.push({
                    id: that.lang[j].id,
                    name: that.lang[j].name,
                    icon: that.lang[j].icon
                  })
                }
              }
              }
            }
            return lang_array
    }
  },
  methods: {
    tableRowClassName ({ row, rowIndex }) {
      if (row.level === '2') {
        return 'success-row'
      }
      return ''
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
    removeAt2(idx) {
      this.pages2.splice(idx, 1);
    },
    removeAt(idx) {
      this.pages.splice(idx, 1);
    },
    get_catalog () {
      var that = this
      var url = DocConfig.server + '/catListGroup'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            that.catalogs = Info
            that.treeData = []
            var duang = function (Info) {
              var treeData = []
              for (var i = 0; i < Info.length; i++) {
                const node = { children: [] }
                node.id = Info[i].cat_id
                node.label = Info[i].cat_name
                node.cid = Info[i].cid
                if (Info[i].sub && Info[i].sub.length > 0) {
                  node.children = duang(Info[i].sub)
                }
                treeData.push(node)
              }
              return treeData
            }
            that.treeData = duang(Info)
          } else {
            that.$message(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    get_lang () {
      var that = this
      var url = DocConfig.server + '/lang'
      var params = new URLSearchParams()
      that.$http.get(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var lang = response.data.data
            // 创建上级目录选项
            var Lang2 = lang.slice(0)
            var no_cat = { ID: 0, name: that.$t('none') }
            Lang2.unshift(no_cat)
            that.lang = Lang2
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    MyFormSubmit () {
      var that = this
      var url = DocConfig.server + '/catalogs/save'
      var params = new URLSearchParams()
      params.append('itemid', that.$route.params.item_id)
      params.append('catid', this.MyForm.cat_id)
      params.append('parentcatid', this.MyForm.parent_cat_id)
      params.append('catname', this.MyForm.cat_name)
      params.append('cid', this.MyForm.cid)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogFormVisible = false
            that.get_catalog()
            that.MyForm = []
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    edit (node, data) {
      this.MyForm = {
        cat_id: data.id,
        parent_cat_id: node.parent.data.id,
        cat_name: data.label,
        cid: data.cid
      }
      this.dialogFormVisible = true
    },
    langlist(id){
      var that = this
      that.get_pages (id,2)
      this.id = id

    },
    langlist1(langid){
      var that = this
      that.get_pages (langid)
       this.langid = langid
    },
    delete_cat (node, data) {
      var that = this
      var cat_id = data.id
      var url = DocConfig.server + `/delcatalogs/${cat_id}`
      this.$confirm(that.$t('confirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('item_id', that.$route.params.item_id)
        that.$http.post(url, params)
          .then(function (response) {
            if (response.data.status === 200) {
              that.get_catalog()
            } else {
              that.$alert(response.data.message)
            }
          })
      })
    },
    get_pages (id , value) {
      var that = this
      var url = DocConfig.server + '/page/list'
      var params = new URLSearchParams()
      params.append('item_id', this.item_id)
      params.append('lang_id', id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            if (value == 2){
              that.pages = response.data.data
            } else {
              that.pages2 = response.data.data
            }
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    resetForm () {
      this.MyForm = {
        cat_id: 0,
        parent_cat_id: '',
        cat_name: '',
        s_number: '',
        cid: ''
      }
    },
    add_cat (node, data) {
      if (node && data.id) {
        this.MyForm = {
          cat_id: '',
          parent_cat_id: data.id,
          cat_name: '',
          s_number: '',
          cid: ''
        }
      } else {
        this.resetForm()
      }
      this.dialogFormVisible = true
    },
    goback () {
      var url = '/' + this.$route.params.item_id
      this.$router.push({ path: url })
    },
    async handleDragEnd (node1, node2, position, evt) {
      const treeData2 = this.dimensionReduction(this.treeData)
      var url = DocConfig.server + '/catalog/batUpdate'
      var formdata = new FormData()
      formdata.append('item_id', this.$route.params.item_id)
      formdata.append('cats', JSON.stringify(treeData2))
      const { data: res } = await this.$http.post(url, formdata)
      if (res.status === 200) {}
    },
    // 将目录数组降维
    dimensionReduction (treeData) {
      const treeData2 = []
      var pushTreeData = function (OneData, parent_cat_id, level, i) {
        treeData2.push({
          id: OneData.id,
          name: OneData.label,
          parent_cat_id: parent_cat_id,
          level: level,
          s_number: i + 1
        })
        if (Object.prototype.hasOwnProperty.call(OneData, 'children')) {
          for (var j = 0; j < OneData.children.length; j++) {
            pushTreeData(OneData.children[j], OneData.id, level + 1, j)
          }
        }
      }
      for (var i = 0; i < treeData.length; i++) {
        pushTreeData(treeData[i], 0, 2, i)
      }
      return treeData2
    },
    handleClick(tab, event) {
        console.log(tab, event);
    },
    // 展示页面排序
    showSortPage (node, data) {
      this.curl_cat_id = data.id
      const childRef = this.$refs.SortPage // 获取子组件
      childRef.show()
    },
    copyCat (node, data) {
      this.curl_cat_id = data.id
      this.copyFormVisible = true
    },
    copyCallback () {
      this.copyFormVisible = false
      this.get_catalog()
    },
    endMove (evt) {
      const data = {}
      for (var i = 0; i < this.pages.length; i++) {
        const key = this.pages[i].ID
        data[key] = i + 1
      }
      this.sort_page(data)
    },
    endMove2 (evt) {
      const data = {}
      for (var i = 0; i < this.pages2.length; i++) {
        const key = this.pages2[i].ID
        data[key] = i + 1
      }
      this.sort_page(data)
    },
    sort_page (data) {
      var that = this
      var url = DocConfig.server + '/page/sortbypage'
      var params = new URLSearchParams()
      params.append('pages', JSON.stringify(data))
      params.append('item_id', this.item_id)
      params.append('lang_id', this.id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          //that.get_pages(this.langid)
          that.$message.success(that.$t('modify_success'))
          // window.location.reload();
        } else {
          that.$alert(response.data.message, '', {
            callback: function () {
              window.location.reload()
            }
          })
        }
      })
    }
  },
  mounted () {
    this.get_catalog()
    this.get_lang()
    this.get_item_info()
    this.item_id = this.$route.params.item_id
  },
  // watch: {
  //   id: function () {
  //     this.get_pages()
  //   },
  // },
  beforeDestroy () {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello{
  text-align: left;
}
.add-cat{
  margin-left: 10px;
}
.center-card{
  text-align: left;
  width: 800px;
  height: 700px;
}
.goback-btn{
  z-index: 999;
  margin-left: 400px;
}
.cat-box {
  background-color: rgb(250, 250, 250);
  width: 100%;
  height: 40px;
  margin-bottom: 10px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
}
.cat-name {
  line-height: 40px;
  margin-left: 20px;
  color: #262626;
}
.tree-node {
  margin-top: 20px;
}
.custom-tree-node {
  width: 100%;
}
.right-bar {
  float: right;
  margin-right: 20px;
}
.tips {
  margin-left: 10px;
  color: #9ea1a6;
  font-size: 11px;
}
.page-card {
  width: 350px;
  height: 570px;
}
.clearfix:before,
.clearfix:after {
    display: table;
    content: "";
}
.clearfix:after {
    clear: both
}
.right {
margin-left: 30px;
}
.dialog-body {
  min-height: 400px;
  max-height: 90%;
  overflow-x: hidden;
  overflow-y: auto;
}
.page-box {
  background-color: rgb(250, 250, 250);
  width: 98%;
  height: 30px;
  margin-top: 10px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  overflow-y: auto
}
.page-name {
  line-height: 30px;
  margin-left: 20px;
  color: #262626;
}
.tips {
  margin-left: 10px;
  color: #9ea1a6;
  font-size: 9px;
}
.close{
  float: right;
  padding-top: 8px;
  padding-bottom: 8px;
  padding-right: 12px;
}
.itxst {
  width: 100%;
  margin: 10px;
  text-align: left;
  height: 450px;
  overflow-y: auto;
  overflow-x: hidden;
}
</style>

<!-- 全局css -->
<style >
.el-table .success-row {
  background: #f0f9eb;
}
</style>
