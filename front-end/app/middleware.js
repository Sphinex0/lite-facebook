import { NextResponse} from "next/server";

async function checkuservalidity(req) {
  const response = await fetch(`${window.location.host}/api/checkuser`,{
    headers: {
      cookie: req.headers.get("session_token") || "",
    },
  })
  
  return response.status == 200
}

export async function middleware(req) {
 // const authToken =  await checkuservalidity(req)
//console.log("sfsgdgdfgdfg");

  if (req.nextUrl.pathname === "/") {
    return NextResponse.redirect(new URL("/login", req.url));
  }

}
/*
export const config = {
  matcher: ["/"],
};
*/