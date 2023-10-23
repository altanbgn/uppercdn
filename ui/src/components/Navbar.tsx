import React from "react"
import Link from "next/link"
import { Bungee_Shade } from "next/font/google"
import { twMerge } from "tailwind-merge"

import DashboardIcon from "@/components/Icons/Dashboard"
import ProjectsIcon from "@/components/Icons/Projects"
import UsersIcon from "@/components/Icons/Users"
import HistoryIcon from "@/components/Icons/History"
import SettingsIcon from "@/components/Icons/Settings"
import LogoutIcon from "@/components/Icons/Logout"

const BUNGEE_SHADE_FONT = Bungee_Shade({ subsets: ["latin"], weight: "400" })

type NavItemProps = {
  className?: string
  href: string
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>
  children?: React.ReactNode
}

function NavItem({ className, href, icon, children }: NavItemProps) {
  return (
    <Link
      href={href}
      className={twMerge(
        "flex justify-start items-center group gap-4 w-full text-left hover:text-red-400 hover:bg-neutral-700 rounded-xl p-4 transition duration-200",
        className
      )}
    >
      {React.createElement(icon, { className: "w-4 h-4 fill-white group-hover:fill-red-400 transition duration-200" })}
      {children}
    </Link>
  )
}

export default function Navbar() {
  return (
    <nav className="hidden lg:flex flex-col justify-start items-center w-[280px] h-screen bg-neutral-900 p-4 gap-2">
      <Link
        href="/"
        className={`text-base lg:text-2xl text-red-400 font-bold my-4 ${BUNGEE_SHADE_FONT.className}`}
      >
        Upperfile
      </Link>
      <NavItem icon={DashboardIcon} href="/">
        Dashboard
      </NavItem>
      <NavItem icon={ProjectsIcon} href="/projects">
        Projects
      </NavItem>
      <NavItem icon={UsersIcon} href="/users">
        Users
      </NavItem>
      <NavItem icon={HistoryIcon} href="/history">
        History
      </NavItem>
      <NavItem icon={SettingsIcon} href="/settings">
        Settings
      </NavItem>
      <NavItem icon={LogoutIcon} href="/logout" className="mt-auto">
        Logout
      </NavItem>
    </nav>
  )
}
