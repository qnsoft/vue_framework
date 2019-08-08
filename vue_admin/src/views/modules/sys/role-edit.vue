<template>
  <el-dialog
    :title="!dataForm.role_id ? '新增' : '修改'"
    :close-on-click-modal="false"
    :visible.sync="visible">
    <el-form :model="dataForm" :rules="dataRule" ref="dataForm" @keyup.enter.native="dataFormSubmit()" label-width="80px">
      <el-form-item label="角色名称" prop="role_name">
        <el-input v-model="dataForm.role_name" placeholder="角色名称"></el-input>
      </el-form-item>
      <el-form-item label="备注" prop="remark">
        <el-input v-model="dataForm.remark" placeholder="备注"></el-input>
      </el-form-item>
      <el-form-item size="mini" label="授权">
        <el-tree
          :data="menuList"
          :props="menuListTreeProps"
          node-key="menu_id"
          ref="menuListTree"
          :default-expand-all="true"
          show-checkbox>
        </el-tree>
      </el-form-item>
    </el-form>
    <span slot="footer" class="dialog-footer">
      <el-button @click="visible = false">取消</el-button>
      <el-button type="primary" @click="dataFormSubmit()">确定</el-button>
    </span>
  </el-dialog>
</template>

<script>
  import { treeDataTranslate } from '@/utils'
  export default {
    data () {
      return {
        visible: false,
        menuList: [],
        menuListTreeProps: {
          label: 'name',
          children: 'children'
        },
        dataForm: {
          role_id: 0,
          role_name: '',
          remark: ''
        },
        dataRule: {
          role_name: [
            { required: true, message: '角色名称不能为空', trigger: 'blur' }
          ]
        },
        tempKey: -666666 // 临时key, 用于解决tree半选中状态项不能传给后台接口问题. # 待优化
      }
    },
    methods: {
      init (id) { 
        this.dataForm.role_id = id || 0
        this.$http({
          url: this.$http.adornUrl('/sys/menu/list'),
          method: 'get',
          params: this.$http.adornParams()
        }).then(({data}) => {
          this.menuList = treeDataTranslate(data.menuList, 'menu_id')
        }).then(() => {
          this.visible = true
          this.$nextTick(() => {
            this.$refs['dataForm'].resetFields()
            this.$refs.menuListTree.setCheckedKeys([])
          })
        }).then(() => { 
          if (this.dataForm.role_id) { 
            this.$http({
              url: this.$http.adornUrl(`/sys/role/info/${this.dataForm.role_id}`),
              method: 'get',
              params: this.$http.adornParams()
            }).then(({data}) => {
              if (data && data.code === 200) { //console.log('获取单条角色信息',data)
                this.dataForm.role_name = data.role.role_name
                this.dataForm.remark = data.role.remark
                let _menu_idlist=JSON.parse(data.role.menu_idlist)
                var idx = _menu_idlist.indexOf(this.tempKey)
                console.log(idx)
                if (idx !== -1) {
                  _menu_idlist.splice(idx, _menu_idlist.length - idx)
                }
                this.$refs.menuListTree.setCheckedKeys(_menu_idlist)
              }
            })
          }
        })
      },
      // 表单提交
      dataFormSubmit () {
        this.$refs['dataForm'].validate((valid) => {
          let _xh_ids=[].concat(this.$refs.menuListTree.getCheckedKeys(true))
         // console.log('获取选中',_xh_ids)
          let _idds=''
          _xh_ids.forEach(function(id){
            if(id>1){
              _idds+=id+','
             }
            })
           _xh_ids = _idds.substring(0, _idds.length - 1)
          if (valid) {
            this.$http({
              url: this.$http.adornUrl(`/sys/role/edit`),
              method: 'post',
              data: this.$http.adornData({
                'role_id':`${this.dataForm.role_id}`,
                'type':!this.dataForm.role_id ? 'save' : 'update',
                'role_name': this.dataForm.role_name,
                'remark': this.dataForm.remark,
                'menu_idlist':_xh_ids
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
                this.$message.error(data.msg)
              }
            })
          }
        })
      }
    }
  }
</script>
