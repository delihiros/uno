package cmd

import (
	"github.com/delihiros/uno/pkg/client"
	"github.com/delihiros/uno/pkg/discord"
	"github.com/delihiros/uno/pkg/jsonutil"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var (
	Prettify bool
	Region   string
	Name     string
	Tag      string
	Filter   string
	MatchID  string
)

var (
	rootCmd = &cobra.Command{
		Use:   "uno",
		Short: "Unofficial Valorant API client",
	}

	wardellCmd = &cobra.Command{
		Use:   "wardell",
		Short: "wardell joins to your discord server!!!",
		RunE: func(cmd *cobra.Command, args []string) error {
			wardell, err := discord.New(viper.Get("discord_token").(string))
			if err != nil {
				return err
			}
			return wardell.Execute()
		},
	}

	accountCmd = &cobra.Command{
		Use:              "account",
		Short:            "account related commands",
		TraverseChildren: true,
	}

	accountInfoCmd = &cobra.Command{
		Use:   "info",
		Short: "get account information by player name and tagline",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			account, err := c.GetAccountByNameTag(Name, Tag)
			if err != nil {
				return err
			}
			return jsonutil.PrintJSON(account, Prettify)
		},
	}

	mmrCmd = &cobra.Command{
		Use: "elo",
	}

	mmrSeasonCmd = &cobra.Command{
		Use:   "season",
		Short: "get season level account elo by player name and tagline",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			mmr, err := c.GetMMRDataByNameTag(Region, Name, Tag)
			if err != nil {
				return err
			}
			return jsonutil.PrintJSON(mmr, Prettify)
		},
	}

	mmrLatestCmd = &cobra.Command{
		Use:   "latest",
		Short: "get latest account elo by player name and tagline",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			history, err := c.GetMMRHistoryByNameTag(Region, Name, Tag)
			if err != nil {
				return err
			}
			return jsonutil.PrintJSON(history, Prettify)
		},
	}

	matchHistoryCmd = &cobra.Command{
		Use:   "matches",
		Short: "get match history",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			matches, err := c.GetMatchHistory(Region, Name, Tag, Filter)
			if err != nil {
				return err
			}
			return jsonutil.PrintJSON(matches, Prettify)
		},
	}

	matchCmd = &cobra.Command{
		Use:   "match",
		Short: "get match information by match ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			c := client.New()
			match, err := c.GetMatchByID(MatchID)
			if err != nil {
				return err
			}
			return jsonutil.PrintJSON(match, Prettify)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.PersistentFlags().BoolVarP(&Prettify, "prettify", "p", false, "prettify output json")

	accountCmd.Flags().StringVarP(&Name, "name", "n", "bobobobobobobo", "player name")
	accountCmd.Flags().StringVarP(&Tag, "tag", "t", "1212", "player tagline")

	matchHistoryCmd.Flags().StringVarP(&Filter, "filter", "f", "", "filter")
	matchHistoryCmd.Flags().StringVarP(&Region, "region", "r", "ap", "region")

	matchCmd.Flags().StringVarP(&MatchID, "match_id", "m", "2aa59334-e53a-415b-bb3d-4832305ee7db", "match ID")

	mmrCmd.AddCommand(mmrSeasonCmd)
	mmrCmd.AddCommand(mmrLatestCmd)

	accountCmd.AddCommand(accountInfoCmd)
	accountCmd.AddCommand(matchHistoryCmd)
	accountCmd.AddCommand(mmrCmd)

	rootCmd.AddCommand(wardellCmd)
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(matchCmd)
}
