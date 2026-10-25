import { Link, useLocation } from "react-router-dom"
import { Star } from "lucide-react"
import { Badge } from "@/components/ui/badge"
import type { MediaSearchItem } from "@/lib/api"

export function MediaCard({ item }: { item: MediaSearchItem }) {
  const location = useLocation()
  const href = item.mediaType === "movie" ? `/movie/${item.id}` : `/show/${item.id}`
  const from = `${location.pathname}${location.search}`

  return (
    <Link
      to={href}
      state={{ from }}
      className="group block overflow-hidden rounded-2xl border border-[var(--color-border)] bg-[var(--color-card)] transition duration-300 hover:-translate-y-1 hover:border-[var(--color-accent)]/40 hover:bg-[var(--color-card-hover)] hover:shadow-[0_20px_50px_-20px_oklch(0.5_0.2_285/0.5)]"
    >
      <div className="relative aspect-[2/3] overflow-hidden bg-white/5">
        {item.posterUrl ? (
          <img
            src={item.posterUrl}
            alt={item.title}
            className="h-full w-full object-cover transition duration-500 group-hover:scale-105"
            loading="lazy"
          />
        ) : (
          <div className="flex h-full items-center justify-center text-sm text-[var(--color-muted)]">
            No poster
          </div>
        )}
        <div className="absolute left-2 top-2">
          <Badge className="border-0 bg-black/50 backdrop-blur-sm">
            {item.mediaType === "movie" ? "Movie" : "Series"}
          </Badge>
        </div>
      </div>
      <div className="space-y-1 p-3">
        <h3 className="line-clamp-1 font-medium leading-snug">{item.title}</h3>
        <div className="flex items-center justify-between text-xs text-[var(--color-muted)]">
          <span>{item.year || "—"}</span>
          <span className="inline-flex items-center gap-1">
            <Star className="size-3 fill-[var(--color-accent)] text-[var(--color-accent)]" />
            {item.voteAverage.toFixed(1)}
          </span>
        </div>
      </div>
    </Link>
  )
}
