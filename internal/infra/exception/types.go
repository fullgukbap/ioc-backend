package exception

import "strings"

const (
	// 웹 라이브라리 에러(WEBSERVER_)
	ErrWebServerInternal = "WEBSERVER_INTERNAL_ERROR" // 웹 서버 내부 에러

	// 인증 관련 에러 (AUTH_)
	ErrAuthInvalid      = "AUTH_INVALID_CREDENTIALS" // 잘못된 인증 정보
	ErrAuthExpired      = "AUTH_TOKEN_EXPIRED"       // 토큰 만료
	ErrAuthUnauthorized = "AUTH_UNAUTHORIZED"        // 인증되지 않은 요청
	ErrAuthForbidden    = "AUTH_FORBIDDEN"           // 권한 없음

	// JWT 관련 에러 (JWT_)
	ErrJWTInvalid   = "JWT_INVALID_TOKEN"     // 유효하지 않은 JWT
	ErrJWTExpired   = "JWT_TOKEN_EXPIRED"     // JWT 만료
	ErrJWTMalformed = "JWT_MALFORMED"         // JWT 형식 오류
	ErrJWTSignature = "JWT_INVALID_SIGNATURE" // JWT 서명 불일치

	// OAuth 공통 에러 (OAUTH_)
	ErrOAuthInvalid = "OAUTH_INVALID_REQUEST" // 잘못된 OAuth 요청
	ErrOAuthState   = "OAUTH_INVALID_STATE"   // OAuth state 불일치
	ErrOAuthScope   = "OAUTH_INVALID_SCOPE"   // 잘못된 스코프

	// Kakao OAuth 에러 (OAUTH_KAKAO_)
	ErrKakaoToken   = "OAUTH_KAKAO_TOKEN_INVALID"  // 카카오 토큰 오류
	ErrKakaoProfile = "OAUTH_KAKAO_PROFILE_FAILED" // 프로필 조회 실패
	ErrKakaoRevoke  = "OAUTH_KAKAO_REVOKE_FAILED"  // 토큰 취소 실패

	// Google OAuth 에러 (OAUTH_GOOGLE_)
	ErrGoogleToken   = "OAUTH_GOOGLE_TOKEN_INVALID"  // 구글 토큰 오류
	ErrGoogleProfile = "OAUTH_GOOGLE_PROFILE_FAILED" // 프로필 조회 실패
	ErrGoogleRevoke  = "OAUTH_GOOGLE_REVOKE_FAILED"  // 토큰 취소 실패

	// Apple OAuth 에러 (OAUTH_APPLE_)
	ErrAppleToken   = "OAUTH_APPLE_TOKEN_INVALID"  // 애플 토큰 오류
	ErrAppleProfile = "OAUTH_APPLE_PROFILE_FAILED" // 프로필 조회 실패
	ErrAppleRevoke  = "OAUTH_APPLE_REVOKE_FAILED"  // 토큰 취소 실패
	ErrAppleJWT     = "OAUTH_APPLE_JWT_FAILED"     // 애플 JWT 검증 실패

	// MySQL 관련 에러 (MYSQL_)
	ErrMysqlInternal   = "MYSQL_INTERNAL_ERROR"       // 내부 MySQL 에러
	ErrMySQLConnection = "MYSQL_CONNECTION_FAILED"    // MySQL 연결 실패
	ErrMySQLDuplicate  = "MYSQL_DUPLICATE_ENTRY"      // 중복 데이터
	ErrMySQLConstraint = "MYSQL_CONSTRAINT_VIOLATION" // 제약조건 위반
	ErrMySQLTimeout    = "MYSQL_TIMEOUT"              // 쿼리 타임아웃
	ErrMySQLDeadlock   = "MYSQL_DEADLOCK"             // 데드락 발생
	ErrMySQLNotFound   = "MYSQL_NOT_FOUND"            // 데이터 없음

	// Redis 관련 에러 (REDIS_)
	ErrRedisConnection   = "REDIS_CONNECTION_FAILED" // Redis 연결 실패
	ErrRedisTimeout      = "REDIS_TIMEOUT"           // Redis 작업 타임아웃
	ErrRedisKeyNotFound  = "REDIS_KEY_NOT_FOUND"     // 키를 찾을 수 없음
	ErrRedisLockFailed   = "REDIS_LOCK_FAILED"       // 락 획득 실패
	ErrRedisScriptFailed = "REDIS_SCRIPT_FAILED"     // Lua 스크립트 실행 실패

	// 유효성 검사 에러 (VAL_)
	ErrValidation         = "VAL_INVALID_INPUT"    // 일반적인 입력값 오류
	ErrValidationEmail    = "VAL_INVALID_EMAIL"    // 이메일 형식 오류
	ErrValidationPassword = "VAL_INVALID_PASSWORD" // 비밀번호 형식 오류
	ErrValidationRequired = "VAL_REQUIRED_FIELD"   // 필수 필드 누락

	// 비즈니스 로직 에러 (BIZ_)
	ErrBizInvalid       = "BIZ_INVALID_OPERATION" // 잘못된 비즈니스 로직
	ErrBizConflict      = "BIZ_CONFLICT"          // 비즈니스 규칙 충돌
	ErrBizLimitExceeded = "BIZ_LIMIT_EXCEEDED"    // 제한 초과

	// 서버 에러 (SRV_)
	ErrServerInternal = "SRV_INTERNAL_ERROR" // 내부 서버 에러
	ErrServerOverload = "SRV_OVERLOADED"     // 서버 과부하
)

func IsMySQLError(err string) bool {
	return strings.HasPrefix(err, "MYSQL_")
}
