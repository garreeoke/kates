package kates

/*
import (
	"fmt"
	"time"
	apiv1 "k8s.io/api/core/v1"
)

// verifyContainers
func verifyContainers(pod apiv1.Pod) []string {

	var msgs []string
	for tries := 1; tries <= verifyAttempts; {
		podStat, _ := k.podStatus(pod.Name)
		if tries < k.Settings.Attempts && len(podStat.ContainerStatuses) != len(pod.Spec.Containers) {
			time.Sleep(time.Duration(k.Settings.Delay) * time.Millisecond)
			tries++
		} else if tries == k.Settings.Attempts && len(podStat.ContainerStatuses) != len(pod.Spec.Containers) {
			for _, c := range pod.Spec.Containers {
				m := fmt.Sprintf("Container %v in Pod %v is Not Ready", c.Name, pod.Name)
				msgs = append(msgs, m)
			}
			tries++
		} else if len(podStat.ContainerStatuses) == len(pod.Spec.Containers) {
			done := make(chan int, len(podStat.ContainerStatuses))
			for _, containerStatus := range podStat.ContainerStatuses {
				go func(container string) {
					tries := 1
					success := 0
					for tries <= k.Settings.Attempts {
						ready := k.containerReady(pod, container)
						k.Log.Printf("Container: %v Status: %v Success: %v", container, ready, success)
						if !ready && tries < k.Settings.Attempts {
							time.Sleep(time.Duration(k.Settings.Delay) * time.Millisecond)
							tries++
							success = 0
						} else if !ready && tries == k.Settings.Attempts {
							m := fmt.Sprintf("Container %v in Pod %v is not ready", container, pod.Name)
							msgs = append(msgs, m)
							tries++
						} else if ready {
							if success == verifySuccessInAaRow {
								tries = k.Settings.Attempts + 1
							} else {
								time.Sleep(time.Duration(k.Settings.Delay) * time.Millisecond)
								success++
								tries++
							}
						}
					}
					done <- 1
				}(containerStatus.Name)

			}
			for len(done) < len(podStat.ContainerStatuses) {
				time.Sleep(channelLoopSleepTime * time.Millisecond)
			}
			tries = k.Settings.Attempts + 1
		}
	}
	return msgs
}
*/