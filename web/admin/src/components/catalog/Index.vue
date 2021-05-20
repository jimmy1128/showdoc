<template>
  <div class="hello">
    <Header> </Header>

    <el-container>
      <el-card class="center-card">
      <el-row>
      <el-button  type="text" class="add-cat" @click="add_cat">{{$t('add_cat')}}</el-button>
      <el-button type="text" class="goback-btn" @click="goback">{{$t('goback')}}</el-button>
      </el-row>
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
      </el-card>

      <el-dialog :visible.sync="dialogFormVisible"  width="300px">
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
                    v-for="item in lang"
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
    </el-container>
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
export default {
  name: 'Login',
  components: {
    SortPage,
    Copy
  },
  data () {
    return {
      MyForm: {
        cat_id: 0,
        parent_cat_id: '',
        cid: '',
        cat_name: '',
        s_number: ''
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
      item_id: '',
      curl_cat_id: ''
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
    }
  },
  methods: {
    tableRowClassName ({ row, rowIndex }) {
      if (row.level === '2') {
        return 'success-row'
      }
      return ''
    },
    get_catalog () {
      var that = this
      var url = this.DocConfig.server + '/catListGroup'
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
                if (Info[i].sub && Info[i].sub.length > 0) {
                  node.children = duang(Info[i].sub)
                }
                treeData.push(node)
              }
              return treeData
            }
            that.treeData = duang(Info)
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    get_lang () {
      var that = this
      var url = this.DocConfig.server + '/lang'
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
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    MyFormSubmit () {
      var that = this
      var url = this.DocConfig.server + '/catalogs/save'
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
        parent_cat_id: node.parent.data.parent_cat_id,
        cat_name: data.label
      }
      this.dialogFormVisible = true
    },
    delete_cat (node, data) {
      var that = this
      var cat_id = data.id
      var url = this.DocConfig.server + `/delcatalogs/${cat_id}`
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
          cat_id: 0,
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
    handleDragEnd (node1, node2, position, evt) {
      const treeData2 = this.dimensionReduction(this.treeData)
      this.request('/catalog/batUpdate', {
        item_id: this.$route.params.item_id,
        cats: JSON.stringify(treeData2)
      })
    },
    // 将目录数组降维
    dimensionReduction (treeData) {
      const treeData2 = []
      var pushTreeData = function (OneData, parent_cat_id, level, i) {
        treeData2.push({
          cat_id: OneData.id,
          cat_name: OneData.label,
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
    }
  },
  mounted () {
    this.get_catalog()
    this.get_lang()
    this.item_id = this.$route.params.item_id
  },
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
</style>

<!-- 全局css -->
<style >
.el-table .success-row {
  background: #f0f9eb;
}
</style>
