import type { Metadata } from "next"
import { Open_Sans } from "next/font/google"
import { cookies } from "next/headers"

import Navbar from "@/components/Navbar"
import Login from "./_views/Login"
import "./globals.css"

const OPEN_SANS_FONT = Open_Sans({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "Upperfile",
  description: "Dashboard - Upperfile",
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={`flex justify-between items-start ${OPEN_SANS_FONT.className}`}>
        {cookies().get("token")
          ? (
            <>
              <Navbar />
              <main className="w-full">{children}</main>
            </>
          ) : <Login />}
      </body>
    </html>
  )
}
