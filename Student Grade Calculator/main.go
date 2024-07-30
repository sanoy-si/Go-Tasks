package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin) 

func recieveInt(inputText string) int64{
	fmt.Print(inputText)
	scanner.Scan()
	input := scanner.Text()
	value, err := strconv.ParseInt(input, 10, 64)

	for err != nil{
		fmt.Print("Please Enter a vaild integer input: ")
		scanner.Scan()
		input := scanner.Text()
		value, err = strconv.ParseInt(input, 10, 32)

	}

	return value
}

func recieveFloat(inputText string) float64{
	fmt.Print(inputText)
	scanner.Scan()
	input := scanner.Text()
	value, err := strconv.ParseFloat(input, 32)
 
	for err != nil{
		fmt.Print("Please Enter a vaild foat input: ")
		scanner.Scan()
		input := scanner.Text()
		value, err = strconv.ParseFloat(input, 32)

	}

	return value
}


func validateScore(score float64) string{
	if score > 100{
		return "A score cannot be greater than a 100. Please enter a valid score: "
	}else if score < 0{
		return "A score cannot be less than 0. Please enter a valid score: "
	}else{
		return "Course added successfully."
	}
}

func calculatTotal(scores map[string]float64) float64{
	var total float64
	for _, score := range scores{
		total += score
	}

	return total
}

func calculateAverage(total float64, numberOfCourses int) float64{
	return total / float64(numberOfCourses)
}

func calculateGrade(averageScore float64) string{
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
		name string
		numberOfCourses int
		scores = map[string]float64{}
	)

	

	fmt.Print("Enter Your Name: ")
	scanner.Scan()
	name = scanner.Text()

	numberOfCourses = int(recieveInt(fmt.Sprintf("\nHello %v, How many courses did you take?\n", name)))

	for i := 1; i <= numberOfCourses; i++{
		var courseName string
		var score float64

		fmt.Printf("\nWhat is the name of the #%v course you took?\n", i)
		scanner.Scan()
		courseName = scanner.Text()

		value, exist := scores[courseName]
		for exist {
			fmt.Printf("\nCourse name %v already added with a score of %v. Please Enter another course.\n", courseName, value) 
			scanner.Scan()
			courseName = scanner.Text()
			value, exist = scores[courseName]
		}

		score = recieveFloat(fmt.Sprintf("\nHow much did you get in %v?(Out of 100)\n", courseName))
		
		message := validateScore(score)
		for message != "Course added successfully."{
			score = recieveFloat(message)
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