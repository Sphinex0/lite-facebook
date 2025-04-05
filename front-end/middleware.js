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
  const api_url = process.env.API_URL || "http://localhost:8080"
  try {
    const authCheck = await fetch(`${api_url}/api/checkuser`, {
      headers: {
        Cookie: request.headers.get('Cookie') || '',
      },
    });

    if (authCheck.status == 429) {
      return NextResponse.next();
    }

    const isAuthenticated = authCheck.status === 200;

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
  } catch (error) {
    console.log(error)
    const errorUrl = new URL('/page500', request.url);
    return NextResponse.rewrite(errorUrl);
  }

}