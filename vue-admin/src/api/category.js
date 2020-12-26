import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/category/list',
    method: 'get',
    params
  })
}

export function add(params){
  return request({
    url: '/category/add',
    method: 'post',
    data:params
  })

}
