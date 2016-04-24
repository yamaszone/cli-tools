#!/usr/bin/env bats

load test_helper

@test "'clean help' tool displays help." {
	run ~/cli-tools/docker/clean help
	assert_contains "$output" "Usage: clean"
}

@test "'clean' tool displays help for invalid parameter." {
	run ~/cli-tools/docker/clean invalid parameter
	assert_contains "$output" "Usage: clean"
}

@test "'clean ca' returns appropriate message when no exited containers." {
	docker pull busybox
	docker run busybox
	docker rm -f $(docker ps -aq)
	run ~/cli-tools/docker/clean ca
	assert_contains "$output" "No containers found."
}

@test "'clean ce' returns appropriate message when no exited containers." {
	docker run busybox
	~/cli-tools/docker/clean fr
	run ~/cli-tools/docker/clean ce
	assert_contains "$output" "No containers found."
}

@test "'clean ca' removes all containers." {
	docker run busybox
	LIST=$(docker ps -aq)
	echo "y" | ~/cli-tools/docker/clean ca
	run docker ps -aq
	refute_contains "$output" "$LIST"
}

@test "'clean ce' removes exited containers." {
	docker run busybox
	LIST=$(docker ps -q -f "status=exited")
	echo "y" | ~/cli-tools/docker/clean ce
	run docker ps -q -f "status=exited"
	refute_contains "$output" "$LIST"
}

@test "'clean fr' removes all exited containers." {
	docker run busybox
	LIST=$(docker ps -q -f "status=exited")
	~/cli-tools/docker/clean fr
	run docker ps -q -f "status=exited"
	refute_contains "$output" "$LIST"
}
