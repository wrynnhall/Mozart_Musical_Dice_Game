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
	// populate dice game with a random measure, gotten by choosing a random value from the variant table
	for i := range musicalDiceGame {
		diceRoll := diceRoll()
		measureSelect := variantTable[i][diceRoll]
		musicalDiceGame[i] = measures[measureSelect-1]
	}

	for i, m := range musicalDiceGame {
		for j, nb := range m.NoteBeat {

			m.NoteBeat[j].Beat = math.Mod(nb.Beat, 3) + float64(3*i)

		}
	}

	f, err := os.Create("result")
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
	return rand.Intn(10)
}

func getVariantTable() [][]int {
	variants := [][]int{}

	file, err := os.Open("variants.txt")
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
				log.Println("Scan complete and reached EOF")
			} else {
				log.Fatal(err)
			}
			eof = true
		} else {
			s := strings.Split(scanner.Text(), ",")
			variant := []int{}
			for i := range s {
				convS, _ := strconv.ParseInt(s[i], 10, 64)
				variant = append(variant, int(convS))
			}
			variants = append(variants, variant)
		}

	}

	return variants
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
