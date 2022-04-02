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

export async function doCheckout(body, options) {
  return request('/api/checkout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}