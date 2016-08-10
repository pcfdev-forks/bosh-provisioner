package updater

import (
	"fmt"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"

	bpeventlog "github.com/cppforlife/bosh-provisioner/eventlog"
	bpapplier "github.com/cppforlife/bosh-provisioner/instance/updater/applier"
)

const updaterLogTag = "Updater"

type Updater struct {
	instanceDesc string

	drainer Drainer
	applier bpapplier.Applier

	eventLog bpeventlog.Log
	logger   boshlog.Logger
}

func NewUpdater(
	instanceDesc string,
	drainer Drainer,
	applier bpapplier.Applier,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) Updater {
	return Updater{
		instanceDesc: instanceDesc,

		drainer: drainer,
		applier: applier,

		eventLog: eventLog,
		logger:   logger,
	}
}

func (u Updater) SetUp() error {
	stage := u.eventLog.BeginStage(fmt.Sprintf("Setting up instance %s", u.instanceDesc), 3)

	task := stage.BeginTask("Applying")

	err := task.End(u.applier.Apply())
	if err != nil {
		return bosherr.WrapError(err, "Applying")
	}

	return nil
}

func (u Updater) TearDown() error {
	stage := u.eventLog.BeginStage(fmt.Sprintf("Tearing down instance %s", u.instanceDesc), 2)

	task := stage.BeginTask("Draining")

	err := task.End(u.drainer.Drain())
	if err != nil {
		return bosherr.WrapError(err, "Draining")
	}

	return nil
}
