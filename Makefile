OUTDIR:=./out
COVER:=$(OUTDIR)/cover


tests:
	@ go test ./... -coverprofile $(COVER)/cover.out
	@ go tool cover -html $(COVER)/cover.out