package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

const weekNumber = 3
const fileName = "dummy data.csv"
const outputFilename = "Project & Cost.csv"

var projects map[string]project

type project struct {
	name     string
	employee []employee
}

type employee struct {
	name  string
	hours string
}

var employeeNames map[int]string
var projectNames []string

func main() {

	projects = make(map[string]project)
	employeeNames = make(map[int]string)
	csvFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for index, projectEntry := range csvLines {

		tempProject := project{}
		tempProject.name = projectEntry[0]

		for i := 1; i < len(projectEntry); i++ {

			//Save employee Names
			if index == 0 {
				employeeNames[i] = projectEntry[i]
				continue
			}

			if projectEntry[i] == "" {
				continue
			}

			// hours, err := strconv.Atoi(projectEntry[i])
			// if err != nil {
			// 	log.Fatal(err)
			// }
			tempProject.employee = append(tempProject.employee, employee{name: employeeNames[i], hours: projectEntry[i]})

		}

		if index != 0 {

			projects[projectEntry[0]] = tempProject
			projectNames = append(projectNames, projectEntry[0])
		}

	}

	fmt.Println("Successfully Read File Data")
	err = WriteToFile(projects)
	if err != nil {
		log.Fatal(err)
	}

}

func WriteToFile(projects map[string]project) (err error) {

	fmt.Println("Writing to Filename", outputFilename)
	csvFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	sort.Strings(projectNames)

	for _, projectName := range projectNames {
		project := projects[projectName]

		err := csvWriter.Write([]string{fmt.Sprintf("Project Code - %s", project.name)})
		if err != nil {
			return err
		}
		err = csvWriter.Write([]string{"Name", "Salary", "Week 1", "Week 2", "Week 3", "Week 4", "Cost/Effort 1", "Cost/effort 2", "Cost/Effort 3", "Cost/Effort 4"})
		if err != nil {
			return err
		}
		for _, employee := range project.employee {

			row := make([]string, 10)
			row[0] = employee.name
			row[weekNumber] = employee.hours

			err = csvWriter.Write(row)
			if err != nil {
				fmt.Println(err)
				continue
			}

		}

		err = csvWriter.WriteAll([][]string{{"total"}, {}})

		if err != nil {
			log.Fatal(err)
		}

	}

	csvWriter.Flush()
	err = csvWriter.Error()
	if err != nil {
		return err
	}
	fmt.Println("Finished Writing to output Filename")

	return nil
}
