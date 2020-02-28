package main

import (
	"fmt"
	"peliculas/repositories"
	"peliculas/usecases"
)

type getMoviesInterface interface {
	Handle()  (string,error)
}

func Handle(moviesRepository getMoviesInterface){
	movies, err := moviesRepository.Handle()
	if err != nil{
		fmt.Println("Error")
	}
	fmt.Println(movies)
}

func main() {
 moviesRepository := repositories.NewMysqlMovie()
 usecase := usecases.NewGetMovies(moviesRepository)
 Handle(usecase)

}
