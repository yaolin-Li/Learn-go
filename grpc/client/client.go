package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	pb "learn-go/grpc/route"
	"log"
	"os"
	"time"
)

func runFirst(client pb.RouteGuideClient) {
	feature, err := client.GetFeature(context.Background(), &pb.Point{
		Latitude:  123,
		Longitude: 123,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(feature)
}

func runThird(client pb.RouteGuideClient) {

	points := []*pb.Point {
		{Latitude: 123, Longitude: 123},
		{Latitude: 234, Longitude: 234},
		{Latitude: 345, Longitude: 345},
	}
	clientStream, err := client.RecordRoute(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	for _, point := range points {
		if err := clientStream.Send(point); err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}
	summary, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(summary)
}
func runSecond(client pb.RouteGuideClient) {
	serverStream, err := client.ListFeatures(context.Background(), &pb.Rectangle{
		Lo: &pb.Point{
			Latitude:  111,
			Longitude: 111,
		},
		Hi: &pb.Point{
			Latitude:  555,
			Longitude: 555,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	for {
		feature, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(feature)
	}
}
func readIntFromCommandLine(reader *bufio.Reader, target *int32) {
	_, err := fmt.Fscanf(reader, "%d\n", target)
	if err != nil {
		log.Fatalln("Cannot scan", err)
	}
}

func runFourth(client pb.RouteGuideClient) {
	stream, err := client.Recommend(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		feature, err2 := stream.Recv()
		if err2 != nil {
			log.Fatalln(err2)
		}
		fmt.Println("Recommended", feature)
	}()

	reader := bufio.NewReader(os.Stdin)
	for {
		request := pb.RecommendationRequest{Point: new(pb.Point)}
		var mode int32
		fmt.Println("Enter Recommendation Mode(0 for farthest, 1 for the nearest)")
		readIntFromCommandLine(reader, &mode)
		fmt.Print("Enter Latitude: ")
		readIntFromCommandLine(reader, &request.Point.Latitude)
		fmt.Print("Enter Longitude: ")
		readIntFromCommandLine(reader, &request.Point.Longitude)
		request.Mode = pb.RecommendationMode(mode)

		err := stream.Send(&request)
		if err != nil {
			log.Fatalln(err)
		}

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("client cannot dial grpc server")
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	runFourth(client)

}
