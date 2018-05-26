# Makefile for the ipar web server.
#
# This is arguably overkill as the build process for Go usually doesn't use
# makefiles, but the author is rather sick of having to remember where he
# can type `make cover` and where not.
all:
	@echo IPAR!
test:
	go test
cover:
	go test -cover
coverhtml:
	go test -coverprofile=cover.out && go tool cover -html=cover.out
bench:
	@echo BENCHMARKING SOON yes we love benchmarking!