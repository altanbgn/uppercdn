import type { ButtonHTMLAttributes } from "react"
import { twMerge } from "tailwind-merge"

export default function Button({ children, ...props }: ButtonHTMLAttributes<HTMLButtonElement>) {
  return (
    <button
      {...props}
      className={twMerge(
        "bg-red-700 hover:bg-red-600 text-white font-bold px-4 py-2 rounded-lg transition duration-200",
        props.className
      )}
    >
      {children}
    </button>
  )
}
