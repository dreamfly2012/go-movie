<template>
  <div class="app-container">
    <el-form :label-position="labelPosition" label-width="80px" >
  <el-form-item label="标题">
    <el-input v-model="name"></el-input>
  </el-form-item>
    <el-form-item>
    <el-button type="primary" @click="search">查询</el-button>
  </el-form-item>
</el-form>
    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.id  }}
        </template>
      </el-table-column>
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.title }}
        </template>
      </el-table-column>

      <el-table-column label="Desc" align="center">
        <template slot-scope="scope">
          {{ scope.row.desc }}
        </template>
      </el-table-column>
      <el-table-column  label="Views" width="110" align="center">
        <template slot-scope="scope">
           {{ scope.row.views }}
        </template>
      </el-table-column>
      <el-table-column align="center" prop="created" label="创建时间" :formatter="dateFormat"></el-table-column>

      <el-table-column align="center" label="operate" >
          <template slot-scope="scope">
            <el-button
          size="mini"
          @click="edit(scope.row)">编辑</el-button>

            <el-button
          size="mini"
          @click="del(scope.row)">删除</el-button>
          </template>
      </el-table-column>
    </el-table>

    <el-pagination
  background
  @current-change="current_change"
  :page-size=pagesize
  layout="prev, pager, next"
  :total="total">
</el-pagination>
  </div>
</template>

<script>
import { getList,searchList,DelList } from '@/api/post'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      name:"",
      list: null,
      total:0,//默认数据总数
      pagesize:10,//每页的数据条数
      currentPage:1,
      listLoading: true
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
     dateFormat:function(row,column){

        var t=new Date(row.created);//row 表示一行数据, updateTime 表示要格式化的字段名称
        return t.getFullYear()+"-"+(t.getMonth()+1)+"-"+t.getDate()+" "+t.getHours()+":"+t.getMinutes()+":"+t.getSeconds()
    },
    current_change(currentPage){
      console.log(currentPage)

       this.currentPage = currentPage;
       this.listLoading = true
      getList({currentPage:this.currentPage}).then(response => {
        this.list = response.data.data
        this.total = response.data.total
        this.listLoading = false
      })
    },
    edit(row){
      this.$router.push("/list/edit?id="+row.id);
    },
    del(row){
      DelList({id:row.id}).then(response=>{
        this.fetchData()
      })
    },
    search(){
      console.log(this.name)
      this.listLoading = true
      searchList({titile:this.name}).then(response => {
        this.list = response.data.data
        this.total = response.data.total
        this.listLoading = false
      })
    },
    fetchData() {
      this.listLoading = true
      getList({currentPage:this.currentPage}).then(response => {
        this.list = response.data.data
        this.total = response.data.total
        this.listLoading = false
      })
    }
  }
}
</script>
