import Link from 'next/link'

export default function Page500() {
  return <>
    <h1>500 - Server-Side Error</h1>;
    <Link href="/">Return Home</Link>
  </>
}