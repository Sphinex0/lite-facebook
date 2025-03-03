"use client";

import { useState } from "react";
import "./signup.css";

export default function SignupPage() {
  const [form, setForm] = useState({
    email: "",
    password: "",
    firstName: "",
    lastName: "",
    dob: "",
    avatar: null,
    nickname: "",
    aboutMe: "",
  });

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleFileChange = (e) => {
    setForm({ ...form, avatar: e.target.files[0] });
  };

  return (
    <div className="container">
      <div className="form-box">
        <h2>Sign Up</h2>

        {["email", "password", "firstName", "lastName", "dob"].map((field) => (
          <div key={field} className="input-group">
            <input
              type={field === "dob" ? "date" : field === "password" ? "password" : "text"}
              name={field}
              placeholder={field.charAt(0).toUpperCase() + field.slice(1)}
              onChange={handleChange}
              className="input-field"
            />
          </div>
        ))}

        {/* Optional Fields */}
        <div className="input-group">
          <label className="file-label">Upload Avatar (Optional)</label>
          <input type="file" name="avatar" className="input-field file-input" onChange={handleFileChange} />
        </div>

        <div className="input-group">
          <input
            type="text"
            name="nickname"
            placeholder="Nickname (Optional)"
            onChange={handleChange}
            className="input-field"
          />
        </div>

        <textarea
          name="aboutMe"
          placeholder="About Me (Optional)"
          className="input-field textarea"
          onChange={handleChange}
        ></textarea>

        <button className="submit-btn">Sign Up</button>
      </div>
    </div>
  );
}
