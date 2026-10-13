package tools


type Tool struct {

	Name string

	Description string

	Function func(string) (string, error)

}

