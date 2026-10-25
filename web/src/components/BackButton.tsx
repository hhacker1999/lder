import { ArrowLeft } from "lucide-react"
import { useLocation, useNavigate } from "react-router-dom"
import { Button } from "@/components/ui/button"

type LocationState = {
  from?: string
}

export function BackButton() {
  const navigate = useNavigate()
  const location = useLocation()
  const from = (location.state as LocationState | null)?.from

  function onBack() {
    if (from) {
      navigate(from)
      return
    }
    if (window.history.length > 1) {
      navigate(-1)
      return
    }
    navigate("/")
  }

  return (
    <Button variant="ghost" size="sm" className="w-fit" type="button" onClick={onBack}>
      <ArrowLeft className="size-4" /> Back
    </Button>
  )
}
