<template>
  <div class="hello">
    <el-form status-icon label-width="10px" class="infoForm" v-model="infoForm">
      <el-form-item>
        <el-radio-group v-model="infoForm.item_type">
          <el-radio label="1">{{$t('regular_item')}}</el-radio>
          <el-radio label="4">{{$t('table')}}</el-radio>
          <el-radio label="2">
            {{$t('single_item')}}
            <el-tooltip
              class="item"
              effect="dark"
              :content="$t('single_item_tips')"
              placement="top"
            >
              <i class="el-icon-question"></i>
            </el-tooltip>
          </el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-tooltip class="item" effect="dark" :content="$t('item_name')" placement="right">
          <el-input
            type="text"
            auto-complete="off"
            v-model="infoForm.item_name"
            :placeholder="$t('item_name')"
          ></el-input>
        </el-tooltip>
      </el-form-item>

      <el-form-item >
        <el-tooltip class="item" effect="dark" :content="$t('item_description')" placement="right">
          <el-input
            type="text"
            auto-complete="off"
            v-model="infoForm.item_description"
            :placeholder="$t('item_description')"
          ></el-input>
        </el-tooltip>
      </el-form-item>

      <el-form-item >
      <el-tooltip class="item" effect="dark" :content="$t('item_language')" placement="right">
      <div >
        <el-checkbox-group v-model="infoForm.lang_list" size="mini" @change="handleClick" :border="true">
        <el-checkbox-button v-for="langs in lang" :label="langs.id" :key="langs.id">{{langs.name}}</el-checkbox-button>
        </el-checkbox-group>
      </div>
      </el-tooltip>
      </el-form-item>

      <el-form-item label>
        <el-radio v-model="isOpenItem" :label="true">{{$t('Open_item')}}</el-radio>
        <el-radio v-model="isOpenItem" :label="false">{{$t('private_item')}}</el-radio>
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
      infoForm: {
        item_name: '',
        item_description: '',
        item_domain: '',
        password: '',
        item_type: '1',
        lang_list: [1]
      },
      isOpenItem: true,
      lang: [],
      list:[]
    }
  },
  methods: {
    FormSubmit () {
      var that = this
      var url = DocConfig.server + '/item/add'
      if (!this.isOpenItem && !this.infoForm.password) {
        that.$alert(that.$t('private_item_passwrod'))
        return false
      }
      if (this.isOpenItem) {
        this.infoForm.password = ''
      }
      var params = new URLSearchParams()
      params.append('types', this.infoForm.item_type)
      params.append('titlename', this.infoForm.item_name)
      params.append('description', this.infoForm.item_description)
      params.append('item_domain', this.infoForm.item_domain)
      params.append('password', this.infoForm.password)
      params.append('langlist', this.infoForm.lang_list)
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          that.$router.push({ path: '/item/index' })
        } else {
          that.$alert(response.data.message)
        }
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
    handleClick(tab) {
      this.infoForm.lang_list = tab
      console.log(tab)
    },
  },
  mounted () {
    this.get_lang()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.center-card a {
  font-size: 12px;
}
.infoForm {
  width: 380px;
  margin-top: 30px;
}
</style>
