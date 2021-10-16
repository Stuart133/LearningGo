# This demonstrates the context switching performed by the go scheduler

go run preemptive.go | cut -c1 | grep '[AB]' | uniq