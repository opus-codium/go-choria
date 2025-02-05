// Copyright (c) 2020-2021, R.I. Pienaar and the Choria Project contributors
//
// SPDX-License-Identifier: Apache-2.0

package filesec

import (
	"crypto/sha256"
	"io"
	"os"
	"regexp"
	"runtime"
	"strings"
)

// MatchAnyRegex checks str against a list of possible regex, if any match true is returned
func MatchAnyRegex(str []byte, regex []string) bool {
	matcher := regexp.MustCompile("^/.+/$")

	for _, reg := range regex {
		if matcher.MatchString(reg) {
			reg = strings.TrimLeft(reg, "/")
			reg = strings.TrimRight(reg, "/")
		}

		if matched, _ := regexp.Match(reg, str); matched {
			return true
		}
	}

	return false
}

func uid() int {
	if useFakeUID {
		return fakeUID
	}

	return os.Geteuid()
}

func runtimeOs() string {
	if useFakeOS {
		return fakeOS
	}

	return runtime.GOOS
}

func fsha256(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := sha256.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
