'use client';

import { useEffect, useState } from "react";

export default function HomePage() {
    const [messages, setMessages] = useState<string[]>([]);

    useEffect(() => {
        const eventSource = new EventSource('http://localhost:8086/events');

        eventSource.onmessage = (event) => {
            setMessages((prev) => [...prev, event.data]);
        };

        eventSource.onerror = (err) => {
            console.error("Event source failed");
            eventSource.close();
        };

        return () => {
            eventSource.close();
        };
    }, []);

    return (
        <main>
            <h1> Live Server </h1>
            <pre>
                {messages.map((msg, idx) => (
                    <div key={idx}>{msg}</div>
                ))}
            </pre>
        </main>
    );
}
