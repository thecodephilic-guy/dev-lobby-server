package helpers

const SelectUserById = `
	SELECT * 
	FROM users 
	WHERE id = $1;
`

const InsertNewOTP = `
	INSERT INTO verifications (user_id, code, expires_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id)
	DO UPDATE SET
		code = EXCLUDED.code,
		attempts = 0,
		expires_at = EXCLUDED.expires_at,
		created_at = NOW()
	RETURNING id, user_id, attempts, expires_at, created_at;
`
