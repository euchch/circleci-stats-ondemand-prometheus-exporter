# circleci-stats-ondemand-prometheus-exporter
CircleCI statistics exporter with external trigger

Very losely based of https://github.com/chaspy/circleci-insights-prometheus-exporter
Biggest change - running a local queue of which repo/branch to update instead of setting it up in advance

## Environment Variable

|name                 |required|default |description|
|---------------------|--------|--------|-----------|
|CIRCLECI_TOKEN       |yes     |-       |[CircleCI API Token](https://app.circleci.com/settings/user/tokens)|
|CIRCLECI_POLING_INTERVAL|no      |300(sec)|Interval second in between polings|
