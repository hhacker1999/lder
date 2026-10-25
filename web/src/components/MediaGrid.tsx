import { MediaCard } from "@/components/MediaCard"
import { Skeleton } from "@/components/ui/skeleton"
import type { MediaSearchItem } from "@/lib/api"

export function MediaGrid({
  items,
  loading,
}: {
  items: MediaSearchItem[]
  loading?: boolean
}) {
  if (loading) {
    return (
      <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
        {Array.from({ length: 10 }).map((_, i) => (
          <div key={i} className="space-y-2">
            <Skeleton className="aspect-[2/3] w-full" />
            <Skeleton className="h-4 w-3/4" />
          </div>
        ))}
      </div>
    )
  }

  if (items.length === 0) {
    return (
      <p className="text-center text-[var(--color-muted)]">No results found.</p>
    )
  }

  return (
    <div className="grid grid-cols-2 gap-4 sm:grid-cols-3 md:grid-cols-4 lg:grid-cols-5">
      {items.map((item) => (
        <MediaCard key={`${item.mediaType}-${item.id}`} item={item} />
      ))}
    </div>
  )
}
