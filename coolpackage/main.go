package coolpackage

import ver "github.com/AislingHPE/simple-go-package/pkg/interfaces"

var v = ver.VersionHolder{
	Version: "v0.0.5",
}

func GetVersion() string {
	return "Version " + v.Version
}
