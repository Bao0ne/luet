// Copyright © 2019-2020 Ettore Di Giacinto <mudler@gentoo.org>,
//                       Daniele Rondina <geaaru@sabayonlinux.org>
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

package spectooling_test

import (
	"github.com/mudler/luet/pkg/api/core/types"

	. "github.com/mudler/luet/pkg/spectooling"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Spec Tooling", func() {
	Context("Conversion1", func() {

		b := types.NewPackage("B", "1.0", []*types.Package{}, []*types.Package{})
		c := types.NewPackage("C", "1.0", []*types.Package{}, []*types.Package{})
		d := types.NewPackage("D", "1.0", []*types.Package{}, []*types.Package{})
		p1 := types.NewPackage("A", "1.0", []*types.Package{b, c}, []*types.Package{d})
		virtual := types.NewPackage("E", "1.0", []*types.Package{}, []*types.Package{})
		virtual.SetCategory("virtual")
		p1.Provides = []*types.Package{virtual}
		p1.AddLabel("label1", "value1")
		p1.AddLabel("label2", "value2")
		p1.SetDescription("Package1")
		p1.SetCategory("cat1")
		p1.SetLicense("GPL")
		p1.AddURI("https://github.com/mudler/luet")
		p1.AddUse("systemd")
		It("Convert pkg1", func() {
			res := NewDefaultPackageSanitized(p1)
			expected_res := &PackageSanitized{
				Name:     "A",
				Version:  "1.0",
				Category: "cat1",
				PackageRequires: []*PackageSanitized{
					&PackageSanitized{
						Name:    "B",
						Version: "1.0",
					},
					&PackageSanitized{
						Name:    "C",
						Version: "1.0",
					},
				},
				PackageConflicts: []*PackageSanitized{
					&PackageSanitized{
						Name:    "D",
						Version: "1.0",
					},
				},
				Provides: []*PackageSanitized{
					&PackageSanitized{
						Name:     "E",
						Category: "virtual",
						Version:  "1.0",
					},
				},
				Labels: map[string]string{
					"label1": "value1",
					"label2": "value2",
				},
				Description: "Package1",
				License:     "GPL",
				Uri:         []string{"https://github.com/mudler/luet"},
				UseFlags:    []string{"systemd"},
			}

			Expect(res).To(Equal(expected_res))
		})

	})
})
