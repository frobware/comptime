test:
	golangci-lint run ./...
	go test -cover ./...

coverage:
	go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
	$(RM) -f cover.out

lint:
	golangci-lint run ./...

benchmark:
	go test -bench=. -benchmem -count=1 -benchtime=1s

benchmark-profile:
	BENCHMARK_PROFILE_PORT=6060 go test -bench=. -benchmem -count=1 -benchtime=1s -cpuprofile=cpu.pprof
	bash -c 'go tool pprof cpu.pprof <<< "list github.com/frobware/comptime"'

clean:
	$(RM) comptimeout comptimeout.test result

.PHONY: test coverage lint benchmark benchmark-profile clean

