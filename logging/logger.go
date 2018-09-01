package logging

var logger Logger

// SetLogger :
func SetLogger(log Logger) {
	logger = log
}

// Debugf :
func Debugf(format string, args ...interface{}) {
	if logger != nil {
		logger.Debugf(format, args...)
	}
}

// Errorf :
func Errorf(format string, args ...interface{}) {
	if logger != nil {
		logger.Errorf(format, args...)
	}
}

// Fatalf :
func Fatalf(format string, args ...interface{}) {
	if logger != nil {
		logger.Fatalf(format, args...)
	}
}

// Infof :
func Infof(format string, args ...interface{}) {
	if logger != nil {
		logger.Infof(format, args...)
	}
}

// Warnf :
func Warnf(format string, args ...interface{}) {
	if logger != nil {
		logger.Warnf(format, args...)
	}
}
