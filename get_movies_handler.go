package main

import (
	"fmt"
	"peliculas/repositories"
	"peliculas/usecases"
)

type getMoviesInterface interface {
	Handle()  (string,error)
}


func main() {
 moviesRepository := repositories.NewMysqlMovie()
 uc := usecases.NewGetMovies(moviesRepository)
 movies, err := uc.Handle()
 if err != nil{
	 fmt.Println("Error")
 }
	fmt.Println(movies)
}
