package controller

type(
	// A annotation on the controller or action
	FunctionalAnnotation struct {
		Name string
		Data map[string]string // The data for the annotation
	}

	// A list of functional annotations
	FunctionalAnnotations []FunctionalAnnotation

)
const (
	FA_ROUTE = 1
	FA_RESULT = 2
)
