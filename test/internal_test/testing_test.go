package internal_test

// Uncomment when needed (vim: 5GvGgcc)

import (
	"testing"

	"github.com/PAiP-Web-Language/stdlib-go/dev"
	"github.com/PAiP-Web-Language/stdlib-go/test"
	. "github.com/PAiP-Web-Language/stdlib-go/test/internal"
)

func TestInterfaces(t *testing.T) {
	t.SkipNow()
	dev.Unreachable(
		"This is mainly for checking if interfaces work according to gopls",
		"Also for testing if interfaces work according to tests (because if not they fail to build)",
	)
	// var t testing.T
	// var b testing.B
	// var f testing.F

	var tp *testing.T = &testing.T{}
	var bp *testing.B = &testing.B{}
	var fp *testing.F = &testing.F{}

	var tb testing.TB = tp


	var ti TI = tp
	var bi BI = bp
	var fi FI = fp
	var tbfi TBFI = tb
	var tbfi1 TBFI = tp
	var tbfi2 TBFI = bp
	var tbfi3 TBFI = fp
	// Errored because of Run types mismatch
	// var tbi1 TBI = tp
	// var tbi2 TBI = bp

	_= ti
	_= bi
	_= fi
	_= tbfi
	_= tbfi1
	_= tbfi2
	_= tbfi3
	// _= tbi1
	// _= tbi2


	var itp *T = test.WrapT(tp)
	var ibp *B = test.WrapB(bp)
	var ifp *F = test.WrapF(fp)
	var itb *TB = test.WrapTB(tb)


	var iti TI = itp
	var ibi BI = ibp
	var ifi FIT = ifp
	var itbfi TBFI = itb
	var itbfi1 TBFI = itp
	var itbfi2 TBFI = ibp
	var itbfi3 TBFI = ifp
	// Errored because of Run types mismatch
	var itbi1 TBI = itp
	var itbi2 TBI = ibp

	_= iti
	_= ibi
	_= ifi
	_= itbfi
	_= itbfi1
	_= itbfi2
	_= itbfi3
	_= itbi1
	_= itbi2
}
