"use client"
import { createContext, useEffect, useRef, useState } from "react";
import Profile from "./_leftSide/profile";
import SideBar from "./_leftSide/sideBar";
import "./main.css";
import { WorkerProvider } from "../_Context/WorkerContext";

export const Context = createContext()


export default function MainLayout({ children }) {

  return (
      <WorkerProvider>
        <main>
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