import { useQuery } from "@tanstack/react-query"
import { Clock, ExternalLink, Star } from "lucide-react"
import { Link, useParams } from "react-router-dom"
import { BackButton } from "@/components/BackButton"
import { CastGrid } from "@/components/detail/CastGrid"
import { DetailSection } from "@/components/detail/DetailSection"
import { MetaGrid } from "@/components/detail/MetaGrid"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Skeleton } from "@/components/ui/skeleton"
import { getMovie } from "@/lib/api"
import { formatCurrency, formatDate, formatLanguage, formatRuntime } from "@/lib/format"

export function MoviePage() {
  const { id } = useParams()
  const movieId = Number(id)

  const { data, isLoading, error } = useQuery({
    queryKey: ["movie", movieId],
    queryFn: () => getMovie(movieId),
    enabled: Number.isFinite(movieId) && movieId > 0,
  })

  if (isLoading) {
    return <DetailSkeleton />
  }

  if (error || !data) {
    return (
      <div className="mx-auto max-w-3xl px-4 py-16 text-center">
        <p className="text-red-200">{(error as Error)?.message ?? "Not found"}</p>
        <Button variant="ghost" className="mt-4" asChild>
          <Link to="/">Home</Link>
        </Button>
      </div>
    )
  }

  const imdbUrl = data.imdbId ? `https://www.imdb.com/title/${data.imdbId}/` : null

  return (
    <div>
      <div className="relative min-h-[50vh] overflow-hidden">
        {data.backdropUrl && (
          <img
            src={data.backdropUrl}
            alt=""
            className="absolute inset-0 h-full w-full object-cover"
          />
        )}
        <div className="absolute inset-0 bg-gradient-to-t from-[var(--color-background)] via-[var(--color-background)]/85 to-black/40" />
        <div className="relative mx-auto flex max-w-6xl flex-col gap-6 px-4 pb-10 pt-8 md:px-6">
          <BackButton />
          <div className="flex flex-1 flex-col gap-6 md:flex-row md:items-end">
            <img
              src={data.posterUrl || data.backdropUrl}
              alt={data.title}
              className="w-40 shrink-0 rounded-2xl border border-white/10 shadow-2xl md:w-52"
            />
            <div className="space-y-3 pb-2">
              <Badge>Movie</Badge>
              <h1 className="text-3xl font-semibold tracking-tight md:text-5xl">{data.title}</h1>
              {data.originalTitle && data.originalTitle !== data.title && (
                <p className="text-sm text-[var(--color-muted)]">{data.originalTitle}</p>
              )}
              {data.tagline && (
                <p className="text-[var(--color-muted)] italic">{data.tagline}</p>
              )}
              <div className="flex flex-wrap gap-3 text-sm text-[var(--color-muted)]">
                <span className="inline-flex items-center gap-1">
                  <Star className="size-4 text-[var(--color-accent)]" />
                  {data.voteAverage.toFixed(1)}
                  <span className="text-xs">({data.voteCount.toLocaleString()} votes)</span>
                </span>
                {data.runtime > 0 && (
                  <span className="inline-flex items-center gap-1">
                    <Clock className="size-4" />
                    {formatRuntime(data.runtime)}
                  </span>
                )}
                {data.releaseDate && <span>{formatDate(data.releaseDate)}</span>}
              </div>
              <div className="flex flex-wrap gap-2">
                {data.genres.map((g) => (
                  <Badge key={g} className="border-0 bg-[var(--color-accent-soft)]">{g}</Badge>
                ))}
              </div>
              <div className="flex flex-wrap gap-2 pt-1">
                {imdbUrl && (
                  <Button variant="outline" size="sm" asChild>
                    <a href={imdbUrl} target="_blank" rel="noreferrer">
                      IMDb <ExternalLink className="size-3" />
                    </a>
                  </Button>
                )}
                {data.homepage && (
                  <Button variant="outline" size="sm" asChild>
                    <a href={data.homepage} target="_blank" rel="noreferrer">
                      Website <ExternalLink className="size-3" />
                    </a>
                  </Button>
                )}
              </div>
            </div>
          </div>
        </div>
      </div>

      <div className="mx-auto max-w-6xl space-y-10 px-4 py-10 md:px-6">
        <DetailSection title="Details">
          <MetaGrid
            items={[
              { label: "Release date", value: formatDate(data.releaseDate) },
              { label: "Runtime", value: formatRuntime(data.runtime) },
              { label: "Status", value: data.status || "—" },
              { label: "Language", value: formatLanguage(data.originalLanguage) },
              { label: "Budget", value: formatCurrency(data.budget) },
              { label: "Revenue", value: formatCurrency(data.revenue) },
              {
                label: "Director",
                value: data.directors.length > 0 ? data.directors.join(", ") : "—",
              },
            ]}
          />
        </DetailSection>

        <DetailSection title="Overview">
          <p className="max-w-3xl leading-relaxed text-[var(--color-muted)]">
            {data.overview || "No overview available."}
          </p>
        </DetailSection>

        {data.cast.length > 0 && (
          <DetailSection title="Cast">
            <CastGrid cast={data.cast} />
          </DetailSection>
        )}

        {data.productionCompanies.length > 0 && (
          <DetailSection title="Production">
            <p className="text-sm text-[var(--color-muted)]">
              {data.productionCompanies.join(" · ")}
            </p>
          </DetailSection>
        )}
      </div>
    </div>
  )
}

function DetailSkeleton() {
  return (
    <div className="mx-auto max-w-6xl space-y-6 px-4 py-10 md:px-6">
      <Skeleton className="h-[40vh] w-full" />
      <Skeleton className="h-6 w-1/3" />
      <Skeleton className="h-24 w-full" />
    </div>
  )
}
