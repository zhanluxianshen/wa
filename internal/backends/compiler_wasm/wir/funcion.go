package wir

import "github.com/wa-lang/wa/internal/backends/compiler_wasm/wir/wtypes"

func (f *Function) Format(indent string) string {
	s := indent + "(func $" + f.Name + " (export \"" + f.Name + "\")"

	for _, param := range f.Params {
		rps := param.Raw()
		for _, rp := range rps {
			s += " (param " + rp.Name() + " " + rp.Type().Name() + ")"
		}
	}

	if !f.Result.Equal(wtypes.Void{}) {
		s += " (result"
		rrs := f.Result.Raw()
		for _, rr := range rrs {
			s += " " + rr.Name()
		}
		s += ")"
	}
	s += "\n"

	for _, local := range f.Locals {
		rls := local.Raw()
		s += indent + " "
		for _, rl := range rls {
			s += " (local " + rl.Name() + " " + rl.Type().Name() + ")"
		}
		s += "\n"
	}

	for _, inst := range f.Insts {
		s += inst.Format(indent+"  ") + "\n"
	}

	s += indent + ") ;;" + f.Name + "\n"
	return s
}
