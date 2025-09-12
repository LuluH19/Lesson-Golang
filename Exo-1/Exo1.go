

package main

import "fmt"

func main() {
var (
    title, writer, artist, publisher, genre string
    year, pageNumber int
    grade float32
)

    // Premier comic
    title = "Mr. GoToSleep"
    writer = "Tracey Hatchet"
    artist = "Jewel Tampson"
    publisher = "DizzyBooks Publishing Inc."
    year = 1997
    pageNumber = 14
    grade = 6.5
    genre = "Horror"

    fmt.Println(title, "-", genre, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages. Grade:", grade)

    // Calcul du prix estimé
    const currentYear = 2025
    age := currentYear - year
    price := float32(age)*0.5 + float32(pageNumber)*0.05 + grade*2.0
    fmt.Println("Estimated price:", price, "€")

    fmt.Println()

    // Deuxième comic
    title = "Epic Vol. 1"
    writer = "Ryan N. Shawn"
    artist = "Phoebe Paperclips"
    publisher = "DizzyBooks Publishing Inc."
    year = 2013
    pageNumber = 160
    grade = 9.0
    genre = "Fantasy"

    fmt.Println(title, "-", genre, "written by", writer, "drawn by", artist, "published by", publisher, "in", year, "with", pageNumber, "pages. Grade:", grade)

    age = 2025 - int(year)
    price = float32(age)*0.5 + float32(pageNumber)*0.05 + grade*2.0
    fmt.Println("Estimated price:", price, "€")
}
