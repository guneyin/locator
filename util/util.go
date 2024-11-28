package util

import (
	"strconv"
	"time"
)

type VersionInfo struct {
	Version    string
	CommitHash string
	BuildTime  string
}

var (
	Version    string
	CommitHash string
	BuildTime  string

	lastRun time.Time
)

func GetVersion() *VersionInfo {
	return &VersionInfo{
		Version:    Version,
		CommitHash: CommitHash,
		BuildTime:  BuildTime,
	}
}

func SetLastRun(t time.Time) {
	lastRun = t
}

func GetLastRun() time.Time {
	return lastRun
}

func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func StrToFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
