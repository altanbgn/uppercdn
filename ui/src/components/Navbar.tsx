"use client"

import Link from "next/link"
import { Bungee_Shade } from "next/font/google"

import NavItem from "@/components/NavItem"
import DashboardIcon from "@/components/Icons/Dashboard"
import ProjectsIcon from "@/components/Icons/Projects"
import UsersIcon from "@/components/Icons/Users"
import HistoryIcon from "@/components/Icons/History"
import SettingsIcon from "@/components/Icons/Settings"
import LogoutIcon from "@/components/Icons/Logout"
import { deleteCookie } from "@/utils/cookie"

const BUNGEE_SHADE_FONT = Bungee_Shade({ subsets: ["latin"], weight: "400" })

export default function Navbar() {
  function handleLogout() {
    deleteCookie("token")
    window.location.replace("/")
  }

  return (
    <nav className="hidden lg:flex flex-col justify-start items-center w-[280px] border-r border-neutral-600 h-screen bg-neutral-900 p-4 gap-1">
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
      <NavItem icon={LogoutIcon} href="#" onClick={handleLogout} className="mt-auto">
        Logout
      </NavItem>
    </nav>
  )
}
