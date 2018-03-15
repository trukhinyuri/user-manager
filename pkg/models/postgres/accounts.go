package postgres

import (
	"errors"

	"context"

	"fmt"

	umtypes "git.containerum.net/ch/json-types/user-manager"
	"git.containerum.net/ch/user-manager/pkg/models"
	"github.com/sirupsen/logrus"
)

func (db *pgDB) GetUserByBoundAccount(ctx context.Context, service umtypes.OAuthResource, accountID string) (*models.User, error) {
	db.log.WithFields(logrus.Fields{
		"service":    service,
		"account_id": accountID,
	}).Infoln("Get bound account")

	switch service {
	case umtypes.GitHubOAuth, umtypes.FacebookOAuth, umtypes.GoogleOAuth:
	default:
		return nil, errors.New("unrecognised service " + string(service))
	}

	var ret models.User

	rows, err := db.qLog.QueryxContext(ctx, fmt.Sprintf(`SELECT users.id, users.login, users.password_hash, users.salt, users.role, users.is_active, users.is_deleted, users.is_in_blacklist
	FROM accounts JOIN users ON accounts.user_id = users.id WHERE accounts.%v = '%v'`, service, accountID))

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, rows.Err()
	}
	err = rows.StructScan(&ret)
	return &ret, err
}

func (db *pgDB) GetUserBoundAccounts(ctx context.Context, user *models.User) (*models.Accounts, error) {
	db.log.Infoln("Get bound accounts for user", user.Login)
	rows, err := db.qLog.QueryxContext(ctx, "SELECT id, github, facebook, google FROM accounts WHERE user_id = $1", user.ID)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, rows.Err()
	}

	ret := models.Accounts{User: user}
	err = rows.Scan(&ret.ID, &ret.Github, &ret.Facebook, &ret.Google)

	return &ret, err
}

func (db *pgDB) BindAccount(ctx context.Context, user *models.User, service umtypes.OAuthResource, accountID string) error {
	db.log.Infof("Bind account %s (%s) for user %s", service, accountID, user.Login)
	switch service {
	case umtypes.GitHubOAuth:
		_, err := db.eLog.ExecContext(ctx, `INSERT INTO accounts (user_id, github, facebook, google) 
													VALUES ($1, $2, '', '')
													ON CONFLICT (user_id) DO UPDATE SET github = $2`, user.ID, accountID)
		return err
	case umtypes.FacebookOAuth:
		_, err := db.eLog.ExecContext(ctx, `INSERT INTO accounts (user_id, github, facebook, google) 
													VALUES ($1, '', $2, '')
													ON CONFLICT (user_id) DO UPDATE SET facebook = $2`, user.ID, accountID)
		return err
	case umtypes.GoogleOAuth:
		_, err := db.eLog.ExecContext(ctx, `INSERT INTO accounts (user_id, github, facebook, google) 
													VALUES ($1, '', '', $2)
													ON CONFLICT (user_id) DO UPDATE SET google = $2`, user.ID, accountID)
		return err
	default:
		return errors.New("unrecognised service " + string(service))
	}
	// see migrations/1515872648_accounts_constraint.up.sql
}

func (db *pgDB) DeleteBoundAccount(ctx context.Context, user *models.User, service umtypes.OAuthResource) error {
	db.log.Infof("Deleting account %s for user %s", service, user.Login)
	switch service {
	case umtypes.GitHubOAuth, umtypes.FacebookOAuth, umtypes.GoogleOAuth:
	default:
		return errors.New("unrecognised service " + string(service))
	}

	_, err := db.eLog.ExecContext(ctx, fmt.Sprintf(`INSERT INTO accounts (user_id, github, facebook, google)
															VALUES ('%v', '', '', '')
															ON CONFLICT (user_id) DO UPDATE SET %v = ''`, user.ID, service))
	return err
}