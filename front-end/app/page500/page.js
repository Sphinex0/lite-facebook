import Link from 'next/link';

export default function NotFound() {
  return (
    <div>
      <h1>404 - Not Found</h1>
      <p>Sorry, we couldn’t find that page.</p>
      <Link href="/">Go Home</Link>
    </div>
  );
}