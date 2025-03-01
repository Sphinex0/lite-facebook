import { NextResponse } from 'next/server';

export function middleware(request, event) {
    console.log("yeeeees");
    
  const valide = event.waitUntil(Checkuservalidity()) 
console.log(valide);

  if (!valide) {
    return NextResponse.redirect(new URL('/login', request.url));
  }
  return NextResponse.next();
}

export const config = {
    matcher: ['/'], // Apply the middleware to specific routes
};