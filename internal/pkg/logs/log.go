package logs

func Init(logLevel string, showConsole bool) {
	SetConsoleMode(showConsole)
	SetLogMode(logLevel)
}

func SetLevel(level LogLevel) {
	std.SetLevel(level)
}

func SetConsoleMode(flag bool) {
	enableConsole = flag
}

func SetWriteRouter(b bool) {
	enableRouter = b
}
