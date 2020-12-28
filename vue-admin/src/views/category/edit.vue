<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="名称">
        <el-input v-model="form.name" />
      </el-form-item>

       <el-form-item label="URL">
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
import { edit,getItem } from '@/api/category'


export default {
  data() {
    return {
      id:"",
      form: {
        name: '',
        url:'',
      },

    }
  },
   created() {

      this.id = parseInt(this.$route.query.id)

      this.fetchData()
    },
  methods: {

    onSubmit() {

      edit({name:this.form.name,id:this.id,url:this.form.url}).then(response => {

        if(response.code !== 20000){
            this.$message(response.message)
        }else{
           this.$router.push("/category/index");

        }

      })
    },

    fetchData(){
        this.listLoading = true
        getItem({id:this.id}).then(response => {
          this.form = response.data
          this.listLoading = false
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

