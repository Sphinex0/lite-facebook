// // Context/WorkerContext.jsx
// "use client";
// import { createContext, useContext, useEffect, useRef, useState } from "react";

// const WorkerContext = createContext();

// export function WorkerProvider({ children }) {
//     const workerRef = useRef(null);
//     const portRef = useRef(null);
//     const [clientWorker, setClientWorker] = useState(null);
//     const [conversations, setConversations] = useState([]);

//     useEffect(() => {
//         workerRef.current = new SharedWorker("/sharedworker.js");
//         portRef.current = workerRef.current.port;
//         portRef.current.start();
//         setClientWorker(workerRef.current);

//         return () => {
//             portRef.current.close();
//             portRef.current = null;
//         };
//     }, []);

//     return (
//         <WorkerContext.Provider value={{ port: portRef.current , clientWorker, conversations, setConversations , portRef }}>
//             {children}
//         </WorkerContext.Provider>
//     );
// }

// export const useWorker = () => useContext(WorkerContext);



// Context/WorkerContext.jsx
"use client";
import { createContext, useContext, useEffect, useRef, useState } from "react";

const WorkerContext = createContext();

export function WorkerProvider({ children }) {
    const workerRef = useRef(null);
    const portRef = useRef(null);
    const [clientWorker, setClientWorker] = useState(null);
    const [conversations, setConversations] = useState([]);

    useEffect(() => {
        const worker = new SharedWorker("/sharedworker.js");
        workerRef.current = worker;
        portRef.current = worker.port;
        portRef.current.start();

        setClientWorker(worker);

        return () => {
            portRef.current?.close();
            portRef.current = null;
            workerRef.current = null;
        };
    }, []);

    return (
        <WorkerContext.Provider
            value={{ portRef, clientWorker, conversations, setConversations }}
        >
            {children}
        </WorkerContext.Provider>
    );
}

export const useWorker = () => useContext(WorkerContext);