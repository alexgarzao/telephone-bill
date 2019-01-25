package infrastructure

type LoggerHandler interface {
	Log(args ...interface{})
}
