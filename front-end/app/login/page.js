'use client';

import { useState } from 'react';
import './globals.css'; 

export default function Home() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    await FetchLogin(email,password);
  };

  return (
    <div className="container">
      <form onSubmit={handleSubmit}>
        {/* Email Input */}
        <label htmlFor="username">User or email:</label>
        <input
          type="text"
          value={email}
          placeholder="user or email"
          required
          id="username"
          onChange={(e) => setEmail(e.target.value)}
        />
        <i className="bx bxs-user"></i>

        {/* Password Input */}
        <label htmlFor="password">Password:</label>
        <input
          type="password"
          value={password}
          placeholder="Password"
          required
          id="password"
          onChange={(e) => setPassword(e.target.value)}
        />
        <i className="bx bxs-lock-alt"></i>

        <button type="submit" id="login-btn">
          Login
        </button>
      </form>
    </div>
  );
}
