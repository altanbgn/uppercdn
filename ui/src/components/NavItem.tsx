import { createElement } from "react"
import Link from "next/link"
import { twMerge } from "tailwind-merge"

type NavItemProps = {
  className?: string
  href: string
  icon: React.ComponentType<React.SVGProps<SVGSVGElement>>
  children?: React.ReactNode
  onClick?: React.MouseEventHandler<HTMLAnchorElement>
}

export default function NavItem({ className, href, icon, children, ...props }: NavItemProps) {
  return (
    <Link
      href={href}
      className={twMerge(
        "flex justify-start items-center group w-full text-left hover:text-red-400 hover:bg-neutral-800 rounded-lg px-4 py-2 gap-4 transition duration-200",
        className
      )}
      {...props}
    >
      {createElement(icon, { className: "w-4 h-4 fill-white group-hover:fill-red-400 transition duration-200" })}
      {children}
    </Link>
  )
}
