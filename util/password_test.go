package util

//func TestPassword(t *testing.T) {
//	password := RandomString(6)
//
//	hashedPassword, err := HashPasword(password)
//	require.NoError(t, err)
//	require.NotEmpty(t, hashedPassword)
//
//	err = CheckPasword(password, hashedPassword)
//	require.NoError(t, err)
//
//	wrongPassword := RandomString(6)
//	err = CheckPasword(wrongPassword, hashedPassword)
//	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
//}
