package username

import (
	"fmt"
	"github.com/gzipchrist/dont_at_me/pkg/social"
	"github.com/gzipchrist/dont_at_me/pkg/style"
	"strings"
	"sync"
)

func CheckAvailabilitySerial(username string) error {
	for i := 0; i < len(social.Platforms); i++ {
		status := social.Platforms[i].GetAvailability(username)
		spacer := style.MaxCharWidth - len(social.PlatformStrings[social.Platforms[i]])
		fmt.Printf("    %s%s%s\n", social.Platforms[i].String(), strings.Repeat(" ", spacer), status.String())
	}

	return nil
}

func CheckAvailabilityConcurrent(username string) {
	wg := sync.WaitGroup{}
	results := make(chan string)

	for i := 0; i < len(social.Platforms); i++ {
		wg.Add(1)
		go func(platform social.Platform) {
			defer wg.Done()
			status := platform.GetAvailability(username)
			results <- fmt.Sprintf("%s%s%s\n", platform.String(), strings.Repeat(" ", style.MaxCharWidth-len(social.PlatformStrings[platform])), status.String())
		}(social.Platforms[i])
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("    %s", result)
	}

	return
}
