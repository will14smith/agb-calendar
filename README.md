# agb-calendar

Copy `config.sh.sample` to `config.sh` and setup the variables

Run `source config.sh` to setup the environment

## Downloading the raw data

Check the start and end dates are correct in `download/main.go`

Clear (or backup) the `data/` folder from previous runs

Run `(cd download/ && go run *.go)`

## Processing the downloaded data

Delete (or backup) `process/tool/output.json`

Run `(cd process/tool/ && go run *.go)`

## Formatting the processed data

Delete (or backup) `format/tool/output.[html,ical]`

Run `(cd format/tool/ && go run *.go)`
