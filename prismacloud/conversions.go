package prismacloud

import (
	"bytes"
	"encoding/base64"
	"strings"

	"github.com/paloaltonetworks/prisma-cloud-go/timerange"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const IdSeparator = ":"

func TwoStringsToId(a, b string) string {
	return strings.Join([]string{a, b}, IdSeparator)
}

func IdToTwoStrings(v string) (string, string) {
	t := strings.Split(v, IdSeparator)
	return t[0], t[1]
}

func ListToStringSlice(v []interface{}) []string {
	if len(v) == 0 {
		return []string{}
	}

	ans := make([]string, len(v))
	for i := range v {
		switch x := v[i].(type) {
		case nil:
			ans[i] = ""
		case string:
			ans[i] = x
		}
	}

	return ans
}

func SetToStringSlice(d *schema.Set) []string {
	list := d.List()
	return ListToStringSlice(list)
}

func StringSliceToSet(list []string) *schema.Set {
	items := make([]interface{}, len(list))
	for i := range list {
		items[i] = list[i]
	}

	return schema.NewSet(schema.HashString, items)
}

func ResourceDataInterfaceMap(d *schema.ResourceData, key string) map[string]interface{} {
	if _, ok := d.GetOk(key); ok {
		if v1, ok := d.Get(key).([]interface{}); ok && len(v1) != 0 {
			if v2, ok := v1[0].(map[string]interface{}); ok && v2 != nil {
				return v2
			}
		}
	}

	return map[string]interface{}{}
}

func ToInterfaceMap(m map[string]interface{}, k string) map[string]interface{} {
	if _, ok := m[k]; ok {
		if v1, ok := m[k].([]interface{}); ok && len(v1) != 0 {
			if v2, ok := v1[0].(map[string]interface{}); ok && v2 != nil {
				return v2
			}
		}
	}

	return map[string]interface{}{}
}

func ParseTimeRange(tr map[string]interface{}) timerange.TimeRange {
	if data := ToInterfaceMap(tr, "absolute"); len(data) != 0 {
		return timerange.TimeRange{
			Value: timerange.Absolute{
				Start: data["start"].(int),
				End:   data["end"].(int),
			},
		}
	} else if data := ToInterfaceMap(tr, "relative"); len(data) != 0 {
		return timerange.TimeRange{
			Value: timerange.Relative{
				Amount: data["amount"].(int),
				Unit:   data["unit"].(string),
			},
		}
	} else if data := ToInterfaceMap(tr, "to_now"); len(data) != 0 {
		return timerange.TimeRange{
			Value: data["unit"].(string),
		}
	}

	return timerange.TimeRange{}
}

func Base64Encode(v []interface{}) string {
	var buf bytes.Buffer

	for i := range v {
		if i != 0 {
			buf.WriteString("\n")
		}
		buf.WriteString(v[i].(string))
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func Base64Decode(v string) []string {
	joined, err := base64.StdEncoding.DecodeString(v)
	if err != nil {
		return nil
	}

	return strings.Split(string(joined), "\n")
}
