package codechallenges

var (
	// IPList - IP's from ServiceDiscovery
	IPList = []string{"10.128.1.1", "10.128.1.2", "10.128.1.3"}

	// Channel - Blocking and Synchrozation
	channel = make(chan string)
)

// Initiate Before Before Take -
// If no go goutine, its a dead lock
func Initiate() {
	// By default, sends and receives block until the other side is ready.
	// This allows goroutines to synchronize without explicit locks or condition variables.
	go func() {
		for {
			for i := 0; i < len(IPList); i++ {
				channel <- IPList[i]
			}
		}
	}()
}

// Take gets a item from the channel - If not avaliable we will block indefinetly
func Take() string {
	return <-channel
}

// ResetIPList - sets a new ip list,
// but the an ip from old list is still waiting to be served, Take() removes that away
func ResetIPList(list []string) {
	IPList = list
	Take()
}
