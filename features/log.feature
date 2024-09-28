Feature: saturn-bot plugin V1 log messages

  Scenario: Send a log message to main process of saturn-bot
    Given the context contains the repository "git.localhost/integration/log"
    When Apply is called
    Then the message "PLUGIN [integration-test stdout] Integration Test" is written to the log
    And the message "PLUGIN [integration-test stderr] Integration Test" is written to the log
