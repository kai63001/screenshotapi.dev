package module

type ImageFormat struct {
	ID      int64  `db:"id" json:"id"`
	Format  string `db:"format" json:"format"`
	Quality int64  `db:"quality" json:"quality"`
	Width   int64  `db:"width" json:"width"`
	Height  int64  `db:"height" json:"height"`
}
