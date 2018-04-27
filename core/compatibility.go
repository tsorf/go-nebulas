// Copyright (C) 2018 go-nebulas authors
//
// This file is part of the go-nebulas library.
//
// the go-nebulas library is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// the go-nebulas library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with the go-nebulas library.  If not, see <http://www.gnu.org/licenses/>.
//

package core

const (
	// TransferFromContractEventCompatibleHeight record event 'TransferFromContractEvent' since this height
	TransferFromContractEventCompatibleHeight uint64 = 200000

	// RandomSeedAvailableCompatibleHeight make 'random' available in contract since this height
	RandomSeedAvailableCompatibleHeight uint64 = 200000

	// CheckRandomSeedCompatibleHeight random seed need be check since this height
	CheckRandomSeedCompatibleHeight = RandomSeedAvailableCompatibleHeight + 1
)
