#!/usr/bin/env bats

load test_helper

@test "HOST: Docker clean tool displays help." {
	run ~/cli-tools/docker/clean help
	assert_contains "$output" "Usage: clean"
}
@test "HOST: Docker clean tool displays help for invalid parameter." {
	run ~/cli-tools/docker/clean invalid parameter
	assert_contains "$output" "Usage: clean"
}
