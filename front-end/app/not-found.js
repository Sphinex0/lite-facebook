// app/not-found.js
import Link from 'next/link';
import "./global.css"

export default function NotFound() {
  return (
    <div className="error-page">
      <h1>404</h1>
      <p>Oops! The page you're looking for doesn't exist.</p>
      <Link href="/">Go back to Homepage</Link>
    </div>
  );
}