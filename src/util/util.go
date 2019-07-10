package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/user/marvel/src/config"
	"strconv"
	"time"
)

func GetAuthParams() string {
	conf := config.New()

	now := time.Now()
	timestamp := strconv.FormatInt(now.UTC().UnixNano(), 10)

	str := string(timestamp) + conf.Marvel.PvtKey + conf.Marvel.PubKey

	hasher := md5.New()
	hasher.Write([]byte(str))
	hash := hex.EncodeToString(hasher.Sum(nil))

	params := "?apikey="+ conf.Marvel.PubKey + "&hash=" + hash + "&ts=" + timestamp

	return params
}
