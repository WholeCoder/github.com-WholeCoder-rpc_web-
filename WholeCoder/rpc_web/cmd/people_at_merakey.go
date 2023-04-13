package cmd

import (
	"fmt"
	"sort"
	"strings"
)

type Person2 struct {
	name                 string
	great_characteristic string
}

type DeservesLove2 interface {
	getTheyDeserveLove() string
}

func (p *Person2) getTheyDeserveLove() string {
	return p.name + " Deserves Love"
}

func populateDeservesLove(people_at_merakey []Person2) []DeservesLove2 {
	dl := []DeservesLove2{}
	for _, p := range people_at_merakey {
		p2 := p
		dl = append(dl, DeservesLove2(&p2))
	}

	return dl
}

func mainPeopleAtMerakey() {
	people_at_merakey := []Person2{Person2{name: "Thomas", great_characteristic: "Great Potential and has a Great Helpful Attitude"},
		Person2{name: "Ally", great_characteristic: "Great Christian Lady"},
		Person2{name: "Justin", great_characteristic: "Inspiring Great Kinistetic Learner."},
		Person2{name: "Shandel", great_characteristic: "Great Lady with Great Hair and Great Attitude"},
		Person2{name: "John", great_characteristic: "Mentally and Physically Strong Poet"},
		Person2{name: "Sam", great_characteristic: "Brilliant Free Spirited Piano Freestyler"},
		Person2{name: "Pradeep", great_characteristic: "Free-Spirit Intelligent and Intuitive"},
		Person2{name: "Andrew", great_characteristic: "Freindly Guy At the Top of His Game"},
		Person2{name: "Emma", great_characteristic: "Great Connection Builder and Great Personality"},
		Person2{name: "Mike", great_characteristic: "Very Intuitive With Lots of Real World Experience"},
		Person2{name: "Nate", great_characteristic: "Great Insight/Intuition On The Group Subjects - Great Guy"},
		Person2{name: "Quan", great_characteristic: "Sensitive Awesome Wresler with a Heart of Gold"},
		Person2{name: "Ruben", great_characteristic: "Caring, Wants to Help People With His Awesome Skills"},
	}
	input := ""
	direction := "d" // will be "u" or "d"
	for input != "q" {
		if input == "n" || input == "w" {
			sort.Slice(people_at_merakey, func(i, j int) bool {
				if direction == "u" {
					return people_at_merakey[i].name < people_at_merakey[j].name
				}
				return people_at_merakey[i].name > people_at_merakey[j].name
			})
		} else if input == "g" {
			sort.Slice(people_at_merakey, func(i, j int) bool {
				if direction == "u" {
					return people_at_merakey[i].great_characteristic < people_at_merakey[j].great_characteristic
				}
				return people_at_merakey[i].great_characteristic > people_at_merakey[j].great_characteristic
			})
		}

		dl := populateDeservesLove(people_at_merakey)

		fmt.Printf("%*s|%*s|%*s\n", 10, "Name", 100, "Great Qualities", 25, "What they Deserve")
		fmt.Printf("-----------------------------------------------------------------------------------------------------------------------------------------\n")
		for i, p := range people_at_merakey {
			fmt.Printf("%*s|%*s|%*s\n", 10, p.name, 100, p.great_characteristic, 25, dl[i].getTheyDeserveLove())
		}

		fmt.Print("Sort By (n, g, q, or w)>  ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if direction == "u" {
			direction = "d"
		} else {
			direction = "u"
		}
		fmt.Println("direction =", direction)

	}
}
