package facebook

import (
	"testing"

	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/test"
)

func TestDownload(t *testing.T) {
	config.InfoOnly = true
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.facebook.com/JackyardsCovers/videos/vb.267832806658747/1215502888558396/",
				Title:   "Jackyards - Great video, beautifully shot down Bournemouth...",
				Quality: "hd",
			},
		},
		{
			name: "normal test",
			args: test.Args{
				URL:     "https://www.facebook.com/groups/314070194112/permalink/10155168902769113/",
				Title:   "Ukrainian Scientists Worldwide Public Group | Facebook",
				Size:    336975453,
				Quality: "hd",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Download(tt.args.URL)
			test.CheckError(t, err)
			test.Check(t, tt.args, data[0])
		})
	}
}
