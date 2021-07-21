<template>
  <div class="hello">
    <Header> </Header>

    <el-container class="container-narrow">

      <el-dialog :title="$t('templ_list')" :visible.sync="dialogTableVisible">
        <el-table :data="content">
          <el-table-column property="CreatedAt" :label="$t('save_time')" width="170" :formatter="dateFormat"></el-table-column>
          <el-table-column property="title" :label="$t('templ_title')" ></el-table-column>
          <el-table-column
            :label="$t('operation')"
            width="150">
            <template slot-scope="scope">
              <el-button @click="insertTemplate(scope.row)" type="text" size="small">{{$t('insert_templ')}}</el-button>
              <el-button type="text" size="small" @click="deleteTemplate(scope.row)">{{$t('delete_templ')}}</el-button>
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
export default {
  props: {
    callback: {}
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
      var url = DocConfig.server + '/admin/template'
      var params = new URLSearchParams()
      // params.append('page_id',  that.$route.params.page_id);
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
              that.$alert(that.$t('no_templ_text'))
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
    insertTemplate (row) {
      this.callback(row.template_content)
      this.dialogTableVisible = false
    },
    deleteTemplate (row) {
      var id = row.id
      var that = this
      var url = DocConfig.server + '/template/delete'
      var params = new URLSearchParams()
      params.append('id', id)
      that.axios.post(url, params)
        .then(function (response) {
          if (response.data.status === 0) {
            that.get_content()
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

  }
}
</script>
