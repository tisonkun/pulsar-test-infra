// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	// Prefix
	InfoPrefix  = "[INFO] "
	ErrorPrefix = "[ERROR] "
	FatalPrefix = "[FATAL] "

	// Color
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"

	// background
	BgRed      = "\033[41m"
	BgLightRed = "\033[101m"
)

func Infoln(v ...interface{}) {
	log.New(os.Stderr, Cyan+InfoPrefix+Reset, log.LstdFlags).Output(2, fmt.Sprintln(v...))
}

func Infof(format string, v ...interface{}) {
	log.New(os.Stderr, Cyan+InfoPrefix+Reset, log.LstdFlags).Output(2, fmt.Sprintf(format, v...))
}

func Errorln(v ...interface{}) {
	log.New(os.Stderr, Red+ErrorPrefix+Reset, log.LstdFlags|log.Llongfile).Output(2, fmt.Sprintln(v...))
}

func Errorf(format string, v ...interface{}) {
	log.New(os.Stderr, Red+ErrorPrefix+Reset, log.LstdFlags|log.Llongfile).Output(2, fmt.Sprintf(format, v...))
}

func Fatalf(format string, v ...interface{}) {
	log.New(os.Stderr, Red+FatalPrefix+Reset, log.LstdFlags|log.Llongfile).Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	log.New(os.Stderr, Red+FatalPrefix+Reset, log.LstdFlags|log.Llongfile).Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}
