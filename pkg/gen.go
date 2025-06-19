package pkg

type Generator interface {
	Name(string, string) string
	Gen(string, string) string
}
