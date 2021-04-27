<template>
  <div class="hello">
    <Header> </Header>

    <el-container>
      <el-card class="center-card">

      <el-button  type="text" class="add-cat" @click="add_cat">{{$t('add_cat')}}</el-button>
      <el-button type="text" class="goback-btn" @click="goback">{{$t('goback')}}</el-button>
       <el-table align="left"
            :data="catalogs"
             height="400"
             :row-class-name="tableRowClassName"
            style="width: 100%">
            <el-table-column
              prop="catname"
              :label="$t('cat_name')"
              width="160">
            </el-table-column>
            <el-table-column
              prop="CreatedAt"
              :label="$t('add_time')"
              width="210">
            </el-table-column>
            <el-table-column
              prop="snumber"
              :label="$t('s_number')">
            </el-table-column>
            <el-table-column
              prop=""
              width="110"
              :label="$t('operation')">
              <template slot-scope="scope">
                <el-button @click="edit(scope.row)" type="text" size="small" :disabled=" scope.row.level > 0 ? false : true ">{{$t('edit')}}</el-button>
                <el-button @click="delete_cat(scope.row.ID)" type="text" size="small" >{{$t('delete')}}</el-button>
              </template>
            </el-table-column>
          </el-table>

            </el-card>
      <el-dialog :visible.sync="dialogFormVisible"  width="300px">
        <el-form >
            <el-form-item :label="$t('cat_name')+' : '" >
              <el-input  :placeholder="$t('input_cat_name')" v-model="MyForm.catname"></el-input>
            </el-form-item>
            <el-form-item :label="$t('parent_cat_name')+' : '" >
              <el-select v-model="MyForm.parentcatid" :placeholder="$t('none')">
                  <el-option
                    v-for="item in catalogs_level_2"
                    :key="item.ID"
                    :label="item.catname"
                    :value="item.ID">
                  </el-option>
                </el-select>
            </el-form-item>
            <el-form-item :label="$t('s_number')" >
              <el-input  :placeholder="$t('s_number_explain')" v-model="MyForm.snumber"></el-input>
            </el-form-item>
        </el-form>

        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisible = false">{{$t('cancel')}}</el-button>
          <el-button type="primary" @click="MyFormSubmit" >{{$t('confirm')}}</el-button>
        </div>
      </el-dialog>
    </el-container>

    <Footer> </Footer>
  </div>
</template>

<script>

export default {
  name: 'Login',
  components: {
  },
  data () {
    return {
      MyForm: {
        ID: 0,
        parentcatid: '',
        catname: '',
        snumber: ''
      },
      catalogs: [],
      catalogs_level_2: [],
      dialogFormVisible: false,
      btnedit: false
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

            // 创建上级目录选项
            var Info2 = Info.slice(0)
            var no_cat = { ID: 0, catname: that.$t('none') }
            Info2.unshift(no_cat)
            that.catalogs_level_2 = Info2
            var cat_array = []
            for (var i = 0; i < Info.length; i++) {
              // Info[i].cat_name = Info[i].cat_name
              cat_array.push(Info[i])
              if (Info[i].pages.length > 0) {
                for (var j = 0; j < Info[i].pages.length; j++) {
                  Info[i].pages[j].catname = ' -- ' + Info[i].pages[j].pagetitle
                  cat_array.push(Info[i].pages[j])
                }
              }
            }
            that.catalogs = cat_array
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
      params.append('catid', this.MyForm.ID)
      params.append('parentcatid', this.MyForm.parentcatid)
      params.append('catname', this.MyForm.catname)
      params.append('snumber', this.MyForm.snumber)
      that.$http.post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogFormVisible = false
            that.get_catalog()
            that.MyForm = []
          } else {
            that.$alert(response.data.status)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },

    edit (row) {
      console.log(row)
      var temp = {}
      temp = JSON.parse(JSON.stringify(row))
      if (temp.catname) {
        temp.catname = temp.catname.replace(' -- ', '')
      }
      if (temp.parentcatid === '0') {
        temp.parentcatid = ''
      }
      this.MyForm = temp
      this.dialogFormVisible = true
    },
    delete_cat (ID) {
      var that = this
      var url = this.DocConfig.server + `/delcatalogs/${ID}`
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
    add_cat () {
      this.MyForm = []
      this.dialogFormVisible = true
    },
    goback () {
      var url = '/' + this.$route.params.item_id
      this.$router.push({ path: url })
    }
  },
  mounted () {
    this.get_catalog()
  }
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
  width: 600px;
  height: 500px;
}
.goback-btn{
  z-index: 999;
  margin-left: 400px;
}
</style>

<!-- 全局css -->
<style >
.el-table .success-row {
  background: #f0f9eb;
}
</style>
