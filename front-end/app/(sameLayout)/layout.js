"use client"
import { createContext, useEffect, useRef, useState } from "react";
import Profile from "./_leftSide/profile";
import SideBar from "./_leftSide/sideBar";
import "./main.css";
import { WorkerProvider } from "../_Context/WorkerContext";
import Navbar from "./_components/navbar/Navbar";

export const Context = createContext()


export default function MainLayout({ children }) {

  return (
      <WorkerProvider>
        <Navbar />
        <main>
        <div style={{display:"none"}} className={`errorPopup error`}></div>
          <div className="container">
            <div className="left">
              <Profile />
              <SideBar />
              <label className="btn btn-primary">Create Post</label>
            </div>
            <div className="middle">
              {children}
            </div>
          </div>
        </main>
      </WorkerProvider>
  );
}