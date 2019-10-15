<template>
<div class="sky">
 <div class="demo-ruleForm login-container login-main">
          <h3 class="login-title">管理员登录</h3>
          <el-form
            :model="dataForm"
            :rules="dataRule"
            ref="dataForm"
            @keyup.enter.native="dataFormSubmit()"
            status-icon
          >
            <el-form-item prop="userName">
              <el-input v-model="dataForm.userName" placeholder="帐号"></el-input>
            </el-form-item>
            <el-form-item prop="password">
              <el-input v-model="dataForm.password" type="password" placeholder="密码"></el-input>
            </el-form-item>
            <el-form-item prop="idkeyd" hidden>
              <el-input v-model="dataForm.idkeyd"></el-input>
            </el-form-item>
            <el-form-item prop="checkcode">
              <el-row :gutter="20">
                <el-col :span="14">
                  <el-input v-model="dataForm.checkcode" placeholder="验证码"></el-input>
                </el-col>
                <el-col :span="8" class="login-checkcode">
                  <img :src="checkcodePath" @click="getCheckcode()" alt />
                </el-col>
              </el-row>
            </el-form-item>
            <el-form-item>
              <el-button class="login-btn-submit" type="primary" @click="dataFormSubmit()">登 录</el-button>
            </el-form-item>
          </el-form>
        </div>
        <div class="cloud"></div>
  </div>
</template>

<script>
import { getUUID } from "@/utils";
export default {
 components: {
			//VueParticles
	},
  data() {
    return {
      dataForm: {
        userName: "",
        password: "",
        uuid: "",
        idkeyd: "",
        checkcode: ""
      },
      dataRule: {
        userName: [
          { required: true, message: "帐号不能为空", trigger: "blur" }
        ],
        password: [
          { required: true, message: "密码不能为空", trigger: "blur" }
        ],
        checkcode: [
          { required: true, message: "验证码不能为空", trigger: "blur" }
        ]
      },
      checkcodePath: ""
    };
  },
  created() {
    this.getCheckcode();
  },
  methods: {
    // 提交表单
    dataFormSubmit() {
      this.$refs["dataForm"].validate(valid => {
        if (valid) {
          this.$http({
            url: this.$http.adornUrl(`/sys/login`),
            method: "post",
            data: this.$http.adornData({
              'username': this.dataForm.userName,
              'password': this.dataForm.password,
              'uuid': this.dataForm.uuid,
              'idkeyd': this.dataForm.idkeyd,
              'checkcode': this.dataForm.checkcode
            })
          }).then(({ data }) => {
            console.log("登录请求返回的数据是：", { data });
            if (data && data.code === 200) {
              this.$cookie.set("token", data.token);//存储token
              this.$cookie.set("login_id", data.user.user_id);//存储登录id
              this.$message({ //弹出登录成功提示框
                  message: '登录成功',
                  type: 'success',
                  duration: 500,
                  onClose: () => {
                    this.visible = false
                    this.$nextTick(() => {
                          this.$router.replace({ name: "home" });
                    })
                  }
                })
            } else {
              this.getCheckcode();
              this.$message.error(data.info);
            }
          });
        }
      });
    },
    // 获取验证码
    getCheckcode() {
      this.dataForm.uuid = getUUID();
      this.$http({
        url: this.$http.adornUrl(`/sys/verifyCode?uuid=${this.dataForm.uuid}`),
        method: 'get'
      }).then(({ data }) => {
        //console.log("验证通过返回的数据是：", { data });
        if (data && data.code === 200) {
          this.dataForm.idkeyd = data.id
          this.checkcodePath = data.src;
        } else {
          this.getCheckcode();
        }
      });
    }
  }
};
</script>

<style lang="scss">
body{
        background: url(~@/assets/img/login_bg.png) repeat-x;
        min-height: 600px;
        position: relative;
  }
  .sky {
    background: url(~@/assets/img/sky.png) repeat;
    width: 100%;
    height: 100%;
    z-index: 1;
    position: absolute;
    top: 0px;
}
.cloud {
    background: url(~@/assets/img/cloud.png) repeat;
    width: 100%;
    height: 356px;
    position:absolute;
    top: 450px;
    -webkit-animation: cloud 60s linear infinite alternate;
    -moz-animation: clouds 60s linear infinite alternate;
    z-index: -2;
 }

 .login-container {
    /*box-shadow: 0 0px 8px 0 rgba(0, 0, 0, 0.06), 0 1px 0px 0 rgba(0, 0, 0, 0.02);*/
    -webkit-border-radius: 5px;
    border-radius: 5px;
    -moz-border-radius: 5px;
    background-clip: padding-box;
    margin: 180px auto;
    width: 350px;
    z-index: 999;
    padding: 35px 35px 15px 35px;
    background: #fff;
    border: 1px solid #eaeaea;
    box-shadow: 0 0 25px #cac6c6;
    .title {
      margin: 0px auto 40px auto;
      text-align: center;
      color: #505458;
    }
    .remember {
      margin: 0px 0px 35px 0px;
    }
  

  .login-main {
    position: absolute;
    top: 0;
    right: 0;
    padding: 150px 60px 180px;
    width: 470px;
    min-height: 100%;
    background-color: #fff;
  }
  .login-title {
    text-align: center;
    font-size: 1.0rem;
  }
  .login-checkcode {
    overflow: hidden;
    > img {
      width: 100%;
      cursor: pointer;
    }
  }
  .login-btn-submit {
    width: 100%;
    margin-top: 38px;
    font-size: 1.0rem;
  }
}


</style>
