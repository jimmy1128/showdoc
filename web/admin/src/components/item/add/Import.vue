<template>
  <div class="hello">
    <p class="tips">
      <span v-html="$t('import_file_tips1')"></span>
    </p>

    <p>
      <el-upload
        class="upload-demo"
        drag
        name="file"
        :action="upload_url"
        :on-success="success"
        :before-upload="beforeUpload"
        :show-file-list="false"
        :with-credentials="true"
        :headers="myHeaders"
      >
        <i class="el-icon-upload"></i>
        <div class="el-upload__text">
          <span v-html="$t('import_file_tips2')"></span>
        </div>
      </el-upload>
    </p>
    <p></p>
    <p></p>
  </div>
</template>

<script>
var token =  `Bearer ${window.sessionStorage.getItem('token')}`
export default {
  name: 'Login',
  components: {},
  data () {
    return {
      api_key: '',
      api_token: '',
      loading: '',
      upload_url: DocConfig.server + '/import/auto',
      myHeaders: {Authorization: token}
    }
  },
  methods: {
    success (data) {
      this.loading.close()
      if (data.status === 200) {
        this.$router.push({ path: '/item/index' })
      } else {
        this.$alert(data.message)
      }
    },
    beforeUpload () {
      this.loading = this.$loading()
    }
  },
  mounted () {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.hello {
  text-align: left;
  margin-left: 10px;
  margin-top: 30px;
}
.tips {
  margin-left: 5px;
  margin-bottom: 20px;
  color: #9ea1a6;
  padding: 10px;
}
</style>
