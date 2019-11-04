// Copyright © 2019 Ettore Di Giacinto <mudler@gentoo.org>
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, see <http://www.gnu.org/licenses/>.

package helpers_test

import (
	. "github.com/mudler/luet/pkg/helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	Context("Exists", func() {
		It("Detect existing and not-existing files", func() {
			Expect(Exists("../../tests/fixtures/buildtree/app-admin/enman/1.4.0/build.yaml")).To(BeTrue())
			Expect(Exists("../../tests/fixtures/buildtree/app-admin/enman/1.4.0/build.yaml.not.exists")).To(BeFalse())
		})
	})
})
