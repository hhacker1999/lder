import { useQuery } from "@tanstack/react-query"
import { Search } from "lucide-react"
import { useEffect, useState } from "react"
import { useNavigate, useSearchParams } from "react-router-dom"
import { MediaGrid } from "@/components/MediaGrid"
import { Input } from "@/components/ui/input"
import { searchMedia } from "@/lib/api"

export function SearchPage() {
  const [params] = useSearchParams()
  const navigate = useNavigate()
  const q = params.get("q")?.trim() ?? ""
  const [draft, setDraft] = useState(q)

  useEffect(() => {
    setDraft(q)
  }, [q])

  useEffect(() => {
    if (!draft.trim()) return
    const t = setTimeout(() => {
      const next = draft.trim()
      if (next !== q) {
        navigate(next ? `/?q=${encodeURIComponent(next)}` : "/", { replace: true })
      }
    }, 300)
    return () => clearTimeout(t)
  }, [draft, q, navigate])

  const { data, isLoading, isFetching, error } = useQuery({
    queryKey: ["search", q],
    queryFn: () => searchMedia(q),
    enabled: q.length > 0,
  })

  const items = [...(data?.movies ?? []), ...(data?.shows ?? [])]

  return (
    <div className="mx-auto max-w-6xl px-4 pb-16 pt-10 md:px-6">
      <div className="mx-auto max-w-2xl text-center">
        <h1 className="text-4xl font-semibold tracking-tight md:text-5xl">
          Discover what to watch next
        </h1>
        <p className="mt-3 text-[var(--color-muted)]">
          Search movies and series in one place.
        </p>
        <div className="relative mt-8">
          <Search
            className="pointer-events-none absolute left-4 top-1/2 size-5 -translate-y-1/2 text-[var(--color-muted)]"
          />
          <Input
            value={draft}
            onChange={(e) => setDraft(e.target.value)}
            placeholder="Try Breaking Bad, Dune, Inception…"
            className="h-14 pl-12 text-lg"
            autoFocus
          />
        </div>
      </div>

      {q && (
        <div className="mt-12 space-y-6">
          <div className="flex items-end justify-between gap-4">
            <h2 className="text-lg font-medium">
              Results for <span className="text-[var(--color-accent)]">{q}</span>
            </h2>
            {data && (
              <span className="text-sm text-[var(--color-muted)]">
                {items.length} titles
              </span>
            )}
          </div>
          {error && (
            <p className="rounded-2xl border border-red-500/30 bg-red-500/10 px-4 py-3 text-sm text-red-200">
              {(error as Error).message}
            </p>
          )}
          <MediaGrid items={items} loading={isLoading || isFetching} />
        </div>
      )}

      {!q && (
        <p className="mt-16 text-center text-sm text-[var(--color-muted)]">
          Start typing to search TMDB for movies and TV shows.
        </p>
      )}
    </div>
  )
}
