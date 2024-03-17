package main

import (
	"github.com/labstack/echo/v4"
	"github.com/rl404/verniy"
)

var v *verniy.Client

func main() {
	e := echo.New()
	v = verniy.New()
	e.GET("/search", searchAnime)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

//package main
//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/rl404/verniy"
//)
//
//// Handle *string to string.
//func ptrToString(str *string) string {
//	if str == nil {
//		return ""
//	}
//	return *str
//}
//
//func main() {
//	v := verniy.New()
//
//	// Get anime One Piece data (with default fields).
//
//	// d, err := c.SearchAnime(verniy.PageParamMedia{
//	// 	Season:     verniy.SeasonWinter,
//	// 	SeasonYear: 2021,
//	// 	TagIn: []string{"Action"},
//	// }, 1, 10)
//	data, err := v.SearchAnime(verniy.PageParamMedia{
//		Search: "Yu Yu Hakusho",
//	}, 1, 50)
//	//data, err := v.GetAnime(21)
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	// Convert the data object to JSON.
//	jsonData, err := json.MarshalIndent(data, "", "  ")
//	if err != nil {
//		fmt.Println("Error marshalling to JSON:", err)
//		return
//	}
//
//	// Print the JSON data.
//	fmt.Println("Anime Data as JSON:")
//	fmt.Println(string(jsonData))
//
//	// Access specific fields using the pointer handler.
//	// fmt.Println("\nAnime Data Fields:")
//	// fmt.Printf("Title (Romaji): %s\n", ptrToString(data.Title.Romaji))
//	// fmt.Printf("Title (English): %s\n", ptrToString(data.Title.English))
//	// fmt.Printf("Title (Native): %s\n", ptrToString(data.Title.Native))
//}
