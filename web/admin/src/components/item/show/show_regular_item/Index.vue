<template>
  <div class="hello" v-if="showComp">
    <Header></Header>
    <div id="header" class="header">
      <div class="masthead" v-if="show_header">
        <div class="flex">
          <el-button type="text">
            <div class="logo-title ">
              <el-avatar size="large" :src="img" v-if="img"></el-avatar>
              <el-link :href="logourl"><h1>{{ title }}</h1></el-link>
            </div>
          </el-button>

      <el-row >
  <el-col :span="8" v-for="(item,index) in header_link" :key="index">
    <div class = header-titles>
      <el-button type="text" class="button" @click="jump_to_item(item.headerUrl)">{{item.headerName}}</el-button>
    </div>
  </el-col>
</el-row>
          <div class="header-language">
            <el-tooltip
              class="item"
              effect="dark"
              :content="$t('language1')"
              placement="top"
            >
              <el-dropdown
                trigger="click"
                class="el-select-lang"
                @command="LangChange"
                placement="top-start"
                v-if="selected != null"
              >
                <span class="el-dropdown-link">
                  <svg
                    class="icons"
                    aria-hidden="true"
                    style="float: left; color: #8492a6; font-size: 10px"
                    v-html="selected"
                  ></svg>
                </span>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item
                    v-for="itemlang in belong_to_lang"
                    :key="itemlang.id"
                    :label="itemlang.name"
                    :value="itemlang"
                    :command="itemlang"
                  >
                    {{ itemlang.name }}
                    <svg
                      class="icon"
                      aria-hidden="true"
                      style="float: left; color: #8492a6; font-size: 13px"
                      v-html="itemlang.icon"
                    ></svg>
                  </el-dropdown-item>
                </el-dropdown-menu>
              </el-dropdown>
            </el-tooltip>
          </div>
        </div>
      </div>
    </div>

    <div class="container doc-container" id="doc-container">
      <div id="left-side" style="width: 300px;">
        <LeftMenu
          ref="leftMenu"
          :get_page_content="get_page_content"
          :keyword="keyword"
          :item_info="item_info"
          :search_item="search_item"
          v-if="item_info"
        ></LeftMenu>
      </div>

      <div id="right-side">
        <div
          id="p-content"
          @mouseenter="showfullPageBtn = true"
          @mouseleave="showfullPageBtn = false"
        >
          <div class="doc-title-box" id="doc-title-box">
            <span id="doc-title-span" class="dn"></span>
            <h2 id="doc-title">{{ page_title }}</h2>
            <i
              class="el-icon-full-screen"
              id="full-page"
              v-show="showfullPageBtn"
              @click="clickFullPage"
            ></i>
            <i
              class="el-icon-upload item"
              id="attachment"
              v-if="attachment_count"
              @click="ShowAttachment"
            ></i>
          </div>
          <div id="doc-body">
            <div id="page_md_content" class="page_content_main">
              <Editormd
                v-bind:content="content"
                v-if="page_id"
                type="html"
                :keyword="keyword"
              ></Editormd>
            </div>
          </div>
        </div>

        <OpBar
          :page_id="page_id"
          :item_id="item_info.id"
          :item_info="item_info"
          :page_info="page_info"
          v-on:itemlangId="item_langId"
        ></OpBar>
      </div>
    </div>
    <BackToTop></BackToTop>
    <Toc v-if="page_id && showToc"></Toc>

    <!-- 附件列表 -->
    <AttachmentList
      callback
      :item_id="page_info.itemid"
      :manage="false"
      :page_id="page_info.ID"
      ref="AttachmentList"
    ></AttachmentList>

    <Footer></Footer>
  </div>
</template>

<script>
import Editormd from "@/components/common/Editormd";
import BackToTop from "@/components/common/BackToTop";
import Toc from "@/components/item/show/show_regular_item/Toc";
import LeftMenu from "@/components/item/show/show_regular_item/LeftMenu";
import OpBar from "@/components/item/show/show_regular_item/OpBar";
import AttachmentList from "@/components/page/edit/AttachmentList";
import { rederPageContent } from "@/models/page";

export default {
  props: {
    item_info: {},
    search_item: {},
    keyword: {}
  },
  data() {
    return {
      content: "###正在加载...",
      page_id: "",
      page_title: "",
      dialogVisible: false,
      share_item_link: "",
      qr_item_link: "",
      page_info: "",
      copyText: "",
      attachment_count: "",
      fullPage: false,
      showfullPageBtn: false,
      showToc: true,
      showComp: true,
      itemlangId: "",
      selected: "",
      itemlang: [],
      langs: [],
      icon: "",
      value: {},
      lang_array: [],
      show_header: true,
      infoProfile:[],
      img:'',
      title:'',
      header_link:[],
      logourl: '/#/'+this.$route.params.item_id
    };
  },
  components: {
    Editormd,
    LeftMenu,
    OpBar,
    BackToTop,
    Toc,
    AttachmentList
  },
  computed: {
    belong_to_lang: function() {
      var that = this;
      if (that.item_info.lang_list === undefined) {
        return [];
      }
      var langInfo = that.item_info.lang_list.split(",");
      var lang_array = [];

      if (that.langs.length > 0) {
        for (var k = 0; k < langInfo.length; k++) {
          if (langInfo[k] === "") {
            return that.langs;
          }
          for (var j = 0; j < that.langs.length; j++) {
            if (that.langs[j].id === parseInt(langInfo[k])) {
              lang_array.push({
                id: that.langs[j].id,
                name: that.langs[j].name,
                icon: that.langs[j].icon
              });
            }
          }
        }
      }
      return lang_array;
    }
  },
  methods: {
    item_langId(value) {
      this.itemlangId = value;
      this.$emit("itemlangId2", this.itemlangId);
      this.item_info.defaultpageid = 0;
    },
    // 获取页面内容
    get_page_content(page_id) {
      if (page_id <= 0) {
        return;
      }
      this.adaptScreen();

      var that = this;
      var url = DocConfig.server + "/admin/page";
      var params = new URLSearchParams();
      params.append("page_id", page_id);
      that.$http.post(url, params).then(function(response) {
        // loading.close()
        if (response.data.status === 200) {
          that.content = rederPageContent(
            response.data.data.pagecontent,
            that.$store.state.item_info.global_param
          );
          that.$store.dispatch("changeOpenCatId", response.data.data.catid);
          that.page_title = response.data.data.pagetitle;
          that.page_info = response.data.data;
          that.attachment_count =
            response.data.attachment_count > 0
              ? response.data.attachment_count
              : "";
          // 切换变量让它重新加载、渲染子组件
          that.page_id = 0;
          that.item_info.defaultpageid = page_id;
          that.$nextTick(() => {
            that.page_id = page_id;
            // 页面回到顶部
            document.body.scrollTop = document.documentElement.scrollTop = 0;
            //document.title = that.page_title +'--'+that.item_info.itemtitle+ '--Doc'
          });
        } else {
          that.$alert(response.data.message);
        }
      });
    },
    // 根据屏幕宽度进行响应(应对移动设备的访问)
    adaptToMobile() {
      const childRef = this.$refs.leftMenu; // 获取子组件
      childRef.hide_menu();
      this.show_page_bar = false;
      var doc_container = document.getElementById("doc-container");
      doc_container.style.width = "95%";
      doc_container.style.padding = "5px";
      var header = document.getElementById("header");
      header.style.height = "10px";
      var docTitle = document.getElementById("doc-title-box");
      docTitle.style.marginTop = "40px";
      this.showToc = false;
    },
    // 根据屏幕宽度进行响应。应对小屏幕pc设备(如笔记本)的访问
    adaptToSmallpc() {
      var doc_container = document.getElementById("doc-container");
      doc_container.style.width = "calc( 95% - 300px )";
      doc_container.style.marginLeft = "300px";
      doc_container.style.padding = "20px";
      var header = document.getElementById("header");
      header.style.height = "20px";
      var docTitle = document.getElementById("doc-title-box");
      docTitle.style.marginTop = "30px";
    },
    // 响应式
    adaptScreen() {
      this.$nextTick(() => {
        // 根据屏幕宽度进行响应(应对移动设备的访问)
        if (this.isMobile() || window.innerWidth < 1300) {
          if (window.innerWidth < 1300 && window.innerWidth > 1100) {
            this.adaptToSmallpc();
          } else {
            this.adaptToMobile();
          }
        }
      });
    },
    onCopy() {
      this.$message(this.$t("copy_success"));
    },
    ShowAttachment() {
      const childRef = this.$refs.AttachmentList; // 获取子组件
      childRef.show();
    },
    clickFullPage() {
      var $ = require("jquery");
      // 点击放大页面。由于历史包袱，只能操作dom。这是不规范的，但是现在没时间重构整块页面
      if (this.fullPage) {
        // 通过v-if指令起到刷新组件的作用
        this.showComp = false;
        this.$nextTick(() => {
          this.showComp = true;
        });
      } else {
        this.adaptToMobile();
        var page_id = this.page_id;
        this.page_id = 0;
        this.$nextTick(() => {
          this.page_id = page_id;
          setTimeout(() => {
            $(".editormd-html-preview").css("font-size", "16px");
          }, 200);
        });
        $("#left-side").hide();
        $(".op-bar").hide();
      }
      this.fullPage = !this.fullPage;
    },
    jump_to_item (urls) {
         const url = 'http://' + urls
         window.open(url)
    },
    LangChange(value) {
      this.item_langId(value.id);
      this.selected = value;
      this.get_item_header(value.id)
      // if (this.$cookies.get('selected') === null) {
      this.$cookies.set("selected", value.icon);
      // }
      if (
        value.name === "English" ||
        value.icon === "#icon-world-flag_-GBR--UnitedKingdom"
      ) {
        this.locale = "EN_US";
        this.lang = true;
        // this.$cookies.set('lng', this.locale === 'EN_US' ? this.locale : this.locale, '1d')
      } else {
        this.lang = false;
        this.locale = "ZH_CN";
        // this.$cookies.set('lng', this.locale === 'ZH_CN' ? this.locale : this.locale, '30d')
      }
    },
    get_lang() {
      var that = this;
      var url = DocConfig.server + "/lang";
      var params = new URLSearchParams();
      that.$http
        .get(url, params)
        .then(function(response) {
          if (response.data.status === 200) {
            that.langs = response.data.data;
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    loadConfig() {
      var item_id = this.$route.params.item_id ? this.$route.params.item_id : 0;
      var that = this;
      var url = DocConfig.server + "/adminSetting/loadLangConfig";
      var params = new URLSearchParams();
      params.append("id", item_id);
      that.$http.post(url, params).then(function(response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return;
          }
          that.selected =
            '<use xlink:href="' + response.data.data.icon + '"></use>';
          that.$cookies.set("selected", that.selected);
          that.value = response.data.data;
          that.value.icon =
            '<use xlink:href="' + response.data.data.icon + '"></use>';

          if (response.data.access === 1) {
            // defaultlang 是全局语言
            localStorage.setItem("defaultlang", that.value.id);
            that.LangChange(that.value);
          }
        } else {
          that.$alert(response.data.message);
        }
      });
    },
    publicloadConfig(value) {
      var that = this;
      var url = DocConfig.server + "/public/lang";
      var params = new URLSearchParams();
      params.append("lang", value);
      that.$http.post(url, params).then(function(response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return;
          }
          that.value = response.data.data;
          that.value.icon =
            '<use xlink:href="' + response.data.data.icon + '"></use>';
          if (response.data.access === 1) {
            that.LangChange(that.value);
            this.selected =
              '<use xlink:href="' + response.data.data.icon + '"></use>';
            this.$cookies.set("selected", that.selected);
          }
        } else {
          this.$alert(response.data.message);
        }
      });
    },
    get_lang1(value) {
      var that = this;
      var url = DocConfig.server + "/lang/info";
      var params = new URLSearchParams();
      params.append("icon", value);
      that.$http.post(url, params).then(function(response) {
        if (response.data.status === 200) {
          if (response.data.data.id === 0) {
            return;
          }
          that.selected =
            '<use xlink:href="' + response.data.data.icon + '"></use>';
          that.$cookies.set("selected", that.selected);
          that.value = response.data.data;
          that.value.icon =
            '<use xlink:href="' + response.data.data.icon + '"></use>';
          that.get_item_header(that.value.id)
        } else {
          that.$alert(response.data.message);
        }
      });
    },
    get_item_info() {
      var that = this;
      var url = DocConfig.server + "/item/detail";
      var params = new URLSearchParams();
      params.append("item_id", that.$route.params.item_id);
      that.$http
        .post(url, params)
        .then(function(response) {
          if (response.data.status === 200) {
            var Info = response.data.data;
            if (Info.password) {
              that.isOpenItem = false;
            }
            if (Info.lang === 1) {
              Info.lang = false;
            } else if (Info.lang === 2) {
              Info.lang = true;
            }
            that.MyForm = Info;
            that.link = "/" + that.MyForm.link;
            that.lang = that.MyForm.lang;
          } else {
            that.$alert(response.data.message);
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    get_item_avatar() {
      var that = this;
      var url = DocConfig.server + "/avatar/profile"
      var params = new URLSearchParams()
      params.append("item_id", that.$route.params.item_id)
      that.$http
        .post(url, params)
        .then(function(response) {
          if (response.data.status === 200) {
            var Info = response.data.data
            that.infoProfile = Info
            that.img = that.infoProfile.avatarUrl
            that.title = that.infoProfile.avatarName
          }
        })
        .catch(function(error) {
          console.log(error);
        });
    },
    get_item_header (value) {
      var that = this
      var url = DocConfig.server + '/avatar/getHeader'
      var params = new URLSearchParams()
      params.append('item_id', that.$route.params.item_id)
      params.append('lang_id', value)
      that.$http
        .post(url, params)
        .then(function (response) {
          if (response.data.status === 200) {
           var Info1 = response.data.data
           console.log(Info1)
           that.header_link = Info1
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
    var that = this;
    var page_id = this.$route.query.page_id ? this.$route.query.page_id : 0;
    this.adaptScreen();
    this.set_bg_grey();
    this.get_item_info();
    this.get_lang();
    this.get_item_avatar()
    //this.get_item_header()
    if (this.$cookies.get("lng") === null) {
      this.locale = DocConfig.lang;
      this.$cookies.set(
        "lng",
        this.locale === "EN_US" ? this.locale : this.locale,
        "1d"
      );
    }
    if (this.$cookies.get("selected") === null) {
      this.loadConfig();
    }
    if (this.$cookies.get("selected") === null) {
      this.publicloadConfig(page_id);
    }

    this.selected = this.$cookies.get("selected");
    if (this.selected !== null) {
      this.get_lang1(this.selected);
      this.$cookies.remove("selected");
    }
    if (
      this.isMobile() ||
      (window.innerWidth < 1300 && !this.item_info.is_login)
    ) {
      this.show_header = false;
    }
    if (
      this.isMobile() ||
      (window.innerWidth < 1300 && this.item_info.is_login)
    ) {
      this.show_header = false;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.flex {
  display: flex;
}
.page_content_main {
  width: 800px;
  margin: 0 auto;
  height: 50%;
  overflow: visible;
}
.editormd-html-preview {
  width: 95%;
  font-size: 16px;
  overflow-y: hidden;
}
#attachment {
  float: right;
  font-size: 25px;
  margin-top: -40px;
  margin-right: 5px;
  cursor: pointer;
  color: #abd1f1;
}
#full-page {
  float: right;
  font-size: 25px;
  margin-top: -50px;
  margin-right: 30px;
  cursor: pointer;
  color: #ccc;
}
#page_md_content {
  padding: 10px 10px 90px 10px;
  overflow: hidden;
  font-size: 11pt;
  line-height: 1.7;
  color: #333;
}
.doc-container {
  position: static;
  background-color: #fff;
  margin-bottom: 20px;
  width: 800px;
  min-height: 750px;
  margin-left: auto;
  margin-right: auto;
  padding: 20px;
}
#header {
  height: 80px;
}
#doc-body {
  width: 95%;
  margin: 0 auto;
  background-color: #fff;
}
.doc-title-box {
  height: auto;
  width: auto;
  border-bottom: 1px solid #ebebeb;
  padding-bottom: 10px;
  width: 100%;
  margin: 10px auto;
  text-align: center;
}
pre ol {
  list-style: none;
}
.markdown-body pre {
  background-color: #f7f7f9;
  border: 1px solid #e1e1e8;
}
.hljs {
  background-color: #f7f7f9;
}
.tool-bar {
  margin-top: -38px;
}
.editormd-html-preview,
.editormd-preview-container {
  padding: 0px;
  font-size: 14px;
}
.text {
  font-size: 14px;
}

.item {
  margin-bottom: 18px;
}

.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}

.box-card {
  width: 295px;
  position: absolute;
  bottom: 0;
}
.empty-tips {
  margin: 5% auto;
  width: 400px;
  text-align: center;
  color: #909399;
}
.empty-tips .icon {
  font-size: 100px;
  margin-left: -50px;
}
.empty-tips .text {
  text-align: left;
}
.empty-tips .links {
  line-height: 2em;
}
.empty-tips .links a {
  color: #909399;
  text-decoration: underline;
}
.masthead {
  width: 1500px;
  margin-top: 10px;
  background: rgb(255, 255, 255);
  color: #fff;
  margin-left: auto;
  margin-right: auto;
}
.logo-title {
  background-color: rgb(255, 255, 255);
  width: 330px;
  border-block-color: black;
  height: 40px;
  border-right-style: solid;
  border-right-color: rgb(212, 218, 223);
  border-right-width: 1px;
  line-height: 40px;
  display: flex;
  justify-content: center;
}
.container-narrow {
  height: 80px;
  margin: auto;
  color: #fff;
}
.header-title {
  background-color: rgb(255, 255, 255);
  font-size: 15px;
  border-right-style: solid;
  border-right-color: rgb(212, 218, 223);
  border-right-width: 1px;
  height: 40px;
  line-height: 40px;
  margin-top: 10px;
  margin-left: 30px;
}
#header-right-btn {
  font-size: 20px;
  top: 15px;
  right: 5%;
  cursor: pointer;
  position: fixed;
}
.el-dropdown-link {
  color: #000;
  font-size: 18px;
  font-weight: bolder;
}
.el-icon-document-copy {
  cursor: pointer;
}
.icon {
  width: 1.9em;
  height: 1.9em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
  padding-top: 5px;
}
.el-select-lang {
  border: none;
  border-radius: 0px;
}
.icons {
  width: 5em;
  height: 5em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
.el-row {
  margin-bottom: 0px;
  display:flex;
  width: 820px;
  border-right-color: rgb(212, 218, 223);
  margin-left: 20px;
}
.el-col {
  border-radius: 0px;
  height: 70px;
  display:flex;
  flex-wrap: wrap;
  align-items: center;
  width: 60px;
  margin-left: 25px;

}
.header-language {
  height: 40px;
  line-height: 40px;
  text-align: center;
  margin-left: 50px;
  margin-top: 10px;
}
.header {
  width: 100%;
  color: #fff;
  border-bottom-color: rgb(212, 218, 223);
  border-bottom-style: solid;
  border-bottom-width: 1px;
  box-shadow:0 3px 8px 0 rgb(116 129 141 / 10%)
}
.clearfix:before,
    .clearfix:after {
        display: table;
        content: "";
  }

.clearfix:after {
        clear: both
}

.header-titles{
border: 0px;
padding: 0px;
text-align: center;
justify-content: center;
height: 30px;
font-size: 20px;

}
.button{
font-size: 20px;
}
</style>
