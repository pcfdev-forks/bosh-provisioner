package main

import (
	"path/filepath"

	boshblob "github.com/cloudfoundry/bosh-agent/blobstore"
	boshlog "github.com/cloudfoundry/bosh-agent/logger"
	boshsys "github.com/cloudfoundry/bosh-agent/system"

	bpdload "github.com/sclevine/bosh-provisioner/downloader"
	bpindex "github.com/sclevine/bosh-provisioner/index"
	bpjobsrepo "github.com/sclevine/bosh-provisioner/instance/templatescompiler/jobsrepo"
	bptplsrepo "github.com/sclevine/bosh-provisioner/instance/templatescompiler/templatesrepo"
	bpcpkgsrepo "github.com/sclevine/bosh-provisioner/packagescompiler/compiledpackagesrepo"
	bppkgsrepo "github.com/sclevine/bosh-provisioner/packagescompiler/packagesrepo"
)

type ReposFactory struct {
	dirPath    string
	fs         boshsys.FileSystem
	downloader bpdload.Downloader
	blobstore  boshblob.Blobstore
	logger     boshlog.Logger
}

func NewReposFactory(
	dirPath string,
	fs boshsys.FileSystem,
	downloader bpdload.Downloader,
	blobstore boshblob.Blobstore,
	logger boshlog.Logger,
) ReposFactory {
	return ReposFactory{
		dirPath:    dirPath,
		fs:         fs,
		downloader: downloader,
		blobstore:  blobstore,
		logger:     logger,
	}
}

func (f ReposFactory) NewJobsRepo() bpjobsrepo.JobsRepository {
	return bpjobsrepo.NewConcreteJobsRepository(
		f.newIndex("jobs"),
		f.logger,
	)
}

func (f ReposFactory) NewTemplateToJobRepo() bpjobsrepo.TemplateToJobRepository {
	return bpjobsrepo.NewConcreteTemplateToJobRepository(
		f.newIndex("templates_to_job"),
		f.logger,
	)
}

func (f ReposFactory) NewRuntimePackagesRepo() bpjobsrepo.RuntimePackagesRepository {
	return bpjobsrepo.NewConcreteRuntimePackagesRepository(
		f.newIndex("runtime_packages"),
		f.logger,
	)
}

func (f ReposFactory) NewTemplatesRepo() bptplsrepo.TemplatesRepository {
	return bptplsrepo.NewConcreteTemplatesRepository(
		f.newIndex("templates"),
		f.logger,
	)
}

func (f ReposFactory) NewPackagesRepo() bppkgsrepo.PackagesRepository {
	return bppkgsrepo.NewConcretePackagesRepository(
		f.newIndex("packages"),
		f.logger,
	)
}

func (f ReposFactory) NewCompiledPackagesRepo() bpcpkgsrepo.CompiledPackagesRepository {
	return bpcpkgsrepo.NewConcreteCompiledPackagesRepository(
		f.newIndex("compiled_packages"),
		f.logger,
	)
}

func (f ReposFactory) newIndex(name string) bpindex.Index {
	return bpindex.NewFileIndex(filepath.Join(f.dirPath, name+".json"), f.fs)
}
