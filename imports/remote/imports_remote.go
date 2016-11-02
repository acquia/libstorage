package remote

import (
	// import to load
	_ "github.com/codedellemc/libstorage/drivers/storage/ebs/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/efs/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/isilon/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/objectivefs/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/scaleio/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/vbox/storage"
	_ "github.com/codedellemc/libstorage/drivers/storage/vfs/storage"
)
