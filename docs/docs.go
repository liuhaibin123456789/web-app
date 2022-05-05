// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "cold bin",
            "url": "https://github.com/liuhaibin123456789/web-app.git"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/community": {
            "get": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "description": "获取前十个标签的id及名字，没有十个则返回所有",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "community"
                ],
                "summary": "获取社区分类信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "community"
                ],
                "summary": "创建社区分类",
                "parameters": [
                    {
                        "maxLength": 128,
                        "minLength": 1,
                        "description": "分类名",
                        "name": "community_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 256,
                        "description": "介绍",
                        "name": "introduction",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/community/{community_id}": {
            "get": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "description": "获取单个分类信息数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "community"
                ],
                "summary": "获取社区分类详细信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询帖子的community_id",
                        "name": "community_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "maxLength": 11,
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 16,
                        "minLength": 8,
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResJson"
                        }
                    }
                }
            }
        },
        "/post": {
            "post": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "创建帖子",
                "parameters": [
                    {
                        "description": "帖子json数据",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ReqPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/post/{post_id}": {
            "get": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "获取帖子详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询帖子的post_id",
                        "name": "post_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/post2": {
            "get": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "description": "支持时间和投票分数排序和查找社区分类下的帖子",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "查询的页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "查询的单页数据",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "只有两个值：` + "`" + `time` + "`" + `表示结果按照时间排序返回；` + "`" + `score` + "`" + `表示结果按照分数排序返回 ",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "为空表示默认不按照社区分类查询；不为空，将按照所给id对应社区分类的帖子返回",
                        "name": "community_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "maxLength": 11,
                        "minLength": 11,
                        "type": "string",
                        "description": "手机号",
                        "name": "phone",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "maxLength": 16,
                        "minLength": 8,
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResJson"
                        }
                    }
                }
            }
        },
        "/tokens": {
            "post": {
                "description": "该api只在access_token失效时使用，并且请求头携带好refresh_token",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "获取token",
                "parameters": [
                    {
                        "type": "string",
                        "description": "header auth",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户user_id",
                        "name": "user_id",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/tool.ResJson"
                        }
                    }
                }
            }
        },
        "/vote": {
            "post": {
                "security": [
                    {
                        "CoreAPI": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "帖子投票",
                "parameters": [
                    {
                        "description": "投票帖子post_id",
                        "name": "post_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "maxLength": 1,
                        "description": "投票帖子，1表示投票赞成，-1表示反对，0表示不投票或取消投票",
                        "name": "direction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/tool.ResJson"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Community": {
            "type": "object",
            "required": [
                "community_name"
            ],
            "properties": {
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "community_name": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "introduction": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "required": [
                "community_id",
                "content",
                "title"
            ],
            "properties": {
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "post_id": {
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "默认未通过审核",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "model.ReqPost": {
            "type": "object",
            "properties": {
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ResCommunity": {
            "type": "object",
            "properties": {
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "community_name": {
                    "type": "string"
                }
            }
        },
        "model.ResPost": {
            "type": "object",
            "required": [
                "community_id",
                "community_name",
                "content",
                "title"
            ],
            "properties": {
                "community_id": {
                    "type": "string",
                    "example": "0"
                },
                "community_name": {
                    "type": "string"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "introduction": {
                    "type": "string"
                },
                "post_id": {
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "description": "默认未通过审核",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "update_time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string",
                    "example": "0"
                },
                "user_name": {
                    "type": "string"
                },
                "vote_score": {
                    "type": "number"
                }
            }
        },
        "model.ResRegister": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "tool.ResJson": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {}
            }
        }
    },
    "securityDefinitions": {
        "CoreAPI": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8085",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "web_app项目接口文档",
	Description:      "投票帖子网站后端接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}