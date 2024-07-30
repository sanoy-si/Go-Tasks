package main

import "fmt"


func validateScore(score float32) string{
	if score > 100{
		return "A score cannot be greater than a 100. Please enter a valid score."
	}else if score < 0{
		return "A score cannot be less than 0. Please enter a valid score."
	}else{
		return "Course added successfully."
	}
}

func calculatTotal(scores map[string]float32) float32{
	var total float32
	for _, score := range scores{
		total += score
	}

	return total
}

func calculateAverage(total float32, numberOfCourses int) float32{
	return total / float32(numberOfCourses)
}

func calculateGrade(averageScore float32) string{
	switch {
		case averageScore >= 90:
			return "A+"

		case averageScore >= 85:
			return "A"
		
		case averageScore >= 80:
			return "A-"
		
		case averageScore >= 75:
			return "B+"
		
		case averageScore >= 70:
			return "B"
		
		case averageScore >= 65:
			return "B-"
		
		case averageScore >= 60:
			return "C+"
		
		case averageScore >= 50:
			return "C"
		
		case averageScore >= 45:
			return "C-"
		
		case averageScore >= 40:
			return "D"
		
		default:
			return "F"

	}
}

func main(){
	fmt.Println("Welcome To Student Grade Calculator")
	
	var(
		firstName string
		lastName string
		numberOfCourses int
		scores = map[string]float32{}
	)
	

	fmt.Println("Enter Your First Name:")
	fmt.Scan(&firstName)

	fmt.Println("\nEnter Your Last Name:")
	fmt.Scan(&lastName)

	fmt.Printf("\nHello %v, How many courses did you take?\n", firstName)
	fmt.Scan(&numberOfCourses)

	for i := 1; i <= numberOfCourses; i++{
		var courseName string
		var score float32

		fmt.Printf("\nWhat is the name of the #%v course you took?\n", i)
		fmt.Scan((&courseName))

		value, exist := scores[courseName]
		for exist {
			fmt.Printf("\nCourse name %v already added with a score of %v. Please Enter another course.\n", courseName, value) 
			fmt.Scan((&courseName))
			value, exist = scores[courseName]
		}

		fmt.Printf("\nHow much did you get in %v?(Out of 100)\n", courseName)
		fmt.Scan(&score)
		
		message := validateScore(score)
		for message != "Course added successfully."{
			fmt.Println(message)
			fmt.Scan(&score)
			message = validateScore(score)
		} 
		scores[courseName] = score
		println(message, "\n")
	}

	fmt.Println("\n\n\n\n\n\n\n\n\\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println("Summary:\n")

	fmt.Println("The courses you took are:")
	for name, score:= range scores{
		fmt.Printf("%v,  score: %v\n", name, score)
	}
	fmt.Println()

	totalScore := calculatTotal(scores)
	averageScore := calculateAverage(totalScore, numberOfCourses)
	averageGrade := calculateGrade(averageScore)

	fmt.Printf("The total sum of your scores is: %v\n\n", totalScore)
	fmt.Printf("The Average of your scores is: %v\n\n", averageScore)
	fmt.Printf("The Average Grade that you got is: %v\n\n", averageGrade)

}