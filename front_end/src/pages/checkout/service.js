import { request } from "umi";
export async function fakeSubmitForm(params) {
  return request("/api/stepForm", {
    method: "POST",
    data: params,
  });
}


export async function queryCartList(options) {
  const temp = request('/api/getCartItems', {
    method: 'GET',
    ...(options || {}),
  });
  return temp;
}