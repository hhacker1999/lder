export function DetailSection({
  title,
  children,
}: {
  title: string
  children: React.ReactNode
}) {
  return (
    <section className="space-y-4">
      <h2 className="text-lg font-medium tracking-tight">{title}</h2>
      {children}
    </section>
  )
}
