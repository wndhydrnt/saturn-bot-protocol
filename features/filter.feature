Feature: saturn-bot plugin V1 Filter function

  Scenario: Filter matches repository
    Given the context contains the repository "git.localhost/integration/test"
    When Filter is called
    Then the response should match JSON:
      """
      {
        "error": "",
        "reply": {
          "match": true
        }
      }
      """

  Scenario: Filter does not match repository
    Given the context contains the repository "git.localhost/other/test"
    When Filter is called
    Then the response should match JSON:
      """
      {
        "error": "",
        "reply": {}
      }
      """

  Scenario: Filter sets run data
    Given the context contains the repository "git.localhost/integration/rundata"
    And the context contains run data:
      """
      {
        "saturn-bot": "set by saturn-bot"
      }
      """
    When Filter is called
    Then the response should match JSON:
      """
      {
        "error": "",
        "reply": {
          "match": true,
          "run_data": {
            "plugin": "set by plugin",
            "saturn-bot": "set by saturn-bot"
          }
        }
      }
      """
