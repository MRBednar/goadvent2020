package s3connect

import (
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"fmt"
)

func Connect(day string) string {
	bucket := "advent2020bednar"
	item := fmt.Sprintf("day%vInput.txt", day)

	fmt.Println(item)

	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	s3Client := s3.New(sess)

	requestInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	}

	result, err := s3Client.GetObject(requestInput)
	if err != nil {
		fmt.Println(err)
	}

	defer result.Body.Close()
	body1, err := ioutil.ReadAll(result.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString1 := string(body1)

	return bodyString1
}
