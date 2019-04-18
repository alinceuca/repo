package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl2 *template.Template

func init() {
	tmpl2 = template.Must(template.ParseGlob("sudoku.html"))
}

func main() {
	http.HandleFunc("/", index2)
	http.HandleFunc("/process", processor2)
	http.ListenAndServe(":9090", nil)

}
func index2(w http.ResponseWriter, r *http.Request) {
	tmpl2.ExecuteTemplate(w, "sudoku.html", nil)
}
func processor2(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var sudoku [9][9]int
	sudoku[0][0], _ = strconv.Atoi(r.FormValue("cell-0"))
	sudoku[0][1], _ = strconv.Atoi(r.FormValue("cell-1"))
	sudoku[0][2], _ = strconv.Atoi(r.FormValue("cell-2"))
	sudoku[0][3], _ = strconv.Atoi(r.FormValue("cell-3"))
	sudoku[0][4], _ = strconv.Atoi(r.FormValue("cell-4"))
	sudoku[0][5], _ = strconv.Atoi(r.FormValue("cell-5"))
	sudoku[0][6], _ = strconv.Atoi(r.FormValue("cell-6"))
	sudoku[0][7], _ = strconv.Atoi(r.FormValue("cell-7"))
	sudoku[0][8], _ = strconv.Atoi(r.FormValue("cell-8"))
	//--
	sudoku[1][0], _ = strconv.Atoi(r.FormValue("cell-9"))
	sudoku[1][1], _ = strconv.Atoi(r.FormValue("cell-10"))
	sudoku[1][2], _ = strconv.Atoi(r.FormValue("cell-11"))
	sudoku[1][3], _ = strconv.Atoi(r.FormValue("cell-12"))
	sudoku[1][4], _ = strconv.Atoi(r.FormValue("cell-13"))
	sudoku[1][5], _ = strconv.Atoi(r.FormValue("cell-14"))
	sudoku[1][6], _ = strconv.Atoi(r.FormValue("cell-15"))
	sudoku[1][7], _ = strconv.Atoi(r.FormValue("cell-16"))
	sudoku[1][8], _ = strconv.Atoi(r.FormValue("cell-17"))
	//--
	sudoku[2][0], _ = strconv.Atoi(r.FormValue("cell-18"))
	sudoku[2][1], _ = strconv.Atoi(r.FormValue("cell-19"))
	sudoku[2][2], _ = strconv.Atoi(r.FormValue("cell-20"))
	sudoku[2][3], _ = strconv.Atoi(r.FormValue("cell-21"))
	sudoku[2][4], _ = strconv.Atoi(r.FormValue("cell-22"))
	sudoku[2][5], _ = strconv.Atoi(r.FormValue("cell-23"))
	sudoku[2][6], _ = strconv.Atoi(r.FormValue("cell-24"))
	sudoku[2][7], _ = strconv.Atoi(r.FormValue("cell-25"))
	sudoku[2][8], _ = strconv.Atoi(r.FormValue("cell-26"))
	//--
	sudoku[3][0], _ = strconv.Atoi(r.FormValue("cell-27"))
	sudoku[3][1], _ = strconv.Atoi(r.FormValue("cell-28"))
	sudoku[3][2], _ = strconv.Atoi(r.FormValue("cell-29"))
	sudoku[3][3], _ = strconv.Atoi(r.FormValue("cell-30"))
	sudoku[3][4], _ = strconv.Atoi(r.FormValue("cell-31"))
	sudoku[3][5], _ = strconv.Atoi(r.FormValue("cell-32"))
	sudoku[3][6], _ = strconv.Atoi(r.FormValue("cell-33"))
	sudoku[3][7], _ = strconv.Atoi(r.FormValue("cell-34"))
	sudoku[3][8], _ = strconv.Atoi(r.FormValue("cell-35"))
	//--
	sudoku[4][0], _ = strconv.Atoi(r.FormValue("cell-36"))
	sudoku[4][1], _ = strconv.Atoi(r.FormValue("cell-37"))
	sudoku[4][2], _ = strconv.Atoi(r.FormValue("cell-38"))
	sudoku[4][3], _ = strconv.Atoi(r.FormValue("cell-39"))
	sudoku[4][4], _ = strconv.Atoi(r.FormValue("cell-40"))
	sudoku[4][5], _ = strconv.Atoi(r.FormValue("cell-41"))
	sudoku[4][6], _ = strconv.Atoi(r.FormValue("cell-42"))
	sudoku[4][7], _ = strconv.Atoi(r.FormValue("cell-43"))
	sudoku[4][8], _ = strconv.Atoi(r.FormValue("cell-44"))
	//--
	//--
	sudoku[5][0], _ = strconv.Atoi(r.FormValue("cell-45"))
	sudoku[5][1], _ = strconv.Atoi(r.FormValue("cell-46"))
	sudoku[5][2], _ = strconv.Atoi(r.FormValue("cell-47"))
	sudoku[5][3], _ = strconv.Atoi(r.FormValue("cell-48"))
	sudoku[5][4], _ = strconv.Atoi(r.FormValue("cell-49"))
	sudoku[5][5], _ = strconv.Atoi(r.FormValue("cell-50"))
	sudoku[5][6], _ = strconv.Atoi(r.FormValue("cell-51"))
	sudoku[5][7], _ = strconv.Atoi(r.FormValue("cell-52"))
	sudoku[5][8], _ = strconv.Atoi(r.FormValue("cell-53"))
	//--
	sudoku[6][0], _ = strconv.Atoi(r.FormValue("cell-54"))
	sudoku[6][1], _ = strconv.Atoi(r.FormValue("cell-55"))
	sudoku[6][2], _ = strconv.Atoi(r.FormValue("cell-56"))
	sudoku[6][3], _ = strconv.Atoi(r.FormValue("cell-57"))
	sudoku[6][4], _ = strconv.Atoi(r.FormValue("cell-58"))
	sudoku[6][5], _ = strconv.Atoi(r.FormValue("cell-59"))
	sudoku[6][6], _ = strconv.Atoi(r.FormValue("cell-60"))
	sudoku[6][7], _ = strconv.Atoi(r.FormValue("cell-61"))
	sudoku[6][8], _ = strconv.Atoi(r.FormValue("cell-62"))
	//--
	sudoku[7][0], _ = strconv.Atoi(r.FormValue("cell-63"))
	sudoku[7][1], _ = strconv.Atoi(r.FormValue("cell-64"))
	sudoku[7][2], _ = strconv.Atoi(r.FormValue("cell-65"))
	sudoku[7][3], _ = strconv.Atoi(r.FormValue("cell-66"))
	sudoku[7][4], _ = strconv.Atoi(r.FormValue("cell-67"))
	sudoku[7][5], _ = strconv.Atoi(r.FormValue("cell-68"))
	sudoku[7][6], _ = strconv.Atoi(r.FormValue("cell-69"))
	sudoku[7][7], _ = strconv.Atoi(r.FormValue("cell-70"))
	sudoku[7][8], _ = strconv.Atoi(r.FormValue("cell-71"))
	//--
	sudoku[8][0], _ = strconv.Atoi(r.FormValue("cell-72"))
	sudoku[8][1], _ = strconv.Atoi(r.FormValue("cell-73"))
	sudoku[8][2], _ = strconv.Atoi(r.FormValue("cell-74"))
	sudoku[8][3], _ = strconv.Atoi(r.FormValue("cell-75"))
	sudoku[8][4], _ = strconv.Atoi(r.FormValue("cell-76"))
	sudoku[8][5], _ = strconv.Atoi(r.FormValue("cell-77"))
	sudoku[8][6], _ = strconv.Atoi(r.FormValue("cell-78"))
	sudoku[8][7], _ = strconv.Atoi(r.FormValue("cell-79"))
	sudoku[8][8], _ = strconv.Atoi(r.FormValue("cell-80"))
	//--

	IsValid(&sudoku, 0)
	Display(sudoku, w)

}
func Display(sudoku [9][9]int, w http.ResponseWriter) {
	fmt.Fprintln(w, "Rezolvarea este: ")
	var x, y int
	for x = 0; x < 9; x++ {
		fmt.Fprintln(w, "")
		if x == 3 || x == 6 {
			fmt.Fprintln(w, " ")
		}
		for y = 0; y < 9; y++ {
			if y == 3 || y == 6 {
				fmt.Fprint(w, "|")
			}
			fmt.Fprint(w, sudoku[x][y])
		}
	}
}
func AbsentOnLine(k int, sudoku [9][9]int, x int) bool {
	var y int
	for y = 0; y < 9; y++ {
		if sudoku[x][y] == k {
			return false
		}
	}
	return true
}
func AbsentOnRow(k int, sudoku [9][9]int, y int) bool {
	var x int
	for x = 0; x < 9; x++ {
		if sudoku[x][y] == k {
			return false
		}
	}
	return true
}
func AbsentOnBloc(k int, sudoku [9][9]int, x int, y int) bool {
	var firstX, firstY int
	firstX = x - (x % 3)
	firstY = y - (y % 3)
	for x = firstX; x < firstX+3; x++ {
		for y = firstY; y < firstY+3; y++ {
			if sudoku[x][y] == k {
				return false
			}
		}
	}
	return true
}
func IsValid(sudoku *[9][9]int, position int) bool {
	if position == 9*9 {
		return true
	}
	var x, y, k int
	x = position / 9
	y = position % 9
	if sudoku[x][y] != 0 {
		return IsValid(sudoku, position+1)
	}
	for k = 1; k <= 9; k++ {
		if AbsentOnLine(k, *sudoku, x) && AbsentOnRow(k, *sudoku, y) && AbsentOnBloc(k, *sudoku, x, y) {
			sudoku[x][y] = k
			if IsValid(sudoku, position+1) {
				return true
			}
		}
	}
	sudoku[x][y] = 0
	return false
}
