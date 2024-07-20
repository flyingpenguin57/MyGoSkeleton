package facade

import (
	"MyApp/facade/request"
	"MyApp/facade/result"
	"log"
)

func Test(request *request.TestRequest) *result.TestResult {
	log.Println(request.Name)
	log.Println(request.Age)

	result := result.NewTestResult()
    result.SetName("hulio")
	result.SetAge(19)
	return result
}
