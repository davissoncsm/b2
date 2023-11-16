package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const MESSAGE = "Message"
const BITMAP = "Bitmap"
const PROCES_CODE = "03 Processing Code"
const TRANSACTION_AMOUNT = "04 Transaction Amoaunt"
const TRANSMISSION_DATE = "07 Transmission Date & Time"
const FREE_AMOUNT = "08 Fee Amount"
const SYSTEM_TRACE = "11 System Trace Audit Number"
const LOCAL_TRANSACTION = "12 Local Transaction Time"
const STTLEMENT_DATE = "15 Settlement Date"
const ACQUIRING_COUNTRY = "19 Acquiring Country Code"
const PAN_COUNTRY = "20 PAN Country Code"
const PAN_SEQUENCE = "23 PAN Sequence Number"
const FUNCTION_CODE = "24 Function Code"
const POS_CAPTURE = "26 POS Capture Code"
const ACQUIRING_IDENTIFICATION = "32 Acquiring Identification"
const TRACK = "35 Track 2"

const INPUT_FILE = "financial_transaction_message.dat"
const OUTPUT_FILE = "output.txt"

var arr []string

func main() {

	inputFile, err := os.Open(INPUT_FILE)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer inputFile.Close()

	fi, err := inputFile.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	if fi.Size() == 0 {
		fmt.Println("The file is empty.")
		return
	}

	outputFile, err := os.Create(OUTPUT_FILE)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return
			}
		}

		split(line)

		for i := 0; i < len(arr); i++ {
			writer.WriteString(arr[i])
		}

	}

	writer.Flush()

	fmt.Println("Finish process!")
}

func split(line string) {

	handleMessage(line[0:4])
	handleBitmap(line[4:12])
	handleProcessCode(line[12:18])
	handleTransactionAmount(line[18:30])
	handleTransmitionDate(line[30:44])
	handleFreeAmount(line[44:56])
	handleSystemTrace(line[56:62])
	handleLocalTransaction(line[62:68])
	handleSttlementDate(line[68:72])
	handleAcquiringCountry(line[72:75])
	handlePanCountry(line[75:78])
	handlePanSequence(line[78:81])
	handleFunctionCode(line[81:84])
	handlePosCapture(line[84:86])
	handleAcquiringIdentification(line[86:91])
	handleTrack(line[91:131])

}

func handleMessage(value string) {
	arr = append(arr, fmt.Sprintf("%s: %s\n", MESSAGE, value))
}

func handleBitmap(value string) {
	bitmap := strToBin(value)
	arr = append(arr, fmt.Sprintf("%s: %s\n", BITMAP, bitmap))
}

func handleProcessCode(value string) {
	arr = append(arr, fmt.Sprintf("%s (%d): %s\n", PROCES_CODE, len(value), value))
}

func handleTransactionAmount(value string) {
	arr = append(arr, fmt.Sprintf("%s (%d): %s\n", TRANSACTION_AMOUNT, len(value), value))
}

func handleTransmitionDate(value string) {
	arr = append(arr, fmt.Sprintf("%s (%d): %s\n", TRANSMISSION_DATE, len(value), value))
}

func handleFreeAmount(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", FREE_AMOUNT, len(value), value))
}

func handleSystemTrace(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", SYSTEM_TRACE, len(value), value))
}

func handleLocalTransaction(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", LOCAL_TRANSACTION, len(value), value))
}

func handleSttlementDate(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", STTLEMENT_DATE, len(value), value))
}

func handleAcquiringCountry(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", ACQUIRING_COUNTRY, len(value), value))
}

func handlePanCountry(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", PAN_COUNTRY, len(value), value))
}

func handlePanSequence(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", PAN_SEQUENCE, len(value), value))
}

func handleFunctionCode(value string) {

	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", FUNCTION_CODE, len(value), value))
}

func handlePosCapture(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", POS_CAPTURE, len(value), value))
}

func handleAcquiringIdentification(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", ACQUIRING_IDENTIFICATION, len(value), value))
}

func handleTrack(value string) {
	arr = append(arr, fmt.Sprintf("%s: (%d): %s\n", TRACK, len(value), value))
}

func strToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}
