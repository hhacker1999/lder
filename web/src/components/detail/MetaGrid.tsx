export type MetaItem = {
  label: string
  value: React.ReactNode
}

export function MetaGrid({ items }: { items: MetaItem[] }) {
  const visible = items.filter((i) => i.value !== "" && i.value != null && i.value !== "—")
  if (visible.length === 0) return null

  return (
    <dl className="grid grid-cols-2 gap-3 sm:grid-cols-3 lg:grid-cols-4">
      {visible.map((item) => (
        <div
          key={item.label}
          className="rounded-2xl border border-[var(--color-border)] bg-white/[0.03] px-4 py-3"
        >
          <dt className="text-xs font-medium uppercase tracking-wide text-[var(--color-muted)]">
            {item.label}
          </dt>
          <dd className="mt-1 text-sm font-medium leading-snug">{item.value}</dd>
        </div>
      ))}
    </dl>
  )
}
