<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.name" />
      </el-form-item>

      <el-form-item label="">
        <el-input v-model="form.url" />
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="onSubmit">保存</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { add } from '@/api/tag'


export default {
  data() {
    return {

      form: {
        name: '',
         url:'',
      },

    }
  },
  methods: {
    onSubmit() {

      add({name:this.form.name,url:this.form.url}).then(response => {

        if(response.code !== 20000){
            this.$message('添加失败')
        }else{
           this.$router.push("/tag/index");

        }

      })
    },


    update(){
        this.input = e.target.value;
    },

    onCancel() {
      this.$message({
        message: 'cancel!',
        type: 'warning'
      })
    }
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

