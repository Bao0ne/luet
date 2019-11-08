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

package compiler

type Artifact interface {
	GetPath() string
	SetPath(string)
}

type PackageArtifact struct {
	Path string
}

func NewPackageArtifact(path string) Artifact {
	return &PackageArtifact{Path: path}
}

func (a *PackageArtifact) GetPath() string {
	return a.Path
}

func (a *PackageArtifact) SetPath(p string) {
	a.Path = p
}
