package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	response := fmt.Sprintf("Image Processing took %v ops/ns",benchmark())

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       response,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}, nil

}

func main() {
	lambda.Start(Handler)
}


/**
Method : Benchmark

This method gets the time taken to execute the factorial 40 times.
In total it loops 80 times.
It takes the last 20 execution times.
Gets the average time
Calculates the throughput as time / 40

Prints out the throughput.

returns: none

*/
func benchmark() float64 {
	listofTime := [20]int64{}

	for j := 0; j <= 5; j++ {
		start := time.Now().UnixNano()
		// Loop 40 times.
		for i := 0; i <= 1; i++ {
			imageProcessing()
		}
		// End time
		end := time.Now().UnixNano()
		// Results
		//if j > 20 {
			difference := end - start
			//listofTime[j-20] = difference
		listofTime[j] = difference
		//}
	}
	// Average Time
	sum := int64(0)
	for i := 0; i < len(listofTime); i++ {
		// adding the values of
		// array to the variable sum
		sum += listofTime[i]
	}
	// avg to find the average
	avg := (float64(sum)) / (float64(len(listofTime)))

	// Throughput Rate
	throughput := avg / 40

	// Response
	return throughput
}

func imageProcessing() image.Image {
	file, err := os.Open("image.jpg")

	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	_ = file.Close()
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(1024, 1000, img, resize.Lanczos3)

	return m

}