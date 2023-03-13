package main

type MyFriends struct {
	name                  string
	MobileNumber          int
	AlternateMobileNumber int
	Address               string
	City                  string
	checkAlternate        bool
}

type friend struct {
	addressbook []MyFriends
}

func (Myobj *friend) AddFriend(obj MyFriends) {

	Myobj.addressbook = append(Myobj.addressbook, obj)
}

func main() {
	Friend1 := MyFriends{
		name:                  "Kuldeep",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend2 := MyFriends{
		name:                  "Ram",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend3 := MyFriends{
		name:                  "Arjun",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend4 := MyFriends{
		name:                  "Sandeep",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend5 := MyFriends{
		name:                  "Vidhyut",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend6 := MyFriends{
		name:                  "Raju",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend7 := MyFriends{
		name:                  "Suraj",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}
	Friend8 := MyFriends{
		name:                  "Lion",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "SunderVan",
		checkAlternate:        true,
	}
	Friend9 := MyFriends{
		name:           "fredie",
		MobileNumber:   9090909090,
		Address:        "Partapur",
		City:           "Meerut",
		checkAlternate: false,
	}
	Friend10 := MyFriends{
		name:                  "Sandeep",
		MobileNumber:          9090909090,
		AlternateMobileNumber: 9087909090,
		Address:               "Partapur",
		City:                  "Meerut",
		checkAlternate:        true,
	}

	myFriendsObj := friend{}

	myFriendsObj.AddFriend(Friend1)
	myFriendsObj.AddFriend(Friend2)
	myFriendsObj.AddFriend(Friend3)
	myFriendsObj.AddFriend(Friend4)
	myFriendsObj.AddFriend(Friend5)
	myFriendsObj.AddFriend(Friend6)
	myFriendsObj.AddFriend(Friend7)
	myFriendsObj.AddFriend(Friend8)
	myFriendsObj.AddFriend(Friend9)
	myFriendsObj.AddFriend(Friend10)
}
