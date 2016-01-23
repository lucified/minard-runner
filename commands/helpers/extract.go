package commands_helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"gitlab.com/gitlab-org/gitlab-ci-multi-runner/common"
	"gitlab.com/gitlab-org/gitlab-ci-multi-runner/helpers"
)

type ExtractCommand struct {
	File    string `long:"file" description:"The file to extract"`
	List    bool   `long:"list" description:"List files in archive"`
	Verbose bool   `long:"verbose" description:"Suppress archiving output"`
}

func (c *ExtractCommand) extractTarArchive() error {
	flags := "-zP"
	if c.List {
		flags += "t"
	} else {
		flags += "x"
	}
	if c.Verbose {
		flags += "v"
	}

	cmd := exec.Command("tar", flags, "-f", c.File)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	logrus.Debugln("Executing command:", strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func (c *ExtractCommand) extractFile(file *zip.File) (err error) {
	if c.Verbose && c.List {
		fmt.Println(helpers.ToJson(*file))
	} else if c.Verbose || c.List {
		fmt.Printf("%v\t%d\t%s\n", file.Mode(), file.UncompressedSize64, file.Name)
		if c.List {
			return
		}
	}

	fi := file.FileInfo()

	switch file.Mode() & os.ModeType {
	case os.ModeDir:
		err = os.MkdirAll(file.Name, fi.Mode().Perm())

	case os.ModeSymlink:
		var data []byte
		in, err := file.Open()
		if err == nil {
			defer in.Close()
			data, err = ioutil.ReadAll(in)
		}
		if err == nil {
			err = os.Symlink(string(data), file.Name)
		}

	case os.ModeNamedPipe, os.ModeSocket, os.ModeDevice:
		// Ignore the files that of these types
		logrus.Warningln("File ignored: %q", file.Name)

	default:
		var out *os.File
		in, err := file.Open()
		if err == nil {
			defer in.Close()
			out, err = os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fi.Mode().Perm())
		}
		if err == nil {
			defer out.Close()
			_, err = io.Copy(out, in)
			out.Close()
		}
		break
	}
	return
}

func (c *ExtractCommand) extractZipArchive() error {
	archive, err := zip.OpenReader(c.File)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, file := range archive.File {
		err = c.extractFile(file)
		if err != nil {
			logrus.Warningf("%s: %s", file.Name, err)
		}
	}

	for _, file := range archive.File {
		err := processZipExtra(file)
		if err != nil {
			logrus.Warningf("%s: %s", file.Name, err)
		}
	}
	return nil
}

func (c *ExtractCommand) extractArchive() error {
	if isTarArchive(c.File) {
		return c.extractTarArchive()
	} else if isZipArchive(c.File) {
		return c.extractZipArchive()
	} else {
		return fmt.Errorf("Unsupported archive format: %q", c.File)
	}
}

func (c *ExtractCommand) Execute(context *cli.Context) {
	if c.File == "" {
		logrus.Fatalln("Missing archive file name!")
	}

	err := c.extractArchive()
	if err != nil {
		logrus.Fatalln("Failed to create archive:", err)
	}
	logrus.Infoln("Done!")
}

func init() {
	common.RegisterCommand2("extract", "extract files from an archive (internal)", &ExtractCommand{})
}
