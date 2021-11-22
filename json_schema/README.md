# json schema

### if then else

```json
{
  "type": "object",
  "properties": {
    "street": {
      "type": "string"
    },
    "country": {
      "enum": ["United States of America", "Canada"]
    }
  },
  "if": {
    "properties": {
      "country": {
        "const": "United States of America"
      }
    }
  },
  "then": {
    "properties": {
      "postal_code": {
        "pattern": "[0-9]{5}(-[0-9]{4})?"
      }
    }
  },
  "else": {
    "properties": {
      "postal_code": {
        "pattern": "[A-Z0-9]{6}"
      }
    }
  }
}
```

```json
{
  "type": "object",
  "properties": {
    "street_address": {
      "type": "string"
    },
    "country": {
      "enum": ["US", "Canada", "Netherlands"]
    }
  },
  "allOf": [
    {
      "if": {
        "properties": {
          "country": {
            "const": "US"
          }
        }
      },
      "then": {
        "properties": {
          "postal_code": {
            "pattern": "[0-9]{5}(-[0-9]{4}?)"
          }
        }
      }
    },{
      "if": {
        "properties": {
          "country": {
            "const": "Netherlands"
          }
        }
      },
      "then": {
        "properties": {
          "postal_code": {
            "pattern": "[0-9]{4} [A-Z]{2}"
          }
        }
      }
    },{
      "if": {
        "properties": {
          "country": {
            "const": "Canada"
          }
        }
      },
      "then": {
        "properties": {
          "postal_code": {
            "pattern": "[A-Z0-9]{6}"
          }
        }
      }
    }
  ]
}
```

## string

- maxLength
- minLength
- enum
- format
  - date
  - time
  - datetime
  - email
  - idn-email
  - hostname
  - idn-hostname
  - ipv4
  - ipv6
  - uri
  - uri-reference
  - iri
  - iri-reference
  - uri-template
  - json-pointer
  - relative-json-pointer
  - regex

## boolean

## integer

- multipleOf
- maximum小于等于
- exclusiveMaximum
- minimum
- exclusiveMinimum

## number

- multipleOf
- maximum
- exclusiveMaximum
- minimum
- exclusiveMinimum

## const

## object

- required
- minProperties
- maxProperties
- patternProperties
  - 验证key满足条件的value是否符合规定
- additionalProperties
  - 验证不匹配条件key的value
- dependencies略
- propertyNames
  - 验证key是否满足该json schema

## array

- maxItems
- minItems
- items
- additionalItems
  - 超出规定位置时有效
  - 当items为一个json schema数组时有效
- uniqueItems
  - true/false
- contains数组中是否包含

## 组合模式

```json
{
  "anyOf": [ 
    {
      "type": "string", 
      "maxLength": 5
    }, 
    {
      "type": "string", 
      "minimum": 0
    }
  ]
}
```

- anyOf对所有子模式有效
- allOf对一个或多个子模似有效
- oneOf仅对一个子模式有效
- not不满足校验条件

## pattern

- ^仅匹配开头
- $仅匹配结尾
- (...|...|...)
- [abc]中括号内任意一个
- [a-z]
- [^abc]匹配未列出的任何字符
- +匹配前一个正则表达式的一个或多个
- *匹配前一个正则表达式的零个或多个
- ?匹配前一个正则表达式的零个或一个
- {x,y}匹配前一个正则表达式x到y次
