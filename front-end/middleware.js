import { NextResponse } from "next/server";

export default async function middleware(request) {
  const publicRoutes = ["/login", "/signup"];
  const { pathname } = request.nextUrl;

  // Exclude static files and API routes
  const excludePaths = [
    '/_next/static',
    '/_next/image',
    '/favicon.ico',
    '/api/',
    '/public/'
  ];

  if (excludePaths.some(path => pathname.startsWith(path))) {
    return NextResponse.next();
  }

  // Check authentication status 
  const authCheck = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/checkuser`, {
    headers: {
      Cookie: request.headers.get('Cookie') || '',
    },
  });
  
  const isAuthenticated = authCheck.status === 200;
  console.log(isAuthenticated)

  // Handle public routes
  if (publicRoutes.includes(pathname)) {
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