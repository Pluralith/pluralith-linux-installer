package assets

type Images struct {
	PluralithIcon []byte
	DownloadBadge []byte
	InstallBadge  []byte
	CompleteBadge []byte
}

// Instantiating asset store
var ImageStore = &Images{}
