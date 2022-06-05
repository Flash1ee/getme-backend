package telegram_checker

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"reflect"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"

	"getme-backend/internal"
	"getme-backend/internal/app/auth/dto"
)

type TelegramChecker struct {
	log       *logrus.Logger
	token     string
	tokenHash hash.Hash
}

func NewTelegramChecker(log *logrus.Logger, authConf internal.TelegramAuth) *TelegramChecker {
	checker := &TelegramChecker{
		log:   log,
		token: authConf.Token,
	}
	checker.tokenHash = sha256.New()
	checker.tokenHash.Write([]byte(authConf.Token))

	return checker
}

func (service *TelegramChecker) Check(user *dto.AuthUsecase) bool {
	secretKey := service.tokenHash.Sum(nil)

	params := &CheckerData{}
	params = params.AuthToChecker(user)

	mapOfParams := structToMap(*params)
	//if err != nil {
	//	return false
	//}
	checkParams := make([]string, 0, len(mapOfParams))

	for k, v := range mapOfParams {
		if k != "hash" {
			checkParams = append(checkParams, fmt.Sprintf("%s=%v", k, v))
		}
	}
	sort.Strings(checkParams)
	checkString := strings.Join(checkParams, "\n")

	receivedHash := hmac.New(sha256.New, secretKey)
	receivedHash.Write([]byte(checkString))
	hashStr := hex.EncodeToString(receivedHash.Sum(nil))
	logrus.Infof("hash expected: %s, hash receive: %s\n", user.Hash, hashStr)
	if hashStr != user.Hash {
		logrus.Infof("hash received not equal expected")
		return false
	}
	logrus.Infof("success check auth")

	return true

}

func structToMap(val interface{}) map[string]interface{} {
	const tagTitle = "json"

	var data = make(map[string]interface{})
	varType := reflect.TypeOf(val)
	if varType.Kind() != reflect.Struct {
		// Provided value is not an interface, do what you will with that here
		fmt.Println("Not a struct")
		return nil
	}

	value := reflect.ValueOf(val)
	for i := 0; i < varType.NumField(); i++ {
		if !value.Field(i).CanInterface() {
			//Skip unexported fields
			continue
		}
		tag, ok := varType.Field(i).Tag.Lookup(tagTitle)
		var fieldName string
		if ok && len(tag) > 0 {
			fieldName = tag
		} else {
			fieldName = varType.Field(i).Name
		}
		if varType.Field(i).Type.Kind() != reflect.Struct {
			data[fieldName] = value.Field(i).Interface()
		} else {
			data[fieldName] = structToMap(value.Field(i).Interface())
		}

	}

	return data
}

//func structToMap(item interface{}) (map[string]interface{}, error) {
//	dataBytes, err := json.Marshal(item)
//	if err != nil {
//		return nil, err
//	}
//	mapData := make(map[string]interface{})
//	d := json.NewDecoder(bytes.NewBuffer(dataBytes))
//	d.UseNumber()
//	err = d.Decode(&mapData)
//	if err != nil {
//		return nil, err
//	}
//	return mapData, nil
//}
