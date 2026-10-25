import type { CastMember } from "@/lib/api"

export function CastGrid({ cast }: { cast: CastMember[] }) {
  if (cast.length === 0) return null

  return (
    <div className="flex gap-4 overflow-x-auto pb-2 [-ms-overflow-style:none] [scrollbar-width:none] [&::-webkit-scrollbar]:hidden">
      {cast.map((person) => (
        <div key={person.id} className="w-28 shrink-0 text-center">
          <div className="mx-auto aspect-square w-24 overflow-hidden rounded-2xl border border-[var(--color-border)] bg-white/5">
            {person.profileUrl ? (
              <img
                src={person.profileUrl}
                alt={person.name}
                className="h-full w-full object-cover"
                loading="lazy"
              />
            ) : (
              <div className="flex h-full items-center justify-center text-xs text-[var(--color-muted)]">
                ?
              </div>
            )}
          </div>
          <p className="mt-2 line-clamp-2 text-sm font-medium leading-tight">{person.name}</p>
          <p className="mt-0.5 line-clamp-2 text-xs text-[var(--color-muted)]">{person.role}</p>
        </div>
      ))}
    </div>
  )
}
