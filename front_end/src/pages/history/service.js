import { request } from 'umi';
export async function fakeSubmitForm(params) {
  return request('/api/advancedForm', {
    method: 'POST',
    data: params,
  });
}


export async function getOrderHis(options) {
  const temp = request('/api/getOrderHistory', {
    method: 'GET',
    ...(options || {}),
  });
  return temp;
}

export async function add2Cart(body, options) {
  return request('/api/addtoCart', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}