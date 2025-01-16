package customer_repository_postgres

const (
	// Customer Queries
	createCustomerQuery = `
	INSERT INTO "Customer" (nik, full_name, legal_name, birth_place, birth_date, salary, id_card_photo, selfie_photo) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING customer_id`

	getCustomerByIDQuery = `
	SELECT customer_id, nik, full_name, legal_name, birth_place, birth_date, salary, id_card_photo, selfie_photo 
	FROM "Customer" WHERE customer_id = $1`

	getCustomerByNikQuery = `
	SELECT nik, full_name, legal_name, birth_place, birth_date, salary, id_card_photo, selfie_photo 
	FROM "Customer" WHERE nik = $1
	`

	updateCustomerQuery = `
	UPDATE "Customer" SET nik = $2, full_name = $3, legal_name = $4, birth_place = $5, birth_date = $6, salary = $7, id_card_photo = $8, selfie_photo = $9 
	WHERE customer_id = $1`

	deleteCustomerQuery = `
	DELETE FROM "Customer" WHERE customer_id = $1`
)
