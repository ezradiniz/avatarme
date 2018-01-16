package main

func main() {

	user := NewUser("user1@mail.com")
	avatar := NewAvatar(user, "./output/user-avatar.png")

	avatar.Draw()
	avatar.Create()
}
