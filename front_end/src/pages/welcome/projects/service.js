import { request } from 'umi';
export async function queryFakeList(params) {
  const res = request('/api/fake_list', {
    params,
  });
  return res;
}

// export function queryItemList() {
//   const res = request('/api/getItems', {
//     method: 'GET',
//   });
//   console.log("items");
//   console.log(res);
//   return res;
// }
export async function queryItemList(options) {
  return request('/api/getItems', {
    method: 'GET',
    ...(options || {}),
  });
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