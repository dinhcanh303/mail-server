package utils

import (
	"log/slog"
	"reflect"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func StringToNullUUID(uuidString string) (uuid.NullUUID, error) {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return uuid.NullUUID{
			UUID:  parsedUUID,
			Valid: false,
		}, errors.Wrap(err, "Parsed UUID failed")
	}
	return uuid.NullUUID{
		UUID:  parsedUUID,
		Valid: true,
	}, nil
}
func StringToUUID(uuidString string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, "Parsed UUID failed")
	}
	return parsedUUID, nil
}

func StringToNullUUIDNormal(uuidString string) uuid.NullUUID {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		slog.Warn("Parsed UUID failed")
	}
	return uuid.NullUUID{
		UUID:  parsedUUID,
		Valid: true,
	}
}
func StringToUUIDNormal(uuidString string) uuid.UUID {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		slog.Warn("Parsed UUID failed")
	}
	return parsedUUID
}
func ConvertArUUIDToArString(uuids []uuid.UUID) []string {
	results := make([]string, 0)
	for _, uuid := range uuids {
		results = append(results, uuid.String())
	}
	return results
}

func ConvertArStringToArUUID(strings []string) ([]uuid.UUID, error) {
	uuids := make([]uuid.UUID, 0)
	for _, str := range strings {
		uuid, err := StringToUUID(str)
		if err != nil {
			return nil, errors.Wrap(err, "Converted UUID failed")
		}
		uuids = append(uuids, uuid)
	}
	return uuids, nil
}
func ConvertArStringToArNullUUID(strings []string) ([]uuid.NullUUID, error) {
	uuids := make([]uuid.NullUUID, 0)
	for _, str := range strings {
		uuid, err := StringToNullUUID(str)
		if err != nil {
			return nil, errors.Wrap(err, "Converted UUID failed")
		}
		uuids = append(uuids, uuid)
	}
	return uuids, nil
}
func LoadFileEnvOnLocal() error {

	err := godotenv.Load("../../.env")
	if err != nil {
		return err
	}
	return nil
}
func Contains[T any](arr []T, x T) bool {
	for _, v := range arr {
		if reflect.ValueOf(v) == reflect.ValueOf(x) {
			return true
		}
	}
	return false
}
func ContainsFunc[T any](arr []T, predicate func(T) bool) bool {
	for _, v := range arr {
		if predicate(v) {
			return true
		}
	}
	return false
}
func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
func FormatInt(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func UniqueSlice[T comparable](inputSlice []T) []T {
	uniqueSlice := make([]T, 0, len(inputSlice))
	seen := make(map[T]bool, len(inputSlice))
	for _, ele := range inputSlice {
		if !seen[ele] {
			uniqueSlice = append(uniqueSlice, ele)
			seen[ele] = true
		}
	}
	return uniqueSlice
}
func RemoveSlice[T comparable](inputSlice *[]T, removeItems ...T) []T {
	results := make([]T, 0, len(*inputSlice))
	// var end bool
	return results
}
