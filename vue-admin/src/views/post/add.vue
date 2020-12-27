<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="文章标题">
        <el-input v-model="form.title" />
      </el-form-item>

       <el-form-item label="文章简介">
         <el-input type="textarea" v-model="form.desc"></el-input>
      </el-form-item>


       <el-form-item label="文章类别">
       <el-select v-model="form.cid" placeholder="请选择">
        <el-option
          v-for="item in options"
          :key="item.ID"
          :label="item.name"
          :value="item.ID">
        </el-option>
      </el-select>
      </el-form-item>


      <el-form-item label="文章内容">

      <mavon-editor v-model="form.content" ref="md" :toolbars="toolbars" @imgAdd="handleEditorImgAdd"
     @imgDel="handleEditorImgDel"/>

      </el-form-item>


      <el-form-item>
        <el-button type="primary" @click="onSubmit">保存</el-button>
        <el-button @click="onCancel">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { getList } from '@/api/category'
import { add,uploadFileRequest } from '@/api/post'



export default {
  data() {
    return {
      imgFile:[],
      options:[],
      form: {
        cid: '',
        title: '',
        content: '',
        desc: ''
      },
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        mark: true, // 标记
        superscript: true, // 上角标
        quote: true, // 引用
        ol: true, // 有序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        help: true, // 帮助
        code: true, // code
        subfield: true, // 是否需要分栏
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        /* 1.3.5 */
        undo: true, // 上一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true // 导航目录
      }
    }
  },
  created() {
    this.fetchCategory()
  },
   computed: {
    compiledMarkdown: function() {
      return marked(this.input, { sanitize: true });
    }
  },
  methods: {
    onSubmit() {

      add({title:this.form.title,desc:this.form.desc,cid:this.form.cid,content:this.form.content}).then(response => {

        if(response.code !== 20000){
            this.$message('添加失败')
        }else{
           this.$router.push("/list/index");

        }

      })
    },
    fetchCategory(){
       getList().then(response => {
        this.options = response.data.data
        this.total = response.data.total
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
    },

    handleEditorImgAdd (pos, $file) {
      console.log(pos)
      console.log($file)
            let formdata = new FormData()
            formdata.append('file', $file)
            this.imgFile[pos] = $file




            uploadFileRequest(formdata).then(res => {
                if (res.code === 20000) {
                  this.$message({
        message: res.msg,
        type: 'success'
      })

                    let url = res.data
                    let name = $file.name
                    if (name.includes('-')) {
                        name = name.replace(/-/g, '')
                    }
                    let content = this.form.content
                    // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)  这里是必须要有的
                    if (content.includes(name)) {
                        let oStr = `(${pos})`
                        let nStr = `(${url})`
                        let index = content.indexOf(oStr)
                        let str = content.replace(oStr, '')
                        let insertStr = (soure, start, newStr) => {
                            return soure.slice(0, start) + newStr + soure.slice(start)
                        }
                        this.form.content = insertStr(str, index, nStr)
                    }
                } else {
                  this.$message({
        message: res.data.msg,
        type: 'error'
      })

                }
            })
        },
        handleEditorImgDel (pos) {
            delete this.imgFile[pos]
        },
  }
}
</script>

<style scoped>
.line{
  text-align: center;
}
</style>

