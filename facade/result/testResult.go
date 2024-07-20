package result

type TestResult struct {
	*BasicResult
	Name string
	Age  int8
}

func NewTestResult() *TestResult {
	return &TestResult{
		BasicResult: NewBasicResult(),
	}
}

func (u *TestResult) SetName(name string) {
    u.Name = name
}

func (u *TestResult) SetAge(age int8) {
    u.Age = age
}