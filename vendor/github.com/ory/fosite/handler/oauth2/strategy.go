/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 *
 */

package oauth2

import (
	"context"

	"github.com/ory/fosite"
)

type CoreStrategy interface {
	AccessTokenStrategy
	RefreshTokenStrategy
	AuthorizeCodeStrategy
}

type JWTStrategy interface {
	ValidateJWT(ctx context.Context, tokenType fosite.TokenType, token string) (requester fosite.Requester, err error)
}

type AccessTokenStrategy interface {
	AccessTokenSignature(token string) string
	GenerateAccessToken(ctx context.Context, requester fosite.Requester) (token string, signature string, err error)
	ValidateAccessToken(ctx context.Context, requester fosite.Requester, token string) (err error)
}

type RefreshTokenStrategy interface {
	RefreshTokenSignature(token string) string
	GenerateRefreshToken(ctx context.Context, requester fosite.Requester) (token string, signature string, err error)
	ValidateRefreshToken(ctx context.Context, requester fosite.Requester, token string) (err error)
}

type AuthorizeCodeStrategy interface {
	AuthorizeCodeSignature(token string) string
	GenerateAuthorizeCode(ctx context.Context, requester fosite.Requester) (token string, signature string, err error)
	ValidateAuthorizeCode(ctx context.Context, requester fosite.Requester, token string) (err error)
}