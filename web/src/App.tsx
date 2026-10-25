import { Route, Routes } from "react-router-dom"
import { AppShell } from "@/components/layout/AppShell"
import { MoviePage } from "@/pages/MoviePage"
import { SearchPage } from "@/pages/SearchPage"
import { ShowPage } from "@/pages/ShowPage"

export default function App() {
  return (
    <AppShell>
      <Routes>
        <Route path="/" element={<SearchPage />} />
        <Route path="/movie/:id" element={<MoviePage />} />
        <Route path="/show/:id" element={<ShowPage />} />
      </Routes>
    </AppShell>
  )
}
