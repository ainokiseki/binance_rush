package common

func GOMultiProcess(n int, f func()) chan struct{} {
	ch := make(chan struct{})
	for i := 0; i < n; i++ {
		go func() {
			<-ch
			f()
		}()
	}
	return ch
}

func GOMultiProcessWithChan(n int, f func(), ch chan struct{}) {
	for i := 0; i < n; i++ {
		go func() {
			<-ch
			f()
		}()
	}
}
