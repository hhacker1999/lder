export function formatRuntime(minutes: number) {
  if (minutes <= 0) return "—"
  const h = Math.floor(minutes / 60)
  const m = minutes % 60
  if (h === 0) return `${m}m`
  if (m === 0) return `${h}h`
  return `${h}h ${m}m`
}

export function formatCurrency(amount: number) {
  if (amount <= 0) return "—"
  return new Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    maximumFractionDigits: 0,
  }).format(amount)
}

export function formatDate(date: string) {
  if (!date) return "—"
  const parsed = new Date(date.includes("T") ? date : `${date}T00:00:00`)
  if (Number.isNaN(parsed.getTime())) return date
  return parsed.toLocaleDateString(undefined, {
    year: "numeric",
    month: "short",
    day: "numeric",
  })
}

export function formatLanguage(code: string) {
  if (!code) return "—"
  try {
    const display = new Intl.DisplayNames(["en"], { type: "language" })
    return display.of(code) ?? code
  } catch {
    return code
  }
}
