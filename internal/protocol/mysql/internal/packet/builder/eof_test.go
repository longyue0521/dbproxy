package builder

import (
	"testing"

	"github.com/meoying/dbproxy/internal/protocol/mysql/internal/flags"
	"github.com/stretchr/testify/assert"
)

func TestEOFPacket_Build(t *testing.T) {
	tests := []struct {
		name    string
		builder *EOFPacket
		want    []byte
	}{
		{
			name: "正常情况",
			builder: func() *EOFPacket {
				return NewEOFPacket(flags.CapabilityFlags(flags.ClientProtocol41), flags.ServerStatusAutoCommit)
			}(),
			want: []byte{
				0x05, 0x00, 0x00, 0x05, // packet header
				0xfe,       // header
				0x00, 0x00, // warnings
				0x02, 0x00, // status_flags
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want[4:], tt.builder.Build()[4:])
		})
	}
}
