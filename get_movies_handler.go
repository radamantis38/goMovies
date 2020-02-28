package main

import (
	"fmt"
	"peliculas/usecases"
)

type getMoviesInterface interface {
	Handle()  (string,error)
}


func main() {
 uc := usecases.NewGetMovies()
 movies, err := uc.Handle()
 if err != nil{
	 fmt.Println("Error")
 }
	fmt.Println(movies)
}
