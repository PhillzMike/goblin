package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Zaida-3dO/goblin/pkg/errs"
)

var color = []string{
	"#512DA8", "#FD9E6C", "#4F98F0", "#F78284",
	"#26B5C6", "#8F9EE1", "#7C9EB0", "#4FB54A",
	"#26BEFC", "#105E62", "#83142C", "#93B5B3",
	"#76DBD1", "#A32F80", "#45454D", "#FF0000",
	"#70416D", "#000272", "#420000", "#CBA1D2",
	"#91B029", "#FF8080", "#E6E56C", "#A8FF3E",
	"#FBDA91", "#216583", "#274EA7", "#FF3F98",
}

type Number interface {
	int | int32 | int64 | float32 | float64
}

func getASCIICode(s string) (int32, error) {
	if len([]rune(s)) != 1 {
		return -1, errors.New("cannot get ascii code of anything but a character")
	}
	return []rune(s)[0], nil
}

func absDiffInt[T Number](x, y T) T {
	if x < y {
		return y - x
	}
	return x - y
}

func AbsInt[T Number](x T) T {
	return absDiffInt(x, 0)
}

func UserDefaultProfileColour(firstName, lastName string) (string, error) {
	fullName := fmt.Sprintf("%s%s", strings.ToLower(firstName), strings.ToLower(lastName))
	val := 0
	for i := 0; i < len([]rune(fullName)); i++ {
		asciiCode, err := getASCIICode(string(fullName[i]))
		if err != nil {
			return "", errors.New("cannot get ascii code of anything but a character")
		}
		val += int(asciiCode) - 97
	}
	val %= len(color)
	return color[AbsInt(val)], nil
}

func MustBePresent(input interface{}, check interface{}, keyBindings []string) ([]error, error) {
	inp := structConv(input)
	chk := structConv(check)

	arr := make([]error, 0, len(keyBindings))
	for _, key := range keyBindings {
		if _, ok := inp[key]; !ok {
			return []error{}, errors.New("invalid data interface for keybindings")
		}
		if inp[key] == chk[key] {
			arr = append(arr, errors.New(key))
		}
	}
	return arr, nil
}

func ValidateHttpRequestsForMissingFields(data interface{}, req interface{}, keyBindings []string) *errs.Err {
	missing, presentErr := MustBePresent(data, req, keyBindings)
	if presentErr != nil {
		return errs.NewInternalServerErr(presentErr.Error(), presentErr)
	}

	var err = errs.NewBadRequestErr("some fields are missing", nil)
	for _, el := range missing {
		err.Add(el)
	}

	if err.HasData() {
		return err
	}

	return nil
}

func structConv(d interface{}) map[string]interface{} {
	var _map map[string]interface{}
	r, _ := json.Marshal(d)
	json.Unmarshal(r, &_map)
	return _map
}
