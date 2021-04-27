<template>
  <div class="hello">
    <p class="tips">{{$t("recycle_tips")}}</p>
    <!-- 页面列表 -->
    <el-table
      align="left"
      class="recycle-table"
      v-if="lists.length>0"
      :data="lists"
      style="width: 100%"
    >
      <el-table-column prop="pagetitle" :label="$t('page_title')"></el-table-column>
      <el-table-column prop="del_name" :label="$t('deleter')"></el-table-column>
      <el-table-column prop="DeletedAt" :label="$t('del_time')"></el-table-column>
      <el-table-column prop :label="$t('operation')">
        <template slot-scope="scope">
          <el-button @click="recover(scope.row.ID)" type="text" size="small">{{$t("recover")}}</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: 'Login',
  components: {},
  data () {
    return {
      MyForm: {
        username: '',
        is_readonly: false
      },
      MyForm2: {
        team_id: ''
      },
      lists: []
    }
  },
  methods: {
    get_list () {
      var that = this
      var url = this.DocConfig.server + '/recycle/getList'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          var Info = response.data.data
          that.lists = Info
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    recover (page_id) {
      var that = this
      var url = this.DocConfig.server + '/recycle/recover'
      this.$confirm(this.$t('recover_tips'), ' ', {
        confirmButtonText: that.$t('confirm'),
        cancelButtonText: that.$t('cancel'),
        type: 'warning'
      }).then(() => {
        var params = new URLSearchParams()
        params.append('item_id', that.$route.params.item_id)
        params.append('page_id', page_id)
        that.$http.post(url, params).then(function (response) {
          if (response.data.status === 200) {
            that.get_list()
          } else {
            that.$alert(response.data.message)
          }
        })
      })
    }
  },
  mounted () {
    this.get_list()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello {
  text-align: left;
}
.tips {
  margin-left: 10px;
  color: #9ea1a6;
}
.recycle-table {
  max-height: 400px;
  overflow: auto;
}
</style>
