import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/admin/list',
    method: 'get',
    params
  })
}

export function getItem(params) {
  return request({
    url: '/post/get',
    method: 'get',
    params
  })
}

export function add(params) {
  return request({
    url: '/post/add',
    method: 'post',
    data: params
  })

}

export function edit(params) {
  return request({
    url: '/post/edit',
    method: 'post',
    data: params
  })
}

export function uploadFileRequest(params)  {
  return request({
    method: 'post',
    url: '/post/upload',
    data: params,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}
