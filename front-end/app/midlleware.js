import { NextResponse } from 'next/server'
 
export async function middleware(request, event) {
    const validity = event.waituntil(Checkuservalidity())
    if (!validity) {
        return NextResponse.redirect(new URL('/login', request.url))
    }
}