package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct{
	Name string
	Grades map[string] float32
} 
func gradeCalculator(commulativeSum float32, numberOfSubject int) float32{
	return commulativeSum/float32(numberOfSubject) 
}

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Your Name:")
	name, _ := reader.ReadString('\n')

	person := Person{
		Name: name,
		Grades: make(map[string] float32),
	}

	fmt.Println("Enter the number of subjects you take:")
	numStr, _ := reader.ReadString('\n')
	numStr = strings.TrimSpace(numStr)
	numberOfSubjects, _ := strconv.Atoi(numStr)
	
	commulativeSum := float32(0.0)
	for i := range numberOfSubjects{
		fmt.Printf("Enter your subject %d: \n", i+1)
		subject,_ := reader.ReadString('\n')
		subject = strings.TrimSpace(subject)

		START:
			fmt.Printf("Enter your Grade for %s \n", subject )
			gradestr, _ := reader.ReadString('\n')
			gradestr = strings.TrimSpace(gradestr)
			grade, _ := strconv.ParseFloat(gradestr, 32)

		if grade < 0 || grade > 100{
			goto START
		}
		commulativeSum += float32(grade)
		person.Grades[subject]= float32(grade)

	}
	totalGrade := gradeCalculator(commulativeSum, numberOfSubjects)

	fmt.Println(".....Your Report.....")
	fmt.Printf("Name: %s", person.Name)
	for subject , score := range person.Grades{
		fmt.Printf("%s: %.2f \n", subject, score)
	}
	fmt.Printf("Your Cummlative Result: %.2f \n" , totalGrade)



}