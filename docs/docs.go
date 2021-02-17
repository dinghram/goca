// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "GoCA API Issues Report",
            "url": "http://github.com/kairoaraujo/goca/issues"
        },
        "license": {
            "name": "MIT",
            "url": "https://opensource.org/licenses/MIT"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/ca": {
            "get": {
                "description": "list all the Certificate Authorities",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA"
                ],
                "summary": "List Certificate Authorities (CA)",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseList"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            },
            "post": {
                "description": "create a new Certificate Authority Root or Intermediate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA"
                ],
                "summary": "Create new Certificate Authorities (CA) or Intermediate Certificate Authorities (ICA)",
                "parameters": [
                    {
                        "description": "Add new Certificate Authority or Intermediate Certificate Authority",
                        "name": "json_payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Payload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCA"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        },
        "/api/v1/ca/{cn}": {
            "get": {
                "description": "list the Certificate Authorities data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA"
                ],
                "summary": "Certificate Authorities (CA) Information based in Common Name",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCA"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        },
        "/api/v1/ca/{cn}/certificates": {
            "get": {
                "description": "list all certificates managed by a certain Certificate Authority (cn)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA/{CN}/Certificates"
                ],
                "summary": "List all Certificates managed by a certain Certificate Authority",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCA"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            },
            "post": {
                "description": "the Certificate Authority issues a new Certificate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA/{CN}/Certificates"
                ],
                "summary": "CA issue new certificate",
                "parameters": [
                    {
                        "description": "Add new Certificate Authority or Intermediate Certificate Authority",
                        "name": "ca",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Payload"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCertificates"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        },
        "/api/v1/ca/{cn}/certificates/{certificate_cn}": {
            "get": {
                "description": "get information about a certificate issued by a certain CA",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA/{CN}/Certificates"
                ],
                "summary": "Get information about a Certificate",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCertificates"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            },
            "delete": {
                "description": "the Certificate Authority revokes a managed Certificate",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA/{CN}/Certificates"
                ],
                "summary": "CA revoke a existent certificate managed by CA",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CABody"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        },
        "/api/v1/ca/{cn}/sign": {
            "post": {
                "description": "create a new certificate signing a Certificate Sigining Request (CSR)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA"
                ],
                "summary": "Certificate Authorities (CA) Signer for Certificate Sigining Request (CSR)",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Attached CSR file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number certificate valid days",
                        "name": "valid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCertificates"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        },
        "/api/v1/ca/{cn}/upload": {
            "post": {
                "description": "Upload a Certificate to a ICA pending certificate",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CA"
                ],
                "summary": "Upload a Certificate to an Intermediate CA",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Attached signed Certificate file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseCA"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "Internal"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "goca.CAData": {
            "type": "object",
            "properties": {
                "certificate": {
                    "description": "Certificate string",
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE-----...-----END CERTIFICATE-----\n"
                },
                "crl": {
                    "description": "Revocation List string",
                    "type": "string",
                    "example": "-----BEGIN X509 CRL-----...-----END X509 CRL-----\n"
                },
                "csr": {
                    "description": "Certificate Signing Request string",
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE REQUEST-----...-----END CERTIFICATE REQUEST-----\n"
                },
                "private_key": {
                    "description": "Private Key string",
                    "type": "string",
                    "example": "-----BEGIN PRIVATE KEY-----...-----END PRIVATE KEY-----\n"
                },
                "public_key": {
                    "description": "Public Key string",
                    "type": "string",
                    "example": "-----BEGIN PUBLIC KEY-----...-----END PUBLIC KEY-----\n"
                }
            }
        },
        "goca.Certificate": {
            "type": "object",
            "properties": {
                "ca_certificate": {
                    "description": "CA Certificate as string",
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE-----...-----END CERTIFICATE-----\n"
                },
                "certificate": {
                    "description": "Certificate certificate string",
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE-----...-----END CERTIFICATE-----\n"
                },
                "csr": {
                    "description": "Certificate Signing Request string",
                    "type": "string",
                    "example": "-----BEGIN CERTIFICATE REQUEST-----...-----END CERTIFICATE REQUEST-----\n"
                },
                "private_key": {
                    "description": "Certificate Private Key string",
                    "type": "string",
                    "example": "-----BEGIN PRIVATE KEY-----...-----END PRIVATE KEY-----\n"
                },
                "public_key": {
                    "description": "Certificate Public Key string",
                    "type": "string",
                    "example": "-----BEGIN PUBLIC KEY-----...-----END PUBLIC KEY-----\n"
                }
            }
        },
        "goca.Identity": {
            "type": "object",
            "properties": {
                "country": {
                    "description": "Country (two letters)",
                    "type": "string",
                    "example": "NL"
                },
                "dns_names": {
                    "description": "DNS Names list",
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "ca.example.com",
                        "root-ca.example.com"
                    ]
                },
                "email": {
                    "description": "Email Address",
                    "type": "string",
                    "example": "sec@company.com"
                },
                "intermediate": {
                    "description": "Intermendiate Certificate Authority (default is false)",
                    "type": "boolean",
                    "example": false
                },
                "key_size": {
                    "description": "Key Bit Size (defaul: 2048)",
                    "type": "integer",
                    "example": 2048
                },
                "locality": {
                    "description": "Locality name",
                    "type": "string",
                    "example": "Noord-Brabant"
                },
                "organization": {
                    "description": "Organization name",
                    "type": "string",
                    "example": "Company"
                },
                "organization_unit": {
                    "description": "Organizational Unit name",
                    "type": "string",
                    "example": "Security Management"
                },
                "province": {
                    "description": "Province name",
                    "type": "string",
                    "example": "Veldhoven"
                },
                "valid": {
                    "description": "Minimum 1 day, maximum 825 days -- Default: 397",
                    "type": "integer",
                    "example": 365
                }
            }
        },
        "models.CABody": {
            "type": "object",
            "properties": {
                "certificates": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "intranet.example.com",
                        "w3.example.com"
                    ]
                },
                "common_name": {
                    "type": "string",
                    "example": "root-ca"
                },
                "csr": {
                    "type": "boolean",
                    "example": false
                },
                "dns_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "ca.example.ca",
                        "root-ca.example.com"
                    ]
                },
                "expire_date": {
                    "type": "string",
                    "example": "2022-01-06 10:31:43 +0000 UTC"
                },
                "files": {
                    "$ref": "#/definitions/goca.CAData"
                },
                "intermediate": {
                    "type": "boolean"
                },
                "issue_date": {
                    "type": "string",
                    "example": "2021-01-06 10:31:43 +0000 UTC"
                },
                "revoked_certificates": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "38188836191244388427366318074605547405",
                        "338255903472757769326153358304310617728"
                    ]
                },
                "serial_number": {
                    "type": "string",
                    "example": "271064285308788403797280326571490069716"
                },
                "status": {
                    "type": "string",
                    "example": "Certificate Authority is ready."
                }
            }
        },
        "models.CertificateBody": {
            "type": "object",
            "properties": {
                "common_name": {
                    "type": "string",
                    "example": "intranet.go-root"
                },
                "dns_names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "w3.intranet.go-root.ca",
                        "intranet.go-root.ca"
                    ]
                },
                "expire_date": {
                    "type": "string",
                    "example": "2022-01-06 10:31:43 +0000 UTC"
                },
                "files": {
                    "$ref": "#/definitions/goca.Certificate"
                },
                "issue_date": {
                    "type": "string",
                    "example": "2021-01-06 10:31:43 +0000 UTC"
                },
                "serial_number": {
                    "type": "string",
                    "example": "338255903472757769326153358304310617728"
                }
            }
        },
        "models.Payload": {
            "type": "object",
            "required": [
                "common_name",
                "identity"
            ],
            "properties": {
                "common_name": {
                    "type": "string",
                    "example": "root-ca"
                },
                "identity": {
                    "$ref": "#/definitions/goca.Identity"
                }
            }
        },
        "models.ResponseCA": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CABody"
                }
            }
        },
        "models.ResponseCertificates": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/models.CertificateBody"
                }
            }
        },
        "models.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error message"
                }
            }
        },
        "models.ResponseList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "cn1",
                        "cn2",
                        "cn3"
                    ]
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{"http", "https"},
	Title:       "GoCA API",
	Description: "GoCA Certificate Authority Management API.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
