package text

import (
	"regexp"
	"strings"
)

var reDomain = regexp.MustCompile(`^[a-z0-9][a-z0-9\.\-]{1,61}[a-z0-9]$`)
var reIPAddress = regexp.MustCompile(`^(\d+\.){3}\d+$`)

func DnsCompatibleBucketName(bucket string) bool {
	return reDomain.MatchString(bucket) &&
		!reIPAddress.MatchString(bucket) &&
		!strings.Contains(bucket, "..")
}
