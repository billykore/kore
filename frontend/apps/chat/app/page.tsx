"use client"

import React, { useEffect, useState } from "react"
import { BackgroundText, Title } from "@repo/ui/typography"
import { TextArea } from "@repo/ui/textarea"
import { SecondaryButton } from "@repo/ui/button"
import { Divider } from "@repo/ui/divider"
import { Card } from "@repo/ui/card"
import { Input } from "@repo/ui/input"

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

interface ChatMessage {
  name: string
  message: string
}

export default function Home() {
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const [name, setName] = useState<string>("")
  const [newMessage, setNewMessage] = useState<string>("")

  useEffect(() => {
    socket.onmessage = (e: MessageEvent) => {
      const msg = JSON.parse(e.data)
      setMessages((prevMessages) => [msg, ...prevMessages])
    }
  }, [])

  const sendMessage = () => {
    const chatMessage: ChatMessage = {
      name: name,
      message: newMessage,
    }
    socket.send(JSON.stringify(chatMessage))
    setName("")
    setNewMessage("")
  }

  return (
    <main className="min-h-screen max-w-2xl py-12 mx-auto">
      <Title>WiChat</Title>

      <Input
        id="name"
        label="Your name"
        value={name}
        onChange={(e) => setName(e.target.value)}
      />
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

      {messages.length > 0 ? messages.map((msg) => (
        <Card
          title={msg.name}
          body={msg.message}
        />
      )) : <BackgroundText>No Message</BackgroundText>}
    </main>
  )
}
