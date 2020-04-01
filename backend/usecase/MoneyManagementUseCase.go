package usecase

import (
	"errors"
	"github.com/BambooTuna/quest-market/backend/aggregate"
	"github.com/BambooTuna/quest-market/backend/model/transaction"
)

type MoneyManagementUseCase struct {
	ManagementAccountId        string
	MoneyTransactionAggregates *aggregate.MoneyTransactionAggregates
}

func (m *MoneyManagementUseCase) GetBalance(accountId string) (int64, error) {
	return m.MoneyTransactionAggregates.GetBalance(accountId)
}

func (m *MoneyManagementUseCase) SendMoney(senderAccountId, receiverAccountId string, money int64) error {
	if money <= 0 {
		return errors.New("金額は正の整数で入力してください")
	}
	senderMoneyTransaction := transaction.ApplyMoneyTransaction(senderAccountId, transaction.Withdraw, money)
	receiverMoneyTransaction := transaction.ApplyMoneyTransaction(receiverAccountId, transaction.Deposit, money)
	err1 := m.MoneyTransactionAggregates.SendTransaction(senderMoneyTransaction)
	err2 := m.MoneyTransactionAggregates.SendTransaction(receiverMoneyTransaction)
	if err1 != nil || err2 != nil {
		return errors.New("SendMoneyTo: 深刻なエラー")
	}
	return nil
}

func (m *MoneyManagementUseCase) ManagementKeeps(fromAccountId string, money int64) error {
	if err := m.SendMoney(fromAccountId, m.ManagementAccountId, money); err != nil {
		return err
	}
	return nil
}

func (m *MoneyManagementUseCase) ManagementPayment(toAccountId string, money int64) error {
	if err := m.SendMoney(m.ManagementAccountId, toAccountId, money); err != nil {
		return err
	}
	return nil
}
