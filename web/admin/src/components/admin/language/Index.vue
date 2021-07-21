<template>
  <div class="hello">
    <el-button type="primary" @click="dialogAddVisible = true">{{
      $t('add_language')
    }}</el-button>
    <el-button type="info" @click="update_lang_config">{{
      $t('default')
    }}</el-button>
    <el-table :data="langList" style="width: 100%">
      <el-table-column
        prop="id"
        label="ID"
        width="200"
      ></el-table-column>
      <el-table-column prop="name" :label="$t('name')"></el-table-column>
      <el-table-column prop="icon" :label="$t('icon')" >
       <template slot-scope="props" >
       <svg  class="icon" aria-hidden="true" v-html="props.row.icon">
      </svg></template></el-table-column>
      <el-table-column prop="item_domain" :label="$t('operation')">
        <template slot-scope="scope">
          <el-button @click="click_edit(scope.row)" type="text" size="small">{{
            $t('edit')
          }}</el-button>
          <el-button
            @click="delete_user(scope.row)"
            type="text"
            size="small"
            >{{ $t('delete') }}</el-button
          >
          <el-button
            @click="update_config(scope.row)"
            type="text"
            size="small"
            >{{ $t('default1') }}</el-button
          >
        </template>
      </el-table-column>

    </el-table>

    <div class="block">
      <span class="demonstration"></span>
      <el-pagination
        @current-change="handleCurrentChange"
        :page-size="count"
        layout="total, prev, pager, next"
        :total="total"
      ></el-pagination>
    </div>

    <el-dialog
      :visible.sync="dialogAddVisible"
      :before-close="resetForm"
      :close-on-click-modal="false"
      width="300px"
    >
      <el-form>
        <el-form-item label>
          <el-input
            type="text"
            :placeholder="$t('language1')"
            v-model="addForm.name"
          ></el-input>
        </el-form-item>
      </el-form>
      <div style="margin: 0 auto;width: 260px">
    <e-icon-picker ref="iconPicker" v-model="addForm.icon" :options="options"/>
    名称：{{ icon }}
    <e-icon :icon-name="icon"/>
  </div>
      <div slot="footer" class="dialog-footer">
        <el-button @click="resetForm">{{ $t('cancel') }}</el-button>
        <el-button type="primary" @click="add_user">{{
          $t('confirm')
        }}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped></style>

<script>
import '../../../../static/fonts/iconfont.js'
import { eIconSymbol } from 'e-icon-picker'
import iconfont from '../../../../static/fonts/iconfont.json'
export default {
  data () {
    return {
      options: {
        FontAwesome: false,
        ElementUI: false,
        eIcon: false,
        eIconSymbol: false,
        addIconList: [],
        removeIconList: []
      },
      langList: [],
      ID: '',
      page: 1,
      count: 10,
      total: 0,
      addForm: {
        id: 0,
        name: '',
        icon: ''
      },
      default_value: '',
      dialogAddVisible: false,
      icon: ''
    }
  },
  methods: {
    addIcon () {
      const icon = eIconSymbol(iconfont)
      this.options.addIconList = icon.list
    },
    get_user_list () {
      var that = this
      var url = DocConfig.server + '/lang'
      var params = new URLSearchParams()
      params.append('page', this.page)
      params.append('count', this.count)
      that.$http.get(url, params).then(function (response) {
        if (response.data.status === 200) {
          // that.$message.success("加载成功");
          var json = response.data.data
          that.langList = json
          that.total = json.total
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    // 跳转到项目
    handleCurrentChange (currentPage) {
      this.page = currentPage
      this.get_user_list()
    },
    delete_user (row) {
      var that = this
      var url = DocConfig.server + '/lang/delete'
      this.$confirm(that.$t('confirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('id', row.id)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            that.$message.success('success')
            that.get_user_list()
            that.username = ''
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    },
    click_edit (row) {
      var lists = row.icon.split('"')
      this.dialogAddVisible = true
      this.addForm = {
        id: row.id,
        name: row.name,
        icon: lists[1]
      }
    },
    add_user () {
      var that = this
      var url = DocConfig.server + '/lang/add'
      var params = new URLSearchParams()
      params.append('id', that.addForm.id)
      params.append('name', that.addForm.name)
      params.append('icon', that.addForm.icon)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogAddVisible = false
            that.addForm.name = ''
            that.icon = ''
            that.$message.success(that.$t('success'))
            that.get_user_list()
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    update_lang_config () {
      var that = this
      var url = DocConfig.server + '/adminSetting/saveLangConfig'
      var params = new URLSearchParams()
      params.append('lang', 0)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          that.$message.success(response.data.message)
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    update_config (row) {
      var that = this
      var url = DocConfig.server + '/adminSetting/saveLangConfig'
      var params = new URLSearchParams()
      params.append('lang', row.id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          that.$message.success(response.data.message)
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    loadConfig () {
      var that = this
      var url = DocConfig.server + '/adminSetting/loadLangConfig'
      var params = new URLSearchParams()
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          if (response.data.data.length === 0) {
            return
          }
          that.default_value = 0
        } else {
          this.$alert(response.data.message)
        }
      })
    },
    resetForm () {
      this.addForm = {
        name: ''
      }
      this.dialogAddVisible = false
    }
  },
  mounted () {
    this.get_user_list()
    this.loadConfig()
    this.addIcon()
    localStorage.removeItem('defaultlang')
  },
  beforeDestroy () {
    this.$message.closeAll()
  }
}
</script>
<style>
.icon {
  width: 4em;
  height: 4em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>
