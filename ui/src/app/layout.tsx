import type { Metadata } from "next"
import { Open_Sans } from "next/font/google"
import "./globals.css"

import Navbar from "@/components/Navbar"

const OPEN_SANS_FONT = Open_Sans({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "Upperfile",
  description: "Dashboard - Upperfile",
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={`flex justify-between items-start ${OPEN_SANS_FONT.className}`}>
        <Navbar />
        <main className="w-full">{children}</main>
      </body>
    </html>
  )
}
