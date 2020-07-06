package tmplfuncs

import (
	"errors"
	"net"
	"os"
	"strconv"
	"strings"
)

var Funcs = map[string]interface{}{
	// built-in functions
	"netResolveIPAddr":  net.ResolveIPAddr,
	"osHostname":        os.Hostname,
	"osUserCacheDir":    os.UserCacheDir,
	"osUserConfigDir":   os.UserConfigDir,
	"osUserHomeDir":     os.UserHomeDir,
	"stringsContains":   strings.Contains,
	"stringsFields":     strings.Fields,
	"stringsIndex":      strings.Index,
	"stringsLastIndex":  strings.LastIndex,
	"stringsHasPrefix":  strings.HasPrefix,
	"stringsHasSuffix":  strings.HasSuffix,
	"stringsRepeat":     strings.Repeat,
	"stringsReplaceAll": strings.ReplaceAll,
	"stringsSplit":      strings.Split,
	"stringsToLower":    strings.ToLower,
	"stringsToUpper":    strings.ToUpper,
	"stringsTrimPrefix": strings.TrimPrefix,
	"stringsTrimSpace":  strings.TrimSpace,
	"stringsTrimSuffix": strings.TrimSuffix,

	// extra functions
	"k8sStatefulSetID": k8sStatefulSetID,
}

func k8sStatefulSetID() (id int64, err error) {
	var hostname string
	if hostname = os.Getenv("HOSTNAME"); hostname == "" {
		if hostname, err = os.Hostname(); err != nil {
			return
		}
	}
	splits := strings.Split(hostname, "-")
	if len(splits) < 2 {
		err = errors.New("invalid stateful-set hostname")
		return
	}
	id, err = strconv.ParseInt(splits[len(splits)-1], 10, 64)
	return
}
