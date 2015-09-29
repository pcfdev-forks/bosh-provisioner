package packagescompiler

import (
	boshblob "github.com/cloudfoundry/bosh-agent/blobstore"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	bpagclient "github.com/sclevine/bosh-provisioner/agent/client"
	bpeventlog "github.com/sclevine/bosh-provisioner/eventlog"
	bpcpkgsrepo "github.com/sclevine/bosh-provisioner/packagescompiler/compiledpackagesrepo"
	bppkgsrepo "github.com/sclevine/bosh-provisioner/packagescompiler/packagesrepo"
)

type ConcretePackagesCompilerFactory struct {
	packagesRepo         bppkgsrepo.PackagesRepository
	compiledPackagesRepo bpcpkgsrepo.CompiledPackagesRepository
	blobstore            boshblob.Blobstore

	eventLog bpeventlog.Log
	logger   boshlog.Logger
}

func NewConcretePackagesCompilerFactory(
	packagesRepo bppkgsrepo.PackagesRepository,
	compiledPackagesRepo bpcpkgsrepo.CompiledPackagesRepository,
	blobstore boshblob.Blobstore,
	eventLog bpeventlog.Log,
	logger boshlog.Logger,
) ConcretePackagesCompilerFactory {
	return ConcretePackagesCompilerFactory{
		packagesRepo:         packagesRepo,
		compiledPackagesRepo: compiledPackagesRepo,
		blobstore:            blobstore,

		eventLog: eventLog,
		logger:   logger,
	}
}

func (f ConcretePackagesCompilerFactory) NewCompiler(agentClient bpagclient.Client) PackagesCompiler {
	return NewConcretePackagesCompiler(
		agentClient,
		f.packagesRepo,
		f.compiledPackagesRepo,
		f.blobstore,
		f.eventLog,
		f.logger,
	)
}
