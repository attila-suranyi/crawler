package main

import "time"

type Config struct {
	WpCfg          WorkerPoolConfig
	RetryNum       int
	RateLimitDelay time.Duration
}

type WorkerPoolConfig struct {
	NumOfWorkers int
	BufferSize   int
}