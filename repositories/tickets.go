package repositories

import (
	"fmt"
	"log"

	"github.com/azura-labs/models"

	dbconnect "github.com/azura-labs/databases"
)

func lastTicketID(id int) (response *models.ResponseTicket) {
	var name, description, createdAt, updatedAt, deletedAt string
	var price float64
	query := fmt.Sprintf(`SELECT id, name, price, description, createdAt, updatedAt, deletedAt FROM tickets WHERE id = %d;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
		&price,
		&description,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if name != "" {
		d := &models.ResponseTicket{
			Success: true,
			Message: "Success",
			Data: &models.Ticket{
				ID:          id,
				Name:        name,
				Price:       price,
				Description: description,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				DeletedAt:   deletedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseTicket{
		Success: false,
		Message: "Ticket Not Found",
	}
	response = d
	return
}

func InsertTicket(request *models.RequestTicket) (response *models.ResponseTicket) {
	var name, description, createdAt, updatedAt, deletedAt string
	var price float64
	var id int
	if request.Name == "" {
		d := &models.ResponseTicket{
			Success: false,
			Message: "Name ticket required",
		}
		response = d
		return
	}
	check := fmt.Sprintf(`SELECT id, name, price, description, createdAt, updatedAt, deletedAt FROM tickets WHERE name = '%s' AND deletedAt is null;`, request.Name)
	dbconnect.QueryRow(check).Scan(
		&id,
		&name,
		&price,
		&description,
		&createdAt,
		&updatedAt,
		&deletedAt,
	)
	if id > 0 {
		d := &models.ResponseTicket{
			Success: false,
			Message: "Name Ticket Already Exist",
			Data: &models.Ticket{
				ID:          id,
				Name:        name,
				Price:       price,
				Description: description,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				DeletedAt:   deletedAt,
			},
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		INSERT INTO tickets (name, price, description, createdAt) VALUES ('%s', %f, '%s', now());
	`,
		request.Name,
		request.Price,
		request.Description,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM tickets WHERE deletedAt is null;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastTicketID(id)
	return
}

func UpdateTicket(request *models.RequestTicket, id int) (response *models.ResponseTicket) {
	if request.Name == "" {
		d := &models.ResponseTicket{
			Success: false,
			Message: "Name ticket required",
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		UPDATE tickets SET name = '%s', price = %f, description = '%s', updatedAt = now() WHERE id = %d;
	`,
		request.Name,
		request.Price,
		request.Description,
		id,
	)
	dbconnect.Exec(query)
	response = lastTicketID(id)
	return
}

func DeleteTicket(id int) (response *models.ResponseTicket) {
	check := fmt.Sprintf(`SELECT id FROM tickets WHERE id = %d AND deletedAt is null;`, id)
	data := dbconnect.QueryRow(check).Scan(
		&id,
	)
	if data != nil {
		d := &models.ResponseTicket{
			Success: false,
			Message: "Ticket Not Found",
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		UPDATE tickets SET deletedAt = now() WHERE id = %d;
	`,
		id,
	)
	dbconnect.Exec(query)
	response = lastTicketID(id)
	return
}

func DetailTicket(id int) (response *models.ResponseTicket) {
	var name, description, createdAt, updatedAt string
	var price float64
	query := fmt.Sprintf(`SELECT id, name, price, description, createdAt, updatedAt FROM tickets WHERE id = %d AND deletedAt is null;`, id)
	dbconnect.QueryRow(query).Scan(
		&id,
		&name,
		&price,
		&description,
		&createdAt,
		&updatedAt,
	)
	if name != "" {
		d := &models.ResponseTicket{
			Success: true,
			Message: "Success",
			Data: &models.Ticket{
				ID:          id,
				Name:        name,
				Price:       price,
				Description: description,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponseTicket{
		Success: false,
		Message: "Ticket Not Found",
	}
	response = d
	return
}

func ListTicket(limit, skip int, q string) (response *models.ResponseTickets) {
	arrData := []models.Ticket{}
	persen := "%"
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM tickets WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	query := fmt.Sprintf(`
		SELECT
			id,
			name,
			price,
			description,
			createdAt
		FROM
			tickets
		WHERE
			deletedAt is null
			AND name LIKE '%s%s%s'
			OR price LIKE '%s%s%s'
			OR description LIKE '%s%s%s'
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
		limit,
		skip,
	)
	rowsQ, err := dbconnect.Query(query)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	var id int
	var name, description, createdAt string
	var price float64
	// var updatedAt, deletedAt interface{}
	for rowsQ.Next() {
		err = rowsQ.Scan(
			&id,
			&name,
			&price,
			&description,
			&createdAt,
			// &updatedAt,
			// &deletedAt,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		rowData := models.Ticket{
			ID:          id,
			Name:        name,
			Price:       price,
			Description: description,
			CreatedAt:   createdAt,
			// UpdatedAt: helper.StringNullable(updatedAt),
			// DeletedAt: helper.StringNullable(deletedAt),
		}
		arrData = append(arrData, rowData)
	}
	d := &models.ResponseTickets{
		Success: true,
		Message: "Success",
		Data: &models.ListTicket{
			Tickets: &arrData,
			Meta: &models.MetaDataTicket{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
