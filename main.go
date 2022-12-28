package main

import (
	"bufio"
	"flag"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
)

var (
	fileName          string
	amountOfTickets   uint
	distributionParam int64
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseCMD() {
	flag.StringVar(&fileName, "file", "", "name of file with students' names")
	flag.UintVar(&amountOfTickets, "numbilets", 0, "amount of tickets to generate")
	flag.Int64Var(&distributionParam, "parameter", 0, "parameter to change distribution")
	flag.Parse()
}

func readInputFile() []string {
	file, err := os.Open(fileName)
	check(err)
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	check(file.Close())
	check(scanner.Err())
	return lines
}

func makeDistribution(names []string) {
	for _, name := range names {
		fmt.Printf("%s: %d\n", name, getTicketNumberForName(name))
	}
}

func getTicketNumberForName(name string) uint32 {
	hash := fnv.New32a()
	_, err := hash.Write(append([]byte(name + strconv.FormatInt(distributionParam, 10))))
	check(err)
	return hash.Sum32()%uint32(amountOfTickets) + 1
}

func main() {
	parseCMD()
	inputData := readInputFile()
	makeDistribution(inputData)
}
