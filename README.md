# kata

kata is an application we'll build, one assignment at a time.

The assignment number will be posted in the #kata channel.

* Create a branch named after the assignment number + your github handle, and open a PR. My branch, for the first assignment, is called

`charliemcelfresh_001` <-- Note the padding in the assignment number.

Check new / old PRs for examples.

* We'll choose one engineer's assignment to merge.

* Please update this README with any ongoing instructions any engineer needs, in order to run existing code, or your new code, as part of each PR you do.

## Install

* Copy .env.sample to .env
* .env entries you need are pinned in the #kata channel
* Install dbmate, and run `dbmate up`, we're going to need the db soon

## Run
### Server

`go build`
`./kata server`

### Worker

`go build`
`./kata worker`