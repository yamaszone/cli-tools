Feature: Jokes Tool
In order to make development process enjoyable,
I should be able to read jokes from terminal

Scenario: Jokes can be read from CLI
	Given CLI jokes tool is installed
	When I execute "joke"
	Then I should see "setup"
	And I should see "punchline"