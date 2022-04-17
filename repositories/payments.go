package repositories

import (
	"fmt"
	"log"

	"github.com/azura-labs/models"

	dbconnect "github.com/azura-labs/databases"
)

func lastPaymentID(id int) (response *models.ResponsePayment) {
	var name, email, hp, addrs, ticket, description, createdAt string
	var usrId, ticId int
	var price float64
	query := fmt.Sprintf(`
		SELECT
			p.id,
			p.user_id,
			u.name,
			u.email,
			u.hp,
			u.address,
			p.ticket_id,
			t.name as ticket,
			t.price,
			t.description,
			p.createdAt
		FROM
			payments p
		INNER JOIN (
			SELECT
				u.id,
				u.name,
				u.email,
				u.hp,
				u.address
			FROM
				users u
			WHERE
				u.deletedAt is null) as u on
			u.id = p.user_id
		INNER JOIN (
			SELECT
				t.id,
				t.name,
				t.price,
				t.description
			FROM
				tickets t
			WHERE
				t.deletedAt is null) as t on
			t.id = p.ticket_id
		WHERE 
			p.id = %d;
	`,
		id,
	)
	dbconnect.QueryRow(query).Scan(
		&id,
		&usrId,
		&name,
		&email,
		&hp,
		&addrs,
		&ticId,
		&ticket,
		&price,
		&description,
		&createdAt,
	)
	if name != "" {
		d := &models.ResponsePayment{
			Success: true,
			Message: "Success",
			Data: &models.Payment{
				ID:          id,
				UserID:      usrId,
				Name:        name,
				Email:       email,
				Hp:          hp,
				Address:     addrs,
				TicketID:    ticId,
				Ticket:      ticket,
				Price:       price,
				Description: description,
				CreatedAt:   createdAt,
			},
		}
		response = d
		return
	}
	d := &models.ResponsePayment{
		Success: false,
		Message: "Payment Not Found",
	}
	response = d
	return
}

func InsertPayment(request *models.RequestPayment) (response *models.ResponsePayment) {
	var id int
	if request.UserID == 0 {
		d := &models.ResponsePayment{
			Success: false,
			Message: "User required",
		}
		response = d
		return
	}
	if request.TicketID == 0 {
		d := &models.ResponsePayment{
			Success: false,
			Message: "Ticket required",
		}
		response = d
		return
	}
	checkUsr := fmt.Sprintf(`SELECT id FROM users WHERE id = %d AND deletedAt is null;`, request.UserID)
	usr := dbconnect.QueryRow(checkUsr).Scan(
		&id,
	)
	if usr != nil {
		d := &models.ResponsePayment{
			Success: false,
			Message: "User Not Found",
		}
		response = d
		return
	}
	checkTic := fmt.Sprintf(`SELECT id FROM tickets WHERE id = %d AND deletedAt is null;`, request.TicketID)
	tic := dbconnect.QueryRow(checkTic).Scan(
		&id,
	)
	if tic != nil {
		d := &models.ResponsePayment{
			Success: false,
			Message: "Ticket Not Found",
		}
		response = d
		return
	}
	query := fmt.Sprintf(`
		INSERT INTO payments (user_id, ticket_id, createdAt) VALUES (%d, %d, now());
	`,
		request.UserID,
		request.TicketID,
	)
	dbconnect.Exec(query)
	last := fmt.Sprintf(`SELECT max(id) as id FROM payments;`)
	dbconnect.QueryRow(last).Scan(&id)
	response = lastPaymentID(id)
	return
}

func DetailPayment(id int) (response *models.ResponsePayment) {
	response = lastPaymentID(id)
	return
}

func ListPayment(limit, skip int, q string) (response *models.ResponsePayments) {
	arrData := []models.Payment{}
	persen := "%"
	count := fmt.Sprintf(`SELECT COUNT(id) AS total FROM tickets WHERE deletedAt is null;`)
	total := dbconnect.QueryCount(count)
	query := fmt.Sprintf(`
		SELECT
			p.id,
			p.user_id,
			u.name,
			u.email,
			u.hp,
			u.address,
			p.ticket_id,
			t.name as ticket,
			t.price,
			t.description,
			p.createdAt
		FROM
			payments p
		INNER JOIN (
			SELECT
				u.id,
				u.name,
				u.email,
				u.hp,
				u.address
			FROM
				users u
			WHERE
				u.deletedAt is null) as u on
			u.id = p.user_id
		INNER JOIN (
			SELECT
				t.id,
				t.name,
				t.price,
				t.description
			FROM
				tickets t
			WHERE
				t.deletedAt is null) as t on
			t.id = p.ticket_id
		WHERE
			u.name LIKE '%s%s%s'
			OR u.email LIKE '%s%s%s'
			OR u.hp LIKE '%s%s%s'
			OR u.address LIKE '%s%s%s'
			OR t.name LIKE '%s%s%s'
			OR t.price LIKE '%s%s%s'
			OR t.description LIKE '%s%s%s'
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
	var name, email, hp, addrs, ticket, description, createdAt string
	var id, usrId, ticId int
	var price float64
	for rowsQ.Next() {
		err = rowsQ.Scan(
			&id,
			&usrId,
			&name,
			&email,
			&hp,
			&addrs,
			&ticId,
			&ticket,
			&price,
			&description,
			&createdAt,
		)

		if err != nil {
			log.Printf(err.Error())
			return
		}

		rowData := models.Payment{
			ID:          id,
			UserID:      usrId,
			Name:        name,
			Email:       email,
			Hp:          hp,
			Address:     addrs,
			TicketID:    ticId,
			Ticket:      ticket,
			Price:       price,
			Description: description,
			CreatedAt:   createdAt,
		}
		arrData = append(arrData, rowData)
	}
	d := &models.ResponsePayments{
		Success: true,
		Message: "Success",
		Data: &models.ListPayment{
			Payments: &arrData,
			Meta: &models.MetaDataPayment{
				Total: total,
				Limit: limit,
				Skip:  skip,
			},
		},
	}

	response = d
	return
}
