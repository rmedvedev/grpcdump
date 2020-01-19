package logger

import "github.com/sirupsen/logrus"

//Init ...
func Init(level string) error {
	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}

	logrus.SetLevel(parsedLevel)
	return nil
}
