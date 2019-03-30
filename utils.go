package kates

const (
	verifyAttempts = 30
	verifyDelayMilli = 1000

)

// EquityLabelSelector ... will build an equity based label selector
func EquityLabelSelector(m map[string]string) string {

	selector := ""
	for key, value := range m {
		if len(selector) == 0 {
			selector = key + "=" + value
		} else {
			selector = selector + "," + key + "=" + value
		}
	}
	return selector
}
