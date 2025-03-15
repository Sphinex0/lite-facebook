"use client";

import { useState } from "react";
import { useRouter } from 'next/navigation'
import "./login.css";

export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    try {
      const response = await fetch("http://localhost:8080/api/login", {
        method: "POST",
        body: JSON.stringify({email,password}), 
      });

      if (response.status == 200) {
        const data = await response.json()
        document.cookie = `session_token=${data.uuid}; Path=/; Max-Age=31536000`;
        localStorage.setItem("user", data);
        console.log(data);
        
        router.push('/');
      } else {
        console.log(response);
      }
    } catch (error) {
      console.error("Error during sign-up:", error);
    }
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