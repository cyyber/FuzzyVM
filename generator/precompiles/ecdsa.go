// Copyright 2020 Marius van der Wijden
// This file is part of the fuzzy-vm library.
//
// The fuzzy-vm library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The fuzzy-vm library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the fuzzy-vm library. If not, see <http://www.gnu.org/licenses/>.

package precompiles

import (
	"crypto/ecdsa"

	"github.com/theQRL/FuzzyVM/filler"
	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/crypto"
	"github.com/theQRL/gozvmlab/program"
)

var ecdsaAddr, _ = common.NewAddressFromString("Z0000000000000000000000000000000000000001")

type ecdsaCaller struct{}

func (*ecdsaCaller) call(p *program.Program, f *filler.Filler) error {
	sk, err := ecdsa.GenerateKey(crypto.S256(), f)
	if err != nil {
		return err
	}
	sig, err := crypto.Sign(f.ByteSlice(32), sk)
	if err != nil {
		return err
	}
	// Sig is in [R | S | V] we need it in components
	c := CallObj{
		Gas:       f.GasInt(),
		Address:   ecdsaAddr,
		InOffset:  0,
		InSize:    uint32(len(sig)),
		OutOffset: 0,
		OutSize:   20,
		Value:     f.BigInt32(),
	}
	p.Mstore(sig, 0)
	CallRandomizer(p, f, c)
	return nil
}
