<template>
  <div class="hello">
    <Header> </Header>
    <el-container>
          <el-card class="center-card">
          <template>
            <el-button type="text" class="goback-btn " ><router-link to="/item/index">{{$t("goback")}}</router-link></el-button>
            <el-tabs  value="first" type="card">
              <el-tab-pane :label="$t('modify_password')" name="first">
                <el-form  status-icon  label-width="0px" class="passwordForm" v-model="passwordForm">
                  <el-form-item label="" >
                    <el-input type="text" auto-complete="off" v-model="passwordForm.email" placeholder="" :disabled="true"></el-input>
                  </el-form-item>
                  <el-form-item label="" >
                    <el-input type="password" auto-complete="off" :placeholder="$t('old_password')" v-model="passwordForm.password"></el-input>
                  </el-form-item>
                  <el-form-item label="" >
                    <el-input type="password" auto-complete="off" v-model="passwordForm.new_password" :placeholder="$t('new_password')"></el-input>
                  </el-form-item>
                   <el-form-item label="" >
                    <el-button type="primary" style="width:100%;" @click="passwordFormSubmit" >{{$t("submit")}}</el-button>
                  </el-form-item>
                </el-form>
            </el-tab-pane>
            </el-tabs>
          </template>
          </el-card>
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
      passwordForm: {
        email: '',
        password: '',
        new_password: ''
      },
      emailForm: {
        status: '',
        email: '',
        password: '',
        submitText: ''
      },
      userInfo: {
      }
    }
  },
  methods: {
    async getUserInfo () {
      var url = DocConfig.server + '/user/info'
      const { data: res } = await this.$http.post(url)
      if (res.status !== 200) {
        this.$router.push('/')
      }
      this.passwordForm.email = res.data.username
      this.emailForm.email = res.data.username
      var status = this.$t('status') + ':'
      if (res.data.username.length > 0) {
        this.emailForm.submitText = this.$t('modify')
        if (res.email_verify > 0) {
          status += this.$t('status_1')
        } else {
          status += this.$t('status_2')
        }
      } else {
        status += this.$t('status_3')
        this.emailForm.submitText = this.$t('binding')
      }
      this.emailForm.status = status
    },
    async passwordFormSubmit () {
      var url = DocConfig.server + '/admin/resetPassword'
      var formdata = new FormData()
      formdata.append('password', this.passwordForm.password)
      formdata.append('new_password', this.passwordForm.new_password)

      const { data: res } = await this.$http.post(url, formdata)
      if (res.status === 200) {
        this.$message.success('添加用户成功')
        this.$router.push('/item/index')
      } else {
        this.$message.error(res.message)
      }
    },
    emailFormSubmit () {
      var that = this
      var url = DocConfig.server + '/api/user/updateEmail'
      var params = new URLSearchParams()
      params.append('email', this.emailForm.email)
      params.append('password', this.emailForm.password)
      that.axios.post(url, params)
        .then(function (response) {
          if (response.data.status === 0) {
            that.$alert(that.$t('update_email_success'))
            this.getUserInfo()
          } else {
            that.$alert(response.data.error_message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  mounted () {
    this.getUserInfo()
    document.getElementsByTagName('body')[0].className = 'grey-bg'
  },
  beforeDestroy () {
    document.body.removeAttribute('class', 'grey-bg')
  }
}
</script>
<!-- Add 'scoped' attribute to limit CSS to this component only -->
<style scoped>
.center-card a {
  font-size: 12px;
}
.center-card{
  text-align: center;
  width: 600px;
  height: 500px;
}
.passwordForm,.emailForm{
  width:300px;
  margin: 0 auto ;
  margin-top: 50px;
}
.goback-btn{
  z-index: 999;
  margin-left: 500px;
}
</style>
