webpackJsonp([1,7,23,24],{TdIe:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=n("cdA+"),a=n("0xDb"),i={data:function(){return{updatePassowrdVisible:!1}},components:{UpdatePassword:s.default},computed:{navbarLayoutType:{get:function(){return this.$store.state.common.navbarLayoutType}},sidebarFold:{get:function(){return this.$store.state.common.sidebarFold},set:function(t){this.$store.commit("common/updateSidebarFold",t)}},mainTabs:{get:function(){return this.$store.state.common.mainTabs},set:function(t){this.$store.commit("common/updateMainTabs",t)}},userName:{get:function(){return this.$store.state.user.name}}},methods:{updatePasswordHandle:function(){var t=this;this.updatePassowrdVisible=!0,this.$nextTick(function(){t.$refs.updatePassowrd.init()})},logoutHandle:function(){var t=this;this.$confirm("确定进行[退出]操作?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then(function(){t.$http({url:t.$http.adornUrl("/sys/logout"),method:"post",data:t.$http.adornData()}).then(function(e){var n=e.data;n&&0===n.code&&(Object(a.a)(),t.$router.push({name:"login"}))})}).catch(function(){})}}},o={render:function(){var t=this,e=t.$createElement,s=t._self._c||e;return s("nav",{staticClass:"site-navbar",class:"site-navbar--"+t.navbarLayoutType},[s("div",{staticClass:"site-navbar__header"},[s("h1",{staticClass:"site-navbar__brand",on:{click:function(e){t.$router.push({name:"home"})}}},[t._m(0),t._v(" "),t._m(1)])]),t._v(" "),s("div",{staticClass:"site-navbar__body clearfix"},[s("el-menu",{staticClass:"site-navbar__menu",attrs:{mode:"horizontal"}},[s("el-menu-item",{staticClass:"site-navbar__switch",attrs:{index:"0"},on:{click:function(e){t.sidebarFold=!t.sidebarFold}}},[s("icon-svg",{attrs:{name:"zhedie"}})],1)],1),t._v(" "),s("el-menu",{staticClass:"site-navbar__menu site-navbar__menu--right",attrs:{mode:"horizontal"}},[s("el-menu-item",{attrs:{index:"1"},on:{click:function(e){t.$router.push({name:"theme"})}}},[s("template",{attrs:{slot:"title"},slot:"title"},[s("el-badge",{attrs:{value:"new"}},[s("icon-svg",{staticClass:"el-icon-setting",attrs:{name:"shezhi"}})],1)],1)],2),t._v(" "),s("el-menu-item",{staticClass:"site-navbar__avatar",attrs:{index:"3"}},[s("el-dropdown",{attrs:{"show-timeout":0,placement:"bottom"}},[s("span",{staticClass:"el-dropdown-link"},[s("img",{attrs:{src:n("zQrT"),alt:t.userName}}),t._v(t._s(t.userName)+"\n          ")]),t._v(" "),s("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[s("el-dropdown-item",{nativeOn:{click:function(e){t.updatePasswordHandle()}}},[t._v("修改密码")]),t._v(" "),s("el-dropdown-item",{nativeOn:{click:function(e){t.logoutHandle()}}},[t._v("退出")])],1)],1)],1)],1)],1),t._v(" "),t.updatePassowrdVisible?s("update-password",{ref:"updatePassowrd"}):t._e()],1)},staticRenderFns:[function(){var t=this.$createElement,e=this._self._c||t;return e("a",{staticClass:"site-navbar__brand-lg",attrs:{href:"javascript:;"}},[e("img",{staticStyle:{"margin-right":"6px"},attrs:{src:n("dLd/")}}),this._v("臻方便后台管理系统")])},function(){var t=this.$createElement,e=this._self._c||t;return e("a",{staticClass:"site-navbar__brand-mini",attrs:{href:"javascript:;"}},[e("img",{attrs:{src:n("dLd/")}})])}]},r=n("46Yf")(i,o,!1,null,null,null);e.default=r.exports},YbVU:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=n("oZaA"),a=n("E4LH"),i={data:function(){return{dynamicMenuRoutes:[]}},components:{SubMenu:s.default},computed:{sidebarLayoutSkin:{get:function(){return this.$store.state.common.sidebarLayoutSkin}},sidebarFold:{get:function(){return this.$store.state.common.sidebarFold}},menuList:{get:function(){return this.$store.state.common.menuList},set:function(t){this.$store.commit("common/updateMenuList",t)}},menuActiveName:{get:function(){return this.$store.state.common.menuActiveName},set:function(t){this.$store.commit("common/updateMenuActiveName",t)}},mainTabs:{get:function(){return this.$store.state.common.mainTabs},set:function(t){this.$store.commit("common/updateMainTabs",t)}},mainTabsActiveName:{get:function(){return this.$store.state.common.mainTabsActiveName},set:function(t){this.$store.commit("common/updateMainTabsActiveName",t)}}},watch:{$route:"routeHandle"},created:function(){this.menuList=JSON.parse(sessionStorage.getItem("menuList")||"[]"),this.dynamicMenuRoutes=JSON.parse(sessionStorage.getItem("dynamicMenuRoutes")||"[]"),this.routeHandle(this.$route)},methods:{routeHandle:function(t){if(t.meta.isTab){var e=this.mainTabs.filter(function(e){return e.name===t.name})[0];if(!e){if(t.meta.isDynamic&&!(t=this.dynamicMenuRoutes.filter(function(e){return e.name===t.name})[0]))return console.error("未能找到可用标签页!");e={menu_id:t.meta.menu_id||t.name,name:t.name,title:t.meta.title,type:Object(a.c)(t.meta.iframeUrl)?"iframe":"module",iframeUrl:t.meta.iframeUrl||""},this.mainTabs=this.mainTabs.concat(e)}this.menuActiveName=e.menu_id+"",this.mainTabsActiveName=e.name}}}},o={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("aside",{staticClass:"site-sidebar",class:"site-sidebar--"+t.sidebarLayoutSkin},[n("div",{staticClass:"site-sidebar__inner"},[n("el-menu",{staticClass:"site-sidebar__menu",attrs:{"default-active":t.menuActiveName||"home",collapse:t.sidebarFold,collapseTransition:!1}},[n("el-menu-item",{attrs:{index:"home"},on:{click:function(e){t.$router.push({name:"home"})}}},[n("icon-svg",{staticClass:"site-sidebar__menu-icon",attrs:{name:"shouye"}}),t._v(" "),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("首页")])],1),t._v(" "),n("el-submenu",{attrs:{index:"demo"}},[n("template",{attrs:{slot:"title"},slot:"title"},[n("icon-svg",{staticClass:"site-sidebar__menu-icon",attrs:{name:"shoucang"}}),t._v(" "),n("span",[t._v("demo")])],1),t._v(" "),n("el-menu-item",{attrs:{index:"demo-echarts"},on:{click:function(e){t.$router.push({name:"demo-echarts"})}}},[n("icon-svg",{staticClass:"site-sidebar__menu-icon",attrs:{name:"tubiao"}}),t._v(" "),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("echarts")])],1),t._v(" "),n("el-menu-item",{attrs:{index:"demo-ueditor"},on:{click:function(e){t.$router.push({name:"demo-ueditor"})}}},[n("icon-svg",{staticClass:"site-sidebar__menu-icon",attrs:{name:"editor"}}),t._v(" "),n("span",{attrs:{slot:"title"},slot:"title"},[t._v("ueditor")])],1)],2),t._v(" "),t._l(t.menuList,function(e){return n("sub-menu",{key:e.menu_id,attrs:{menu:e,dynamicMenuRoutes:t.dynamicMenuRoutes}})})],2)],1)])},staticRenderFns:[]},r=n("46Yf")(i,o,!1,null,null,null);e.default=r.exports},"dLd/":function(t,e){t.exports="data:image/png;base64,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"},"sRz/":function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=n("TdIe"),a=n("YbVU"),i=n("xzNW"),o={data:function(){return{loading:!0}},components:{MainNavbar:s.default,MainSidebar:a.default,MainContent:i.default},computed:{documentClientHeight:{get:function(){return this.$store.state.common.documentClientHeight},set:function(t){this.$store.commit("common/updateDocumentClientHeight",t)}},sidebarFold:{get:function(){return this.$store.state.common.sidebarFold}},userId:{get:function(){return this.$store.state.user.id},set:function(t){this.$store.commit("user/updateId",t)}},userName:{get:function(){return this.$store.state.user.name},set:function(t){this.$store.commit("user/updateName",t)}}},created:function(){this.getUserInfo()},mounted:function(){this.resetDocumentClientHeight()},methods:{resetDocumentClientHeight:function(){var t=this;this.documentClientHeight=document.documentElement.clientHeight,window.onresize=function(){t.documentClientHeight=document.documentElement.clientHeight}},getUserInfo:function(){var t=this;this.$http({url:this.$http.adornUrl("/sys/user/info/3"),method:"get",params:this.$http.adornParams()}).then(function(e){var n=e.data;n&&200===n.code&&(t.loading=!1,t.userId=n.user.UserId,t.userName=n.user.Username)})}}},r={render:function(){var t=this.$createElement,e=this._self._c||t;return e("div",{directives:[{name:"loading",rawName:"v-loading.fullscreen.lock",value:this.loading,expression:"loading",modifiers:{fullscreen:!0,lock:!0}}],staticClass:"site-wrapper",class:{"site-sidebar--fold":this.sidebarFold},attrs:{"element-loading-text":"拼命加载中"}},[this.loading?this._e():[e("main-navbar"),this._v(" "),e("main-sidebar"),this._v(" "),e("div",{staticClass:"site-content__wrapper",style:{"min-height":this.documentClientHeight+"px"}},[e("main-content")],1)]],2)},staticRenderFns:[]},m=n("46Yf")(o,r,!1,null,null,null);e.default=m.exports},xzNW:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var s=n("E4LH"),a={data:function(){return{}},computed:{documentClientHeight:{get:function(){return this.$store.state.common.documentClientHeight}},menuActiveName:{get:function(){return this.$store.state.common.menuActiveName},set:function(t){this.$store.commit("common/updateMenuActiveName",t)}},mainTabs:{get:function(){return this.$store.state.common.mainTabs},set:function(t){this.$store.commit("common/updateMainTabs",t)}},mainTabsActiveName:{get:function(){return this.$store.state.common.mainTabsActiveName},set:function(t){this.$store.commit("common/updateMainTabsActiveName",t)}},siteContentViewHeight:function(){var t=this.documentClientHeight-50-30-2;return this.$route.meta.isTab?(t-=40,Object(s.c)(this.$route.meta.iframeUrl)?{height:t+"px"}:{minHeight:t+"px"}):{minHeight:t+"px"}}},methods:{selectedTabHandle:function(t){(t=this.mainTabs.filter(function(e){return e.name===t.name})).length>=1&&this.$router.push({name:t[0].name})},removeTabHandle:function(t){var e=this;this.mainTabs=this.mainTabs.filter(function(e){return e.name!==t}),this.mainTabs.length>=1?t===this.mainTabsActiveName&&this.$router.push({name:this.mainTabs[this.mainTabs.length-1].name},function(){e.mainTabsActiveName=e.$route.name}):(this.menuActiveName="",this.$router.push({name:"home"}))},tabsCloseCurrentHandle:function(){this.removeTabHandle(this.mainTabsActiveName)},tabsCloseOtherHandle:function(){var t=this;this.mainTabs=this.mainTabs.filter(function(e){return e.name===t.mainTabsActiveName})},tabsCloseAllHandle:function(){this.mainTabs=[],this.menuActiveName="",this.$router.push({name:"home"})},tabsRefreshCurrentHandle:function(){var t=this,e=this.mainTabsActiveName;this.removeTabHandle(e),this.$nextTick(function(){t.$router.push({name:e})})}}},i={render:function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("main",{staticClass:"site-content",class:{"site-content--tabs":t.$route.meta.isTab}},[t.$route.meta.isTab?n("el-tabs",{attrs:{closable:!0},on:{"tab-click":t.selectedTabHandle,"tab-remove":t.removeTabHandle},model:{value:t.mainTabsActiveName,callback:function(e){t.mainTabsActiveName=e},expression:"mainTabsActiveName"}},[n("el-dropdown",{staticClass:"site-tabs__tools",attrs:{"show-timeout":0}},[n("i",{staticClass:"el-icon-arrow-down el-icon--right"}),t._v(" "),n("el-dropdown-menu",{attrs:{slot:"dropdown"},slot:"dropdown"},[n("el-dropdown-item",{nativeOn:{click:function(e){t.tabsCloseCurrentHandle(e)}}},[t._v("关闭当前标签页")]),t._v(" "),n("el-dropdown-item",{nativeOn:{click:function(e){t.tabsCloseOtherHandle(e)}}},[t._v("关闭其它标签页")]),t._v(" "),n("el-dropdown-item",{nativeOn:{click:function(e){t.tabsCloseAllHandle(e)}}},[t._v("关闭全部标签页")]),t._v(" "),n("el-dropdown-item",{nativeOn:{click:function(e){t.tabsRefreshCurrentHandle(e)}}},[t._v("刷新当前标签页")])],1)],1),t._v(" "),t._l(t.mainTabs,function(e){return n("el-tab-pane",{key:e.name,attrs:{label:e.title,name:e.name}},[n("el-card",{attrs:{"body-style":t.siteContentViewHeight}},["iframe"===e.type?n("iframe",{attrs:{src:e.iframeUrl,width:"100%",height:"100%",frameborder:"0",scrolling:"yes"}}):n("keep-alive",[e.name===t.mainTabsActiveName?n("router-view"):t._e()],1)],1)],1)})],2):n("el-card",{attrs:{"body-style":t.siteContentViewHeight}},[n("keep-alive",[n("router-view")],1)],1)],1)},staticRenderFns:[]},o=n("46Yf")(a,i,!1,null,null,null);e.default=o.exports},zQrT:function(t,e,n){t.exports=n.p+"static/img/avatar.c58e465.png"}});