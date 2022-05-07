package setting

import (
	"strconv"

	harvesterv1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	"github.com/sirupsen/logrus"
)

func (h *Handler) setSRIOVVirtualFunctionCount(setting *harvesterv1.Setting) error {
	numvfs, err := strconv.Atoi(setting.Value)
	if err != nil {
		return err
	}

	logrus.Infof("set SR-IOV VFs Count to %s", numvfs)
	return nil
}
