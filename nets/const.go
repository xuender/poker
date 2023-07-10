package nets

import (
	"crypto/sha1" // nolint: gosec

	"github.com/samber/lo"
	"github.com/xtaci/kcp-go/v5"
	"github.com/xuender/kit/base"
	"golang.org/x/crypto/pbkdf2"
)

const (
	Pass          = "nets pass"
	Salt          = "nets salt"
	Port          = "3880"
	_dataShards   = 10
	_parityShards = 3
	_bufSize      = 4096
)

// nolint: gochecknoglobals
var _block kcp.BlockCrypt

// nolint: gochecknoinits
func init() {
	key := pbkdf2.Key([]byte(Pass), []byte(Salt), base.Kilo, base.ThirtyTwo, sha1.New)
	_block = lo.Must1(kcp.NewAESBlockCrypt(key))
}
