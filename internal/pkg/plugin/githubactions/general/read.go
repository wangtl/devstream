package general

import (
	"github.com/devstream-io/devstream/internal/pkg/configmanager"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/ci"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/ci/cifile"
	"github.com/devstream-io/devstream/pkg/util/log"
)

func Read(options map[string]interface{}) (map[string]interface{}, error) {
	// Initialize Operator with Operations
	operator := &plugininstaller.Operator{
		PreExecuteOperations: plugininstaller.PreExecuteOperations{
			ci.SetDefault(ciType),
			validate,
		},
		GetStatusOperation: cifile.GetCIFileStatus,
	}

	// Execute all Operations in Operator
	status, err := operator.Execute(configmanager.RawOptions(options))
	if err != nil {
		return nil, err
	}
	log.Debugf("Return map: %v", status)
	return status, nil
}
