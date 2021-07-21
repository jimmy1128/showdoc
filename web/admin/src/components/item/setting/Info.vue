<template>
  <div class="hello">
    <el-form status-icon label-width="100px" class="infoForm" v-model="infoForm">
      <el-form-item>
        <el-input type="text" auto-complete="off" v-model="infoForm.titlename" placeholder></el-input>
      </el-form-item>

      <el-form-item>
        <el-input
          type="text"
          auto-complete="off"
          v-model="infoForm.description"
          :placeholder="$t('item_description')"
        ></el-input>
      </el-form-item>

      <el-form-item label>
        <el-radio v-model="isOpenItem" :label="true">{{$t('Open_item')}}</el-radio>
        <el-radio v-model="isOpenItem" :label="false">{{$t('private_item')}}</el-radio>
      </el-form-item>

      <el-form-item>
      <div>
        <el-checkbox-group v-model="dataIntArr" size="mini" @change="handleClick" :border="true">
        <el-checkbox-button v-for="langs in lang" :label="langs.id" :key="langs.id">{{langs.name}}</el-checkbox-button>
        </el-checkbox-group>
      </div>
      </el-form-item>

      <el-form-item v-show="!isOpenItem">
        <el-input
          type="password"
          auto-complete="off"
          v-model="infoForm.password"
          :placeholder="$t('visit_password')"
        ></el-input>
      </el-form-item>

      <el-form-item label>
        <el-button type="primary" style="width:100%;" @click="FormSubmit">{{$t('submit')}}</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'Login',
  components: {},
  data () {
    return {
      infoForm: {},
      isOpenItem: true,
      lang: [],
      checkboxGroup3:[],
      list:[],
      dataIntArr:[]
    }
  },
  methods: {
    get_item_info () {
      var that = this
      var url = DocConfig.server + '/item/detail'
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
            that.infoForm = Info
            var langInfo =that.infoForm.lang_list.split(',')
            langInfo.forEach(item => {
            that.dataIntArr.push(+item)
            })
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    handleClick(tab) {
      this.list = tab
    },
    FormSubmit () {
      var that = this
      var url = DocConfig.server + '/item/update'
      if (!this.isOpenItem && !this.infoForm.password) {
        that.$alert(that.$t('private_item_passwrod'))
        return false
      }
      if (this.isOpenItem) {
        this.infoForm.password = ''
      }
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('title', this.infoForm.titlename)
      params.append('description', this.infoForm.description)
      // params.append('item_domain', this.infoForm.item_domain)
      params.append('password', this.infoForm.password)
      params.append('langlist', this.list)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            that.$message.success(response.data.message)
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    get_lang () {
      var that = this
      var url = DocConfig.server + '/lang'
      var params = new URLSearchParams()
      that.$http.get(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var lang = response.data.data
            // 创建上级目录选项
            var Lang2 = lang.slice(0)
            that.lang = Lang2

          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
  },
  mounted () {
    this.get_item_info()
    this.get_lang()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.center-card a {
  font-size: 12px;
}
.center-card {
  text-align: center;
  width: 600px;
  height: 500px;
}
.infoForm {
  width: 350px;
  margin-left: 20px;
  margin-top: 60px;
}
.goback-btn {
  z-index: 999;
  margin-left: 500px;
}
</style>
