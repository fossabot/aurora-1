// Copyright 2016 - 2018 The aurora Authors. All rights reserved. Use of this
// source code is governed by a MIT license that can be found in the LICENSE
// file.
//
// The aurora is a web-based Beanstalk queue server console written in Go
// and works on macOS, Linux and Windows machines. Main idea behind using Go
// for backend development is to utilize ability of the compiler to produce
// zero-dependency binaries for multiple platforms. aurora was created as an
// attempt to build very simple and portable application to work with local or
// remote Beanstalk server.

package main

// currentTube call currentTubeJobs by given server and tube config.
func currentTube(server string, tube string) string {
	return currentTubeJobs(server, tube)
}
