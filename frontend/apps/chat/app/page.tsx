"use client"

import React, { useEffect, useState } from "react"
import { BackgroundText, Title } from "@repo/ui/typography"
import { TextArea } from "@repo/ui/textarea"
import { SecondaryButton } from "@repo/ui/button"
import { Divider } from "@repo/ui/divider"
import { Card } from "@repo/ui/card"

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
      <Title>WiChat</Title>

      <TextArea
        id="message"
        label="Your message"
        value={newMessage}
        onChange={(e) => setNewMessage(e.target.value)}
      />
      <SecondaryButton onClick={() => sendMessage()}>
        Send
      </SecondaryButton>

      <Divider/>

      {messages.length > 0 ? messages.map((msg, index) => (
        <Card id={index}>{msg}</Card>
      )) : <BackgroundText>No Message</BackgroundText>}
    </main>
  )
}
