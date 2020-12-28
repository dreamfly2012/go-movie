import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admin/tag/list',
    method: 'get',
    params
  })
}


export function getItem(params) {
  return request({
    url: '/admin/tag/get',
    method: 'get',
    params
  })
}

export function add(params){
  return request({
    url: '/admin/tag/add',
    method: 'post',
    data:params
  })

}


export function edit(params){
  return request({
    url: '/admin/tag/edit',
    method: 'post',
    data:params
  })

}
