package repositories

import (
	"fmt"
	"log"

	"github.com/azura-labs/models"

	dbconnect "github.com/azura-labs/databases"
)

func lastUserID(id int) (response *models.ResponseUser) {
	var name, email, hp, address, createdAt, updatedAt, deletedAt string
	query := fmt.Sprintf(`SELECT id, name, email, hp, address, createdAt, updatedAt, deletedAt FROM users WHERE id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
		&email,
		&hp,
		&address,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if hp != "" {
		d := &models.ResponseUser{
			Success: true,
			Message: "Success",
			Data: &models.User{
				ID:        id,
				Name:      name,
				Email:     email,
				Hp:        hp,
				Address:   address,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
				DeletedAt: deletedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseUser{
		Success: false,
		Message: "User Not Found",
	}
	response = d
	return
}

func InsertUser(request *models.RequestUser) (response *models.ResponseUser) {
	var name, email, hp, address, createdAt, updatedAt, deletedAt string
	var id int
	if request.Name == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Name required",
		}
		response = d
		return
	} else if request.Email == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Email required",
		}
		response = d
		return
	} else if request.Hp == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "No. Hp required",
		}
		response = d
		return
	} else if request.Address == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Address required",
		}
		response = d
		return
	}
	checkEmail := fmt.Sprintf(`SELECT id, name, email, hp, address, createdAt, updatedAt, deletedAt FROM users WHERE email = '%s' AND deletedAt is null;`, request.Email)
	dbconnect.QueryRow(checkEmail).Scan(
		&id,
		&name,
		&email,
		&hp,
		&address,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if id > 0 {
		d := &models.ResponseUser{
			Success: false,
			Message: "Email Already Exist",
			Data: &models.User{
				ID:        id,
				Name:      name,
				Email:     email,
				Hp:        hp,
				Address:   address,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	checkHp := fmt.Sprintf(`SELECT id, name, email, hp, address, createdAt, updatedAt, deletedAt FROM users WHERE hp = '%s' AND deletedAt is null;`, request.Hp)
	dbconnect.QueryRow(checkHp).Scan(
		&id,
		&name,
		&email,
		&hp,
		&address,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if id > 0 {
		d := &models.ResponseUser{
			Success: false,
			Message: "No. Hp Already Exist",
			Data: &models.User{
				ID:        id,
				Name:      name,
				Email:     email,
				Hp:        hp,
				Address:   address,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		INSERT INTO users (name, email, hp, address, createdAt) VALUES ('%s', '%s', '%s', '%s', now());
	`,
		request.Name,
		request.Email,
		request.Hp,
		request.Address,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM users WHERE deletedAt is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastUserID(id)
	return
}

func UpdateUser(request *models.RequestUser, id int) (response *models.ResponseUser) {
	if request.Name == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Name required",
		}
		response = d
		return
	} else if request.Email == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Email required",
		}
		response = d
		return
	} else if request.Hp == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "No. Hp required",
		}
		response = d
		return
	} else if request.Address == "" {
		d := &models.ResponseUser{
			Success: false,
			Message: "Address required",
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		UPDATE users SET name = '%s', email = '%s', hp = '%s', address = '%s', updatedAt = now() WHERE id = %d;
	`,
		request.Name,
		request.Email,
		request.Hp,
		request.Address,
		id,
	)
	dbconnect.Exec(query)
	response = lastUserID(id)
	return
}

func DeleteUser(id int) (response *models.ResponseUser) {
	check := fmt.Sprintf(`SELECT id FROM users WHERE id = %d AND deletedAt is null;`, id)
	data := dbconnect.QueryRow(check).Scan(
		&id,
	)
	if data != nil {
		d := &models.ResponseUser{
			Success: false,
			Message: "User Not Found",
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		UPDATE users SET deletedAt = now() WHERE id = %d;
	`,
		id,
	)
	dbconnect.Exec(query)
	response = lastUserID(id)
	return
}

func DetailUser(id int) (response *models.ResponseUser) {
	var name, email, hp, address, createdAt, updatedAt string
	query := fmt.Sprintf(`SELECT id, name, email, hp, address, createdAt, updatedAt FROM users WHERE id = %d AND deletedAt is null;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
		&email,
		&hp,
		&address,
		&createdAt,
		&updatedAt,
	)
	if name != "" {
		d := &models.ResponseUser{
			Success: true,
			Message: "Success",
			Data: &models.User{
				ID:        id,
				Name:      name,
				Email:     email,
				Hp:        hp,
				Address:   address,
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseUser{
		Success: false,
		Message: "User Not Found",
	}
	response = d
	return
}

func ListUser(limit, skip int, q string) (response *models.ResponseUsers) {
	arrData := []models.User{}
	persen := "%"
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM users WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			email,
			hp,
			address,
			createdAt
		FROM
			users
		WHERE
			deletedAt is null
			AND name LIKE '%s%s%s'
			OR email LIKE '%s%s%s'
			OR hp LIKE '%s%s%s'
			OR address LIKE '%s%s%s'
		LIMIT %d OFFSET %d;
	`,
		persen,
		q,
		persen,
		persen,
		q,
		persen,
		persen,
		q,
		persen,
		persen,
		q,
		persen,
		limit,
		skip,
	)
	rowsQ, err := dbconnect.Query(query)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	var id int
	var name, email, hp, address, createdAt string
	// var updatedAt, deletedAt interface{}
	for rowsQ.Next() {
		err = rowsQ.Scan(
			&id,
			&name,
			&email,
			&hp,
			&address,
			&createdAt,
			// &updatedAt,
			// &deletedAt,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		rowData := models.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Hp:        hp,
			Address:   address,
			CreatedAt: createdAt,
			// UpdatedAt: helper.StringNullable(updatedAt),
			// DeletedAt: helper.StringNullable(deletedAt),
		}
		arrData = append(arrData, rowData)
	}
	d := &models.ResponseUsers{
		Success: true,
		Message: "Success",
		Data: &models.ListUser{
			Users: &arrData,
			Meta: &models.MetaDataUser{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
