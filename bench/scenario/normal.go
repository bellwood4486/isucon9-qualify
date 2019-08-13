package scenario

import (
	"time"

	"github.com/isucon/isucon9-qualify/bench/asset"
	"github.com/isucon/isucon9-qualify/bench/fails"
	"github.com/isucon/isucon9-qualify/bench/session"
)

const (
	CorrectCardNumber = "AAAAAAAA"
	FailedCardNumber  = "FA10AAAA"
	IsucariShopID     = "11"
)

func sellAndBuy(user1, user2 asset.AppUser) error {
	s1, err := session.NewSession()
	if err != nil {
		return err
	}

	s2, err := session.NewSession()
	if err != nil {
		return err
	}

	seller, err := s1.Login(user1.AccountName, user1.Password)
	if err != nil {
		return err
	}

	if !user1.Equal(seller) {
		return fails.NewError(nil, "ログインが失敗しています")
	}

	err = s1.SetSettings()
	if err != nil {
		return err
	}

	buyer, err := s2.Login(user2.AccountName, user2.Password)
	if err != nil {
		return err
	}

	if !user2.Equal(buyer) {
		return fails.NewError(nil, "ログインが失敗しています")
	}

	err = s2.SetSettings()
	if err != nil {
		return err
	}

	targetItemID, err := s1.Sell("abcd", 100, "description description", 32)
	if err != nil {
		return err
	}
	token, err := s2.PaymentCard(CorrectCardNumber, IsucariShopID)
	if err != nil {
		return err
	}
	_, err = s2.Buy(targetItemID, token)
	if err != nil {
		return err
	}

	apath, err := s1.Ship(targetItemID)
	if err != nil {
		return err
	}

	surl, err := s1.DecodeQRURL(apath)
	if err != nil {
		return err
	}

	s3, err := session.NewSession()
	if err != nil {
		return err
	}

	err = s3.ShipmentAccept(surl)
	if err != nil {
		return err
	}

	err = s1.ShipDone(targetItemID)
	if err != nil {
		return err
	}

	<-time.After(6 * time.Second)

	err = s2.Complete(targetItemID)
	if err != nil {
		return err
	}

	return nil
}

func bump(user1, user2 asset.AppUser) error {
	s1, err := session.NewSession()
	if err != nil {
		return err
	}

	s2, err := session.NewSession()
	if err != nil {
		return err
	}

	seller, err := s1.Login(user1.AccountName, user1.Password)
	if err != nil {
		return err
	}

	if !user1.Equal(seller) {
		return fails.NewError(nil, "ログインが失敗しています")
	}

	err = s1.SetSettings()
	if err != nil {
		return err
	}

	buyer, err := s2.Login(user2.AccountName, user2.Password)
	if err != nil {
		return err
	}

	if !user2.Equal(buyer) {
		return fails.NewError(nil, "ログインが失敗しています")
	}

	err = s2.SetSettings()
	if err != nil {
		return err
	}

	err = s1.Bump(1)
	if err != nil {
		return err
	}

	return nil
}