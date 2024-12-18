// Code generated by protoc-gen-cobra. DO NOT EDIT.

package pactus

import (
	client "github.com/NathanBaulch/protoc-gen-cobra/client"
	flag "github.com/NathanBaulch/protoc-gen-cobra/flag"
	iocodec "github.com/NathanBaulch/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
)

func WalletClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("Wallet"),
		Short: "Wallet service client",
		Long:  "Define the Wallet service with various RPC methods for wallet management.",
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_WalletCreateWalletCommand(cfg),
		_WalletRestoreWalletCommand(cfg),
		_WalletLoadWalletCommand(cfg),
		_WalletUnloadWalletCommand(cfg),
		_WalletGetTotalBalanceCommand(cfg),
		_WalletSignRawTransactionCommand(cfg),
		_WalletGetValidatorAddressCommand(cfg),
		_WalletGetNewAddressCommand(cfg),
		_WalletGetAddressHistoryCommand(cfg),
		_WalletSignMessageCommand(cfg),
		_WalletGetTotalStakeCommand(cfg),
		_WalletGetAddressInfoCommand(cfg),
		_WalletSetAddressLabelCommand(cfg),
		_WalletListWalletCommand(cfg),
		_WalletGetWalletInfoCommand(cfg),
		_WalletListAddressCommand(cfg),
	)
	return cmd
}

func _WalletCreateWalletCommand(cfg *client.Config) *cobra.Command {
	req := &CreateWalletRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("CreateWallet"),
		Short: "CreateWallet RPC client",
		Long:  "CreateWallet creates a new wallet with the specified parameters.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "CreateWallet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &CreateWalletRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.CreateWallet(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the new wallet.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "The password for securing the wallet.")

	return cmd
}

func _WalletRestoreWalletCommand(cfg *client.Config) *cobra.Command {
	req := &RestoreWalletRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("RestoreWallet"),
		Short: "RestoreWallet RPC client",
		Long:  "RestoreWallet restores an existing wallet with the given mnemonic.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "RestoreWallet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &RestoreWalletRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.RestoreWallet(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to restore.")
	cmd.PersistentFlags().StringVar(&req.Mnemonic, cfg.FlagNamer("Mnemonic"), "", "The mnemonic for wallet recovery.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "The password for securing the wallet.")

	return cmd
}

func _WalletLoadWalletCommand(cfg *client.Config) *cobra.Command {
	req := &LoadWalletRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("LoadWallet"),
		Short: "LoadWallet RPC client",
		Long:  "LoadWallet loads an existing wallet with the given name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "LoadWallet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &LoadWalletRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.LoadWallet(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to load.")

	return cmd
}

func _WalletUnloadWalletCommand(cfg *client.Config) *cobra.Command {
	req := &UnloadWalletRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("UnloadWallet"),
		Short: "UnloadWallet RPC client",
		Long:  "UnloadWallet unloads a currently loaded wallet with the specified name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "UnloadWallet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &UnloadWalletRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.UnloadWallet(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to unload.")

	return cmd
}

func _WalletGetTotalBalanceCommand(cfg *client.Config) *cobra.Command {
	req := &GetTotalBalanceRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetTotalBalance"),
		Short: "GetTotalBalance RPC client",
		Long:  "GetTotalBalance returns the total available balance of the wallet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetTotalBalance"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetTotalBalanceRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetTotalBalance(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to get the total balance.")

	return cmd
}

func _WalletSignRawTransactionCommand(cfg *client.Config) *cobra.Command {
	req := &SignRawTransactionRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SignRawTransaction"),
		Short: "SignRawTransaction RPC client",
		Long:  "SignRawTransaction signs a raw transaction for a specified wallet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "SignRawTransaction"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &SignRawTransactionRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.SignRawTransaction(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet used for signing.")
	cmd.PersistentFlags().StringVar(&req.RawTransaction, cfg.FlagNamer("RawTransaction"), "", "The raw transaction data to be signed.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "The password for unlocking the wallet for signing.")

	return cmd
}

func _WalletGetValidatorAddressCommand(cfg *client.Config) *cobra.Command {
	req := &GetValidatorAddressRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetValidatorAddress"),
		Short: "GetValidatorAddress RPC client",
		Long:  "GetValidatorAddress retrieves the validator address associated with a\n public key.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetValidatorAddress"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetValidatorAddressRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetValidatorAddress(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.PublicKey, cfg.FlagNamer("PublicKey"), "", "The public key for which the validator address is requested.")

	return cmd
}

func _WalletGetNewAddressCommand(cfg *client.Config) *cobra.Command {
	req := &GetNewAddressRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetNewAddress"),
		Short: "GetNewAddress RPC client",
		Long:  "GetNewAddress generates a new address for the specified wallet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetNewAddress"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetNewAddressRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetNewAddress(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to generate a new address.")
	flag.EnumVar(cmd.PersistentFlags(), &req.AddressType, cfg.FlagNamer("AddressType"), "The type of address to generate.")
	cmd.PersistentFlags().StringVar(&req.Label, cfg.FlagNamer("Label"), "", "A label for the new address.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "Password for the new address. It's required when address_type is ADDRESS_TYPE_ED25519_ACCOUNT.")

	return cmd
}

func _WalletGetAddressHistoryCommand(cfg *client.Config) *cobra.Command {
	req := &GetAddressHistoryRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetAddressHistory"),
		Short: "GetAddressHistory RPC client",
		Long:  "GetAddressHistory retrieves the transaction history of an address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetAddressHistory"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetAddressHistoryRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetAddressHistory(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet.")
	cmd.PersistentFlags().StringVar(&req.Address, cfg.FlagNamer("Address"), "", "The address to retrieve the transaction history for.")

	return cmd
}

func _WalletSignMessageCommand(cfg *client.Config) *cobra.Command {
	req := &SignMessageRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SignMessage"),
		Short: "SignMessage RPC client",
		Long:  "SignMessage signs an arbitrary message.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "SignMessage"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &SignMessageRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.SignMessage(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "The password for unlocking the wallet for signing.")
	cmd.PersistentFlags().StringVar(&req.Address, cfg.FlagNamer("Address"), "", "The account address associated with the private key.")
	cmd.PersistentFlags().StringVar(&req.Message, cfg.FlagNamer("Message"), "", "The arbitrary message to be signed.")

	return cmd
}

func _WalletGetTotalStakeCommand(cfg *client.Config) *cobra.Command {
	req := &GetTotalStakeRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetTotalStake"),
		Short: "GetTotalStake RPC client",
		Long:  "GetTotalStake return total stake of wallet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetTotalStake"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetTotalStakeRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetTotalStake(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet.")

	return cmd
}

func _WalletGetAddressInfoCommand(cfg *client.Config) *cobra.Command {
	req := &GetAddressInfoRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetAddressInfo"),
		Short: "GetAddressInfo RPC client",
		Long:  "GetAddressInfo return address information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetAddressInfo"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetAddressInfoRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetAddressInfo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to generate a new address.")
	cmd.PersistentFlags().StringVar(&req.Address, cfg.FlagNamer("Address"), "", "")

	return cmd
}

func _WalletSetAddressLabelCommand(cfg *client.Config) *cobra.Command {
	req := &SetLabelRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("SetAddressLabel"),
		Short: "SetAddressLabel RPC client",
		Long:  "SetAddressLabel set label for given address.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "SetAddressLabel"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &SetLabelRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.SetAddressLabel(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "The name of the wallet to generate a new address.")
	cmd.PersistentFlags().StringVar(&req.Password, cfg.FlagNamer("Password"), "", "The password for unlocking the wallet for signing.")
	cmd.PersistentFlags().StringVar(&req.Address, cfg.FlagNamer("Address"), "", "")
	cmd.PersistentFlags().StringVar(&req.Label, cfg.FlagNamer("Label"), "", "")

	return cmd
}

func _WalletListWalletCommand(cfg *client.Config) *cobra.Command {
	req := &ListWalletRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListWallet"),
		Short: "ListWallet RPC client",
		Long:  "ListWallet return list wallet name.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "ListWallet"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &ListWalletRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListWallet(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	return cmd
}

func _WalletGetWalletInfoCommand(cfg *client.Config) *cobra.Command {
	req := &GetWalletInfoRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("GetWalletInfo"),
		Short: "GetWalletInfo RPC client",
		Long:  "GetWalletInfo return wallet information.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "GetWalletInfo"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &GetWalletInfoRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetWalletInfo(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "")

	return cmd
}

func _WalletListAddressCommand(cfg *client.Config) *cobra.Command {
	req := &ListAddressRequest{}

	cmd := &cobra.Command{
		Use:   cfg.CommandNamer("ListAddress"),
		Short: "ListAddress RPC client",
		Long:  "ListAddress return list address in wallet.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "Wallet", "ListAddress"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewWalletClient(cc)
				v := &ListAddressRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListAddress(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.WalletName, cfg.FlagNamer("WalletName"), "", "")

	return cmd
}
