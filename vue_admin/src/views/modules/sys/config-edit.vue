<template>
  <el-dialog
    :title="!dataForm.id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="参数名" prop="param_key">
        <el-input v-model="dataForm.param_key" placeholder="参数名"></el-input>
      </el-form-item>
      <el-form-item label="参数值" prop="param_value">
        <el-input v-model="dataForm.param_value" placeholder="参数值"></el-input>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="dataForm.remark" placeholder="备注"></el-input>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
  </el-dialog>
</template>

<script>
  export default {
    data () {
      return {
        visible: false,
        dataForm: {
          id: 0,
          param_key: '',
          param_value: '',
          remark: ''
        },
        dataRule: {
          param_key: [
            { required: true, message: '参数名不能为空', trigger: 'blur' }
          ],
          param_value: [
            { required: true, message: '参数值不能为空', trigger: 'blur' }
          ]
        }
      }
    },
    methods: {
      init (id) {
        this.dataForm.id = id || 0
        this.visible = true
        this.$nextTick(() => {
          this.$refs['dataForm'].resetFields()
          if (this.dataForm.id) {
            this.$http({
              url: this.$http.adornUrl(`/sys/config/info/${this.dataForm.id}`),
              method: 'get',
              params: this.$http.adornParams()
            }).then(({data}) => { 
              if (data && data.code === 200) { //console.log('获取到的配置信息',data)
                this.dataForm.id = data.config.id
                this.dataForm.param_key = data.config.param_key
                this.dataForm.param_value = data.config.param_value
                this.dataForm.remark = data.config.remark
              }
            })
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          if (valid) {
            this.$http({
              url: this.$http.adornUrl(`/sys/config/edit`),
              method: 'post',
              data: this.$http.adornData({
                'id': `${this.dataForm.id}`,
                'type':!this.dataForm.id ? 'save' : 'update',
                'param_key': this.dataForm.param_key,
                'param_value': this.dataForm.param_value,
                'remark': this.dataForm.remark
              })
            }).then(({data}) => {
              if (data && data.code === 200) {
                this.$message({
                  message: '操作成功',
                  type: 'success',
                  duration: 1500,
                  onClose: () => {
                    this.visible = false
                    this.$emit('refreshDataList')
                  }
                })
              } else {
                this.$message.error(data.info)
              }
            })
          }
        })
      }
    }
  }
</script>
