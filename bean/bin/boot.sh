#!/usr/bin/env sh

function main() {
  install_deps

  migrate

  seed

  run
}

function install_deps() {
  # tern for migration and seed
  go install github.com/jackc/tern/v2@latest

  # air for live reload
  go install github.com/cosmtrek/air@latest
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
