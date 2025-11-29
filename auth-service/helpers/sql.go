package helpers

const InsertNewUser = `
	INSERT INTO users (email, password) 
	VALUES ($1, $2) 
	RETURNING id, email, is_email_verified, created_at;	
`
const FindUser = `
	SELECT *
	FROM users
	WHERE email = $1;
`
