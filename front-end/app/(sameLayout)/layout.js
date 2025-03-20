"use client"
import { createContext, useEffect, useRef, useState } from "react";
import Profile from "./_leftSide/profile";
import SideBar from "./_leftSide/sideBar";
import "./main.css";
import { WorkerProvider } from "../_Context/WorkerContext";

export const Context = createContext()


export default function MainLayout({ children }) {


  // const [clientWorker, setClientWorker] = useState(null);
  // const workerPortRef = useRef(null);
  // const [conversations, setConversations] = useState([]);

  // // Initialize SharedWorker
  // useEffect(() => {
  //   const worker = new SharedWorker("/sharedworker.js");
  //   workerPortRef.current = worker.port;
  //   setClientWorker(worker);

  //   return () => {
  //     worker.port.close();
  //     workerPortRef.current = null;
  //   };
  // }, []);

  // useEffect(() => {
  //   console.log("conversations", conversations)
  // }, [conversations])



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