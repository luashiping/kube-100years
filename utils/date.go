package utils

func GetDateString() string {
	return Eval("bash", "-c", "date +%Y%m%d")
}
