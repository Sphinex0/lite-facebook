"use client";

import { useState } from "react";
import "./login.css";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    //Login()
  };

  return (
    <div className="container">
      <div className="form-box">
        <h2>Login</h2>

        <form onSubmit={handleSubmit}>
          {/* Email Input */}
          <div className="input-group">
            <input
              type="text"
              value={email}
              placeholder="User or Email"
              required
              id="username"
              onChange={(e) => setEmail(e.target.value)}
              className="input-field"
            />
          </div>

          {/* Password Input */}
          <div className="input-group">
            <input
              type="password"
              value={password}
              placeholder="Password"
              required
              id="password"
              onChange={(e) => setPassword(e.target.value)}
              className="input-field"
            />
          </div>

          {/* Submit Button */}
          <button type="submit" className="submit-btn">Login</button>
        </form>
      </div>
    </div>
  );
}