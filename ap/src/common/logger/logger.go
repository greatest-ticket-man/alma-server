package logger

import "log"

// logをwrapしておく、今後StackDriverに出すようになったときに、指定した形式でログが出せるようにしておく

// Debug .
func Debug(msg ...interface{}) {
	log.Println(msg...)
}

// Debugf .
func Debugf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Info .
func Info(msg ...interface{}) {
	log.Println(msg...)
}

// Infof .
func Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Error .
func Error(err error, msg ...interface{}) {
	log.Fatalln(err, msg)
}

// ErrorAll .
func ErrorAll(msg ...interface{}) {
	log.Fatalln(msg...)
}

// Errorf .
func Errorf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
