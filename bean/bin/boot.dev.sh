#!/usr/bin/env sh

function main() {
  migrate

  seed

  run
}

function migrate() {
  tern migrate -c $TERN_CONFIG -m $TERN_MIGRATIONS
}

function seed() {
  tern migrate -c $SEED_TERN_CONFIG -m $SEED_TERN_MIGRATIONS
}

function run() {
  air -c ./bean/config/.air.toml
}

main
