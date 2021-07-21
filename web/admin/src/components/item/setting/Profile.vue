<template>
  <div class="hello">
    <el-upload
      class="avatar-uploader"
      :action="upUrl"
      :show-file-list="false"
      :on-success="handleAvatarSuccess"
      :before-upload="beforeAvatarUpload"
    >
      <img v-if="infoProfile.avatarUrl" :src="infoProfile.avatarUrl" class="avatar" />
      <i v-else class="el-icon-plus avatar-uploader-icon"></i>
    </el-upload>
    <el-form
      status-icon
      label-width="100px"
      class="infoForm"
      v-model="infoProfile"
    >
      <el-form-item>
        <el-input
          type="text"
          auto-complete="off"
          v-model="infoProfile.avatarName"
          placeholder
        ></el-input>
      </el-form-item>

      <el-form-item>
        <el-button
          type="info"
          style="width:100%;"
          @click="addlistVisible = true"
          >{{ $t("setting") }}</el-button
        >
      </el-form-item>

      <el-form-item label>
        <el-button type="primary" style="width:100%;" @click="FormSubmit">{{
          $t("submit")
        }}</el-button>
      </el-form-item>
    </el-form>
    <el-dialog
      :title="$t('add_header')"
      :modal-append-to-body = false
      :visible.sync="addlistVisible"
      width="520px"
      :before-close="beforeClose"
    >
      <div >
        <el-form :model="infoHeader">
          <div>
            <el-form-item :label-width="formLabelWidth">
              <el-input
                style="width:100px"
                :placeholder="$t('page_menu')"
                autocomplete="off"
                @input="change($event)"
                v-model="infoHeader.headerName"
              ></el-input>
              <el-input
                style="width:180px"
                :placeholder="$t('dest_addr')"
                autocomplete="off"
                @input="change($event)"
                v-model="infoHeader.headerUrl"
              ></el-input>
               <el-select v-model="infoHeader.cid" :placeholder="$t('language1')" style="width:100px" @input="change($event)">
                  <el-option
                    v-for="item in belong_to_lang"
                    :key="item.id"
                    style="width:100px"
                    :label="item.name"
                    :value="item.id">
                  </el-option>
                </el-select>
          <el-button @click="jump_to_item(infoHeader.headerUrl)" type="text" size="small">{{$t('link')}}</el-button>
        <el-button @click="delete_header(infoHeader.headerId)" type="text" size="small">{{$t('delete')}}</el-button>
            </el-form-item>
          </div>
          <div v-for="(ddd,index) in counter" :key="index">
                            <el-form-item :label-width="formLabelWidth">
                                <!--这里v-model绑定要 ddd.zyname   才能确保动态添加的所绑定的数值不一样-->
                                <el-input style="width:100px" v-model="ddd.headerName" :placeholder="$t('page_menu')" autocomplete="off"></el-input>
                                <el-input style="width:180px" v-model="ddd.headerUrl" :placeholder="$t('dest_addr')" autocomplete="off"></el-input>
                                <el-select v-model="ddd.cid" :placeholder="$t('language1')" style="width:100px">


                  <el-option
                    v-for="item in belong_to_lang"
                    :key="item.id"
                    style="width:100px"
                    :label="item.name"
                    :value="item.id">
                  </el-option>
                </el-select>
                <el-button @click="jump_to_item(ddd.headerUrl)" type="text" size="small">{{$t('link')}}</el-button>
                <el-button @click="remove(index)" type="text" size="small">{{$t('delete')}}</el-button>
                            </el-form-item>
            </div>

          <div class="ujm">
            <el-button class="addstyle" type="info" @click="addinput"
              >{{$t('add_link')}}</el-button
            >
          </div>
        </el-form>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="goback">{{$t('cancel')}}</el-button>
        <el-button type="primary" @click="determineadd">{{$t('confirm')}}</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: "Login",
  components: {},
  data() {
    return {
      infoProfile:[],
      infoHeader:{},
      infoForm: {},
      checkboxGroup3: [],
      list: [],
      dataIntArr: [],
      imageUrl: "",
      counter: [],
      addlistVisible: false,
      handleClose2: "",
      formLabelWidth: "0px",
      upUrl: DocConfig.server + "/uploadImg",
      lang: [],
    };
  },
  computed:{
    belong_to_lang: function () {
      var that = this
      if (that.infoForm.lang_list === undefined){
              return []
            }
            var langInfo =that.infoForm.lang_list.split(',')
            var lang_array = []
            if (that.lang.length > 0){
              for (var k=0; k < langInfo.length; k++){
                if (langInfo[k] === ""){
                  return that.lang
                }
              for (var j=0; j < that.lang.length; j++){
                if (that.lang[j].id === parseInt(langInfo[k])){
                  lang_array.push({
                    id: that.lang[j].id,
                    name: that.lang[j].name,
                    icon: that.lang[j].icon
                  })
                }
              }
              }
            }
            return lang_array
    }
  },
  methods: {
    get_item_avatar() {
      var that = this;
      var url = DocConfig.server + "/avatar/profile";
      var params = new URLSearchParams();
      params.append("item_id", that.$route.params.item_id);
      that.$http
        .post(url, params)
        .then(function(response) {
          if (response.data.status === 200) {
            var Info = response.data.data;

            that.infoProfile = Info;
          } else {
            that.$alert(response.data.message);
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    handleClick(tab) {
      this.list = tab;
    },
    FormSubmit() {
      var that = this;
      var url = DocConfig.server + "/avatar/update";
      var params = new URLSearchParams();
      params.append("item_id", that.$route.params.item_id);
      params.append("avatar_name", this.infoProfile.avatarName);
      params.append("uri", this.infoProfile.avatarUrl);
      that.$http
        .post(url, params)
        .then(function(response) {
          if (response.data.status === 200) {
            that.$message.success(response.data.message);
          } else {
            that.$alert(response.data.message);
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    handleRemove(file) {
      console.log(file);
    },
    handlePictureCardPreview(file) {
      this.dialogImageUrl = file.url;
      this.dialogVisible = true;
    },
    handleDownload(file) {
      console.log(file);
    },
    addinput() {
      this.counter.push({ headerName:'', headerUrl:'' ,cid:'',headerId:0});
    },
    remove(index) {
      //splice 操作数组的方法
      this.counter.splice(index, 1)
    },
    change (e) {
      this.$forceUpdate()
    },
    delete_header(headerId) {
      //splice 操作数组的方法
      var that = this
      var url = DocConfig.server + '/avatar/deleteHeader'
      var params = new URLSearchParams()
      if (headerId == undefined) {
          that.infoHeader.headerName=''
          that.infoHeader.headerUrl='' ,
          that.infoHeader.cid='',
          that.infoHeader.headerId=0
      }
      params.append('header_id', headerId)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            this.get_item_header()
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })

    },
    determineadd() {
      var that = this
      var ui = {
        itemId : that.$route.params.item_id,
        data : [{
                  headerName : that.infoHeader.headerName,
                  headerUrl : that.infoHeader.headerUrl,
                  cid : that.infoHeader.cid,
                  headerId : that.infoHeader.headerId
                }]
      }
      var url = DocConfig.server + '/avatar/header'
      var params = new URLSearchParams()
      console.log( ui.data.push.apply(ui.data, that.counter) )
      params.append('item_id',ui.itemId)
      params.append('data',JSON.stringify(ui.data))
      that.$http.post(url, params).then(function (response) {
        if (response.data.status === 200) {
          var json = response.data.data
        } else {
          that.$alert(response.data.message)
        }
      })
      that.addlistVisible = false;
    },
    jump_to_item (urls) {
         const url = 'http://' + urls
         window.open(url)
    },
    handleAvatarSuccess(res, file) {
      console.log(res.url);
      this.infoProfile.avatarUrl = res.url;
    },
    beforeAvatarUpload(file) {
      const isJPG = file.type === "image/jpeg";
      const isLt2M = file.size / 1024 / 1024 < 2;
      const isPNG = file.type === "image/png";

      if (!isJPG) {
        this.$message.error("上传头像图片只能是 JPG 格式!");
      }
      if (!isLt2M) {
        this.$message.error("上传头像图片大小不能超过 2MB!");
      }
      return isJPG && isLt2M;
    },
    beforeClose(done) {
      this.dialogVisible = false;
      done();
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
            var no_cat = { ID: 0, name: that.$t('none') }
            Lang2.unshift(no_cat)
            that.lang = Lang2
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    get_item_info () {
      var that = this
      var url = DocConfig.server + '/item/detail'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
            var Info1 = response.data.data
            that.infoForm = Info1
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    goback () {
       var that = this
        that.addlistVisible= false
        that.get_item_header()
    },
    get_item_header () {
      var that = this
      var url = DocConfig.server + '/avatar/getHeader'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('lang_id', 1)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
           var Info = response.data.data
           that.counter = response.data.data
           that.infoHeader.headerName = that.counter[0].headerName
           that.infoHeader.headerUrl = that.counter[0].headerUrl
           that.infoHeader.cid = that.counter[0].cid
           that.infoHeader.headerId = that.counter[0].headerId
           that.counter.shift()
          } else {
            that.$alert(response.data.message)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
  },
  mounted() {
    this.get_item_avatar()
    this.get_lang()
    this.get_item_info()
    this.get_item_header()
  }
};
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
.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}
.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  line-height: 178px;
  text-align: center;
}
.avatar {
  width: 178px;
  height: 178px;
  display: block;
}
</style>
