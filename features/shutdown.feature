Feature: saturn-bot plugin V1 shutdown function

  Scenario: Call shutdown function
    When Shutdown is called
    Then the message "PLUGIN [integration-test stdout] Shutdown called" is written to the log
