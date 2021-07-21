<template>
  <div class="hello">
    <Header></Header>

    <el-container>
      <el-card class="center-card">
        <el-form status-icon label-width="0px" class="demo-ruleForm">
          <h2></h2>
          <el-form-item label>
            <el-radio-group v-model="export_format">
              <el-radio-button label="word">{{$t('export_format_word')}}</el-radio-button>
              <el-radio-button label="markdown">{{$t('export_format_markdown')}}</el-radio-button>
            </el-radio-group>
          </el-form-item>

          <el-form-item label v-if="export_format == 'word'">
            <el-radio v-model="export_type" label="1">{{$t('export_all')}}</el-radio>
            <el-radio v-model="export_type" label="2">{{$t('export_cat')}}</el-radio>
          </el-form-item>

          <el-form-item label v-if="export_format == 'markdown'">
            <p class="markdown-tips">{{$t('export_markdown_tips')}}</p>
          </el-form-item>

          <el-form-item label v-if="export_format == 'word' && export_type == 2">
            <el-select
              :placeholder="$t('catalog')"
              class="cat"
              v-model="cat_id"
              v-if="computed_catalogs"
              @change="get_pages"
            >
              <el-option
                v-for="cat in computed_catalogs "
                :key="cat.cat_name"
                :label="cat.cat_name"
                :value="cat.cat_id"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label v-if="export_format == 'word' && export_type == 2">
            <el-select class="cat" v-model="page_id" v-if="pages">
              <el-option
                v-for="page in pages "
                :key="page.pagetitle"
                :label="page.pagetitle"
                :value="page.ID"
              ></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label>
            <el-button type="primary" style="width:100%;" @click="onSubmit">{{$t('begin_export')}}</el-button>
          </el-form-item>

          <el-form-item label>
            <el-button type="text" @click="goback" class="goback-btn">{{$t('goback')}}</el-button>
          </el-form-item>
        </el-form>
      </el-card>
    </el-container>

    <Footer></Footer>
  </div>
</template>

<script>
export default {
  name: 'Login',
  components: {},
  data () {
    return {
      catalogs: [],
      cat_id: '',
      export_type: '1',
      item_id: 0,
      export_format: 'word',
      pages: [{
        ID: '0',
        pagetitle: this.$t('all_pages')
      }],
      ID: '0'
    }
  },
  computed: {
    // 新建/编辑页面时供用户选择的归属目录列表
    computed_catalogs: function () {
      var Info = this.catalogs.slice(0)
      var cat_array = []
      for (var i = 0; i < Info.length; i++) {
        cat_array.push(Info[i])
        var sub = Info[i].sub

        if (sub !== null) {
          for (var j = 0; j < sub.length; j++) {
            cat_array.push({
              cat_id: sub[j].cat_id,
              cat_name: Info[i].cat_name + ' / ' + sub[j].cat_name
            })
            var sub_sub = sub[j].sub
            if (sub_sub !== null) {
              for (var k = 0; k < sub_sub.length; k++) {
                cat_array.push({
                  cat_id: sub_sub[k].cat_id,
                  cat_name:
                    Info[i].cat_name +
                    ' / ' +
                    sub[j].cat_name +
                    ' / ' +
                    sub_sub[k].cat_name
                })
              }
            }
          }
        }
      }
      var no_cat = { cat_id: '', cat_name: this.$t('none') }
      cat_array.unshift(no_cat)
      return cat_array
    }
  },
  methods: {
    // 获取所有目录
    get_catalog (item_id) {
      var that = this
      var url = DocConfig.server + '/catListGroup'
      var params = new URLSearchParams()
      params.append('item_id', item_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            that.catalogs = Info
          } else {
            that.$message(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    onSubmit () {
      if (this.export_type === 1) {
        this.cat_id = ''
      }
      var url =
        DocConfig.server +
        '/export/word?item_id=' +
        this.item_id +
        '&cat_id=' +
        this.cat_id + '&page_id=' + this.ID
      if (this.export_format === 'markdown') {
        url = DocConfig.server + '/export/markdown/' + this.item_id
      }
      window.location.href = url
    },
    goback () {
      this.$router.go(-1)
    },
    // 获取某目录下的所有页面
    get_pages (cat_id) {
      var that = this
      var url = DocConfig.server + '/catalog/getPagesBycat'
      var params = new URLSearchParams()
      params.append('item_id', this.item_id)
      params.append('cat_id', cat_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var pages = response.data.data
            pages.unshift({
              ID: '0',
              pagetitle: that.$t('all_pages')
            })
            that.pages = pages
            that.ID = '0'
          } else {
            that.$alert(response.data.message)
          }
        })
    }
  },
  mounted () {
    this.get_catalog(this.$route.params.item_id)
    this.item_id = this.$route.params.item_id
  },
  beforeDestroy () {}
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.center-card a {
  font-size: 12px;
}
.center-card {
  text-align: center;
  width: 400px;
}
.markdown-tips {
  text-align: left;
  margin-left: 25px;
  font-size: 12px;
}
</style>
