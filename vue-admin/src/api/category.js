import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admin/category/list',
    method: 'get',
    params
  })
}


export function getItem(params) {
  return request({
    url: '/admin/category/get',
    method: 'get',
    params
  })
}

export function add(params){
  return request({
    url: '/admin/category/add',
    method: 'post',
    data:params
  })

}


export function edit(params){
  return request({
    url: '/admin/category/edit',
    method: 'post',
    data:params
  })

}
