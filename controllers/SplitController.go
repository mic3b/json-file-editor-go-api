package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"translations/editor/models"

	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

// Dont mind about "Id", our primary key is "Key"!!!

func SplitController(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err.Error()))
		return
	}
	f, err := file.Open()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err.Error()))
		return
	}
	defer f.Close()

	var data map[string]interface{}
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err.Error()))
		return
	}

	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error: %s", err.Error()))
		return
	}

	dataS := string(jsonData)
	ioutil.WriteFile("temp/file.json", jsonData, 0o644)

	_, elements := getKeys("", data)

	for i, v := range elements {
		result := gjson.Get(dataS, v.Key)

		if result.IsArray() {
			var arr []string
			var arrInt []int
			numCond := true

			for _, v := range result.Array() {
				if v.Type.String() == "Number" {
					str, _ := strconv.Atoi(v.Raw)
					arrInt = append(arrInt, str)
					elements[i].Value = arrInt
				} else {
					numCond = false
					break
				}
			}

			if !numCond {
				for _, v := range result.Array() {
					if v.Type.String() == "Number" {
						arr = append(arr, string(v.Raw))
					} else {
						arr = append(arr, v.Str)
					}
				}

				elements[i].Value = arr
			}
		}
		if result.Type.String() == "String" {
			elements[i].Value = result.Str
		}

		if result.Type.String() == "Number" {
			elements[i].Value, _ = strconv.Atoi(result.Raw)
		}

	}

	m := make(map[int]int)
	for i, obj := range elements {
		if _, ok := m[obj.Id]; !ok {
			m[obj.Id] = i
		}
	}

	newArr := make([]models.Element, 0, len(m))
	for _, i := range m {
		newArr = append(newArr, elements[i])
	}

	fmt.Println(newArr)
	c.JSON(http.StatusOK, gin.H{
		"result": newArr,
		"error":  nil,
	})
}

func getKeys(prefix string, data interface{}) ([]string, []models.Element) {
	keys := []string{}
	Elements := []models.Element{}

	switch value := data.(type) {
	case map[string]interface{}:
		for k, v := range value {
			subkeys, _ := getKeys(fmt.Sprintf("%s%s.", prefix, k), v)
			keys = append(keys, subkeys...)
			splitedKey := strings.Split(keys[len(keys)-1], ".")
			name := splitedKey[len(splitedKey)-1]
			for i, v := range keys {
				Elem := models.Element{
					Id:    i,
					Name:  name,
					Value: "",
					Key:   v,
				}
				Elements = append(Elements, Elem)
			}

		}
	default:
		keys = append(keys, prefix[:len(prefix)-1])
	}

	return keys, Elements
}
