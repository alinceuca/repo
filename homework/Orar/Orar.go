package main

import (
	"bufio"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type searchOption int

const fileName = "Orar_sem_II_2018-2019.xlsx"
const urlAdress = "https://mateinfo.unitbv.ro/images/studenti/orar/sem2_2018-2019/Orar_sem_II_2018-2019_V4.xlsx"

const (
	dupaZi searchOption = iota
	dupaMaterie
	dupaSala
	dupaAn
	dupaSpecializare
	dupaProfesor
	dupaTipActivitate
	dupaGrupa
	dupaSemiGrupa
	dupaOra
)

type course struct {
	day             string
	subject         string
	classroom       string
	year            string
	specialisation  string
	teacherFullName string
	activityType    string
	group           string
	semiGroup       string
	hours           string
}

var n int

func main() {

	if !existsFile(fileName) {
		if err := downloadFile(fileName, urlAdress); err != nil {
			panic(err)
		}
	}

	courses := getCourses()
	n = len(courses)

	scanner := bufio.NewScanner(os.Stdin)
	i := 0
	for i != 12 {
		showMenu(scanner, courses, i)
	}
}

func showMenu(scanner *bufio.Scanner, courses []course, i int){
	fmt.Println("Search:")
	fmt.Println("1.  dupa zi")
	fmt.Println("2.  dupa materie ")
	fmt.Println("3.  dupa clasa")
	fmt.Println("4.  dupa an ")
	fmt.Println("5.  dupa specializare")
	fmt.Println("6.  dupa numele intreg al profesorului")
	fmt.Println("7.  dupa activitate")
	fmt.Println("8.  dupa grupa")
	fmt.Println("9.  dupa semi-grupa")
	fmt.Println("10. dupa ora")
	fmt.Println("11. dupa un curs specific")
	fmt.Println("12. Iesire")

	fmt.Printf("Optiunea Dvs.: ")
	scanner.Scan()
	optionNumber := scanner.Text()
	i, err := strconv.Atoi(optionNumber)
	if err != nil {
		panic(err)
	}

	switch i {
	case 1:
		search("Zi: ", scanner, courses, dupaZi)
	case 2:
		search("Materie: ", scanner, courses, dupaMaterie)
	case 3:
		search("Sala: ", scanner, courses, dupaSala)
	case 4:
		search("An: ", scanner, courses, dupaAn)
	case 5:
		search("Specializare: ", scanner, courses, dupaSpecializare)
	case 6:
		search("Numele profesorului: ", scanner, courses, dupaProfesor)
	case 7:
		search("Tip de activitate: ", scanner, courses, dupaTipActivitate)
	case 8:
		search("Grupa: ", scanner, courses, dupaGrupa)
	case 9:
		search("Semi-grupa: ", scanner, courses, dupaSemiGrupa)
	case 10:
		search("Ore: ", scanner, courses, dupaOra)
	case 11:
		optionsFields := getCourseFields(scanner)
		course := course{
			day:             optionsFields[0],
			subject:         optionsFields[1],
			classroom:       optionsFields[2],
			year:            optionsFields[3],
			specialisation:  optionsFields[4],
			teacherFullName: optionsFields[5],
			activityType:    optionsFields[6],
			group:           optionsFields[7],
			semiGroup:       optionsFields[8],
			hours:           optionsFields[9],
		}
		if searchForASpecificCourse(courses, course) {
			fmt.Println("Cursul specificat exista!")
		} else {
			fmt.Println("Cursul specificat NU exista!")
		}
	case 12:
		os.Exit(0)
	default:
		fmt.Println("Optiune eronata!")
	}
	fmt.Println()
}

func search(message string, scanner *bufio.Scanner, courses []course, option searchOption) {
	fmt.Printf(message)
	scanner.Scan()
	coursesSearched := searchForCourses(courses, scanner.Text(), option)
	displayArray(coursesSearched)
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func existsFile(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getCourses() []course {
	unmerged, _ := xlsx.FileToSliceUnmerged(fileName)
	content := unmerged[0]

	courses := []course{}

	for index, row := range content {
		if (row[0] == "1" || row[0] == "2" || row[0] == "3") && row[1] != "Codul Orarului: " {
			year := row[0]
			specialisation := row[1]
			group := row[2]
			semiGroup := row[3]

			for i := 4; i < len(row); i++ {
				subject, activityType, classroom, teacherFullName := getElements(row[i])
				if subject != "" {
					courses = append(courses, course{
						day:             getDay(i),
						subject:         subject,
						classroom:       classroom,
						year:            year,
						specialisation:  specialisation,
						teacherFullName: teacherFullName,
						activityType:    activityType,
						group:           group,
						semiGroup:       semiGroup,
						hours:           getHours(index),
					})
				}
			}
		}

	}
	return courses
}

func getElements(row string) (string, string, string, string) {
	if row != "" {
		split := strings.Split(row, ",")
		if len(split) == 4 {
			return split[0], split[1], split[2], split[3][1:]
		}
	}
	return "", "", "", ""
}

func getDay(index int) string {
	dayIndex := (index - 4) % 7

	switch dayIndex {
	case 0:
		return "Luni"
	case 1:
		return "Marti"
	case 2:
		return "Miercuri"
	case 3:
		return "Joi"
	case 4:
		return "Vineri"
	case 5:
		return "Sambata"
	case 6:
		return "Duminica"
	default:
		return "unknown"
	}
}

func getHours(index int) string {
	dayIndex := (index - 4) % 7

	switch dayIndex {
	case 0:
		return "8,00-9,50"
	case 1:
		return "10,00-11,50"
	case 2:
		return "12,00-13,50"
	case 3:
		return "14,00-15,50"
	case 4:
		return "16,00-17,50"
	case 5:
		return "18,00-19,50"
	case 6:
		return "20,00-21,50"
	default:
		return "unknown"
	}
}

func searchForCourses(courses []course, specificationSearched string, option searchOption) []course {
	var result []course

	for i := 0; i < n; i++ {
		found := false
		switch option {
		case dupaZi:
			{
				if courses[i].day == specificationSearched {
					found = true
				}
			}
		case dupaMaterie:
			{
				if courses[i].subject == specificationSearched {
					found = true
				}
			}
		case dupaSala:
			{
				if courses[i].classroom == specificationSearched {
					found = true
				}
			}
		case dupaAn:
			{
				if courses[i].year == specificationSearched {
					found = true
				}
			}
		case dupaSpecializare:
			{
				if courses[i].specialisation == specificationSearched {
					found = true
				}
			}
		case dupaProfesor:
			{
				if courses[i].teacherFullName == specificationSearched {
					found = true
				}
			}
		case dupaTipActivitate:
			{
				if courses[i].activityType == specificationSearched {
					found = true
				}
			}
		case dupaGrupa:
			{
				if courses[i].group == specificationSearched {
					found = true
				}
			}
		case dupaSemiGrupa:
			{
				if courses[i].semiGroup == specificationSearched {
					found = true
				}
			}
		case dupaOra:
			{
				if courses[i].hours == specificationSearched {
					found = true
				}
			}
		}
		if found == true {
			result = append(result, courses[i])
		}
	}

	return result
}

func searchForASpecificCourse(courses []course, course course) bool {
	for i := 0; i < n; i++ {
		if courses[i] == course {
			return true
		}
	}
	return false
}

func displayArray(array []course) {
	if len(array) == 0 {
		fmt.Println("Nu s-au gasit cursuri...")
	} else {
		for i := 0; i < len(array); i++ {
			displayCourse(array[i])
			fmt.Println()
		}
	}
}

func displayCourse(cell course) {
	fmt.Printf("Zi:            %v\n", cell.day)
	fmt.Printf("Materie:        %v\n", cell.subject)
	fmt.Printf("Sala:      %v\n", cell.classroom)
	fmt.Printf("An:           %v\n", cell.year)
	fmt.Printf("Specializare: %v\n", cell.specialisation)
	fmt.Printf("Profesor:        %v\n", cell.teacherFullName)
	fmt.Printf("Tip de activitate:           %v\n", cell.activityType)
	fmt.Printf("Grupa:          %v\n", cell.group)
	fmt.Printf("Semi-grupa:      %v\n", cell.semiGroup)
	fmt.Printf("Ore:          %v\n", cell.hours)
}

func getCourseFields(scanner *bufio.Scanner) []string {
	var optionsFields []string

	fmt.Printf("Zi: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Materie: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Sala: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("An: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Specializare: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Numele profesorului: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Tip de activitate: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Grupe: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Semi-grupa: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	fmt.Printf("Ore: ")
	scanner.Scan()
	optionsFields = append(optionsFields, scanner.Text())

	return optionsFields
}