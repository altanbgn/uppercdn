import type { InputHTMLAttributes } from "react"
import { twMerge } from "tailwind-merge"

export default function InputForm(props: InputHTMLAttributes<HTMLInputElement>) {
  return (
    <input
      {...props}
      className={twMerge(
        "w-full border border-neutral-600 focus:border-red-500 outline-none bg-transparent rounded-lg p-2 transition duration-200",
        props.className
      )}
      autoComplete="off"
    />
  )
}
