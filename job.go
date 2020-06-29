package JRYGJOB

type Job interface {
	Execute(params string) (string, error)
}
