package impl

import (
	"context"

	"time"

	"git.containerum.net/ch/grpc-proto-files/auth"
	"git.containerum.net/ch/grpc-proto-files/common"
	"git.containerum.net/ch/json-types/errors"
	mttypes "git.containerum.net/ch/json-types/mail-templater"
	umtypes "git.containerum.net/ch/json-types/user-manager"
	"git.containerum.net/ch/user-manager/models"
	"git.containerum.net/ch/user-manager/server"
	"git.containerum.net/ch/user-manager/utils"
)

func (u *serverImpl) ChangePassword(ctx context.Context, request umtypes.PasswordChangeRequest) (*auth.CreateTokenResponse, error) {
	userID := server.MustGetUserID(ctx)
	u.log.WithField("user_id", userID).Info("Changing password")

	user, err := u.svc.DB.GetUserByID(ctx, userID)
	if err := u.handleDBError(err); err != nil {
		return nil, userGetFailed
	}
	if err := u.loginUserChecks(ctx, user); err != nil {
		return nil, err
	}

	if !utils.CheckPassword(user.Login, request.CurrentPassword, user.Salt, user.PasswordHash) {
		return nil, &server.AccessDeniedError{Err: errors.New(invalidPassword)}
	}

	var tokens *auth.CreateTokenResponse

	err = u.svc.DB.Transactional(ctx, func(ctx context.Context, tx models.DB) error {
		user.PasswordHash = utils.GetKey(user.Login, request.NewPassword, user.Salt)
		if err := tx.UpdateUser(ctx, user); err != nil {
			return err
		}

		_, err := u.svc.AuthClient.DeleteUserTokens(ctx, &auth.DeleteUserTokensRequest{
			UserId: &common.UUID{Value: user.ID},
		})
		if err != nil {
			return tokenDeleteFailed
		}

		tokens, err = u.createTokens(ctx, user)
		return err
	})
	err = u.handleDBError(err)

	go func() {
		err := u.svc.MailClient.SendPasswordChangedMail(ctx, &mttypes.Recipient{
			ID:        user.ID,
			Name:      user.Login,
			Email:     user.Login,
			Variables: map[string]interface{}{},
		})
		if err != nil {
			u.log.WithError(err).Error("password change email send failed")
		}
	}()

	return tokens, err
}

func (u *serverImpl) ResetPassword(ctx context.Context, request umtypes.PasswordResetRequest) error {
	u.log.WithField("login", request.Username).Info("resetting password")
	user, err := u.svc.DB.GetUserByLogin(ctx, request.Username)
	if err := u.handleDBError(err); err != nil {
		return userGetFailed
	}
	if err := u.loginUserChecks(ctx, user); err != nil {
		return err
	}

	var link *models.Link
	err = u.svc.DB.Transactional(ctx, func(ctx context.Context, tx models.DB) error {
		var err error
		link, err = tx.CreateLink(ctx, umtypes.LinkTypePwdChange, 24*time.Hour, user)
		return err
	})
	if err := u.handleDBError(err); err != nil {
		return linkCreateFailed
	}

	go func() {
		err := u.svc.MailClient.SendPasswordResetMail(ctx, &mttypes.Recipient{
			ID:        user.ID,
			Name:      user.Login,
			Email:     user.Login,
			Variables: map[string]interface{}{"TOKEN": link.Link},
		})
		if err != nil {
			u.log.WithError(err).Error("password reset email send failed")
		}
	}()

	return nil
}

func (u *serverImpl) RestorePassword(ctx context.Context, request umtypes.PasswordRestoreRequest) (*auth.CreateTokenResponse, error) {
	u.log.WithField("link", request.Link).Info("Restore password")
	link, err := u.svc.DB.GetLinkFromString(ctx, request.Link)
	if err := u.handleDBError(err); err != nil {
		return nil, linkGetFailed
	}
	if link == nil {
		return nil, &server.NotFoundError{Err: errors.Format(linkNotFound, request.Link)}
	}
	if link.Type != umtypes.LinkTypePwdChange {
		return nil, &server.AccessDeniedError{Err: errors.Format(linkNotForPassword, request.Link)}
	}

	var tokens *auth.CreateTokenResponse

	err = u.svc.DB.Transactional(ctx, func(ctx context.Context, tx models.DB) error {
		link.User.PasswordHash = utils.GetKey(link.User.Login, request.NewPassword, link.User.Salt)
		if err := tx.UpdateUser(ctx, link.User); err != nil {
			return userUpdateFailed
		}
		link.IsActive = false

		_, err := u.svc.AuthClient.DeleteUserTokens(ctx, &auth.DeleteUserTokensRequest{
			UserId: &common.UUID{Value: link.User.ID},
		})
		if err != nil {
			return oneTimeTokenDeleteFailed
		}

		if err := tx.UpdateLink(ctx, link); err != nil {
			return linkUpdateFailed
		}

		tokens, err = u.createTokens(ctx, link.User)
		return err
	})
	if err := u.handleDBError(err); err != nil {
		return nil, err
	}

	go func() {
		err := u.svc.MailClient.SendPasswordChangedMail(ctx, &mttypes.Recipient{
			ID:        link.User.ID,
			Name:      link.User.Login,
			Email:     link.User.Login,
			Variables: map[string]interface{}{},
		})
		if err != nil {
			u.log.WithError(err).Error("password changed email send failed")
		}
	}()

	return tokens, nil
}