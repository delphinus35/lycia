package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
)

var Commands = []cli.Command{
	commandOpen,
}

var openFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "root",
		Value: ".",
		Usage: "Specify root dir for repository",
	},
	cli.StringFlag{
		Name:  "ref, r",
		Value: "master",
		Usage: "Ref to open such as branch, tag and hash",
	},
	cli.IntFlag{
		Name:  "from, f",
		Usage: "Line to highlight from",
	},
	cli.IntFlag{
		Name:  "to, t",
		Usage: "Line to highlight to",
	},
}

var commandOpen = cli.Command{
	Name:    "open",
	Aliases: []string{"o"},
	Usage:   "Open github repository page",
	Description: `
Open URL link that can be identified by git-remote command.
If you do not specify the path, you can see the top page.
And if you specify, you can see the path on Web with highlighted lines.
`,
	Action: doOpen,
	Flags:  openFlags,
}

func doOpen(c *cli.Context) {
	argPath := c.Args().Get(0)
	root := c.String("root")
	ref := c.String("ref")
	from := c.Int("from")
	to := c.Int("to")

	url, err := RemoteURL(root, ref, argPath, from, to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "remote url not found: %s\n", err)
	} else {
		fmt.Printf("opening url: \"%s\"...\n", url.String())
		cmd := exec.Command("open", url.String())
		cmd.Run()
	}
}
