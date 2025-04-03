"use client";
import { createContext, useContext, useEffect, useRef } from "react";

export const Ctx = createContext();

export function CtxP({ children }) {
    const userRef = useRef({});

    useEffect(() => {
        const storedUser = localStorage.getItem("user");
        userRef.current = storedUser ? JSON.parse(storedUser) : {};
        console.log("userRef", userRef.current);
    }, []);

    return (
        <Ctx.Provider value={{userRef}}>
            {children}
        </Ctx.Provider>
    );
}

export const useCtx = () => useContext(Ctx);