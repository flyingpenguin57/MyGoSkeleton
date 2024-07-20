package result

type BasicResult struct {
	Success      bool
	ErrorMessage string
}

func NewBasicResult() *BasicResult {
	return &BasicResult{
		Success: true,
	}
}
