package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryPaymentTemplate            = "queryPaymentTemplate"
	QueryPaymentContract            = "queryPaymentContract"
	QueryPaymentContractsByIdPrefix = "queryPaymentContractsByIdPrefix"
	QuerySubscription               = "querySubscription"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryPaymentTemplate:
			return queryPaymentTemplate(ctx, path[1:], k, legacyQuerierCdc)
		case QueryPaymentContract:
			return queryPaymentContract(ctx, path[1:], k, legacyQuerierCdc)
		case QueryPaymentContractsByIdPrefix:
			return queryPaymentContractsByIdPrefix(ctx, path[1:], k, legacyQuerierCdc)
		case QuerySubscription:
			return querySubscription(ctx, path[1:], k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown payments query endpoint")
		}
	}
}

func queryPaymentTemplate(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	templateId := path[0]

	template, err := k.GetPaymentTemplate(ctx, templateId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"payment template '%s' does not exist", templateId)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, template)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal JSON: %s", err.Error())
	}

	return res, nil
}

func queryPaymentContract(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	contractId := path[0]

	contract, err := k.GetPaymentContract(ctx, contractId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "payment contract '%s' does not exist", contractId)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, contract)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal JSON: %s", err.Error())
	}

	return res, nil
}

func queryPaymentContractsByIdPrefix(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	contractIdPrefix := path[0]

	contracts := k.GetPaymentContractsByPrefix(ctx, contractIdPrefix)

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, contracts)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal JSON: %s", err.Error())
	}

	return res, nil
}

func querySubscription(ctx sdk.Context, path []string, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	subscriptionId := path[0]

	subscription, err := k.GetSubscription(ctx, subscriptionId)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "subscription '%s' does not exist", subscriptionId)
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, subscription)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal JSON: %s", err.Error())
	}

	return res, nil
}
