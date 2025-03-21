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
  const publicRoutes = ["/login", "/signup"];
  const { pathname } = request.nextUrl;

  // Check authentication status
  const authCheck = await fetch("http://localhost:8080/api/checkauth", {
    credentials: "include",
  });
  
  const isAuthenticated = authCheck.status === 200;

  // Handle public routes
  if (publicRoutes.includes(pathname)) {
    // Redirect authenticated users away from public routes
    if (isAuthenticated) {
      return NextResponse.redirect(new URL("/", request.url));
    }
    return NextResponse.next();
  }

  // Protect private routes
  if (!isAuthenticated) {
    return NextResponse.redirect(new URL("/login", request.url));
  }

  return NextResponse.next();
}