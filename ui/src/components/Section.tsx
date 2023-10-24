import { twMerge } from "tailwind-merge"

type Props = React.BaseHTMLAttributes<HTMLDivElement>

export default function Section({ children, className, ...props }: Props) {
  return (
    <div
      className={twMerge(
        "border border-neutral-600 rounded-lg bg-transparent p-4",
        className
      )}
      {...props}
    >
      {children}
    </div>
  )
}
