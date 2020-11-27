package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zlingqu/go-template-tool/service"
)

var (
	config string
	src    string
	dest   string
)

// NewGenetateYmlCommand cobra实例化
func NewGenetateYmlCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-template-tool",
		Short: "生成yml文件",
		Long:  "go-template-tool 用于生成使用模板生成yml文件",
		// Args:  cobra.ExactValidArgs(3), //三个参数必须都有
		Example: `go-template-tool --config /tmp/svc.config --src src_dir  --dest dest_dir
go-template-tool --config /tmp/svc.config --src a.tmpl  --dest a.yml
go-template-tool --config ./svc.config --src a.tmpl`,
		Run: func(cmd *cobra.Command, args []string) {
			service.GenetateYml(config, src, dest)
		},
	}
	rootCmd.Flags().StringVarP(&config, "config", "C", "", "配置文件路径,例如 ./")
	rootCmd.Flags().StringVarP(&src, "src", "S", "", "模板文件路径，写成目录或者文件，目录表示所有。文件名必须使用.tmpl结尾")
	rootCmd.Flags().StringVarP(&dest, "dest", "D", "", "生成文件的路径，可以写成目录或则文件，文件名和模板文件名相同，并以.yml结尾")
	return rootCmd
}
