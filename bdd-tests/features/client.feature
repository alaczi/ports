Feature: Retrieving port data from the service

  Scenario Outline:
    Given the initial database of the ports was uploaded successfully
    When I request the port "<PortID>" from the service
    Then I receive a response with HTTP status <HTTPStatus>
    And the response content should be "<ContentType>" and the body should should match <ResponseBody>

    Examples:
      | PortID | HTTPStatus | ContentType      | ResponseBody                                                                                                                                                                                                                 |
      | BDMGL  | 200        | application/json | {"id":"BDMGL","code":"53800","name":"Mongla","city":"Mongla","country":"Bangladesh","timezone":"Asia/Dhaka","province":"Khulna Division","coordinates":[89.601616,22.494219],"alias":null,"regions":null,"unlocs":["BDMGL"]} |
      | NOPORT | 404        | application/json | {"code": 404, "error": "port not found"}                                                                                                                                                                                       |