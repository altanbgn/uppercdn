import { twMerge } from "tailwind-merge"

type Props = {
  children: React.ReactNode
  className?: string
}

export default function Container({ children, className }: Props) {
  const mergedClassName = twMerge(
    "container flex justify-between items-center px-8 lg:p-16 mx-auto",
    className
  )

  return (
    <div className={mergedClassName}>{children}</div>
  )
}
