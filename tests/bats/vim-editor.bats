#!/usr/bin/env bats

load test_helper

@test "Dependency check: ack, curl, git, man, screen, tree, vim, wget" {
	run docker exec test_editor ack
	assert_contains "$output" "Usage:"

	run docker exec test_editor curl --help
	assert_contains "$output" "Usage:"

	run docker exec test_editor git
	assert_contains "$output" "usage:"

	run docker exec test_editor man --help
	assert_contains "$output" "Usage:"

	run docker exec test_editor screen --help
	assert_contains "$output" "Use:"

	run docker exec test_editor tree --help
	assert_contains "$output" "usage:"

	run docker exec test_editor vim --help
	assert_contains "$output" "VIM"

	run docker exec test_editor wget --help
	assert_contains "$output" "Usage:"
}

