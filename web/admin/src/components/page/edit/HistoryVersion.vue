<template>
  <div class="hello">
    <Header> </Header>

    <el-container class="container-narrow">

      <el-dialog :title="$t('history_version')" :modal="is_modal" :visible.sync="dialogTableVisible" :close-on-click-modal="false">
        <el-table :data="content">
          <el-table-column property="CreatedAt" :label="$t('update_time')" width="170" :formatter="dateFormat"></el-table-column>
          <el-table-column property="authorusername" :label="$t('update_by_who')" ></el-table-column>
          <el-table-column property="pagecomments" :label="$t('remark')">
            <template slot-scope="scope">
              {{ scope.row.page_comments }}
              <el-button v-if="is_show_recover_btn" @click="editComments(scope.row)" type="text" size="small">{{ $t('edit') }}</el-button>
            </template>
          </el-table-column>
          <el-table-column
            label="操作"
            width="150">
            <template slot-scope="scope">
              <el-button @click="preview_diff(scope.row)" type="text" size="small">{{$t('overview')}}</el-button>
              <el-button v-if="is_show_recover_btn" type="text" size="small" @click="recover(scope.row)">{{$t('recover_to_this_version')}}</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-dialog>

      </el-container>
    <Footer> </Footer>
    <div class=""></div>
  </div>
</template>

<style>
</style>

<script>
const moment = require('moment')
export default {
  props: {
    callback: '',
    page_id: '',
    is_modal: true,
    is_show_recover_btn: true
  },
  data () {
    return {
      currentDate: new Date(),
      content: [],
      dialogTableVisible: false
    }
  },
  components: {
  },
  methods: {
    dateFormat: function (row, column) {
      var date = row[column.property]
      if (date === undefined) {
        return ''
      }
      return moment(date).format('YYYY-MM-DD HH:mm:ss')
    },
    get_content () {
      var that = this
      var url = this.DocConfig.server + '/page/history'
      var params = new URLSearchParams()
      params.append('page_id', that.page_id)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            // that.$message.success("加载成功");
            var json = response.data.data
            if (json.length > 0) {
              that.content = response.data.data
              that.dialogTableVisible = true
            } else {
              that.dialogTableVisible = false
              that.$alert('no data')
            }
          } else {
            that.dialogTableVisible = false
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    show () {
      this.get_content()
    },
    recover (row) {
      this.callback(row.page_content, true)
      this.dialogTableVisible = false
    },
    preview_diff (row) {
      var page_history_id = row.page.page_history_id
      var page_id = this.$route.params.page_id
      var url = '#/page/diff/' + page_id + '/' + page_history_id
      window.open(url)
    }
  },
  mounted () {

  }
}
</script>
