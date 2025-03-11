import { NextResponse } from "next/server";

export function middleware(req) {
  const authToken = req.cookies.get("session_token");
console.log(authToken);

  if (!authToken) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: ["/"], // Protect dashboard routes
};
