package repo

//type authRepo struct {
//	firestore *firestore.Client
//}
//
//func NewAuthRepository(firestore *firestore.Client) repo.AuthRepository {
//	return &authRepo{firestore: firestore}
//}
//
//func (r *authRepo) Login(ctx context.Context, auth *model.Auth) error {
//	return r.saveLoginActivity(ctx, auth)
//}
//
//func (r *authRepo) saveLoginActivity(ctx context.Context, auth *model.Auth) error {
//	err := r.autoLogout(ctx, auth)
//	if err != nil {
//		return err
//	}
//	_, err = r.firestore.Collection(db.LoginActivitiesCollectionPath).Doc(auth.Id).Set(ctx, auth)
//	return err
//}
//
//// autoLogout will logout all login activities.
//func (r *authRepo) autoLogout(ctx context.Context, auth *model.Auth) error {
//	iter := r.firestore.Collection(db.LoginActivitiesCollectionPath).
//		Where("Username", "==", auth.Username).
//		Where("IsLoggedOut", "==", false).
//		Documents(ctx)
//
//	docs, err := iter.GetAll()
//	if err != nil {
//		return err
//	}
//
//	if docs != nil && len(docs) > 0 {
//		var auths []*model.Auth
//		for _, doc := range docs {
//			a := new(model.Auth)
//			err = doc.DataTo(a)
//			if err != nil {
//				return err
//			}
//			auths = append(auths, a)
//		}
//		for _, a := range auths {
//			err := r.updateLogoutData(ctx, a)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
//
//func (r *authRepo) Logout(ctx context.Context, auth *model.Auth) error {
//	return r.saveLogoutActivity(ctx, auth)
//}
//
//func (r *authRepo) saveLogoutActivity(ctx context.Context, auth *model.Auth) error {
//	iter := r.firestore.Collection(db.LoginActivitiesCollectionPath).
//		Where("Username", "==", auth.Username).
//		Where("Id", "==", auth.Id).
//		Limit(1).
//		Documents(ctx)
//
//	docs, err := iter.GetAll()
//	if err != nil {
//		return err
//	}
//	if docs == nil || len(docs) < 1 {
//		return errors.New("login activity not found")
//	}
//
//	data := new(model.Auth)
//	if err = docs[0].DataTo(data); err != nil {
//		return err
//	}
//	if data.IsLoggedOut {
//		return perrors.ErrAlreadyLoggedOut
//	}
//
//	return r.updateLogoutData(ctx, auth)
//}
//
//func (r *authRepo) updateLogoutData(ctx context.Context, auth *model.Auth) error {
//	_, err := r.firestore.Collection(db.LoginActivitiesCollectionPath).Doc(auth.Id).
//		Update(ctx, []firestore.Update{
//			{Path: "LogoutTime", Value: time.Now()},
//			{Path: "IsLoggedOut", Value: true},
//		})
//	return err
//}
