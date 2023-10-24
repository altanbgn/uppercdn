import { type BaseHTMLAttributes } from "react"
import { twMerge } from "tailwind-merge"

type Props = BaseHTMLAttributes<HTMLDivElement> & {
  onClose?: () => void
}

export default function Alert({
  className,
  children,
  onClose,
  ...props
}: Props) {
  return (
    <div
      className={twMerge(
        "flex justify-between items-center bg-red-400 text-white font-bold px-4 py-2 rounded-lg transition duration-200",
        className
      )}
      {...props}
    >
      {children}
      <button onClick={onClose}>X</button>
    </div>
  )
}
