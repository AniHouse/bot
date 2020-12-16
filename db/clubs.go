package db

import (
	"errors"
	"time"

	"github.com/jackc/pgx"

	"github.com/brianvoe/gofakeit/v5"
)

type Club struct {
	ID          uint       `db:"id"`
	InsertedAt  *time.Time `db:"inserted_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	OwnerID     string     `db:"owner_id"`
	RoleID      *string    `db:"role_id"`
	ChannelID   *string    `db:"channel_id"`
	Title       string     `db:"title"`
	Description *string    `db:"description"`
	Symbol      string     `db:"symbol"`
	IconURL     *string    `db:"icon_url"`
	XP          uint64     `db:"xp"`
	ExpiredAt   *time.Time `db:"expired_at"`
	Verified    bool       `db:"verified"`
}

func (c *Club) randomize() {
	desc := gofakeit.Paragraph(1, 1, 10, "")
	chid := gofakeit.Numerify("test##############")
	rlid := gofakeit.Numerify("test##############")

	c.OwnerID = gofakeit.Numerify("test##############")
	c.ChannelID = &chid
	c.RoleID = &rlid
	c.Title = gofakeit.Word()
	c.Symbol = gofakeit.Emoji()
	c.Description = &desc
}

func (c *Club) AddMember(tx *pgx.Tx, memberID string) (err error) {
	_, err = tx.Exec(`
		INSERT INTO "club_members"("club_id","user_id") 
		VALUES($1, $2)
		ON CONFLICT DO NOTHING
	`,
		c.ID,
		memberID,
	)
	return
}

func (c *Club) DeleteMember(memberID string) (err error) {
	_, err = pgxconn.Exec(`
		DELETE FROM club_members 
		WHERE user_id = $1
	`,
		memberID,
	)
	return
}

func (c *Club) DeleteMembers() (err error) {
	_, err = pgxconn.Exec(`
		DELETE FROM club_members 
		WHERE club_id = $1
	`,
		c.ID,
	)
	return
}

func (c *Club) HasMember(memberID string) (result bool, err error) {
	err = pgxconn.QueryRow(`
		SELECT EXISTS(SELECT 1 FROM club_members WHERE user_id = $1)
	`,
		memberID,
	).Scan(
		&result,
	)
	return
}

func (c *Club) Delete() (err error) {
	_, err = pgxconn.Exec(`
		DELETE FROM clubs WHERE id = $1;
		DELETE FROM club_members WHERE club_id = $1;
	`,
		c.ID,
	)
	return
}

type ClubMember struct {
	ClubID     uint       `db:"club_id"`
	UserID     string     `db:"user_id"`
	InsertedAt *time.Time `db:"inserted_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	XP         uint64     `db:"xp"`
}

type clubs struct{}

func (c *clubs) Create(club *Club) (err error) {
	err = pgxconn.QueryRow(`
		WITH club_id as (
			INSERT INTO clubs(owner_id, title, symbol, expired_at)
			VALUES ($1, $2, $3, $4)
			RETURNING id
		), owner_user as (
			INSERT INTO club_members(club_id, user_id) VALUES
			((SELECT id FROM club_id LIMIT 1), $1)
		)
		SELECT id FROM club_id
	`,
		club.OwnerID,
		club.Title,
		club.Symbol,
		club.ExpiredAt,
	).Scan(&club.ID)
	return
}

func (c *clubs) DeleteByOwner(ownerID string) (err error) {
	_, err = pgxconn.Exec(`
		WITH club_id as (
			DELETE FROM clubs WHERE owner_id = $1
			RETURNING id
		), owner_user as (
			DELETE FROM club_members
			WHERE club_id = (SELECT id FROM club_id LIMIT 1)
		)
		SELECT * FROM club_id
	`,
		ownerID,
	)
	return
}

func (c *clubs) GetClubByUser(userID string) (club *Club, err error) {
	club = new(Club)
	err = pgxconn.QueryRow(`
		SELECT
			c.id,
			c.inserted_at,
			c.updated_at,
			c.owner_id,
			c.role_id,
			c.channel_id,
			c.title,
			c.description,
			c.symbol,
			c.icon_url,
			c.xp,
			c.expired_at,
			c.verified
		FROM clubs c
		JOIN club_members cm on c.id = cm.club_id
		WHERE cm.user_id = $1
	`,
		userID,
	).Scan(
		&club.ID,
		&club.InsertedAt,
		&club.UpdatedAt,
		&club.OwnerID,
		&club.RoleID,
		&club.ChannelID,
		&club.Title,
		&club.Description,
		&club.Symbol,
		&club.IconURL,
		&club.XP,
		&club.ExpiredAt,
		&club.Verified,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return club, err
}

func (c *clubs) GetExpired() ([]Club, error) {
	rows, err := pgxconn.Query(`
		SELECT
			id,
			inserted_at,
			updated_at,
			owner_id,
			role_id,
			channel_id,
			title,
			description,
			symbol,
			icon_url,
			xp,
			expired_at,
			verified
		FROM clubs
		WHERE NOT verified
			AND localtimestamp >= expired_at
	`)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var clubs []Club
	for rows.Next() {
		var club Club
		err = rows.Scan(
			&club.ID,
			&club.InsertedAt,
			&club.UpdatedAt,
			&club.OwnerID,
			&club.RoleID,
			&club.ChannelID,
			&club.Title,
			&club.Description,
			&club.Symbol,
			&club.IconURL,
			&club.XP,
			&club.ExpiredAt,
			&club.Verified,
		)

		if err != nil {
			return nil, err
		}
		clubs = append(clubs, club)
	}
	return clubs, nil
}

func (c *clubs) RemoveExpired() ([]Club, error) {
	rows, err := pgxconn.Query(`
		WITH club_id as (
			DELETE FROM clubs WHERE NOT verified
				AND localtimestamp >= expired_at
			RETURNING *
		), owner_user as (
			DELETE FROM club_members
			WHERE club_id IN (SELECT id FROM club_id)
		)
		SELECT 
			id,
			inserted_at,
			updated_at,
			owner_id,
			role_id,
			channel_id,
			title,
			description,
			symbol,
			icon_url,
			xp,
			expired_at,
			verified
		FROM club_id
	`)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var clubs []Club
	for rows.Next() {
		var club Club
		err = rows.Scan(
			&club.ID,
			&club.InsertedAt,
			&club.UpdatedAt,
			&club.OwnerID,
			&club.RoleID,
			&club.ChannelID,
			&club.Title,
			&club.Description,
			&club.Symbol,
			&club.IconURL,
			&club.XP,
			&club.ExpiredAt,
			&club.Verified,
		)

		if err != nil {
			return nil, err
		}
		clubs = append(clubs, club)
	}
	return clubs, nil
}
