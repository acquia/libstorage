package handlers

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/akutz/gotil"
	"github.com/emccode/libstorage/api/types"
	"github.com/emccode/libstorage/api/types/context"
	apihttp "github.com/emccode/libstorage/api/types/http"
	"github.com/emccode/libstorage/api/utils"
)

// localDevicesHandler is a global HTTP filter for grokking the local devices
// from the headers
type localDevicesHandler struct {
	handler apihttp.APIFunc
}

// NewLocalDevicesHandler returns a new global HTTP filter for grokking the
// local devices from the headers
func NewLocalDevicesHandler() apihttp.Middleware {
	return &localDevicesHandler{}
}

func (h *localDevicesHandler) Name() string {
	return "local-devices-handler"
}

func (h *localDevicesHandler) Handler(m apihttp.APIFunc) apihttp.APIFunc {
	return (&localDevicesHandler{m}).Handle
}

var (
	locDevRX = regexp.MustCompile(`^(.+?)=((?:(?:.+?)=(?:.+?)(?:,\s*)?)+)$`)
)

// Handle is the type's Handler function.
func (h *localDevicesHandler) Handle(
	ctx context.Context,
	w http.ResponseWriter,
	req *http.Request,
	store types.Store) error {

	headers := utils.GetHeader(req.Header, apihttp.LocalDevicesHeader)
	ctx.Log().WithField(
		apihttp.LocalDevicesHeader, headers).Debug("http header")

	valMap := map[string]map[string]string{}
	for _, v := range headers {

		locDevM := locDevRX.FindStringSubmatch(v)
		driver := locDevM[1]

		devMntPairs := strings.Split(locDevM[2], ",")
		for _, dmp := range devMntPairs {
			dmpParts := strings.Split(dmp, "=")
			device := gotil.Trim(dmpParts[0])
			mountPoint := gotil.Trim(dmpParts[1])
			if _, ok := valMap[driver]; !ok {
				valMap[driver] = map[string]string{}
			}
			valMap[driver][device] = mountPoint
		}
	}

	if len(valMap) > 0 {
		ctx = ctx.WithLocalDevicesByService(valMap)
	}

	return h.handler(ctx, w, req, store)
}