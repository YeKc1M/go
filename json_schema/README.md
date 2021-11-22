# json schema

## string

- maxLength
- minLength
- enum

## boolean

## integer

## number

## const

## object

- required
- minProperties
- maxProperties

## array

- maxItems
- minItems

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
