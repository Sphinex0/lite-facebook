import Link from 'next/link';
import "../global.css"

export default function NotFound() {
    return (
        <div className="error-page">
            <h1>405</h1>
            <p>Status Method Not Allowed</p>
            <Link href="/">Go back to Homepage</Link>
        </div>
    );
}