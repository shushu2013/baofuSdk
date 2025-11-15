package tool

type SendRobotWarningFunc func(msg string, err error) error

var sendRobotWarning SendRobotWarningFunc

func SetSendRobotWarning(f SendRobotWarningFunc) {
	sendRobotWarning = f
}

func SendRobotWarning(msg string, err error) {
	if sendRobotWarning != nil {
		sendRobotWarning(msg, err)
	}
}
