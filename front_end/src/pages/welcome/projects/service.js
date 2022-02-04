import { request } from 'umi';
export async function queryFakeList(params) {
  const res = request('/api/fake_list', {
    params,
  });
  return res;
}

export function queryItemList(params) {
  const res = request('/api/getItems', {
    params,
  });
  console.log(res);
  return res;
}
