"use client"

import { useState } from "react"
import { Bungee_Shade } from "next/font/google"

// Local
import Button from "@/components/Button"
import InputForm from "@/components/InputForm"
import Container from "@/components/Container"
import axios from "@/utils/axios"
import { setCookie } from "@/utils/cookie"
import Alert from "@/components/Alert"

const BUNGEE_SHADE_FONT = Bungee_Shade({ subsets: ["latin"], weight: "400" })

export default function Login() {
  const [message, setMessage] = useState("")

  function handleLogin(event: any) {
    event.preventDefault()

    axios.post("/auth/login", {
      username: event.target.username.value,
      password: event.target.password.value
    })
      .then((res) => {
        setCookie("token", res.data.data);
        setMessage("")
        window.location.reload()
      })
      .catch((err) => {
        setMessage(err.response?.data?.message || "Something went wrong")
      })
  }

  return (
    <Container className="h-screen flex flex-col justify-center items-center">
      <p className={`text-4xl text-red-400 font-bold mb-12 ${BUNGEE_SHADE_FONT.className}`}>
        Upperfile
      </p>
      {message && (
        <Alert className="w-[400px] mb-4" onClose={() => setMessage("")}>
          {message}
        </Alert>
      )}
      <form
        onSubmit={handleLogin}
        className="w-[400px] flex flex-col justify-center items-center rounded-lg bg-neutral-800 border border-neutral-600 gap-3 px-4 py-6"
      >
        <div className="w-full">
          <label htmlFor="username">Username</label>
          <InputForm id="username" type="username" className="mt-1" />
        </div>
        <div className="w-full">
          <label htmlFor="password">Password</label>
          <InputForm id="password" type="password" className="mt-1" />
        </div>
        <Button type="submit" className="w-full">Login</Button>
      </form>
    </Container>
  )
}
