import Link from "next/link"
import Image from "next/image"
import { Bungee_Shade } from "next/font/google"

// Local
import Container from "@/components/Container"

const BungeeShade = Bungee_Shade({ subsets: ["latin"], weight: "400" })

export default function Navbar() {
  return (
    <nav>
      <Container className="h-28">
        <Link
          href="/"
          className={`text-base lg:text-3xl text-red-400 font-bold ${BungeeShade.className}`}
        >
          Upperfile
        </Link>
        <div className="hidden lg:flex justify-end items-center gap-10">
          <Link href="https://github.com/altanbgn/upperfile.com">
            <Image
              className="invert"
              src="/github.svg"
              alt="Github"
              width={24}
              height={24}
              priority
            />
          </Link>
          <Link href="https://discord.com">
            <Image
              className="invert"
              src="/discord.svg"
              alt="Github"
              width={24}
              height={24}
              priority
            />
          </Link>
        </div>
      </Container>
    </nav>
  )
}
