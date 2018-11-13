package auth

const (
	enrollQuery = `
mutation EnrollUser {
  enroll(
    userId: "%s",
    email: "%s",
    username: "%s",
    password: "%s"
  )
}
`
	authenticateQuery = `
query authenticateUser{
  authenticate(
    userId: "%s",
    email: "%s",
    password: "%s"
  ) {
    token,
    userId,
    email,
    role
  }
}
`
)
