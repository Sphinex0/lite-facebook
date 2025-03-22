// app/not-found.js
import Link from 'next/link';
import { notFound } from 'next/navigation';

export default function NotFound() {
  return (
    <div>
      <h1>404 - Not Found</h1>
      <p>Sorry, we couldn’t find that page.</p>
      <Link href="/">Go Home</Link>
    </div>
  );
}