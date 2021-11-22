package json_schema

import (
	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
	"io/ioutil"
	"json_schema/model"
	"sigs.k8s.io/yaml"
	"testing"
)

func Test_string(t *testing.T) {
	schema_str := `{
		"type":"object",
		"properties":{
			"name":{
				"type":"string"
			}
		},
		"required": ["name"]
	}`
	schema := gojsonschema.NewStringLoader(schema_str)
	json_str := `{}`
	json := gojsonschema.NewStringLoader(json_str)
	res, err := gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_str = `{"age":123}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_str = `{"name":123}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_str = `{"name":"string"}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
}

func Test_map(t *testing.T) {
	m := map[string]interface{}{
		"type":       "object",
		"properties": map[string]interface{}{"name": "string"},
		"required":   []string{"name"},
	}
	schema := gojsonschema.NewGoLoader(m)
	json_map := map[string]interface{}{}
	json := gojsonschema.NewGoLoader(json_map)
	res, err := gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_map = map[string]interface{}{
		"age": 123,
	}
	json = gojsonschema.NewGoLoader(json_map)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_map = map[string]interface{}{
		"name": 123,
	}
	json = gojsonschema.NewGoLoader(json_map)
	assert.Nil(t, err)
	assert.False(t, res.Valid())

	json_map = map[string]interface{}{
		"name": "123",
	}
	json = gojsonschema.NewGoLoader(json_map)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
}

func Test_struct(t *testing.T) {
	schema_str := `{
		"type": "object",
		"properties":{
			"cluster":{
				"type":"object",
				"properties":{
					"name":{
						"type":"string"
					},
					"kubeconfig":{
						"type":"string"
					}
				},
				"required":["name"]
			},
			"namespace":{
				"type":"string"
			},
			"name":{
				"type":"string"
			},
			"database":{
				"type":"array",
				"items":{
					"type":"string"
				}
			}
		},
		"required":["cluster", "namespace", "name", "database"]
	}`
	schema := gojsonschema.NewStringLoader(schema_str)
	app := model.App{
		Cluster: model.Cluster{
			Name:       "cluster",
			Kubeconfig: "kubeconfig",
		},
		Namespace: "namespace",
		Name:      "name",
		Database:  []string{"postgresql", "redis", "mysql"},
	}
	json := gojsonschema.NewGoLoader(app)
	res, err := gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
}

func Test_Pattern(t *testing.T) {
	bs, err := ioutil.ReadFile("./pattern.json")
	assert.Nil(t, err)
	schema := gojsonschema.NewStringLoader(string(bs))
	json_str := `{
		"data":"abc"
	}`
	json := gojsonschema.NewStringLoader(json_str)
	res, err := gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
	json_str = `{"data":"abc65213"}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
	json_str = `{"data":"a423"}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.False(t, res.Valid())
	json_str = `{
		"data":"abc",
		"url":"https://123"
	}`
	json = gojsonschema.NewStringLoader(json_str)
	res, err = gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
}

func Test_Kubeconfig(t *testing.T) {
	bs, err := ioutil.ReadFile("./file/kubeconfig.yaml")
	assert.Nil(t, err)
	m := map[string]interface{}{}
	err = yaml.Unmarshal(bs, &m)
	t.Log(m)
	json := gojsonschema.NewGoLoader(m)

	schemaBs, err := ioutil.ReadFile("./file/kubeschema.json")
	assert.Nil(t, err)
	schema := gojsonschema.NewStringLoader(string(schemaBs))

	res, err := gojsonschema.Validate(schema, json)
	assert.Nil(t, err)
	assert.True(t, res.Valid())
	t.Log(res.Errors())
}
