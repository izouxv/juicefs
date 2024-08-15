//go:build nofuse

/*
 * JuiceFS, Copyright 2023 Juicedata, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import (
	"github.com/juicedata/juicefs/pkg/meta"
	"github.com/juicedata/juicefs/pkg/object"
	"github.com/juicedata/juicefs/pkg/vfs"
	"github.com/urfave/cli/v2"
)

func mountFlags() []cli.Flag {
	return []cli.Flag{}
}

func makeDaemon(c *cli.Context, conf *vfs.Config) error {
	logger.Warnf("Fuse NA")
	return nil
}

func makeDaemonForSvc(c *cli.Context, m meta.Meta, metaUrl, listenAddr string) error {
	logger.Warnf("Fuse NA")
	return nil
}

func getDaemonStage() int {
	return 0
}

func mountMain(v *vfs.VFS, c *cli.Context) {
	logger.Warnf("Fuse NA")
}

func checkMountpoint(name, mp, logPath string, background bool) {
	logger.Warnf("Fuse NA")
}

func prepareMp(mp string) {}

func setFuseOption(c *cli.Context, format *meta.Format, vfsConf *vfs.Config) {
	logger.Warnf("Fuse NA")
}

func launchMount(mp string, conf *vfs.Config) error {
	logger.Warnf("Fuse NA")
	return nil
}

func installHandler(mp string, v *vfs.VFS, blob object.ObjectStorage) {
	logger.Warnf("Fuse NA")
}
