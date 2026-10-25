package tmdb

import "fmt"

const imageBaseURL = "https://image.tmdb.org/t/p"

func ImageURL(path *string, size string) string {
	if path == nil || *path == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s%s", imageBaseURL, size, *path)
}

func PosterURL(path *string) string {
	return ImageURL(path, "w500")
}

func BackdropURL(path *string) string {
	return ImageURL(path, "w1280")
}
