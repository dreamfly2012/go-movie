<template>
  <div class="app-container">
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
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="Name">
        <template slot-scope="scope">
          {{ scope.row.name }}
        </template>
      </el-table-column>

      <el-table-column label="URL">
        <template slot-scope="scope">
          {{ scope.row.url }}
        </template>
      </el-table-column>

      <el-table-column align="center" label="operate" >
          <template slot-scope="scope">
             <el-button size="mini" @click="edit(scope.row)">编辑</el-button>
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
import { getList } from '@/api/tag'

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
      console.log(row)
      this.$router.push("/tag/edit?id="+row.ID);
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
