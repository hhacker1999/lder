import { useQuery } from "@tanstack/react-query"
import { Clock, ExternalLink, Star, Tv } from "lucide-react"
import { Link, useParams } from "react-router-dom"
import { BackButton } from "@/components/BackButton"
import { CastGrid } from "@/components/detail/CastGrid"
import { DetailSection } from "@/components/detail/DetailSection"
import { MetaGrid } from "@/components/detail/MetaGrid"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Skeleton } from "@/components/ui/skeleton"
import { getShow } from "@/lib/api"
import { formatDate, formatLanguage } from "@/lib/format"

function episodeRuntimeLabel(data: {
  averageRuntime: number | null
  episodeRunTime: number[]
}) {
  if (data.averageRuntime != null && data.averageRuntime > 0) {
    return `${data.averageRuntime} min / episode`
  }
  if (data.episodeRunTime.length > 0) {
    const avg = Math.round(
      data.episodeRunTime.reduce((a, b) => a + b, 0) / data.episodeRunTime.length,
    )
    return `${avg} min / episode`
  }
  return "—"
}

export function ShowPage() {
  const { id } = useParams()
  const tmdbId = Number(id)

  const { data, isLoading, error } = useQuery({
    queryKey: ["show", tmdbId],
    queryFn: () => getShow(tmdbId),
    enabled: Number.isFinite(tmdbId) && tmdbId > 0,
  })

  if (isLoading) {
    return (
      <div className="mx-auto max-w-6xl space-y-6 px-4 py-10 md:px-6">
        <Skeleton className="h-[40vh] w-full" />
        <Skeleton className="h-6 w-1/3" />
        <Skeleton className="h-24 w-full" />
      </div>
    )
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

  const hero = data.backdropUrl || data.imageUrl
  const imdbUrl = data.imdbId ? `https://www.imdb.com/title/${data.imdbId}/` : null
  const seriesStatus = data.status || data.tmdbStatus

  return (
    <div>
      <div className="relative min-h-[50vh] overflow-hidden">
        {hero && (
          <img src={hero} alt="" className="absolute inset-0 h-full w-full object-cover" />
        )}
        <div className="absolute inset-0 bg-gradient-to-t from-[var(--color-background)] via-[var(--color-background)]/85 to-black/40" />
        <div className="relative mx-auto flex max-w-6xl flex-col gap-6 px-4 pb-10 pt-8 md:flex-row md:items-end md:px-6">
          <BackButton />
          <div className="flex flex-1 flex-col gap-6 md:flex-row md:items-end">
            {data.imageUrl && (
              <img
                src={data.imageUrl}
                alt={data.name}
                className="w-40 shrink-0 rounded-2xl border border-white/10 object-cover shadow-2xl md:w-52"
              />
            )}
            <div className="space-y-3 pb-2">
              <Badge>Series</Badge>
              <h1 className="text-3xl font-semibold tracking-tight md:text-5xl">{data.name}</h1>
              {data.originalName && data.originalName !== data.name && (
                <p className="text-sm text-[var(--color-muted)]">{data.originalName}</p>
              )}
              <div className="flex flex-wrap gap-3 text-sm text-[var(--color-muted)]">
                {data.tmdbVoteAverage > 0 && (
                  <span className="inline-flex items-center gap-1">
                    <Star className="size-4 text-[var(--color-accent)]" />
                    {data.tmdbVoteAverage.toFixed(1)}
                    <span className="text-xs">({data.tmdbVoteCount.toLocaleString()} votes)</span>
                  </span>
                )}
                {data.score > 0 && (
                  <span className="inline-flex items-center gap-1">
                    TVDB {data.score.toFixed(1)}
                  </span>
                )}
                <span className="inline-flex items-center gap-1">
                  <Tv className="size-4" />
                  {data.numberOfSeasons} seasons · {data.numberOfEpisodes} episodes
                </span>
                <span className="inline-flex items-center gap-1">
                  <Clock className="size-4" />
                  {episodeRuntimeLabel(data)}
                </span>
              </div>
              <div className="flex flex-wrap gap-2">
                {data.genres.map((g) => (
                  <Badge key={g} className="border-0 bg-[var(--color-accent-soft)]">{g}</Badge>
                ))}
              </div>
              {imdbUrl && (
                <Button variant="outline" size="sm" asChild>
                  <a href={imdbUrl} target="_blank" rel="noreferrer">
                    IMDb <ExternalLink className="size-3" />
                  </a>
                </Button>
              )}
            </div>
          </div>
        </div>
      </div>

      <div className="mx-auto max-w-6xl space-y-10 px-4 py-10 md:px-6">
        <DetailSection title="Details">
          <MetaGrid
            items={[
              { label: "Series status", value: seriesStatus || "—" },
              { label: "First aired", value: formatDate(data.firstAired) },
              { label: "Last aired", value: formatDate(data.lastAired) },
              { label: "Next episode", value: formatDate(data.nextAired) },
              { label: "Premiere year", value: data.year || "—" },
              { label: "Seasons", value: data.numberOfSeasons > 0 ? data.numberOfSeasons : "—" },
              { label: "Episodes", value: data.numberOfEpisodes > 0 ? data.numberOfEpisodes : "—" },
              { label: "Episode runtime", value: episodeRuntimeLabel(data) },
              { label: "Country", value: data.originalCountry || "—" },
              { label: "Language", value: formatLanguage(data.originalLanguage) },
              { label: "Type", value: data.showType || "—" },
              {
                label: "Created by",
                value: data.createdBy.length > 0 ? data.createdBy.join(", ") : "—",
              },
              {
                label: "Network",
                value: data.networks.length > 0 ? data.networks.join(", ") : "—",
              },
              {
                label: "Last updated",
                value: data.lastUpdated ? formatDate(data.lastUpdated) : "—",
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
      </div>
    </div>
  )
}
