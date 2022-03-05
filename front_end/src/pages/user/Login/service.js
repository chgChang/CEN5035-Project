import { request } from "umi";
export async function fakeRegister(params) {
  console.log("hhh");
  return request("/api/register", {
    method: "POST",
    data: params,
  });
}
