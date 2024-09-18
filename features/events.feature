Feature: saturn-bot plugin V1 event functions

  Scenario: On Pull Request Closed
    Given the context contains the repository "git.localhost/integration/test"
    And the temporary file "saturn-bot_OnPrClosed.txt" is deleted
    And the plugin configuration:
      """
      {
        "event_out_tmp_file_path": "saturn-bot_OnPrClosed.txt"
      }
      """
    When OnPrClosed is called
    Then the content of temporary file "saturn-bot_OnPrClosed.txt" matches:
      """
      Integration Test OnPrClosed
      """

  Scenario: On Pull Request Created
    Given the context contains the repository "git.localhost/integration/test"
    And the temporary file "saturn-bot_OnPrCreated.txt" is deleted
    And the plugin configuration:
      """
      {
        "event_out_tmp_file_path": "saturn-bot_OnPrCreated.txt"
      }
      """
    When OnPrCreated is called
    Then the content of temporary file "saturn-bot_OnPrCreated.txt" matches:
      """
      Integration Test OnPrCreated
      """

  Scenario: On Pull Request Merged
    Given the context contains the repository "git.localhost/integration/test"
    And the temporary file "saturn-bot_OnPrMerged.txt" is deleted
    And the plugin configuration:
      """
      {
        "event_out_tmp_file_path": "saturn-bot_OnPrMerged.txt"
      }
      """
    When OnPrMerged is called
    Then the content of temporary file "saturn-bot_OnPrMerged.txt" matches:
      """
      Integration Test OnPrMerged
      """
