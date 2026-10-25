export type MediaSearchItem = {
  id: number
  title: string
  overview: string
  year: string
  posterUrl: string
  voteAverage: number
  mediaType: "movie" | "show"
}

export type SearchResponse = {
  movies: MediaSearchItem[]
  shows: MediaSearchItem[]
  page: number
  totalPages: number
}

export type CastMember = {
  id: number
  name: string
  role: string
  profileUrl: string
}

export type MovieDetail = {
  id: number
  title: string
  originalTitle: string
  overview: string
  tagline: string
  releaseDate: string
  runtime: number
  status: string
  posterUrl: string
  backdropUrl: string
  voteAverage: number
  voteCount: number
  genres: string[]
  imdbId: string
  homepage: string
  originalLanguage: string
  budget: number
  revenue: number
  productionCompanies: string[]
  directors: string[]
  cast: CastMember[]
}

export type ShowDetail = {
  id: number
  tmdbId: number
  tvdbId: number
  name: string
  originalName: string
  overview: string
  status: string
  tmdbStatus: string
  firstAired: string
  lastAired: string
  nextAired: string
  year: string
  score: number
  tmdbVoteAverage: number
  tmdbVoteCount: number
  imageUrl: string
  backdropUrl: string
  averageRuntime: number | null
  episodeRunTime: number[]
  numberOfSeasons: number
  numberOfEpisodes: number
  originalCountry: string
  originalLanguage: string
  imdbId: string
  genres: string[]
  networks: string[]
  createdBy: string[]
  lastUpdated: string
  showType: string
  cast: CastMember[]
}

async function getJSON<T>(path: string): Promise<T> {
  const res = await fetch(path)
  if (!res.ok) {
    const body = await res.json().catch(() => ({}))
    throw new Error((body as { error?: string }).error ?? res.statusText)
  }
  return res.json() as Promise<T>
}

export function searchMedia(q: string, page = 1) {
  const params = new URLSearchParams({ q, page: String(page) })
  return getJSON<SearchResponse>(`/api/search?${params}`)
}

export function getMovie(id: number) {
  return getJSON<MovieDetail>(`/api/movies/${id}`)
}

export function getShow(tmdbId: number) {
  return getJSON<ShowDetail>(`/api/shows/${tmdbId}`)
}
