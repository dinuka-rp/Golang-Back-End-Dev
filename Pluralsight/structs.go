package main

import "fmt"

func main(){
	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}


	// 2 ways below initializes with zero values
	//var DockerDeepDive courseMeta

	//DockerDeepDive := new(courseMeta)		// gives a pointer

	// composite literal way
	DockerDeepDive := courseMeta{
		Author: "Nigel Poulton",
		Level:  "Intermediate",
		Rating: 5,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("\nDocker Deep Dive author is:", DockerDeepDive.Author)

	// modifying field values in structs
	DockerDeepDive.Rating=1
	fmt.Println(DockerDeepDive)
}
