import { Link, useLocation, useNavigate, useSearchParams } from "react-router-dom"
import { Search } from "lucide-react"
import { useEffect, useState } from "react"
import { Input } from "@/components/ui/input"

export function AppShell({ children }: { children: React.ReactNode }) {
  const location = useLocation()
  const navigate = useNavigate()
  const [params] = useSearchParams()
  const isHome = location.pathname === "/"
  const [query, setQuery] = useState(params.get("q") ?? "")

  useEffect(() => {
    setQuery(params.get("q") ?? "")
  }, [params])

  function onSubmit(e: React.FormEvent) {
    e.preventDefault()
    const trimmed = query.trim()
    navigate(trimmed ? `/?q=${encodeURIComponent(trimmed)}` : "/")
  }

  return (
    <div className="min-h-screen mesh-bg">
      <header className="sticky top-0 z-50 border-b border-white/5 bg-[oklch(0.13_0.02_280/0.75)] backdrop-blur-xl">
        <div className="mx-auto flex max-w-6xl items-center gap-4 px-4 py-4 md:px-6">
          <Link to="/" className="shrink-0 text-xl font-semibold tracking-tight">
            lder
          </Link>
          {!isHome && (
            <form onSubmit={onSubmit} className="relative max-w-md flex-1">
              <Search
                className="pointer-events-none absolute left-3 top-1/2 size-4 -translate-y-1/2 text-[var(--color-muted)]"
              />
              <Input
                value={query}
                onChange={(e) => setQuery(e.target.value)}
                placeholder="Search movies and shows…"
                className="h-10 pl-10"
              />
            </form>
          )}
        </div>
      </header>
      <main>{children}</main>
    </div>
  )
}
