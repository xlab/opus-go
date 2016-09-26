all:
	cgogen opus.yml

clean:
	rm -f opus/cgo_helpers.go opus/cgo_helpers.h opus/cgo_helpers.c
	rm -f opus/doc.go opus/types.go opus/const.go
	rm -f opus/opus.go

test:
	cd opus && go build
