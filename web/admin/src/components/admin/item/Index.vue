<template>
  <div class="hello">
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label>
        <el-input v-model="item_name" :placeholder="$t('item_name')"></el-input>
      </el-form-item>
      <el-form-item label>
        <el-input v-model="username" :placeholder="$t('owner')"></el-input>
      </el-form-item>
      <!--   <el-form-item label="活动区域">
    <el-select v-model="formInline.region" placeholder="活动区域">
      <el-option label="区域一" value="shanghai"></el-option>
      <el-option label="区域二" value="beijing"></el-option>
    </el-select>
      </el-form-item>-->
      <el-form-item>
        <el-button @click="onSubmit">{{$t('search')}}</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="itemList" style="width: 100%">
      <el-table-column prop="titlename" :label="$t('item_name')" width="140"></el-table-column>
      <el-table-column prop="description" :label="$t('item_description')" width="140"></el-table-column>
      <el-table-column prop="password" :label="$t('privacy')" :formatter="formatPrivacy" width="80"></el-table-column>

      <el-table-column prop="id" :label="$t('link')" width="100">
        <template slot-scope="scope">
          <el-button @click="jump_to_item(scope.row)" type="text" size="small">{{$t('link')}}</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="username" :label="$t('owner')" width="160"></el-table-column>
      <el-table-column prop="member_num" :label="$t('memberCount')" width="80"></el-table-column>
      <el-table-column prop="created_at" :label="$t('add_time')" width="160" :formatter="dateFormat"></el-table-column>
      <el-table-column prop="item_domain" :label="$t('operation')">
        <template slot-scope="scope">
          <el-button @click="click_attorn_item(scope.row)" type="text" size="small">{{$t('attorn')}}</el-button>
          <el-button @click="delete_item(scope.row)" type="text" size="small">{{$t('delete')}}</el-button>
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

    <el-dialog :visible.sync="dialogAttornVisible" :close-on-click-modal="false" width="300px">
      <el-form>
        <el-form-item label>
          <el-input :placeholder="$t('attorn_username')" v-model="attornForm.username"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogAttornVisible = false">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="attorn">{{$t('attorn')}}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<style scoped>
</style>

<script>
const moment = require('moment')
export default {
  data () {
    return {
      page: 1,
      count: 10,
      item_name: '',
      username: '',
      itemList: [],
      total: 0,
      dialogAttornVisible: false,
      attornForm: {
        username: ''
      },
      attorn_item_id: ''
    }
  },
  methods: {
    dateFormat: function (row, column) {
      var date = row[column.property]
      if (date === undefined) {
        return ''
      }
      return moment(date).format('YYYY-MM-DD HH:mm:ss')
    },
    get_item_list () {
      var that = this
      var url = this.DocConfig.server + '/adminItem/getList'
      var params = new URLSearchParams()
      params.append('item_name', this.item_name)
      params.append('username', this.username)
      params.append('page', this.page)
      params.append('count', this.count)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          // that.$message.success("加载成功");
          var json = response.data.data
          that.itemList = json
          that.total = response.data.total
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    formatPrivacy (row, column) {
      if (row) {
        if (row.password.length > 0) {
          return this.$t('private')
        } else {
          return this.$t('public')
        }
      }
    },
    // 跳转到项目
    jump_to_item (row) {
      const url = '#/' + row.id
      window.open(url)
    },
    handleCurrentChange (currentPage) {
      this.page = currentPage
      this.get_item_list()
    },
    onSubmit () {
      this.page = 1
      this.get_item_list()
    },
    delete_item (row) {
      var that = this
      var url = this.DocConfig.server + '/adminItem/deleteItem'
      this.$confirm(that.$t('confirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('item_id', row.id)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            that.$message.success('删除成功')
            that.get_item_list()
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    },
    click_attorn_item (row) {
      this.dialogAttornVisible = true
      this.attorn_item_id = row.id
    },
    attorn () {
      var that = this
      var url = this.DocConfig.server + '/adminItem/attorn'
      var params = new URLSearchParams()
      params.append('item_id', that.attorn_item_id)
      params.append('username', this.attornForm.username)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogAttornVisible = false
            that.$message.success(that.$t('success'))
            that.get_item_list()
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  mounted () {
    this.get_item_list()
  },
  beforeDestroy () {
    this.$message.closeAll()
  }
}
</script>
