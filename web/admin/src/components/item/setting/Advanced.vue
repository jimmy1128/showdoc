<template>
  <div class="hello">
    <p style="height:40px;"></p>
    <p>
      <el-tooltip :content="$t('attorn_tips')" placement="top-start">
        <el-button class="a_button" @click="dialogAttornVisible = true">{{$t('attorn')}}</el-button>
      </el-tooltip>
    </p>
    <p>
      <el-tooltip :content="$t('archive_tips')" placement="top-start">
        <el-button class="a_button" @click="dialogArchiveVisible = true">{{$t('archive')}}</el-button>
      </el-tooltip>
    </p>

    <p>
      <el-tooltip :content="$t('delete_tips')" placement="top-start">
        <el-button class="a_button" @click="dialogDeleteVisible = true">{{$t('delete')}}</el-button>
      </el-tooltip>
    </p>

     <p>
      <el-tooltip :content="$t('link_tips')" placement="top-start">
        <el-button class="a_button" @click="dialogLinkVisible = true">{{$t('link_item')}}</el-button>
      </el-tooltip>
    </p>

    <el-dialog
      :visible.sync="dialogAttornVisible"
      :modal="false"
      width="300px"
      :close-on-click-modal="false"
    >
      <el-form>
        <el-form-item label>
          <el-input :placeholder="$t('attorn_username')" v-model="attornForm.username"></el-input>
        </el-form-item>
        <el-form-item label>
          <el-input
            type="password"
            :placeholder="$t('input_login_password')"
            v-model="attornForm.password"
          ></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogAttornVisible = false">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="attorn">{{$t('attorn')}}</el-button>
      </div>
    </el-dialog>

    <el-dialog
      :visible.sync="dialogArchiveVisible"
      :modal="false"
      width="300px"
      :close-on-click-modal="false"
    >
      <el-form>
        <el-form-item label>
          <el-input
            type="password"
            :placeholder="$t('input_login_password')"
            v-model="archiveForm.password"
          ></el-input>
        </el-form-item>
      </el-form>

      <p class="tips">{{$t('archive_tips2')}}</p>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogArchiveVisible = false">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="archive">{{$t('archive')}}</el-button>
      </div>
    </el-dialog>

    <el-dialog
      :visible.sync="dialogDeleteVisible"
      :modal="false"
      width="300px"
      :close-on-click-modal="false"
    >
      <el-form>
        <el-form-item label>
          <el-input
            type="password"
            :placeholder="$t('input_login_password')"
            v-model="deleteForm.password"
          >></el-input>
        </el-form-item>
      </el-form>

      <p class="tips">
        <el-tag type="danger">{{$t('delete_tips')}}</el-tag>
      </p>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogDeleteVisible = false">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="deleteItem">{{$t('delete')}}</el-button>
      </div>
    </el-dialog>
    <el-dialog
      :visible.sync="dialogLinkVisible"
      :modal="false"
      width="300px"
      :close-on-click-modal="false"
    >
      <el-form v-model="MyForm">
        <el-form-item label>
           <el-radio v-model="MyForm.lang" :label="1">中文</el-radio>
           <el-radio v-model="MyForm.lang" :label="2">英文</el-radio>
        </el-form-item>
        <el-form-item label>
          <el-form>
          <el-select v-model="MyForm.link" :placeholder="$t('please_choose')">
            <el-option
              v-for="item in itemList"
              :key="item.ID"
              :label="item.titlename"
              :value="item.ID"
            ></el-option>
          </el-select>
        </el-form>
        </el-form-item>
      </el-form>
      <p class="tips">{{$t('link_tips')}}</p>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogLinkVisible = false">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="linkItem">{{$t('link_item')}}</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'Login',
  components: {},
  data () {
    return {
      MyForm: {
      },
      dialogAttornVisible: false,
      dialogArchiveVisible: false,
      dialogDeleteVisible: false,
      dialogLinkVisible: false,
      itemList: [
      ],
      attornForm: {
        username: '',
        password: ''
      },
      archiveForm: {
        password: ''
      },
      deleteForm: {
        password: ''
      },
      linkForm: {
        lang: '2',
        link: ''
      }
    }
  },
  methods: {
    deleteItem () {
      var that = this
      var url = this.DocConfig.server + '/item/delete'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('password', this.deleteForm.password)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogDeleteVisible = false
            that.$message.success(that.$t('success_jump'))
            setTimeout(function () {
              that.$router.push({ path: '/item/index' })
            }, 2000)
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    archive () {
      var that = this
      var url = this.DocConfig.server + '/item/archive'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('password', this.archiveForm.password)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogArchiveVisible = false
            that.$message.success(that.$t('success_jump'))
            setTimeout(function () {
              that.$router.push({ path: '/item/index' })
            }, 2000)
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    getItemList () {
      var that = this
      var url = this.DocConfig.server + '/admin/list'
      var params = new URLSearchParams()
      that.$http.get(url, params).then(function (response) {
        if (response.data.status === 200) {
          var json = response.data.data
          that.itemList = json
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    attorn () {
      var that = this
      var url = this.DocConfig.server + '/item/attorn'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('username', this.attornForm.username)
      params.append('password', this.attornForm.password)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.dialogAttornVisible = false
            that.$message.success(that.$t('success_jump'))
            setTimeout(function () {
              that.$router.push({ path: '/item/index' })
            }, 2000)
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    linkItem () {
      var that = this
      var url = this.DocConfig.server + '/item/link'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('id', this.MyForm.link)
      params.append('lang', this.MyForm.lang)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          that.dialogLinkVisible = false
          that.$message.success(that.$t('success_jump'))
          setTimeout(function () {
            that.$router.push({ path: '/item/index' })
          }, 2000)
        } else {
          that.$alert(response.data.message)
        }
      })
    },
    get_item_info () {
      var that = this
      var url = this.DocConfig.server + '/item/detail'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            if (Info.password) {
              that.isOpenItem = false
            }
            that.MyForm = Info
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
    this.getItemList()
    this.get_item_info()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.a_button {
  width: 30%;
}
.a_button:first-child {
  margin-top: 30px;
}
</style>
