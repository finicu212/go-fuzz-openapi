package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go_fuzz_openapi/pkg/loadmaker"
	"go_fuzz_openapi/pkg/xtraflag"
	"strings"
	"time"
)

const (
	flagProxiesToOperation = "proxies"
	flagAttackAtDuration   = "delay"
	flagAttackAtTimeStamp  = "delayRFC"
	//flagRequestBody        = "body"
)

func loadTestCmd() *cobra.Command {
	var proxiesToOperation map[string][]string

	cmd := &cobra.Command{
		Use:   fmt.Sprintf("loadtest [-pb] [--%s | --%s] <time>", flagAttackAtDuration, flagAttackAtTimeStamp),
		Short: `Test rate limit`,
		Long: `Test the API's rate limiting feature by synchronizing an attack on one or all of the endpoints using Proxy APIs.
It is mandatory that the proxy APIs provided should forward the request to the targeted API asynchronously, with the ability to delay the request.`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			cmd.MarkFlagsMutuallyExclusive(flagAttackAtDuration, flagAttackAtTimeStamp)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			delay, _ := cmd.Flags().GetDuration(flagAttackAtDuration)
			delayRFC, _ := cmd.Flags().GetString(flagAttackAtTimeStamp)

			if delayRFC != "" {
				t, err := time.Parse("15:04:05Z", delayRFC)
				if err != nil {
					return fmt.Errorf("could not parse %s as RFC3339 format: %w", delayRFC, err)
				}
				delay = time.Until(t)
			}

			//url, _ := cmd.Flags().GetString(flagUrl)
			//swagger, _ := cmd.Flags().GetString(flagSpec)
			//body, _ := cmd.Flags().GetString(flagRequestBody)

			proxyMaster := loadmaker.NewProxyCoordinator(delay)
			for p, ops := range proxiesToOperation {
				for _, o := range ops {
					method, endpoint, ok := strings.Cut(o, ":")
					if !ok {
						return fmt.Errorf("%s is not separated by `:` for --%s", flagProxiesToOperation, ops)
					}
					var err error
					proxyMaster, err = proxyMaster.AddLoadMaker(p, endpoint, method)
					if err != nil {
						return err
					}
				}
			}
			_ = proxyMaster

			return nil
		},
	}

	cmd.Flags().VarP(xtraflag.NewValue(
		map[string][]string{"localhost:3000": {"POST:/pet"}},
		&proxiesToOperation,
		xtraflag.StringToStringSliceParser(",", "=", "|")),
		flagProxiesToOperation, "p", "Map a proxy API to one or multiple operations. E.g. localhost:3000=POST:/pet|GET:/pet,localhost:3001=POST:/store|GET:/user/{username}")

	cmd.Flags().DurationP(flagAttackAtDuration, "d", 3*time.Second, "The time that approximately all the requests will be sent at, as a duration from now. e.g 20s for twenty seconds")
	cmd.Flags().StringP(flagAttackAtTimeStamp, "", "", "The time that the requests will all be sent at in RFC3339 format, e.g 11:19:04Z")
	//cmd.Flags().StringP(flagRequestBody, "b", "", "The method to use on the target endpoint, e.g. POST, PUT, GET")

	cmd.Flags().SortFlags = false
	cmd.SilenceUsage = true

	return cmd
}

func init() {
	rootCmd.AddCommand(loadTestCmd())

}
