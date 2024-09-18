Feature: saturn-bot plugin V1 Apply function

  Scenario: Modify repository
    Given the plugin configuration:
      """
      {
        "content": "Integration Test"
      }
      """
    And the context contains the repository "git.localhost/integration/test"
    And the context contains run data:
      """
      {
        "dynamic": "Dynamic Data"
      }
      """
    When Apply is called
    Then the response should match JSON:
      """
      {
        "error": "",
        "reply": {
          "run_data": {
            "dynamic": "Dynamic Data"
          }
        }
      }
      """
    And the file "integration-test.txt" exists with content:
      """
      Integration Test
      Dynamic Data
      """
