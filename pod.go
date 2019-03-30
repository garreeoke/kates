package kates

const (
	channelLoopSleepTime  = 500 // ms
	verifySuccessInAaRow  = 3
)

/*
// verifyPods
func verifyPods(input *Input, output *Output, query string, expected int) (bool, error) {

	podList, err := syncPodList(input, output, expected, query)
	if err != nil {
		return false, err
	}

	ch1 := make(chan []string)                 // Container msgs
	ch2 := make(chan []string)                 // Pod msgs
	done := make(chan int, len(podList.Items)) // Counter channel

	go func() {
		for _, pod := range podList.Items {
			go func(p apiv1.Pod) {
				doneC := make(chan int, 1)
				doneP := make(chan int, 1)
				// Verify containers
				go func() {
					ch1 <- verifyContainers(p)
					doneC <- 1
				}()
				// Verify Pod itself
				go func() {
					ch2 <- verifyPod(input, output, p, query, expected)
					doneP <- 1
				}()

				for len(doneC) < 1 || len(doneP) < 1 {
					time.Sleep(channelLoopSleepTime * time.Millisecond)
				}
				done <- 1
			}(pod)
		}
	}()

	// Merge messages
	go func() {
		for msgLists := range types.MergeChannels(ch1, ch2) {
			for _, msg := range msgLists {
				msgs = append(msgs, msg)
			}
		}
	}()

	for len(done) < len(podList.Items) {
		time.Sleep(channelLoopSleepTime * time.Millisecond)
	}

	close(ch1)
	close(ch2)
	close(done)

	k.Log.Println("Done with pod verification: ", query)
	return msgs

	return output, nil
}
*/

/*
// verifyPod
func verifyPod(input *Input, output *Output, pod apiv1.Pod, query string, expected int) error {

	// Verify pod gets a running status on 3 successive tries
	tries := 1
	success := 0
	retries := 0
	for tries <= verifyAttempts {
		ps,_ := podStatus(input, output, pod.Name)
		ready := "False"
		for _, c := range ps.Conditions {
			if c.Type == "Ready" {
				ready = string(c.Status)
			}
		}
		if ready == "False" && tries < verifyAttempts {
			// Get list of events to see if the reason is backoff
			if retries <= 3 {
				eventList, err := input.Client.CoreV1().Events(input.Namespace).List(meta_v1.ListOptions{
					FieldSelector: "involvedObject.name="+pod.Name,
				})
				if err != nil {
					return err
				}
				for _,e := range eventList.Items {
					if e.Reason == "BackOff" || e.Reason == "Failed" {
						// Delete the pod
						retries++
						err := input.Client.CoreV1().Pods(input.Namespace).Delete(pod.Name, &meta_v1.DeleteOptions{
						})
						if err != nil {
							break
						}
						time.Sleep(5 * time.Second)
						podList, err := syncPodList(input, output, expected, query)
						if err != nil {
							break
						}
						pod = podList.Items[0]
					}
				}
			}
			time.Sleep(time.Duration(verifyDelayMilli) * time.Millisecond)
			tries++
			success = 0 // If fails at all reset
		} else if ready == "False" && tries == verifyAttempts {
			err := eventMessages(input, output, meta_v1.ListOptions{
				FieldSelector: "involvedObject.name="+pod.Name,
			})
			if err != nil {
				return err
			}
			tries++
		} else if ready == "True" {
			if success == verifySuccessInAaRow {
				tries = verifyAttempts + 1
			} else {
				time.Sleep(time.Duration(verifyDelayMilli) * time.Millisecond)
				success++
				tries++
			}
		}
	}
	return nil

}
*/

/*
func syncPodList(input *Input, output *Output, expected int, query string) (*apiv1.PodList, error) {

	podList := &apiv1.PodList{
		Items: []apiv1.Pod{},
	}
	// Go through on cycle so podlist can update ... if this is after an update
	for tries := 1; len(podList.Items) != expected && tries < verifyAttempts; tries++ {
		time.Sleep(time.Duration(verifyDelayMilli) * time.Millisecond)
		podList, _ = getPodList(input, output, query)
	}

	if len(podList.Items) != expected {
		return podList, errors.New(fmt.Sprintf("Unable to verify pods, number of expected pods %v does not match current amount %v", expected, len(podList.Items)))
	}
	return podList, nil
}
*/

/*
// GetPodList get a list of pods based on query string
func getPodList(input *Input, output *Output, query string) (*apiv1.PodList, error) {
	podList, err := input.Client.CoreV1().Pods(input.Namespace).List(meta_v1.ListOptions{
		LabelSelector: query,
	})
	if err != nil {
		return podList,err
	}

	return podList, nil

}
*/

/*
// podStatus get status object of a pod
func podStatus(input *Input, output *Output, podName string) (apiv1.PodStatus, error) {

	pod, err := input.Client.CoreV1().Pods(input.Namespace).Get(podName, meta_v1.GetOptions{})
	if err != nil {
		return pod.Status, err
	}
	return pod.Status, nil
}
*/