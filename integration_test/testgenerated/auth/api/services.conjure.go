// This file was generated by Conjure and should not be manually edited.

package api

import (
	"context"
	"fmt"

	"github.com/palantir/conjure-go-runtime/conjure-go-client/httpclient"
	"github.com/palantir/pkg/bearertoken"
)

type BothAuthServiceClient interface {
	Default(ctx context.Context, authHeader bearertoken.Token) (string, error)
	Cookie(ctx context.Context, cookieToken bearertoken.Token) error
	None(ctx context.Context) error
	WithArg(ctx context.Context, authHeader bearertoken.Token, argArg string) error
}

type bothAuthServiceClient struct {
	client httpclient.Client
}

func NewBothAuthServiceClient(client httpclient.Client) BothAuthServiceClient {
	return &bothAuthServiceClient{client: client}
}

func (c *bothAuthServiceClient) Default(ctx context.Context, authHeader bearertoken.Token) (string, error) {
	var defaultReturnVal string
	var returnVal *string
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Default"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint(authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/default"))
	requestParams = append(requestParams, httpclient.WithJSONResponse(&returnVal))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return defaultReturnVal, err
	}
	_ = resp
	if returnVal == nil {
		return defaultReturnVal, fmt.Errorf("returnVal cannot be nil")
	}
	return *returnVal, nil
}

func (c *bothAuthServiceClient) Cookie(ctx context.Context, cookieToken bearertoken.Token) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Cookie"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("P_TOKEN", fmt.Sprint(cookieToken)))
	requestParams = append(requestParams, httpclient.WithPathf("/cookie"))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *bothAuthServiceClient) None(ctx context.Context) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("None"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithPathf("/none"))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

func (c *bothAuthServiceClient) WithArg(ctx context.Context, authHeader bearertoken.Token, argArg string) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("WithArg"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("POST"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint(authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/withArg"))
	requestParams = append(requestParams, httpclient.WithJSONRequest(argArg))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

type BothAuthServiceClientWithAuth interface {
	Default(ctx context.Context) (string, error)
	Cookie(ctx context.Context) error
	None(ctx context.Context) error
	WithArg(ctx context.Context, argArg string) error
}

func NewBothAuthServiceClientWithAuth(client BothAuthServiceClient, authHeader bearertoken.Token, cookieToken bearertoken.Token) BothAuthServiceClientWithAuth {
	return &bothAuthServiceClientWithAuth{client: client, authHeader: authHeader, cookieToken: cookieToken}
}

type bothAuthServiceClientWithAuth struct {
	client      BothAuthServiceClient
	authHeader  bearertoken.Token
	cookieToken bearertoken.Token
}

func (c *bothAuthServiceClientWithAuth) Default(ctx context.Context) (string, error) {
	return c.client.Default(ctx, c.authHeader)
}

func (c *bothAuthServiceClientWithAuth) Cookie(ctx context.Context) error {
	return c.client.Cookie(ctx, c.cookieToken)
}

func (c *bothAuthServiceClientWithAuth) None(ctx context.Context) error {
	return c.client.None(ctx)
}

func (c *bothAuthServiceClientWithAuth) WithArg(ctx context.Context, argArg string) error {
	return c.client.WithArg(ctx, c.authHeader, argArg)
}

type HeaderAuthServiceClient interface {
	Default(ctx context.Context, authHeader bearertoken.Token) (string, error)
}

type headerAuthServiceClient struct {
	client httpclient.Client
}

func NewHeaderAuthServiceClient(client httpclient.Client) HeaderAuthServiceClient {
	return &headerAuthServiceClient{client: client}
}

func (c *headerAuthServiceClient) Default(ctx context.Context, authHeader bearertoken.Token) (string, error) {
	var defaultReturnVal string
	var returnVal *string
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Default"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("Authorization", fmt.Sprint(authHeader)))
	requestParams = append(requestParams, httpclient.WithPathf("/default"))
	requestParams = append(requestParams, httpclient.WithJSONResponse(&returnVal))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return defaultReturnVal, err
	}
	_ = resp
	if returnVal == nil {
		return defaultReturnVal, fmt.Errorf("returnVal cannot be nil")
	}
	return *returnVal, nil
}

type HeaderAuthServiceClientWithAuth interface {
	Default(ctx context.Context) (string, error)
}

func NewHeaderAuthServiceClientWithAuth(client HeaderAuthServiceClient, authHeader bearertoken.Token) HeaderAuthServiceClientWithAuth {
	return &headerAuthServiceClientWithAuth{client: client, authHeader: authHeader}
}

type headerAuthServiceClientWithAuth struct {
	client     HeaderAuthServiceClient
	authHeader bearertoken.Token
}

func (c *headerAuthServiceClientWithAuth) Default(ctx context.Context) (string, error) {
	return c.client.Default(ctx, c.authHeader)
}

type CookieAuthServiceClient interface {
	Cookie(ctx context.Context, cookieToken bearertoken.Token) error
}

type cookieAuthServiceClient struct {
	client httpclient.Client
}

func NewCookieAuthServiceClient(client httpclient.Client) CookieAuthServiceClient {
	return &cookieAuthServiceClient{client: client}
}

func (c *cookieAuthServiceClient) Cookie(ctx context.Context, cookieToken bearertoken.Token) error {
	var requestParams []httpclient.RequestParam
	requestParams = append(requestParams, httpclient.WithRPCMethodName("Cookie"))
	requestParams = append(requestParams, httpclient.WithRequestMethod("GET"))
	requestParams = append(requestParams, httpclient.WithHeader("P_TOKEN", fmt.Sprint(cookieToken)))
	requestParams = append(requestParams, httpclient.WithPathf("/cookie"))
	resp, err := c.client.Do(ctx, requestParams...)
	if err != nil {
		return err
	}
	_ = resp
	return nil
}

type CookieAuthServiceClientWithAuth interface {
	Cookie(ctx context.Context) error
}

func NewCookieAuthServiceClientWithAuth(client CookieAuthServiceClient, cookieToken bearertoken.Token) CookieAuthServiceClientWithAuth {
	return &cookieAuthServiceClientWithAuth{client: client, cookieToken: cookieToken}
}

type cookieAuthServiceClientWithAuth struct {
	client      CookieAuthServiceClient
	cookieToken bearertoken.Token
}

func (c *cookieAuthServiceClientWithAuth) Cookie(ctx context.Context) error {
	return c.client.Cookie(ctx, c.cookieToken)
}
