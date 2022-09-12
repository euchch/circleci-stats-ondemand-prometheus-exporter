package main

import (
	"bufio"
	"euchch/circleci-stats-ondemand-prometheus-exporter/pkg/polingqueue"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	pq := polingqueue.NewPolingQueue()
	interval, err := getInterval()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/metrics", promhttp.Handler())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("Running ticker every ", interval, " seconds")
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			fmt.Println("***Tick!***")
			fmt.Println(pq)
			rb := pq.Dequeue()
			if rb != nil {
				fmt.Println("Dealing with: ", polingqueue.RepoBranch{Repo: rb.Repo, Branch: rb.Branch})
			}
			fmt.Println("***********")

			select {
			case <-ticker.C:
				continue
			case <-interrupt:
				ticker.Stop()
				os.Exit(0)
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "." {
			break
		}
		rb := strings.Split(line, ",")
		pq.EnqueueUniqe(rb[0], rb[1])
	}
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func getInterval() (int, error) {
	const defaultCircleCIAPIIntervalSecond = 5
	circleciAPIInterval := os.Getenv("CIRCLECI_POLING_INTERVAL")
	if len(circleciAPIInterval) == 0 {
		return defaultCircleCIAPIIntervalSecond, nil
	}

	integerCircleCIAPIInterval, err := strconv.Atoi(circleciAPIInterval)
	if err != nil {
		return 0, fmt.Errorf("failed to read CircleCI Config: %w", err)
	}

	return integerCircleCIAPIInterval, nil
}
