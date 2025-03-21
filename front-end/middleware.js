// import { NextResponse} from "next/server";

// async function checkuservalidity(req) {
//   const response = await fetch(`${window.location.host}/api/checkuser`,{
//     headers: {
//       cookie: req.headers.get("session_token") || "",
//     },
//   })

//   return response.status == 200
// }

// export async function middleware(req) {
//  // const authToken =  await checkuservalidity(req)
// //console.log("sfsgdgdfgdfg");

//   if (req.nextUrl.pathname === "/") {
//     return NextResponse.redirect(new URL("/login", req.url));
//   }

// }
/*
export const config = {
  matcher: ["/"],
};
*/

import { NextResponse } from 'next/server';

const redirect = true

export default async function middleware(request) {
  const allowed = ["/login", "/signup"];
  const res = await fetch("http://localhost:8080/api/checkuser", {
    credentials: "include",
  })
  // if (res.status == 200) {
  //   // console.log("res", res)
  //   return NextResponse.next();
  // }
  // // console.log("res", res)
  // // if (user) {
  // // return NextResponse.next();
  // // }
  // console.log(res.status)
  // if (res.status == 401 && !allowed.includes(request.nextUrl.pathname)) {
  //   return NextResponse.redirect('http://localhost:3000/login');
  // }
  // const response = NextResponse.redirect('/login');
    return NextResponse.next();
}