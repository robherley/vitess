/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mysqlctl

import (
	"testing"
)

type testcase struct {
	versionString string
	version       ServerVersion
	flavor        MySQLFlavor
}

func TestParseVersionString(t *testing.T) {

	var testcases = []testcase{

		{
			versionString: "mysqld  Ver 5.7.27-0ubuntu0.19.04.1 for Linux on x86_64 ((Ubuntu))",
			version:       ServerVersion{5, 7, 27},
			flavor:        FlavorMySQL,
		},
		{
			versionString: "mysqld  Ver 5.6.43 for linux-glibc2.12 on x86_64 (MySQL Community Server (GPL))",
			version:       ServerVersion{5, 6, 43},
			flavor:        FlavorMySQL,
		},
		{
			versionString: "mysqld  Ver 5.7.26 for linux-glibc2.12 on x86_64 (MySQL Community Server (GPL))",
			version:       ServerVersion{5, 7, 26},
			flavor:        FlavorMySQL,
		},
		{
			versionString: "mysqld  Ver 8.0.16 for linux-glibc2.12 on x86_64 (MySQL Community Server - GPL)",
			version:       ServerVersion{8, 0, 16},
			flavor:        FlavorMySQL,
		},
		{
			versionString: "mysqld  Ver 5.7.26-29 for Linux on x86_64 (Percona Server (GPL), Release 29, Revision 11ad961)",
			version:       ServerVersion{5, 7, 26},
			flavor:        FlavorPercona,
		},
		{
			versionString: "mysqld  Ver 10.0.38-MariaDB for Linux on x86_64 (MariaDB Server)",
			version:       ServerVersion{10, 0, 38},
			flavor:        FlavorMariaDB,
		},
		{
			versionString: "mysqld  Ver 10.1.40-MariaDB for Linux on x86_64 (MariaDB Server)",
			version:       ServerVersion{10, 1, 40},
			flavor:        FlavorMariaDB,
		},
		{
			versionString: "mysqld  Ver 10.2.25-MariaDB for Linux on x86_64 (MariaDB Server)",
			version:       ServerVersion{10, 2, 25},
			flavor:        FlavorMariaDB,
		},
		{
			versionString: "mysqld  Ver 10.3.16-MariaDB for Linux on x86_64 (MariaDB Server)",
			version:       ServerVersion{10, 3, 16},
			flavor:        FlavorMariaDB,
		},
		{
			versionString: "mysqld  Ver 10.4.6-MariaDB for Linux on x86_64 (MariaDB Server)",
			version:       ServerVersion{10, 4, 6},
			flavor:        FlavorMariaDB,
		},
		{
			versionString: "mysqld  Ver 5.6.42 for linux-glibc2.12 on x86_64 (MySQL Community Server (GPL))",
			version:       ServerVersion{5, 6, 42},
			flavor:        FlavorMySQL,
		},
		{
			versionString: "mysqld  Ver 5.6.44-86.0 for Linux on x86_64 (Percona Server (GPL), Release 86.0, Revision eba1b3f)",
			version:       ServerVersion{5, 6, 44},
			flavor:        FlavorPercona,
		},
		{
			versionString: "mysqld  Ver 8.0.15-6 for Linux on x86_64 (Percona Server (GPL), Release 6, Revision 63abd08)",
			version:       ServerVersion{8, 0, 15},
			flavor:        FlavorPercona,
		},
	}

	for _, testcase := range testcases {
		f, v, err := ParseVersionString(testcase.versionString)
		if v != testcase.version || f != testcase.flavor || err != nil {
			t.Errorf("ParseVersionString failed for: %#v, Got: %#v, %#v Expected: %#v, %#v", testcase.versionString, v, f, testcase.version, testcase.flavor)
		}
	}

}
