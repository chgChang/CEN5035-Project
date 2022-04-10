import { request } from "umi";
export async function queryFakeList(params) {
  return request("/api/get_list", {
    params,
  });
}

export async function removeFakeList(params) {
  return request("/api/post_fake_list", {
    method: "POST",
    data: { ...params, method: "delete" },
  });
}
export async function addFakeList(params) {
  return request("/api/post_fake_list", {
    method: "POST",
    data: { ...params, method: "post" },
  });
}
export async function updateFakeList(params) {
  return request("/api/post_fake_list", {
    method: "POST",
    data: { ...params, method: "update" },
  });
}


export async function queryCartList(options) {
  const temp = request('/api/getCartItems', {
    method: 'GET',
    ...(options || {}),
  });
  return temp;
}

export async function updateCart(body, options) {
  return request('/api/updateCart', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteCartByItemId(body, options) {
  return request('/api/deleteCartByItemId', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function removeCart(body, options) {
  return request('/api/removeCart', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
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