<template>
  <div class="hello">
    <Header></Header>

    <el-container>
      <el-card class="center-card">
        <el-button type="text" class="add-cat" @click="addTeamMember">{{$t('add_member')}}</el-button>
        <el-button type="text" class="goback-btn" @click="goback">{{$t('back_to_team')}}</el-button>
        <el-table align="left" :data="list" height="400" style="width: 100%">
          <el-table-column prop="member_username" :label="$t('member_username')"></el-table-column>
          <el-table-column prop="name" :label="$t('name')" width="150"></el-table-column>
          <el-table-column prop="CreatedAt" :label="$t('Join_time')" :formatter="dateFormat"></el-table-column>

          <el-table-column prop :label="$t('operation')" width="110">
            <template slot-scope="scope">
              <el-button
                @click="deleteTeamMember(scope.row.ID)"
                type="text"
                size="small"
              >{{$t('delete')}}</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-card>

      <el-dialog :visible.sync="dialogFormVisible" width="300px" :close-on-click-modal="false">
        <el-form>
          <el-form-item :label="$t('member_username')+':'">
            <el-select
              v-model="MyForm.member_username"
              multiple
              filterable
              reserve-keyword
              placeholder
              :loading="loading"
            >
              <el-option
                v-for="item in memberOptions"
                :key="item.ID"
                :label="item.label"
                :value="item.value"
              ></el-option>
            </el-select>
          </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">{{$t('cancel')}}</el-button>
          <el-button type="primary" @click="MyFormSubmit">{{$t('confirm')}}</el-button>
        </div>
      </el-dialog>
    </el-container>

    <Footer></Footer>
  </div>
</template>

<script>
const moment = require('moment')
export default {
  components: {},
  data () {
    return {
      MyForm: {
        id: '',
        member_username: ''
      },
      list: [],
      loading: false,
      dialogFormVisible: false,
      team_id: '',
      memberOptions: []
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
    getList () {
      var that = this
      var url = this.DocConfig.server + '/member/list '
      var params = new URLSearchParams()
      params.append('teamid', this.team_id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          var Info = response.data.data
          that.list = Info
          that.getAllUser()
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    MyFormSubmit () {
      var that = this
      var url = this.DocConfig.server + '/member/save'
      var params = new URLSearchParams()
      params.append('teamid', this.team_id)
      params.append('member_username', this.MyForm.member_username)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          that.dialogFormVisible = false
          that.getList()
          that.MyForm = {}
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    deleteTeamMember (id) {
      var that = this
      var url = this.DocConfig.server + '/member/delete'
      this.$confirm(that.$t('confirm_delete'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('id', id)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            that.getList()
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    },
    addTeamMember () {
      this.MyForm = []
      this.dialogFormVisible = true
    },
    goback () {
      this.$router.push({ path: '/team/index' })
    },
    getAllUser () {
      var that = this
      var url = this.DocConfig.server + '/admin/user'
      var params = new URLSearchParams()
      params.append('username', '')
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          var Info = response.data.data
          var newInfo = []
          // 过滤掉已经是成员的用户
          for (var i = 0; i < Info.length; i++) {
            const isMember = that.isMember(Info[i].value)
            if (!isMember) {
              newInfo.push(Info[i])
            }
          }
          that.memberOptions = []
          for (let index = 0; index < newInfo.length; index++) {
            that.memberOptions.push({
              value: newInfo[index].value,
              label: newInfo[index].value,
              key: newInfo[index].value
            })
          }
          console.log(that.memberOptions)
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    // 判断某个用户是否已经是会员
    isMember (username) {
      const list = this.list
      for (var i = 0; i < list.length; i++) {
        if (list[i].member_username === username) {
          return true
        }
      }
      return false
    }
  },
  mounted () {
    this.team_id = this.$route.params.team_id
    this.getList()
    this.getAllUser()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello {
  text-align: left;
}
.add-cat {
  margin-left: 10px;
}
.center-card {
  text-align: left;
  width: 800px;
  height: 600px;
}
.goback-btn {
  z-index: 999;
  margin-left: 550px;
}
</style>

<!-- 全局css -->
<style >
.el-table .success-row {
  background: #f0f9eb;
}
</style>
