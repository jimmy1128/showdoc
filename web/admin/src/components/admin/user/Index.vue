<template>
  <div class="hello">
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item label>
        <el-input v-model="username" :placeholder="$t('username')"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="onSubmit">{{$t('search')}}</el-button>
      </el-form-item>
    </el-form>
    <el-button type="primary" @click="dialogAddVisible = true">{{$t('add_user')}}</el-button>
    <el-table :data="itemList" style="width: 100%">
      <el-table-column prop="username" :label="$t('username')" width="250"></el-table-column>
      <el-table-column prop="name" :label="$t('name')" width="300"></el-table-column>
      <el-table-column prop="role" :label="$t('userrole')" :formatter="formatGroup" width="150"></el-table-column>
      <el-table-column prop="CreatedAt" :label="$t('reg_time')" width="190" :formatter="dateFormat"></el-table-column>
      <el-table-column prop="last_login" :label="$t('last_login_time')" width="190" :formatter="dateFormat"></el-table-column>
      <el-table-column prop="item_domain" :label="$t('operation')">
        <template slot-scope="scope">
          <el-button @click="click_edit(scope.row)" type="text" size="small">{{$t('edit')}}</el-button>
          <el-button
            @click="delete_user(scope.row)"
            v-if="scope.row.role != 1"
            type="text"
            size="small"
          >{{$t('delete')}}</el-button>
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
            :placeholder="$t('username')"
            :readonly="addForm.uid > 0"
            v-model="addForm.username"
          ></el-input>
        </el-form-item>
        <el-form-item label>
          <el-input type="text" :placeholder="$t('name')" v-model="addForm.name"></el-input>
        </el-form-item>

        <el-form-item label v-if="addForm.uid <= 0">
          <el-input type="password" :placeholder="$t('password')" v-model="addForm.password"></el-input>
        </el-form-item>

        <el-form-item label v-if="addForm.uid > 0">
          <el-input type="password" :placeholder="$t('update_pwd_tips')" v-model="addForm.password"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="resetForm">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="add_user">{{$t('confirm')}}</el-button>
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
      itemList: [],
      username: '',
      page: 1,
      count: 10,
      total: 0,
      addForm: {
        username: '',
        password: '',
        uid: 0,
        name: ''
      },
      dialogAddVisible: false
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
    get_user_list () {
      var that = this
      var url = this.DocConfig.server + '/adminUser/getList'
      var params = new URLSearchParams()
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
    formatGroup (row, column) {
      if (row) {
        if (row.role === 1) {
          return this.$t('administrator')
        } else if (row.role === 2) {
          return this.$t('ordinary_users')
        } else {
          return ''
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
      this.get_user_list()
    },
    onSubmit () {
      this.page = 1
      this.get_user_list()
    },
    delete_user (row) {
      var that = this
      var url = this.DocConfig.server + '/adminUser/deleteUser'
      this.$confirm(that.$t('confirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('uid', row.ID)
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
      this.dialogAddVisible = true
      this.addForm = {
        uid: row.ID,
        name: row.name,
        username: row.username,
        password: ''
      }
    },
    add_user () {
      var that = this
      var url = this.DocConfig.server + '/adminUser/addUser'
      var params = new URLSearchParams()
      params.append('username', that.addForm.username)
      params.append('password', this.addForm.password)
      params.append('uid', this.addForm.uid)
      params.append('name', this.addForm.name)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogAddVisible = false
            that.addForm.password = ''
            that.addForm.username = ''
            that.addForm.uid = 0
            that.addForm.name = ''
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
    resetForm () {
      this.addForm = {
        uid: 0,
        name: '',
        username: '',
        password: ''
      }
      this.dialogAddVisible = false
    }
  },
  mounted () {
    this.get_user_list()
  },
  beforeDestroy () {
    this.$message.closeAll()
  }
}
</script>
