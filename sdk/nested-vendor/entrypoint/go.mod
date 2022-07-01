module entrypoint

go 1.18

require (
	github.com/LucasRoesler/openfaas-examples-and-tests/sdk/nested-vendor v0.0.0-20220701134155-836b1ee47691
	handler v0.0.0-00010101000000-000000000000
)

replace handler => ./handler
