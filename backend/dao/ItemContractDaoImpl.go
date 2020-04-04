package dao

import (
	"fmt"
	"github.com/BambooTuna/quest-market/backend/model/item"
	"github.com/BambooTuna/quest-market/backend/settings"
	"gopkg.in/gorp.v1"
)

type ItemContractDaoImpl struct {
	DBSession *gorp.DbMap
}

func (i ItemContractDaoImpl) JoinWhere() string {
	return "select item_details.item_id,item_details.title,item_details.detail,item_details.price,contract_details.purchaser_account_id,contract_details.seller_account_id,contract_details.state,contract_details.created_at,contract_details.updated_at from item_details,contract_details where item_details.item_id = contract_details.item_id"
}

func (i ItemContractDaoImpl) Publishable(q settings.QuantityLimit) []*item.ContractDetails {
	var result []*item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.state != '%s' and contract_details.state != '%s' ORDER BY contract_details.updated_at desc Limit %d,%d", i.JoinWhere(), item.Deleted, item.Draft, q.Drop(), q.Limit)
	println("Publishable: " + sql)
	i.DBSession.Select(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolvePublishableItemByItemId(itemId string) *item.ContractDetails {
	var result *item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.item_id = '%s' and contract_details.state != '%s' and contract_details.state != '%s'", i.JoinWhere(), itemId, item.Deleted, item.Draft)
	println("ResolvePublishableItemByItemId: " + sql)
	i.DBSession.SelectOne(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolvePrivateItemByItemId(itemId, practitioner string) *item.ContractDetails {
	var result *item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.item_id = '%s' and contract_details.state != '%s' and ((contract_details.seller_account_id = '%s') or (contract_details.state != '%s'))", i.JoinWhere(), itemId, item.Deleted, practitioner, item.Draft)
	println("ResolvePrivateItemByItemId: " + sql)
	i.DBSession.SelectOne(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolveByAccountId(q settings.QuantityLimit, accountId string) []*item.ContractDetails {
	var result []*item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.state != '%s' and contract_details.state != '%s' and (contract_details.purchaser_account_id = '%s' or contract_details.seller_account_id = '%s') ORDER BY contract_details.updated_at desc Limit %d,%d", i.JoinWhere(), item.Deleted, item.Draft, accountId, accountId, q.Drop(), q.Limit)
	println("ResolveByAccountId: " + sql)
	i.DBSession.Select(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolveOpenItemByItemId(itemId string) *item.ContractDetails {
	var result *item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.item_id = '%s' and contract_details.state = '%s'", i.JoinWhere(), itemId, item.Open)
	println("ResolveOpenItemByItemId: " + sql)
	i.DBSession.SelectOne(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolveUnpaidItemByItemId(itemId, purchaserAccountId string) *item.ContractDetails {
	var result *item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.item_id = '%s' and contract_details.state = '%s' and contract_details.purchaser_account_id = '%s'", i.JoinWhere(), itemId, item.Unpaid, purchaserAccountId)
	println("ResolveUnpaidItemByItemId: " + sql)
	i.DBSession.SelectOne(&result, sql)
	return result
}

func (i ItemContractDaoImpl) ResolveSentItemByItemId(itemId, purchaserAccountId string) *item.ContractDetails {
	var result *item.ContractDetails
	sql := fmt.Sprintf("%s and contract_details.item_id = '%s' and contract_details.state = '%s' and contract_details.purchaser_account_id = '%s'", i.JoinWhere(), itemId, item.Sent, purchaserAccountId)
	println("ResolveSentItemByItemId: " + sql)
	i.DBSession.SelectOne(&result, sql)
	return result
}

func (i ItemContractDaoImpl) Insert(record *item.ContractDetails) error {
	itemDetails := ItemDetails.FromContractDetails(ItemDetails{}, record)
	contractDetails := ContractDetails.FromContractDetails(ContractDetails{}, record)
	if err := i.DBSession.Insert(itemDetails); err != nil {
		return err
	} else if err := i.DBSession.Insert(contractDetails); err != nil {
		return err
	} else {
		return nil
	}
}

func (i ItemContractDaoImpl) UpdateItemDetails(record *item.ContractDetails) error {
	itemDetails := ItemDetails.FromContractDetails(ItemDetails{}, record)
	if i, err := i.DBSession.Update(itemDetails); err != nil || i == 0 {
		return err
	} else {
		return nil
	}
}

func (i ItemContractDaoImpl) UpdateContractDetails(record *item.ContractDetails) error {
	contractDetails := ContractDetails.FromContractDetails(ContractDetails{}, record)
	if i, err := i.DBSession.Update(contractDetails); err != nil || i == 0 {
		return err
	} else {
		return nil
	}
}
