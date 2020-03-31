pull_request_health_check "pull_request" {
  metric_factor "age" {
    metric_name = "PR age"
    scale {
      from = 72
      to = 24
    }
  }

  metric_factor "activity" {
    metric_name = "PR Activity level"
    scale {
      from = 100
      to = 0
    }
  }
}