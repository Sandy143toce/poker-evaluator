#!/bin/sh
atlas schema fmt --config file://schema.hcl && atlas schema apply \
  --url "postgresql://postgres:postgres@localhost:5432/poker_evaluator?sslmode=disable" \
  --to "file://schema.hcl"