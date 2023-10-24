import Container from "@/components/Container"
import Section from "@/components/Section"

export default function Projects() {
  return (
    <Container className="flex gap-8">
      <Section className="w-full">
        <p className="text-lg font-bold mb-4">Projects</p>
        <p>Pog</p>
      </Section>
      <Section className="w-full lg:max-w-[400px]">
        <p className="text-lg font-bold mb-4">Change Logs</p>
      </Section>
    </Container>
  )
}
