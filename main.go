package main

import (
	"bufio"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type NoteBeat struct {
	Note       string
	Beat       float64
	BeatLength float64
}

type Measure struct {
	NoteBeat      []NoteBeat
	MeasureNumber int
}

func main() {
	measures := getMeasures()
	variantTable := getVariantTable()
	musicalDiceGame := [16]Measure{}
	for i, _ := range musicalDiceGame {
		diceRoll := diceRoll()
		measureSelect := variantTable[i][diceRoll]
		musicalDiceGame[i] = measures[measureSelect-1]
	}
	for i, m := range musicalDiceGame {
		for j, nb := range m.NoteBeat {

			m.NoteBeat[j].Beat = math.Mod(nb.Beat, 3) + float64(3*i)

		}
	}

	f, err := os.Create("dat2")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for _, m := range musicalDiceGame {
		for _, nb := range m.NoteBeat {
			_, err := f.WriteString(nb.Note + " " + strconv.FormatFloat(nb.Beat, 'f', 1, 64) + " " + strconv.FormatFloat(nb.BeatLength, 'f', 1, 64) + "\n")
			if err != nil {
				panic(err)
			}
		}
	}
}

func diceRoll() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(12-2) + 2
}

func getNoteBeats() []NoteBeat {
	noteBeats := []NoteBeat{}

	file, err := os.Open("mozart-dice-starting.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	eof := false
	for eof != true {

		success := scanner.Scan()

		if success == false {
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
			} else {
				log.Fatal(err)
			}
			eof = true
		} else {
			s := strings.Split(scanner.Text(), " ")
			note := s[0]
			beat, _ := strconv.ParseFloat(s[1], 64)
			beatLength, _ := strconv.ParseFloat(s[2], 64)

			noteBeat := NoteBeat{Note: note, Beat: beat, BeatLength: beatLength}
			noteBeats = append(noteBeats, noteBeat)
		}

	}
	return noteBeats
}

func getMeasures() []Measure {
	measures := []Measure{}

	measureNum := 1
	nbs := getNoteBeats()
	measure := Measure{MeasureNumber: measureNum, NoteBeat: []NoteBeat{}}
	for i, nb := range nbs {
		if (measureNum * 3) <= (int(nb.Beat)) {
			measures = append(measures, measure)
			measureNum++
			measure = Measure{MeasureNumber: measureNum, NoteBeat: []NoteBeat{}}
			measure.NoteBeat = append(measure.NoteBeat, nb)
		} else if i == len(nbs)-1 {
			measure.NoteBeat = append(measure.NoteBeat, nb)
			measures = append(measures, measure)
		} else {
			measure.NoteBeat = append(measure.NoteBeat, nb)
		}

	}

	return measures
}

func getVariantTable() [][]int {

	row1 := []int{96, 32, 69, 40, 148, 104, 152, 119, 98, 3, 54}
	row2 := []int{22, 6, 95, 17, 74, 157, 60, 84, 142, 87, 130}
	row3 := []int{141, 128, 158, 113, 163, 27, 171, 114, 42, 165, 10}
	row4 := []int{41, 63, 13, 85, 45, 167, 53, 50, 156, 61, 103}
	row5 := []int{105, 146, 153, 161, 80, 154, 99, 140, 75, 135, 28}
	row6 := []int{122, 46, 55, 2, 97, 68, 133, 86, 129, 47, 37}
	row7 := []int{11, 134, 110, 159, 36, 118, 21, 169, 62, 147, 106}
	row8 := []int{30, 81, 24, 100, 107, 91, 127, 94, 123, 33, 5}
	row9 := []int{70, 117, 66, 90, 25, 138, 16, 120, 65, 102, 35}
	row10 := []int{121, 39, 136, 176, 143, 71, 155, 88, 77, 4, 20}
	row11 := []int{26, 126, 15, 7, 64, 150, 57, 48, 19, 31, 108}
	row12 := []int{9, 56, 132, 34, 125, 29, 175, 166, 82, 164, 92}
	row13 := []int{112, 174, 73, 67, 76, 101, 43, 51, 137, 144, 12}
	row14 := []int{49, 18, 58, 160, 136, 162, 168, 115, 38, 59, 124}
	row15 := []int{109, 116, 145, 52, 1, 23, 89, 72, 149, 173, 44}
	row16 := []int{14, 83, 79, 170, 93, 151, 172, 111, 8, 78, 131}

	values := [][]int{row1, row2, row3, row4, row5, row6, row7, row8,
		row9, row10, row11, row12, row13, row14, row15, row16}

	return values
}
