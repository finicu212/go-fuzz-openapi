package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

const (
	flagProxies           = "proxies"
	flagEndpoint          = "endpoint"
	flagRequestBody       = "body"
	flagOperation         = "operation"
	flagAttackAtDuration  = "delay"
	flagAttackAtTimeStamp = "delayRFC"
)

var loadTestCmd = &cobra.Command{
	Use:   fmt.Sprintf("loadtest --%s <IP_Proxy1,IP_Proxy2> --%s <operation> --%s <endpoint> --%s <request_minified_json> [--%s | --%s] <time>", flagProxies, flagOperation, flagEndpoint, flagRequestBody, flagAttackAtDuration, flagAttackAtTimeStamp),
	Short: `Test rate limit`,
	Long: `Test the API's rate limiting feature by synchronizing an attack on one or all of the endpoints using Proxy APIs.
It is mandatory that the proxy APIs provided should forward the request to the targeted API asynchronously, with the ability to delay the request.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.MarkFlagsMutuallyExclusive(flagAttackAtDuration, flagAttackAtTimeStamp)
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		url, _ := cmd.Flags().GetString(flagUrl)
		swagger, _ := cmd.Flags().GetString(flagSpec)
	},
}

func init() {
	rootCmd.AddCommand(loadTestCmd)

	fifteenSecs, _ := time.ParseDuration("15s")
	loadTestCmd.Flags().DurationP(flagAttackAtDuration, "d", fifteenSecs, "The time that the requests will all be sent at, as a duration from now. e.g 20s for twenty seconds")
	loadTestCmd.Flags().StringP(flagAttackAtTimeStamp, "", "", "The time that the requests will all be sent at in RFC3339 format, e.g 11:19:04Z")
	loadTestCmd.Flags().StringSliceP(flagProxies, "p", []string{"localhost:3000"}, "Slice of strings containing the URL or IP on which to reach the Proxies")
	loadTestCmd.Flags().StringP(flagEndpoint, "e", "pet", "The endpoint to forward all the requests to")
	loadTestCmd.Flags().StringP(flagOperation, "o", "POST", "The operation to use on the target endpoint")

	loadTestCmd.Flags().SortFlags = false
	loadTestCmd.SilenceUsage = true
}
