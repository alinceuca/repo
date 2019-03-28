package main
import (
	"fmt"
	"html/template"
	"math"
	"net/http"
	"strconv"
)
type point struct{
	coordX float64
	coordY float64
}
type pointPair struct{
	firstPoint point
	secondPoint point
	distance float64
}
func main(){
	tmpl := template.Must(template.ParseFiles("form.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		nrPoints, _ := strconv.ParseInt(r.FormValue("nrOfPoints"),10,64)
		pointList:= []point{}
		for i:= 0;i<int(nrPoints);i++ {

			coordX, _ := strconv.ParseFloat(r.FormValue("coordX"+strconv.Itoa(i)),64)
			coordY, _ := strconv.ParseFloat(r.FormValue("coordY"+strconv.Itoa(i)),64)
			pointList = append(pointList,createPoint(coordX,coordY))
		}
		farthestPointPair := getFarthestPointPair(pointList)
		distancePoint:=[]float64{}
		for _, element  := range pointList {
			distancePoint=append(distancePoint,getDistanceBetweenPoints(getMiddlePoint(farthestPointPair.firstPoint,farthestPointPair.secondPoint),element ))
		}
		fmt.Fprintf(w,"<!DOCTYPE html> <html lang='en'> <head> <meta charset='UTF-8'> <title>Test</title> </head> <body><script>\n function drawCircle(){\n"+
		"var c = document.getElementById('myCanvas');\n"+
		"var ctx = c.getContext('2d');\n"+
		"var coordCenterX=document.getElementById('coordCenterX').value;\n"+
		"var coordCenterY=document.getElementById('coordCenterY').value;\n"+
		"var centerRadius=document.getElementById('centerRadius').value;\n"+
		"var distanceList=document.getElementById('distanceList').value;\n"+
		"ctx.beginPath();\n"+
		"ctx.arc(coordCenterX, coordCenterY, centerRadius, 0, 2 * Math.PI);\n"+
		"ctx.stroke();\n" +
		"var distanceList =distanceList.replace('[','');\n"+
		"var distanceList =distanceList.replace(']','');\n"+
		"var elmentsList =distanceList.split(' ');\n"+
		"for(var j=0;j<elmentsList.length;j++){\n"+
		"drawPoint(66,parseFloat(elmentsList[j]),'P'+j,coordCenterX,coordCenterY);\n"+
		"}}\n"+
		"function drawPoint(angle,distance,label,coordCenterX,coordCenterY,centerRadius){\n"+
		"var x = coordCenterX + centerRadius * Math.cos(-angle*Math.PI/180) * distance;\n"+
		"var y = coordCenterY + centerRadius * Math.sin(-angle*Math.PI/180) * distance;\n"+
		"var c = document.getElementById('myCanvas');\n"+
		"var ctx = c.getContext('2d');\n"+
		"ctx.beginPath();\n"+
		"ctx.arc(x, y, 4, 0, 2 * Math.PI);\n"+
		"ctx.fill();\n"+
		"ctx.font = '20px';\n"+
		"ctx.fillText(label,x + 10,y);\n"+
		"ctx.stroke();\n" +
		"}\n"+
	"</script>")
		fmt.Fprintf(w,"<input type='text' id='distanceList' hidden name='distanceList' value='%f'> \n",distancePoint)
		fmt.Fprintf(w, "<p>NrOfPoints= %d </p> \n", nrPoints)
		fmt.Fprintf(w,"<p>Circle center point: %f </p>  \n",getMiddlePoint(farthestPointPair.firstPoint,farthestPointPair.secondPoint))
		fmt.Fprintf(w,"<p>Circle radius: %f </p>\n",farthestPointPair.distance/2)
		fmt.Fprintf(w,"<input type='number' id='coordCenterX' hidden name='coordCenterX' value='%f'> \n",getMiddlePoint(farthestPointPair.firstPoint,farthestPointPair.secondPoint).coordX)
		fmt.Fprintf(w,"<input type='number' id='coordCenterY' hidden name='coordCenterY' value='%f'> \n",getMiddlePoint(farthestPointPair.firstPoint,farthestPointPair.secondPoint).coordY)
		fmt.Fprintf(w,"<input type='number' id='centerRadius' hidden name='centerRadius' value='%f'> \n",farthestPointPair.distance/2)
		fmt.Fprintf(w,"<button name='generateFormBtn' onclick='drawCircle()'>Draw Circle</button><br>\n")
		fmt.Fprintf(w,"<canvas id='myCanvas' width='800' height='800' style='border:1px solid #d3d3d3;'></canvas>"+
		"</body> </html>")

	})
	http.ListenAndServe(":8080", nil)
}

func createPoint(coordX float64,coordY float64) point{
	return point{
		coordX:coordX,
		coordY:coordY,
	}
}

func getMiddlePoint(firstPonit,secondPoint point) point{
	return point{
		coordX:(firstPonit.coordX+secondPoint.coordX)/2,
		coordY:(firstPonit.coordY+secondPoint.coordY)/2,
	}
}
func getPointPair(firstPoint,secondPoint point) pointPair{
	return pointPair{
		firstPoint:firstPoint,
		secondPoint:secondPoint,
		distance:getDistanceBetweenPoints(firstPoint,secondPoint),
	}
}

func getDistanceBetweenPoints(firstPoint,secondPoint point) float64 {
	xPart := math.Pow(firstPoint.coordX-secondPoint.coordX,2)
	yPart := math.Pow(firstPoint.coordY-secondPoint.coordY,2)
	return math.Pow(xPart+yPart,0.5)
}

func getFarthestPointPair(pointList []point) pointPair{
	farthestPointPair := getPointPair(pointList[0],pointList[1])
	for _,firstPoint := range pointList {
		for _,secondPoint := range pointList {
			pointPair := getPointPair(firstPoint,secondPoint)
			if(farthestPointPair.distance < pointPair.distance ){
				farthestPointPair = pointPair
			}
		}
	}
	return farthestPointPair
}