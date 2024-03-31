"use client"

import React, { useEffect, useState } from "react"

function connectWebSocket(url: string): WebSocket {
  const socket = new WebSocket(url)

  socket.onopen = () => {
    console.log("Connected to websocket")
  }
  socket.onerror = (e) => {
    console.log("Error:", e)
  }

  return socket
}

const socket = connectWebSocket("ws://localhost:8000/chat")

export default function Home() {
  const [messages, setMessages] = useState<string[]>([])
  const [newMessage, setNewMessage] = useState<string>("")

  useEffect(() => {
    socket.onmessage = (e: MessageEvent<string>) => {
      setMessages((prevMessages) => [e.data, ...prevMessages])
    }
  }, [])

  const sendMessage = () => {
    socket.send(newMessage)
    setNewMessage("")
  }

  return (
    <main className="min-h-screen max-w-4xl py-12 mx-auto">
      <h1 className="text-4xl mb-12">WiChat</h1>

      <div className="mt-4">
        <label
          htmlFor="message"
          className="block mb-2 text-sm font-medium text-gray-900"
        >
          Your message
        </label>
        <textarea
          id="message"
          rows={4}
          className="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border
            border-gray-300 focus:ring-blue-500 focus:border-blue-500"
          placeholder="Write your thoughts here..."
          value={newMessage}
          onChange={(e) => setNewMessage(e.target.value)}></textarea>
      </div>

      <button
        type="button"
        className="mt-4 text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100
          focus:ring-4 focus:ring-gray-100 font-medium rounded-lg text-sm px-5 py-2.5 me-2 mb-2"
        onClick={() => sendMessage()}
      >
        Send
      </button>

      <div className="my-12 border border-t-gray-50"></div>

      <div>
        {messages.length > 0 ? messages.map((msg, index) => (
          <div key={index} className="block p-6 bg-white border border-gray-200 rounded-lg my-2">
            <p className="font-normal text-xl text-gray-700">
              {msg}
            </p>
          </div>
        )) : <h1 className="text-2xl text-center text-gray-400">No Message</h1>}
      </div>
    </main>
  );
}
