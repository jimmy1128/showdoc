<template>
  <div class="hello">
    <Header> </Header>
    <el-container>
          <el-card class="center-card">
            <el-form  status-icon  label-width="0px" class="demo-ruleForm">
              <h2>{{$t("login")}}</h2>
              <el-form-item label="" >
                <el-input type="text" auto-complete="off" :placeholder="$t('username_description')" v-model="username"></el-input>
              </el-form-item>
              <el-form-item label="" >
                <el-input type="password" auto-complete="off" v-model="password" :placeholder="$t('password')"></el-input>
              </el-form-item>
               <el-form-item label="" >
                <el-button type="primary" style="width:100%;" @click="onSubmit" >{{$t("login")}}</el-button>
              </el-form-item>
              <el-form-item label="" >
                  <router-link to="/user/register">{{$t("register_new_account")}}</router-link>
                  &nbsp;&nbsp;&nbsp;
              </el-form-item>
            </el-form>
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
      username: '',
      password: ''
    }
  },
  methods: {
    async onSubmit () {
      // this.$message.success(this.username);
      var url = this.DocConfig.server + '/login'
      const { data: res } = await this.$http.post(url, {
        username: this.username,
        password: this.password
      })
      if (res.status !== 200) return this.$message.error(res.message)
      window.sessionStorage.setItem('token', res.token)
      this.$router.push('/item/index')
    }
  },
  mounted () {
    // 给body添加类，设置背景色
    document.getElementsByTagName('body')[0].className = 'grey-bg'
  },
  beforeDestroy () {
    // 去掉添加的背景色
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
}
</style>
