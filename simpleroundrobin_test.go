package codechallenges

import "testing"

func Test_Take(t *testing.T) {

	Initiate()
	if IPList[0] != Take() {
		t.Fail()
	}
	if IPList[1] != Take() {
		t.Fail()
	}
	if IPList[2] != Take() {
		t.Fail()
	}

	for i := 0; i < 100; i++ {
		go func() {
			if "" == Take() {
				t.Fail()
			}
		}()
	}

	ResetIPList([]string{"10.128.1.5", "10.128.1.6", "10.128.1.7"})
	ip := Take()
	if !(IPList[0] != ip || IPList[1] != ip || IPList[2] != ip) {
		t.Fail()
	}

}
