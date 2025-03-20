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

// import { NextResponse } from 'next/server';

// export default async function middleware(request) {
//   console.log("Middleware running for:", request.nextUrl.pathname); // Log the pathname
//   const response = await NextResponse.next();
//   console.log("Response status:", response.status); // Log the status
//   if (response.status === 404) {
//     console.log("Redirecting to /login");
//     return NextResponse.redirect(new URL('/login', request.url));
//   }
//   return response;
// }