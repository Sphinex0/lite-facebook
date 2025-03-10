"use client"
import { useEffect, useState } from "react";

export default function Chat() {
    const [data, setData] = useState(null)
    console.log("kk")
    useEffect(() => {
        fetch("http://localhost:8080/api/posts", {
            method: "POST",
            body: "{}"
        })
            .then(res => {
                console.log("response => ", res)
                return res.json()
            })
            .then((content) => {
                setData(content)
                console.log("data => ", data)
            })
    }, [])
    return (
        <>
            <h1>Chat</h1>
        </>
    )
}